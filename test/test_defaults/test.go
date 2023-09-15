// SPDX-FileCopyrightText: Alvar Penning
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux
// +build linux

package main

import (
	syscallset "github.com/oxzi/syscallset-go"
)

func main() {
	if err := syscallset.LimitTo(""); err != nil {
		panic(err)
	}
}

// TEST_ERRORS=0
