// SPDX-FileCopyrightText: Alvar Penning
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux

package main

import (
	"fmt"

	syscallset "github.com/oxzi/syscallset-go"
)

func main() {
	if err := syscallset.LimitTo("@system-service ~@process ~@setuid ~@basic-io"); err != nil {
		panic(err)
	}

	fmt.Println("I'm still alive.. NOT")
}

// TEST_ERRORS=1
