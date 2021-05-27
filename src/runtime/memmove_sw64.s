// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "textflag.h"

// func memmove(to, from unsafe.Pointer, n uintptr)
TEXT runtime·memmove(SB), NOFRAME|NOSPLIT, $0-24
	LDL	R0, to+0(FP)
	LDL	R1, from+8(FP)
	LDL	R2, n+16(FP)
	BNE	R2, check
	RET

check:
	CMPULE	R0,R1,R4
	BEQ	R4,backward
	
	ADDL	R0,R2,R6 // end pointer
	
	 // if the two pointers are not of same alignments, do byte copying
	SUBL	R0,R1,R4
	AND	R4,$7,R4
	BNE	R4, out
	
	 // if less than 8 bytes, do byte copying
	CMPULT	R2,$8,R4
	BNE	R4,out
	// do one byte at a time until 8-aligned
	AND	R0,$7,R5
	BEQ	R5,words
	LDBU	R3, (R1)
	ADDL	R1, $1, R1
	
	STB	R3, (R0)
	ADDL	R0,$1,R0
	JMP	-6(PC)

words:
	// do 8 bytes at a time if there is room
	SUBL	R6,$7,R7
	CMPULE	R7,R0,R8
	BNE	R8,out
	LDL	R3,(R1)
	ADDL	R1,$8,R1
	
	STL	R3,(R0)
	ADDL	R0,$8,R0
	JMP	-6(PC)

out:
	CMPEQ	R0,R6,R8
	BNE	R8,done
	LDBU	R3, (R1)
	ADDL	R1, $1, R1
	
	STB	R3, (R0)
	ADDL	R0,$1,R0
	JMP	-6(PC)

done:
	RET

backward:
	ADDL	R1,R2,R4 // from-end pointer
	ADDL	R0,R2,R5  // to-end pointer
	
	// if the two pointers are not of same alignments, do byte copying
	SUBL	R4,R5,R7
	AND	R7,$7,R7
	BNE	R7,out1
	
	 // if less than 8 bytes, do byte copying
	CMPULT	R2,$8,R7
	BNE	R7,out1
	
	// do one byte at a time until 8-aligned
	AND	R5,$7,R6
	BEQ	R6,words1
	SUBL	R4,$1,R4
	LDBU	R3, (R4)
	SUBL	R5,$1,R5
	STB	R3,(R5)
	JMP	-6(PC)

words1:
	// do 8 bytes at a time if there is room
	ADDL	R0,$7,R3
	
	CMPULE	R5,R3,R7
	BNE	R7,out1
	SUBL	R4,$8,R4
	LDL	R7,(R4)
	SUBL	R5,$8,R5
	STL	R7,(R5)
	JMP	-6(PC)


out1:
	CMPEQ	R0,R5,R7
	BNE	R7,done1
	SUBL	R4,$1,R4
	LDBU	R3, (R4)
	SUBL	R5,$1,R5
	STB	R3,(R5)
	JMP	-6(PC)

done1:
	RET
