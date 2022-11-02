// SPDX-FileCopyrightText: 2022 Alvar Penning
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

// This file transforms systemd's syscall sets through Golang's template engine
// to another format, e.g., Go code.
//
// Usage: go run generator/gen.go TEMPLATE_PATH
//        go run generator/gen.go generator/go.tmpl

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

// syscallFilterParser for `systemd-analyze syscall-filter`.
type syscallFilterParser struct {
	scanner *bufio.Scanner

	err  error
	sets map[string]map[string]bool
}

// nextLine for processing or an error. Commented lines are being skipped.
func (parser *syscallFilterParser) nextLine() (string, error) {
	reComment := regexp.MustCompile(`^\s*#.*`)

	for {
		if parser.scanner.Scan() {
			line := parser.scanner.Text()
			if !reComment.MatchString(line) {
				return line, nil
			}
		} else {
			err := parser.scanner.Err()
			if err == nil {
				err = io.EOF
			}
			return "", err
		}
	}
}

// parseSet from the output and store it within the parser's sets.
func (parser *syscallFilterParser) parseSet() {
	var setName string
	var setSyscalls map[string]bool

	reSetName := regexp.MustCompile(`^@([a-z-]+)$`)
	reSyscall := regexp.MustCompile(`^\s+(@?[a-z0-9-_]+)$`)

	// Store the parsed data afterwards.
	defer func() {
		if parser.err != nil {
			return
		}
		parser.sets[setName] = setSyscalls
	}()

	// First, parse the set's name.
	for {
		if nameLine, err := parser.nextLine(); err != nil {
			parser.err = err
			return
		} else if nameLine == "" {
			continue
		} else if name := reSetName.FindStringSubmatch(nameLine); len(name) == 0 {
			parser.err = fmt.Errorf("expected set name, got %q", nameLine)
			return
		} else {
			setName = name[1]
			break
		}
	}

	setSyscalls = make(map[string]bool)

	// Now parse all syscalls.
	for {
		if syscallLine, err := parser.nextLine(); err != nil {
			parser.err = err
			return
		} else if syscallLine == "" {
			// Reached an empty line; end of set.
			return
		} else if syscall := reSyscall.FindStringSubmatch(syscallLine); len(syscall) == 0 {
			parser.err = fmt.Errorf("expected syscall, got %q", syscallLine)
			return
		} else {
			setSyscalls[syscall[1]] = true
		}
	}
}

// syscallFilters from `systemd-analyze syscall-filter` as a map of the set's
// name pointing to an array of syscalls and/or other sets.
func syscallFilters() (map[string]map[string]bool, error) {
	out, err := exec.Command("systemd-analyze", "syscall-filter").Output()
	if err != nil {
		return nil, err
	}

	parser := &syscallFilterParser{
		scanner: bufio.NewScanner(bytes.NewBuffer(out)),
		sets:    make(map[string]map[string]bool),
	}

	for parser.err == nil {
		parser.parseSet()
	}
	if parser.err != io.EOF {
		return nil, fmt.Errorf("cannot parse output, %w", parser.err)
	}

	return parser.sets, nil
}

// syscallSetFlatten removes the internal references from one set to another.
// Thus, the output map's value will only contain the names of syscalls.
func syscallSetFlatten(in map[string]map[string]bool) (out map[string]map[string]bool, err error) {
	out = make(map[string]map[string]bool)

	for setName, setSyscalls := range in {
		tmpSyscalls := setSyscalls

		// If one set refers to another, do another check as this could also point
		// to a third one. This code might result in an infinite loop, but I am
		// pretty sure the systemd people have prevented this.. Otherwise, this is
		// just the generator.
		for nestedCheck := true; nestedCheck; {
			nestedCheck = false

			outSyscalls := make(map[string]bool)
			for syscall, _ := range tmpSyscalls {
				if strings.HasPrefix(syscall, "@") {
					if nestedSetSyscalls, nestedSetExist := in[syscall[1:]]; !nestedSetExist {
						err = fmt.Errorf("referenced set %q does not exist", syscall)
						return
					} else {
						for nestedSyscall, _ := range nestedSetSyscalls {
							outSyscalls[nestedSyscall] = true
						}
						nestedCheck = true
					}
				} else {
					outSyscalls[syscall] = true
				}
			}
			tmpSyscalls = outSyscalls
		}

		out[setName] = tmpSyscalls
	}

	return
}

// syscallSetMoveExecve from the default set to process.
// This syscall is obviously needed for systemd to spawn another process, but
// otherwise would not be expected in a default set.
func syscallSetMoveExecve(syscallSets map[string]map[string]bool) error {
	defaultSet, defaultOk := syscallSets["default"]
	processSet, processOk := syscallSets["process"]

	if !defaultOk || !processOk {
		return fmt.Errorf("set is missing; default = %t, process = %t", defaultOk, processOk)
	}

	delete(defaultSet, "execve")
	processSet["execve"] = true

	syscallSets["default"] = defaultSet
	syscallSets["process"] = processSet
	return nil
}

// systemdVersion as a multi line string to be displayed in the generated file.
func systemdVersion() (string, error) {
	out, err := exec.Command("systemd-analyze", "--version").Output()
	if err != nil {
		return "", err
	}

	return strings.Split(string(out), "\n")[0], nil
}

func main() {
	// Fetch syscall sets from systemd
	syscallSets, err := syscallFilters()
	if err != nil {
		panic(err)
	}
	syscallSets, err = syscallSetFlatten(syscallSets)
	if err != nil {
		panic(err)
	}
	err = syscallSetMoveExecve(syscallSets)
	if err != nil {
		panic(err)
	}

	// Fetch systemd version string
	systemdVersionStr, err := systemdVersion()
	if err != nil {
		panic(err)
	}

	// Fetch template from .tmpl file, given as first argument
	if len(os.Args) != 2 {
		panic("Usage: generator/gen.go TEMPLATE_PATH")
	}
	tmpl, err := template.ParseFiles(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Generate output and write back to stdout
	tmplVars := struct {
		SystemdVersionStr string
		SyscallSets       map[string]map[string]bool
	}{
		SystemdVersionStr: systemdVersionStr,
		SyscallSets:       syscallSets,
	}
	err = tmpl.Execute(os.Stdout, tmplVars)
	if err != nil {
		panic(err)
	}
}
