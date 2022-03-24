// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux
// +build sw64

#include "textflag.h"

//
// System calls for sw64, Linux
//

#define SYSCALL SYS_CALL_B $131

// func Syscall(trap int64, a1, a2, a3 int64) (r1, r2, err int64);
TEXT ·Syscall(SB), NOSPLIT, $0-56
	CALL	runtime·entersyscall(SB)
	LDL	R16, a1+8(FP)
	LDL	R17, a2+16(FP)
	LDL	R18, a3+24(FP)
	LDI	R19, ZERO
	LDI	R20, ZERO
	LDI	R21, ZERO
	LDL	R0, trap+0(FP)
	SYSCALL
	BEQ	R19, ok

	LDI	R1, $-1
	STL	R1, r1+32(FP)
	STL	ZERO, r2+40(FP)
	STL	R0, err+48(FP)
	CALL	runtime·exitsyscall(SB)
	RET
ok:
	STL	R0, r1+32(FP)
	STL	R20, r2+40(FP)
	STL	ZERO, err+48(FP)
	CALL	runtime·exitsyscall(SB)
	RET

// func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
TEXT ·Syscall6(SB), NOSPLIT, $0-80
	CALL	runtime·entersyscall(SB)
	LDL	R16, a1+8(FP)
	LDL	R17, a2+16(FP)
	LDL	R18, a3+24(FP)
	LDL	R19, a4+32(FP)
	LDL	R20, a5+40(FP)
	LDL	R21, a6+48(FP)
	LDL	R0, trap+0(FP)
	SYSCALL
	BEQ	R19, ok

	LDI	R1, $-1
	STL	R1, r1+56(FP)
	STL	ZERO, r2+64(FP)
	STL	R0, err+72(FP)
	CALL	runtime·exitsyscall(SB)
	RET
ok:
	STL	R0, r1+56(FP)
	STL	R20, r2+64(FP)
	STL	ZERO, err+72(FP)
	CALL	runtime·exitsyscall(SB)
	RET

TEXT ·RawSyscall(SB), NOSPLIT, $0-56
	LDL	R16, a1+8(FP)
	LDL	R17, a2+16(FP)
	LDL	R18, a3+24(FP)
	LDI	R19, ZERO
	LDI	R20, ZERO
	LDI	R21, ZERO
	LDL	R0, trap+0(FP)
	SYSCALL
	BEQ	R19, ok

	LDI	R1, $-1
	STL	R1, r1+32(FP)
	STL	ZERO, r2+40(FP)
	STL	R0, err+48(FP)
	RET
ok:
	STL	R0, r1+32(FP)
	STL	R20, r2+40(FP)
	STL	ZERO, err+48(FP)
	RET

TEXT ·RawSyscall6(SB),NOSPLIT, $0-80
	LDL	R16, a1+8(FP)
	LDL	R17, a2+16(FP)
	LDL	R18, a3+24(FP)
	LDL	R19, a4+32(FP)
	LDL	R20, a5+40(FP)
	LDL	R21, a6+48(FP)
	LDL	R0, trap+0(FP)
	SYSCALL
	BEQ	R19, ok

	LDI	R1, $-1
	STL	R1, r1+56(FP)
	STL	ZERO, r2+64(FP)
	STL	R0, err+72(FP)
	RET
ok:
	STL	R0, r1+56(FP)
	STL	R20, r2+64(FP)
	STL	ZERO, err+72(FP)
	RET

TEXT ·rawSyscallNoError(SB),NOSPLIT,$0-48
	LDL	R16, a1+8(FP)
	LDL	R17, a2+16(FP)
	LDL	R18, a3+24(FP)
	LDI	R19, ZERO
	LDI	R20, ZERO
	LDI	R21, ZERO
	LDL	R0, trap+0(FP)  // syscall entry
	SYSCALL
	STL	R0, r1+32(FP)
	STL	R20, r2+40(FP)
	RET
