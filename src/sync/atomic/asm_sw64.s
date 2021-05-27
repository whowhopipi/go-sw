// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "textflag.h"

TEXT ·SwapInt32(SB), NOFRAME|NOSPLIT,$0-20
	JMP	·SwapUint32(SB)

TEXT ·SwapUint32(SB), NOFRAME|NOSPLIT,$0-20
	LDL	R3, addr+0(FP)
	LDW	R4, new+8(FP)
	MEMB
	LLDW	R5, 0(R3)
	LDI	R6, 1
	WR_F	R6
	BIS	R4, R31, R6
	LSTW	R6, 0(R3)
	RD_F	R6
	BEQ	R6, -6(PC)
	STW	R5, old+16(FP)
	RET

TEXT ·SwapInt64(SB), NOFRAME|NOSPLIT, $0-24
	JMP	·SwapUint64(SB)

TEXT ·SwapUintptr(SB), NOFRAME|NOSPLIT, $0-24
	JMP	·SwapUint64(SB)

TEXT ·SwapUint64(SB), NOFRAME|NOSPLIT, $0-24
	LDL	R3, addr+0(FP)
	LDL	R4, new+8(FP)
	MEMB
	LLDL	R5, 0(R3)
	LDI	R6, 1
	WR_F	R6
	BIS	R4, R31, R6
	LSTL	R6, 0(R3)
	RD_F	R6
	BEQ	R6, -6(PC)
	STL	R5, old+16(FP)
	RET

TEXT ·CompareAndSwapInt32(SB), NOFRAME|NOSPLIT, $0-17
	JMP	·CompareAndSwapUint32(SB)

TEXT ·CompareAndSwapUint32(SB), NOFRAME|NOSPLIT, $0-17
	LDL	R3, addr+0(FP)
	LDW	R4, old+8(FP)
	LDW	R5, new+12(FP)
	MEMB
	LLDW	R6, 0(R3)
	CMPEQ	R6, R4, R7
	WR_F	R7
	BIS	R5, R31, R6
	LSTW	R6, 0(R3)
	RD_F	R6
	BEQ	R7, 2(PC)
	BEQ	R6, -7(PC)
	STB	R6, swapped+16(FP)
	RET

TEXT ·CompareAndSwapUintptr(SB), NOFRAME|NOSPLIT, $0-25
	JMP	·CompareAndSwapUint64(SB)

TEXT ·CompareAndSwapInt64(SB), NOFRAME|NOSPLIT, $0-25
	JMP	·CompareAndSwapUint64(SB)

TEXT ·CompareAndSwapUint64(SB), NOFRAME|NOSPLIT, $0-25
	LDL	R3, addr+0(FP)
	LDL	R4, old+8(FP)
	LDL	R5, new+16(FP)
	MEMB
	LLDL	R6, 0(R3)
	CMPEQ	R6, R4, R7
	WR_F	R7
	BIS	R5, R31, R6
	LSTL	R6, 0(R3)
	RD_F	R6
	BEQ	R7, 2(PC)
	BEQ	R6, -7(PC)
	STB	R6, swapped+24(FP)
	RET

TEXT ·AddInt32(SB), NOFRAME|NOSPLIT, $0-20
	JMP	·AddUint32(SB)

TEXT ·AddUint32(SB), NOFRAME|NOSPLIT, $0-20
	LDL	R3, addr+0(FP)
	LDW	R4, delta+8(FP)
	MEMB
	LLDW	R5, 0(R3)
	LDI	R6, 1
	WR_F	R6
	ADDW	R5, R4, R6
	LSTW	R6, 0(R3)
	RD_F	R6
	BEQ	R6, -6(PC)
	ADDW	R5, R4, R5
	STW	R5, new+16(FP)
	RET

TEXT ·AddUintptr(SB), NOFRAME|NOSPLIT, $0-24
	JMP	·AddUint64(SB)

TEXT ·AddInt64(SB), NOFRAME|NOSPLIT, $0-24
	JMP	·AddUint64(SB)

TEXT ·AddUint64(SB), NOFRAME|NOSPLIT, $0-24
	LDL	R3, addr+0(FP)
	LDL	R4, delta+8(FP)
	MEMB
	LLDL	R5, 0(R3)
	LDI	R6, 1
	WR_F	R6
	ADDL	R5, R4, R6
	LSTL	R6, 0(R3)
	RD_F	R6
	BEQ	R6, -6(PC)
	ADDL	R5, R4, R5
	STL	R5, new+16(FP)
	RET

TEXT ·StoreInt32(SB), NOFRAME|NOSPLIT, $0-12
	JMP	·StoreUint32(SB)

TEXT ·StoreUint32(SB), NOFRAME|NOSPLIT, $0-12
	LDL	R3, addr+0(FP)
	LDW	R4, val+8(FP)
	MEMB
	STW	R4, 0(R3)
	MEMB
	RET

TEXT ·StoreInt64(SB), NOFRAME|NOSPLIT, $0-16
	JMP	·StoreUint64(SB)

TEXT ·StoreUintptr(SB), NOFRAME|NOSPLIT, $0-16
	JMP	·StoreUint64(SB)

TEXT ·StoreUint64(SB), NOFRAME|NOSPLIT, $0-16
	LDL	R3, addr+0(FP)
	LDL	R4, val+8(FP)
	MEMB
	STL	R4, 0(R3)
	MEMB
	RET

TEXT ·LoadInt32(SB), NOFRAME|NOSPLIT, $0-12
	JMP	·LoadUint32(SB)

TEXT ·LoadUint32(SB), NOFRAME|NOSPLIT, $0-12
	LDL	R3, addr+0(FP)
	MEMB
	LDW	R3, 0(R3)
	MEMB
	STW	R3, val+8(FP)
	RET

TEXT ·LoadInt64(SB), NOFRAME|NOSPLIT, $0-16
	JMP	·LoadUint64(SB)

TEXT ·LoadUintptr(SB), NOFRAME|NOSPLIT,$0-16
	JMP	·LoadPointer(SB)

TEXT ·LoadPointer(SB), NOFRAME|NOSPLIT,$0-16
	JMP	·LoadUint64(SB)

TEXT ·LoadUint64(SB), NOFRAME|NOSPLIT, $0-16
	LDL	R3, addr+0(FP)
	MEMB
	LDL	R3, 0(R3)
	MEMB
	STL	R3, val+8(FP)
	RET
