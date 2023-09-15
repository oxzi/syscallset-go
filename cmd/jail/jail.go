// SPDX-FileCopyrightText: Alvar Penning
//
// SPDX-License-Identifier: BSD-3-Clause

// Demo program to demonstrate how to limit another program.
//
// Usage: ./jail [-log] FILTER CMD ARGS...

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	syscallset "github.com/oxzi/syscallset-go"
)

// exit in case of an error with an error message and exit code 1.
func exit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

// execute the jailed command with redirection of stdin, stderr, and stdout.
func execute(cmd []string) {
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin

	if err := c.Run(); err != nil {
		exit("Running %q failed: %v", strings.Join(cmd, " "), err)
	}
}

func main() {
	if len(os.Args) < 3 || os.Args[1] == "-log" && len(os.Args) < 4 {
		exit("Usage: %s [-log] FILTER CMD ARGS...", os.Args[0])
	}

	// Additional syscall groups are required to spawn the jailed process.
	filter := "@basic-io @file-system @process @ipc @signal "

	var f func(string) error
	var cmd []string

	if os.Args[1] == "-log" {
		f = syscallset.LimitAndLog
		filter += os.Args[2]
		cmd = os.Args[3:]
	} else {
		f = syscallset.LimitTo
		filter += os.Args[1]
		cmd = os.Args[2:]
	}

	if !syscallset.IsSupported() {
		exit("Filtering not supported on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
	if err := f(filter); err != nil {
		exit("Cannot apply filter: %v", err)
	}

	execute(cmd)
}
