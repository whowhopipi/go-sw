// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

// If !iscgo, this is a no-op.
//
// NOTE: mcall() assumes this clobbers only R28 and R0
TEXT runtime·save_g(SB), NOFRAME|NOSPLIT, $0-0
	LDBU	R28, runtime·iscgo(SB)
	BEQ	R28, nocgo
	
	LDI	R28, R0
	STL	g, runtime·tls_g(SB) // TLS relocation clobbers R0
	LDI	R0, R28
	SETFPEC1
nocgo:
	RET

TEXT runtime·load_g(SB), NOFRAME|NOSPLIT, $0-0
	LDI	R28, R0
	LDL	g, runtime·tls_g(SB) // TLS relocation clobbers R0
	LDI	R0, R28
	SETFPEC1
	RET

GLOBL	runtime·tls_g(SB), TLSBSS, $8
