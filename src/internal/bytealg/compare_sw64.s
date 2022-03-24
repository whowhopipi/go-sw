// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "go_asm.h"
#include "textflag.h"

TEXT ·Compare(SB),NOFRAME|NOSPLIT,$0-56
	LDL	R2, a_base+0(FP)
	LDL	R0, a_len+8(FP)
	LDL	R3, b_base+24(FP)
	LDL	R1, b_len+32(FP)
	CMPEQ	R2, R3, R4
	BNE	R4, samebytes
	CMPULT	R0, R1, R4
	SELEQ	R4, R1, R0, R5    // R5 is min(R0,R1)

	ADDL	R5, R2, R5
loop:
	CMPEQ	R5, R2, R8
	BNE	R8, samebytes   // all compared bytes were the same; compare lengths
	LDBU	R6, (R2)
	ADDL	R2, $1, R2
	LDBU	R7, (R3)
	ADDL	R3, $1, R3
	CMPEQ	R6, R7, R8
	BNE	R8, loop
	// bytes differed
	CMPULT	R7, R6, R8
	LDI	R6, $-1
	SELEQ	R8, R6, R8, R8
	JMP	cmp_ret
samebytes:
	CMPULT	R1, R0, R6
	CMPULT	R0, R1, R7
	SUBL	R6, R7, R8
cmp_ret:
	STL	R8, ret+48(FP)
	RET

TEXT runtime·cmpstring(SB),NOFRAME|NOSPLIT,$0-40
	LDL	R2, a_base+0(FP)
	LDL	R0, a_len+8(FP)
	LDL	R3, b_base+16(FP)
	LDL	R1, b_len+24(FP)
	CMPEQ	R2, R3, R4
	BNE	R4, samebytes
	CMPULT	R0, R1, R4
	SELEQ	R4, R1, R0, R5    // R5 is min(R0,R1)

	ADDL	R5, R2, R5
loop:
	CMPEQ	R5, R2, R8
	BNE	R8, samebytes   // all compared bytes were the same; compare lengths

	LDBU	R6, (R2)
	ADDL	R2, $1, R2
	LDBU	R7, (R3)
	ADDL	R3, $1, R3
	CMPEQ	R6, R7, R8
	BNE	R8, loop
	// bytes differed
	CMPULT	R7, R6, R8
	LDI	R6, $-1
	SELEQ	R8, R6, R8, R8
	JMP	cmp_ret
samebytes:
	CMPULT	R1, R0, R6
	CMPULT	R0, R1, R7
	SUBL	R6, R7, R8
cmp_ret:
	STL	R8, ret+32(FP)
	RET
