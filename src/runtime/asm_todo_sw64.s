// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

#define SYSCALL(n) \
	LDI	R0, $n \
	SYS_CALL_B	$131
#define SYS_write 4

#define WARN_TODO

#ifdef WARN_TODO

DATA	todo_str+0(SB)/8,  $"\033[30;31m"
DATA	todo_str+8(SB)/8,  $"\nTODO: "
DATA	todo_str+16(SB)/8, $"implemen"
DATA	todo_str+24(SB)/8, $"t the fu"
DATA	todo_str+32(SB)/8, $"nc\u2192  "
DATA	todo_str+40(SB)/8, $"runtime·"
GLOBL	todo_str(SB), NOPTR, $48

DATA	todo_str2+0(SB)/8,  $"\033[0m\n"
GLOBL	todo_str2(SB), NOPTR, $8

#define TODO(FN, dstr, name) \
	DATA	dstr+0(SB)/8, $name \
	GLOBL	dstr(SB), NOPTR, $8 \
	\
	TEXT FN(SB), NOSPLIT, $8-0 \
	LDI	R16, $2 \
	SYMADDR	R17, todo_str(SB) \
	LDI	R18, $48 \
	SYSCALL(SYS_write) \
	\
	LDI	R16, $2 \
	SYMADDR	R17, dstr(SB) \
	LDI	R18, $8 \
	SYSCALL(SYS_write) \
	\
	LDI	R16, $2 \
	SYMADDR	R17, todo_str2(SB) \
	LDI	R18, $8 \
	SYSCALL(SYS_write) \
	\
	RET

#else

#define TODO(FN, dstr, name) \
	TEXT FN(SB), NOSPLIT, $0-0 \
	RET

#endif

// check that SP is in range [g->stack.lo, g->stack.hi)
TODO(runtime·stackcheck, todo_stackcheck, "stck_chc")

TODO(runtime·memeqbody, todo_memeqbody, "memEqbod")

TODO(bytes·countByte, todo_bytes_countbyte, "b·cntByt")
TODO(strings·countByte, todo_strings_countByte, "s·cntByt")

// input:
//   SI: data
//   BX: data len
//   AL: byte sought
//   R8: address to put result
// This requires the POPCNT instruction
TODO(runtime·countByte, todo_countbyte, "cntByt")

// This is called from .init_array and follows the platform, not Go, ABI.
TODO(·addmoduledata, todo_addmoduledata, "addMD")
