// SPDX-FileCopyrightText: Alvar Penning
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build !linux

// This file is a stub for non-supported platforms.

package syscallset

import "errors"

func IsSupported() bool {
	return false
}

func LimitTo(syscallFilter string) error {
	return errors.New("unsupported platform")
}

func LimitAndLog(syscallset string) error {
	return errors.New("unsupported platform")
}
