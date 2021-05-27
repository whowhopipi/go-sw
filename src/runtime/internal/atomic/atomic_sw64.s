// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "textflag.h"

// uint32 runtime∕internal∕atomic·Load(uint32 volatile* ptr)
TEXT ·Load(SB), NOFRAME|NOSPLIT, $0-12
	LDL	R3, ptr+0(FP)
	MEMB
	LDW	R3, 0(R3)
   	MEMB
	STW	R3, ret+8(FP)
	RET

// uint64 runtime∕internal∕atomic·Load64(uint64 volatile* ptr)
TEXT ·Load64(SB), NOFRAME|NOSPLIT, $0-16
	LDL	R3, ptr+0(FP)
	MEMB
	LDL	R3, 0(R3)
	MEMB
	STL	R3, ret+8(FP)
	RET

// void *runtime∕internal∕atomic·Loadp(void *volatile *ptr)
TEXT ·Loadp(SB), NOFRAME|NOSPLIT, $0-16
	LDL	R3, ptr+0(FP)
	MEMB
	LDL	R3, 0(R3)
	MEMB
	STL	R3, ret+8(FP)
	RET

//zxw new add
// uint8 runtime∕internal∕atomic·Load8(uint8 volatile* ptr)
TEXT ·Load8(SB), NOFRAME|NOSPLIT, $0-9
	LDL	R3, ptr+0(FP)
	MEMB
	LDBU	R3, 0(R3)
   	MEMB
	STB	R3, ret+8(FP)
	RET

// uint32 runtime∕internal∕atomic·LoadAcq(uint32 volatile* ptr)
TEXT ·LoadAcq(SB),NOSPLIT|NOFRAME,$0-12
	JMP	·Load(SB)
