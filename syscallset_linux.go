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

	for _, obj := range strings.Split(strings.TrimSpace(syscallFilter), " ") {
		// If an object is prefixed by "~", it will be removed from the previously
		// allowed syscalls. By default, new syscalls are added to the set.
		isAdditive := true
		if obj[0] == '~' {
			isAdditive = false
			obj = obj[1:]
		}

		var tmpSyscalls []string

		if obj[0] == '@' {
			// For a set, filter unsupported syscalls first.
			set, setExist := syscallSets[obj[1:]]
			if !setExist {
				return nil, fmt.Errorf("syscall set %q does not exist", obj)
			}

			if tmpSyscalls, err = filterUnsupportedSyscalls(set); err != nil {
				return
			}
		} else {
			// Do not filter for single syscalls, as this should error.
			tmpSyscalls = []string{obj}
		}

		for _, syscall := range tmpSyscalls {
			if isAdditive {
				syscallMap[syscall] = struct{}{}
			} else {
				delete(syscallMap, syscall)
			}
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
	syscalls, err := unwrapSyscalls("@default " + syscallFilter)
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

// IsSupported returns true if filtering syscalls through seccomp-bpf is
// possible on this platform.
func IsSupported() bool {
	if !seccomp.Supported() {
		return false
	}

	_, infoErr := arch.GetInfo(runtime.GOARCH)
	return infoErr == nil
}

// LimitTo a subset of the available Linux syscalls using a systemd system call
// filter string.
//
// A filter string might contain both syscall sets, prefixed by an at sign (@),
// as well as single syscalls by their name. The list of syscall sets is either
// available in this package's main documentation or can be  fetched from
// systemd's exec documentation:
//
//	https://www.freedesktop.org/software/systemd/man/systemd.exec.html#System%20Call%20Filtering
//
// The filter acts as an allow list. Thus, every other syscall results in the
// termination of the process and its children. One can remove single syscalls
// or smaller sets again by prefixing them with a tilde (~). As the set of
// allowed syscalls is created by parsing the words from left to right, one
// should start with building the allow list and reducing it afterwards.
//
// A small subset of syscalls (@default) is always allowed. Thus, when calling
// with an empty string, a very strict filter is applied, not even allowing
// using stdin or stdout.
//
// A simple example with systemd's wide @system-service might be:
//
//	@system-service
//
// Allowing some IO and file system access might be achieved through:
//
//	@basic-io @file-system @io-event
//
// To restrict a wider set might be used like the following:
//
//	@system-service ~@process ~@setuid
func LimitTo(syscallFilter string) error {
	return limit(syscallFilter, seccomp.ActionKillProcess)
}

// LimitAndLog acts like LimitTo; however, non allowed syscalls are being logged
// instead of resulting in aborting the process. This might be useful for
// testing the application.
func LimitAndLog(syscallFilter string) error {
	return limit(syscallFilter, seccomp.ActionLog)
}
