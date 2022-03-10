// SPDX-FileCopyrightText: 2022 Alvar Penning
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"

	syscallset "github.com/oxzi/syscallset-go"
)

func main() {
	if err := syscallset.LimitTo("@basic-io"); err != nil {
		panic(err)
	}

	// Opening files is not allowed; check out @file-system
	_, _ = os.Open(os.Args[0])

	fmt.Println("hello restricted world")
}

// TEST_ERRORS=1
