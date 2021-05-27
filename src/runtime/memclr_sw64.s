// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "textflag.h"

// func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
TEXT runtime·memclrNoHeapPointers(SB),NOSPLIT,$0-16
	LDL	R1, ptr+0(FP)
	LDL	R2, n+8(FP)
	ADDL	R1, R2, R4
	
	// if less than 8 bytes, do one byte at a time
	CMPULT	R2, $8, R3
	BNE	R3, out
	
	// do one byte at a time until 8-aligned
	AND	R1, $7, R3
	BEQ	R3, words
	
	STB	R31, (R1)
	ADDL    R1, $1, R1
	JMP	-4(PC)

words:
	// do 8 bytes at a time if there is room
	ADDL	R4, $-7, R2
	
	CMPULE	R2, R1, R3
	BNE	R3, out
	
	STL	R31, (R1)
	ADDL	R1, $8, R1
	JMP	-4(PC)

out:
	CMPEQ	R1, R4, R3
	BNE	R3, done
	
	STB	R31, (R1)
	ADDL	R1, $1, R1
	JMP	-4(PC)
done:
	RET
