// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "textflag.h"

// bool cas(uint32 *ptr, uint32 old, uint32 new)
// Atomically:
//    if(*val == old){
//    *val = new;
//    return 1;
//    } else
//    return 0;
TEXT runtime∕internal∕atomic·Cas(SB), NOSPLIT|NOFRAME, $0-17
	LDL	R3, ptr+0(FP)
	LDW	R4, old+8(FP)
	LDW	R5, new+12(FP)

	MEMB
	LLDW	R6, 0(R3)
	CMPEQ	R6, R4, R6
	WR_F	R6
	BIS	R5, R31, R7
	LSTW	R7, 0(R3)
	RD_F	R7
	BEQ	R6, 2(PC)
	BEQ	R7, -7(PC)
	STB	R7, ret+16(FP)
	RET

// bool    runtime∕internal∕atomic·Cas64(uint64 *ptr, uint64 old, uint64 new)
// Atomically:
//    if(*val == *old){
//    *val = new;
//    return 1;
//    } else {
//    return 0;
//    }
TEXT runtime∕internal∕atomic·Cas64(SB), NOFRAME|NOSPLIT, $0-25
	LDL	R3, ptr+0(FP)
	LDL	R4, old+8(FP)
	LDL	R5, new+16(FP)

	MEMB
	LLDL	R6, 0(R3)
	CMPEQ	R6, R4, R6
	WR_F	R6
	BIS	R5, R31, R7
	LSTL	R7, 0(R3)
	RD_F	R7
	BEQ	R6, 2(PC)
	BEQ	R7, -7(PC)
	STB	R7, ret+24(FP)
	RET

TEXT runtime∕internal∕atomic·Casuintptr(SB), NOFRAME|NOSPLIT, $0-25
	JMP	runtime∕internal∕atomic·Cas64(SB)

TEXT runtime∕internal∕atomic·Loaduintptr(SB),  NOFRAME|NOSPLIT, $0-16
	JMP	runtime∕internal∕atomic·Load64(SB)

TEXT runtime∕internal∕atomic·Loaduint(SB), NOFRAME|NOSPLIT, $0-16
	JMP	runtime∕internal∕atomic·Load64(SB)

TEXT runtime∕internal∕atomic·Storeuintptr(SB), NOFRAME|NOSPLIT, $0-16
	JMP	runtime∕internal∕atomic·Store64(SB)

TEXT runtime∕internal∕atomic·Xadduintptr(SB), NOFRAME|NOSPLIT, $0-24
	JMP	runtime∕internal∕atomic·Xadd64(SB)

TEXT runtime∕internal∕atomic·Loadint64(SB), NOFRAME|NOSPLIT, $0-16
	JMP	runtime∕internal∕atomic·Load64(SB)

TEXT runtime∕internal∕atomic·Xaddint64(SB), NOFRAME|NOSPLIT, $0-24
	JMP	runtime∕internal∕atomic·Xadd64(SB)

// bool casp(void **val, void *old, void *new)
// Atomically:
//    if(*val == old){
//    *val = new;
//    return 1;
//    } else
//    return 0;
TEXT runtime∕internal∕atomic·Casp1(SB), NOFRAME|NOSPLIT, $0-25
	JMP	runtime∕internal∕atomic·Cas64(SB)

// uint32 xadd(uint32 volatile *ptr, int32 delta)
// Atomically:
//    *val += delta;
//    return *val;
TEXT runtime∕internal∕atomic·Xadd(SB), NOFRAME|NOSPLIT, $0-20
	LDL	R4, ptr+0(FP)
	LDW	R5, delta+8(FP)

	MEMB
	LLDW	R6, 0(R4)
	LDI	R7, 1
	WR_F	R7
	ADDW	R6, R5, R7
	LSTW	R7, 0(R4)
	RD_F	R7
	BEQ	R7, -6(PC)

	ADDW	R6, R5, R7
	STW	R7, ret+16(FP)
	RET

TEXT runtime∕internal∕atomic·Xadd64(SB), NOFRAME|NOSPLIT, $0-24
	LDL	R4, ptr+0(FP)
	LDL	R5, delta+8(FP)

	MEMB
	LLDL	R6, 0(R4)
	LDI	R7, 1
	WR_F	R7
	ADDL	R6, R5, R7
	LSTL	R7, 0(R4)
	RD_F	R7
	BEQ	R7, -6(PC)

	ADDL	R6, R5, R7
	STL	R7, ret+16(FP)
	RET

TEXT runtime∕internal∕atomic·Xchg(SB), NOFRAME|NOSPLIT, $0-20
	LDL	R4, ptr+0(FP)
	LDW	R5, new+8(FP)

	MEMB
	LLDW	R6, 0(R4)
	LDI	R7, 1
	WR_F	R7
	BIS	R5, R31, R7
	LSTW	R7, 0(R4)
	RD_F	R7
	BEQ	R7, -6(PC)

	STW	R6, ret+16(FP)
	RET

