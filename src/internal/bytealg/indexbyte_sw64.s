// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "go_asm.h"
#include "textflag.h"

TEXT ·IndexByte(SB),NOFRAME|NOSPLIT,$0-40
	LDL	R1, b_base+0(FP)
	LDL	R2, b_len+8(FP)
	BEQ	R2, notfound
	LDBU	R3, c+24(FP)
	LDI	R4, R1
	ADDL	R1, R2, R2 // end
	SUBL	R1, $1, R1

loop:
	ADDL	R1, $1, R1
	CMPEQ	R1, R2, R5
	BNE	R5, notfound
	
	LDBU	R6, (R1)
	CMPEQ	R6, R3, R6
	BEQ	R6, loop
	
	SUBL	R1, R4, R1 // remove base
	STL	R1, ret+32(FP)
	RET

notfound:
	LDI	R1, $-1
	STL	R1, ret+32(FP)
	RET

TEXT ·IndexByteString(SB),NOFRAME|NOSPLIT,$0-32
	LDL	R1, s_base+0(FP)
	LDL	R2, s_len+8(FP)
	BEQ	R2, notfound
	LDBU	R3, c+16(FP)
	LDI	R4, R1
	ADDL	R1, R2, R2 // end
	SUBL	R1, $1, R1

loop:
	ADDL	R1, $1, R1
	CMPEQ	R1, R2, R5
	BNE	R5, notfound
	
	LDBU	R6, (R1)
	CMPEQ	R6, R3, R6
	BEQ	R6, loop
	
	SUBL	R1, R4, R1 // remove base
	STL	R1, ret+24(FP)
	RET

notfound:
	LDI	R1, $-1
	STL	R1, ret+24(FP)
	RET
