// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "textflag.h"
#include "funcdata.h"

// makeFuncStub is the code half of the function returned by MakeFunc.
// See the comment on the declaration of makeFuncStub in makefunc.go
// for more details.
// No arg size here, runtime pulls arg map out of the func value.
//TEXT ·makeFuncStub(SB), (NOSPLIT|WRAPPER), $16
TEXT ·makeFuncStub(SB), (NOSPLIT|WRAPPER), $32
	NO_LOCAL_POINTERS
	STL	REGCTXT, 8(SP)
	LDI	R1, $argframe+0(FP)
	STL	R1, 16(SP)
	STB	ZERO, 32(SP)
	ADDL	SP, $32, R1
	STL	R1, 24(SP)
	CALL	·callReflect(SB)
	RET

// methodValueCall is the code half of the function returned by makeMethodValue.
// See the comment on the declaration of methodValueCall in makefunc.go
// for more details.
// No arg size here; runtime pulls arg map out of the func value.
//TEXT ·methodValueCall(SB), (NOSPLIT|WRAPPER), $16
TEXT ·methodValueCall(SB), (NOSPLIT|WRAPPER), $32
	NO_LOCAL_POINTERS
	LDI	R1, $argframe+0(FP)
	STL	R1, 16(SP)
	STL	REGCTXT, 8(SP)
        STB     ZERO, 32(SP)
	ADDL	SP, $32, R1
        STL     R1, 24(SP)
	CALL	·callMethod(SB)
	RET
