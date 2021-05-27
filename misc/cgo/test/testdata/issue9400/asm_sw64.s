// Copyright 2016 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64
// +build !gccgo

#include "textflag.h"

TEXT ·RewindAndSetgid(SB), NOSPLIT|NOFRAME, $0-0
	// Rewind stack pointer so anything that happens on the stack
	// will clobber the test pattern created by the caller
	LDI	SP, $(1024*8)(SP)
	
	// Ask signaller to setgid
	LDI	R1, $1
	MEMB
	STW	R1, ·Baton(SB)
	MEMB
	
	// Wait for setgid completion
loop:
	MEMB
	LDW	R1, ·Baton(SB)
	NOOP  // hint that we're in a spin loop
	BNE    R1, loop
	MEMB
	
	LDI	SP, $-(1024*8)(SP)
	RET
