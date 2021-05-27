// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "go_asm.h"
#include "textflag.h"

// memequal(a, b unsafe.Pointer, size uintptr) bool
TEXT runtime·memequal(SB),NOSPLIT,$0-25
	LDL	R2, $size+16(FP)
	BEQ	R2, eq
	LDL	R0, $p+0(FP)
	LDL	R1, $q+8(FP)
	CMPEQ	R0, R1, R3
	BNE	R3, eq
loop:
	LDBU	R3, (R0)
	LDBU	R4, (R1)
	
	CMPEQ	R3, R4, R3
	BEQ	R3, ne
	
	ADDL	R0, $1, R0
	ADDL	R1, $1, R1
	SUBL	R2, $1, R2
	BNE	R2, loop
eq:
	LDI	R0, $1
	STB	R0, $ret+24(FP)
	RET
ne:
	STB	ZERO, $ret+24(FP)
	RET

// memequal_varlen(a, b unsafe.Pointer) bool
TEXT runtime·memequal_varlen(SB),NOSPLIT,$40-17
	LDL	R1, a+0(FP)
	LDL	R2, b+8(FP)
	CMPEQ	R1, R2, R11
	BNE	R11, eq
	LDL	R3, 8(REGCTXT) // compiler stores size at offset 8 in the closure
	STL	R1, 8(SP)
	STL	R2, 16(SP)
	STL	R3, 24(SP)
	
	CALL	runtime·memequal(SB)
	LDBU	R1, 32(SP)
	STB 	R1, ret+16(FP)
	RET
eq:
	LDI	R1, $1
	STB	R1, ret+16(FP)
	RET

