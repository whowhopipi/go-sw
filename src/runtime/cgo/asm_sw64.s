// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "textflag.h"

/*
 * void crosscall2(void (*fn)(void*, int32, uintptr), void*, int32, uintptr)
 * Save registers and call fn with two arguments.
 */
TEXT crosscall2(SB), NOFRAME|NOSPLIT, $0
	/*
	 * We still need to save all callee save register as before, and then
	 *  push 3 args for fn (R17, R18, R19).
	 * Also note that at procedure entry in gc world, 8(SP) will be the
	 *  first arg.
	 */
	LDI	SP, $-8*23(SP)

	STL	R17, (8*1)(SP) // void*
	STL	R18, (8*2)(SP) // int32
	STL	R19, (8*3)(SP) // uintptr

	STL	R9, (8*4)(SP)
	STL	R10, (8*5)(SP)
	STL	R11, (8*6)(SP)
	STL	R12, (8*7)(SP)
	STL	R13, (8*8)(SP)
	STL	R14, (8*9)(SP)

	STL	RSB, (8*12)(SP)
	STL	g, (8*13)(SP)
	STL	R26, (8*14)(SP)

	FSTD	F2, (8*15)(SP)
	FSTD	F3, (8*16)(SP)
	FSTD	F4, (8*17)(SP)
	FSTD	F5, (8*18)(SP)
	FSTD	F6, (8*19)(SP)
	FSTD	F7, (8*20)(SP)
	FSTD	F8, (8*21)(SP)
	FSTD	F9, (8*22)(SP)

	CALL	runtime·load_g(SB)
	LDI	R27, R16
	CALL	R26, (R27)

	LDL	R9, (8*4)(SP)
	LDL	R10, (8*5)(SP)
	LDL	R11, (8*6)(SP)
	LDL	R12, (8*7)(SP)
	LDL	R13, (8*8)(SP)
	LDL	R14, (8*9)(SP)

	LDL	RSB, (8*12)(SP)
	LDL	g, (8*13)(SP)
	LDL	R26, (8*14)(SP)

	FLDD	F2, (8*15)(SP)
	FLDD	F3, (8*16)(SP)
	FLDD	F4, (8*17)(SP)
	FLDD	F5, (8*18)(SP)
	FLDD	F6, (8*19)(SP)
	FLDD	F7, (8*20)(SP)
	FLDD	F8, (8*21)(SP)
	FLDD	F9, (8*22)(SP)


	LDI	SP, $8*23(SP)
	RET
