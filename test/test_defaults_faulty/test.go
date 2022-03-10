// SPDX-FileCopyrightText: 2022 Alvar Penning
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux
// +build linux

package main

import (
	"fmt"

	syscallset "github.com/oxzi/syscallset-go"
)

func main() {
	if err := syscallset.LimitTo(""); err != nil {
		panic(err)
	}

	fmt.Println("oh noes")
}

// TEST_ERRORS=1