TEXT runtime∕internal∕atomic·Xchg64(SB), NOFRAME|NOSPLIT, $0-24
	LDL	R4, ptr+0(FP)
	LDL	R5, new+8(FP)

	MEMB
	LLDL	R6, 0(R4)
	LDI	R7, 1
	WR_F	R7
	BIS	R5, R31, R7
	LSTL	R7, 0(R4)
	RD_F	R7
	BEQ	R7, -6(PC)

	STL	R6, ret+16(FP)
	RET

TEXT runtime∕internal∕atomic·Xchguintptr(SB), NOFRAME|NOSPLIT, $0-24
	JMP	runtime∕internal∕atomic·Xchg64(SB)


TEXT runtime∕internal∕atomic·StorepNoWB(SB), NOFRAME|NOSPLIT, $0-16
	LDL	R3, ptr+0(FP)
	LDL	R4, val+8(FP)
	MEMB
	STL	R4, 0(R3)
	MEMB
	RET

TEXT runtime∕internal∕atomic·Store(SB), NOFRAME|NOSPLIT, $0-12
	LDL	R3, ptr+0(FP)
	LDW	R4, val+8(FP)
	MEMB
	STW	R4, 0(R3)
	MEMB
	RET

TEXT runtime∕internal∕atomic·Store64(SB), NOFRAME|NOSPLIT, $0-16
	LDL	R3, ptr+0(FP)
	LDL	R4, val+8(FP)
	MEMB
	STL	R4, 0(R3)
	MEMB
	RET

// void runtime∕internal∕atomic·And8(byte volatile*, byte);
TEXT runtime∕internal∕atomic·And8(SB), NOFRAME|NOSPLIT, $0-9
	LDL	R1, ptr+0(FP)
	LDBU	R2, val+8(FP)
	LDI	R3, $-4
	AND	R3, R1, R3
	AND	R1, $3, R4
	SLL	R4, $3, R4
	LDI	R5, 0xFF
	SLL	R2, R4, R2
	SLL	R5, R4, R5
	ORNOT	R31, R5, R5
	BIS	R2, R5, R2

	MEMB
	LLDW	R4, (R3)
	LDI	R1, 1
	WR_F	R1
	AND	R4, R2, R1
	LSTW	R1, (R3)
	RD_F	R1
	BEQ	R1, -6(PC)
	RET

// void runtime∕internal∕atomic·Or8(byte volatile*, byte);
TEXT runtime∕internal∕atomic·Or8(SB), NOFRAME|NOSPLIT, $0-9
	LDL	R1, ptr+0(FP)
	LDBU	R2, val+8(FP)
	LDI	R3, $-4
	AND	R3, R1, R3
	AND	R1, $3, R4
	SLL	R4, $3, R4
	SLL	R2, R4, R2

	MEMB
	LLDW	R4, (R3)
	LDI	R1, 1
	WR_F	R1
	BIS	R4, R2, R1
	LSTW	R1, (R3)
	RD_F	R1
	BEQ	R1, -6(PC)
	RET

// void func And(addr *uint32, v uint32)
TEXT runtime∕internal∕atomic·And(SB), NOFRAME|NOSPLIT, $0-12
	LDL	R1, ptr+0(FP)
	LDW	R2, val+8(FP)

	MEMB
	LLDW	R4, (R1)
	LDI	R3, 1
	WR_F	R3
	AND	R4, R2, R4
	LSTW	R4, (R1)
	RD_F	R3
	BEQ	R3, -6(PC)
	RET

// func Or(addr *uint32, v uint32)
TEXT runtime∕internal∕atomic·Or(SB), NOFRAME|NOSPLIT, $0-12
	LDL	R1, ptr+0(FP)
	LDW	R2, val+8(FP)

	MEMB
	LLDW	R4, (R1)
	LDI	R3, 1
	WR_F	R3
	BIS	R4, R2, R4
	LSTW	R4, (R1)
	RD_F	R3
	BEQ	R3, -6(PC)
	RET
//zxw new add
TEXT runtime∕internal∕atomic·Store8(SB), NOFRAME|NOSPLIT, $0-9
	LDL	R3, ptr+0(FP)
	LDBU	R4, val+8(FP)
	MEMB
	STB	R4, 0(R3)
	MEMB
	RET

TEXT runtime∕internal∕atomic·StoreRel(SB), NOSPLIT|NOFRAME, $0-12
	JMP	runtime∕internal∕atomic·Store(SB)

TEXT runtime∕internal∕atomic·CasRel(SB), NOSPLIT|NOFRAME, $0-17
	JMP	runtime∕internal∕atomic·Cas(SB)

