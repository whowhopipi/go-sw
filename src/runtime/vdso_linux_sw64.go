// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux
// +build sw64

package runtime

const (
	// vdsoArrayMax is the byte-size of a maximally sized array on this architecture.
	// See cmd/compile/internal/sw64/galign.go arch.MAXWIDTH initialization.
	vdsoArrayMax = 1<<50 - 1
)

var vdsoLinuxVersion = vdsoVersionKey{"LINUX_2.6.39", 0x75fcb89}

// The symbol name is not __kernel_clock_gettime as suggested by the manpage;
// according to Linux source code it should be __vdso_clock_gettime instead.
var vdsoSymbolKeys = []vdsoSymbolKey{
	{"__vdso_clock_gettime", 0xd35ec75, 0x6e43a318, &vdsoClockgettimeSym},
}

// initialize to fall back to syscall
var (
	vdsoClockgettimeSym uintptr = 0
)
