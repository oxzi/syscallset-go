// SPDX-FileCopyrightText: 2022 Alvar Penning
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux
// +build linux

package syscallset

import (
	"fmt"
	"runtime"
	"strings"

	seccomp "github.com/elastic/go-seccomp-bpf"
	"github.com/elastic/go-seccomp-bpf/arch"
)

// filterUnsupportedSyscalls from other architectures as the current one.
func filterUnsupportedSyscalls(in []string) (out []string, err error) {
	info, infoErr := arch.GetInfo(runtime.GOARCH)
	if infoErr != nil {
		err = infoErr
		return
	}

	for _, syscall := range in {
		if _, exist := info.SyscallNames[syscall]; !exist {
			continue
		}
		out = append(out, syscall)
	}

	return
}

// unwrapSyscalls from the given syscall filter string to an array of syscalls.
func unwrapSyscalls(syscallFilter string) (syscalls []string, err error) {
	syscallMap := map[string]struct{}{}

	for _, obj := range strings.Split(syscallFilter, " ") {
		if strings.HasPrefix(obj, "@") {
			// For a set, filter unsupported syscalls first.
			set, setExist := syscallSets[obj[1:]]
			if !setExist {
				return nil, fmt.Errorf("syscall set %q does not exist", obj)
			}

			setSupported, setSupportedErr := filterUnsupportedSyscalls(set)
			if setSupportedErr != nil {
				return nil, setSupportedErr
			}

			for _, syscall := range setSupported {
				syscallMap[syscall] = struct{}{}
			}
		} else {
			// Do not filter for single syscalls, as this should error.
			syscallMap[obj] = struct{}{}
		}
	}

	for k := range syscallMap {
		syscalls = append(syscalls, k)
	}
	return
}

// limit the usable syscalls by applying a seccomp-bpf filter. The given action
// will be performed for unspecified syscalls.
func limit(syscallFilter string, action seccomp.Action) error {
	syscalls, err := unwrapSyscalls(syscallFilter)
	if err != nil {
		return err
	}

	return seccomp.LoadFilter(seccomp.Filter{
		NoNewPrivs: true,
		Flag:       seccomp.FilterFlagTSync,
		Policy: seccomp.Policy{
			DefaultAction: action,
			Syscalls: []seccomp.SyscallGroup{
				{
					Action: seccomp.ActionAllow,
					Names:  syscalls,
				},
			},
		},
	})
}

// LimitTo a subset of the available Linux syscalls using a systemd system call
// filter string.
//
// A filter string might contain both syscall sets, prefixed by an at sign (@),
// as well as single syscalls by their name. The list of syscall sets can be
// fetched from systemd's exec documentation:
//
//   https://www.freedesktop.org/software/systemd/man/systemd.exec.html#System%20Call%20Filtering
//
// The filter acts as an allow list. Thus, every other syscall results in the
// whole processes being killed.
func LimitTo(syscallFilter string) error {
	return limit(syscallFilter, seccomp.ActionKillProcess)
}

// LimitAndLog acts like LimitTo; however, non allowed syscalls are being logged
// instead of resulting in aborting the process. This might be useful for
// testing the application.
func LimitAndLog(syscallFilter string) error {
	return limit(syscallFilter, seccomp.ActionLog)
}
