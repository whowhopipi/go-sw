// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include <asm/regdef.h>
/*
 * void crosscall1(void (*fn)(void), void (*setg_gcc)(void *g), void *g)
 *
 * Calling into the gc tool chain, where all registers are caller save.
 * Called from standard SW_64 ABI, where $9-$14, are callee-save,
 * so they must be saved explicitly, along with $26 (RA).
 */

#if defined(__sw_64_sw6__)
    #define LDD ldt
    #define STD stt
    #define CALL jsr
#else
    #define LDD fldd
    #define STD fstd
    #define CALL call
#endif

	.set	noat
	.align	2
	.globl	crosscall1
	.ent	crosscall1
crosscall1:

	.frame	fp, 16, ra, 0
	.cfi_startproc
	ldi	sp, -160(sp)
	.cfi_def_cfa_offset	160

	stl	$26, 0(sp)
	stl	$15, 56(sp)
	stl	$29, 64(sp)
	.cfi_offset	26, -16
	.cfi_offset	15, -56
	mov	sp, $15
	.cfi_def_cfa_register	15
	.prologue	0

	stl	$9, 8(sp)
	stl	$10, 16(sp)
	stl	$11, 24(sp)
	stl	$12, 32(sp)
	stl	$13, 40(sp)
	stl	$14, 48(sp)

	STD	$f2, 88 (sp)
	STD	$f3, 96 (sp)
	STD	$f4, 104(sp)
	STD	$f5, 112(sp)
	STD	$f6, 120(sp)
	STD	$f7, 128(sp)
	STD	$f8, 136(sp)
	STD	$f9, 144(sp)

	#a0=fn, a1=setg_gcc, a2=g
	stl	a0, 72(sp)

	ldi	$27, 0(a1)
	ldi	a0, 0(a2)
	ldi	a1, 0(zero)
	ldi	a2, 0(zero)
	CALL	$26, ($27), 0  #call setg_gcc (clobbers A0 ?)

	ldl	$27, 72(sp)
	ldi	a0, 0(zero)
	ldi	a1, 0(zero)
	ldi	a2, 0(zero)
	CALL	$26, ($27), 0  #call fn

	ldl	$9, 8(sp)
	ldl	$10, 16(sp)
	ldl	$11, 24(sp)
	ldl	$12, 32(sp)
	ldl	$13, 40(sp)
	ldl	$14, 48(sp)

	ldl	$15, 56(sp)
	.cfi_restore	15

	LDD	$f2, 88 (sp)
	LDD	$f3, 96 (sp)
	LDD	$f4, 104(sp)
	LDD	$f5, 112(sp)
	LDD	$f6, 120(sp)
	LDD	$f7, 128(sp)
	LDD	$f8, 136(sp)
	LDD	$f9, 144(sp)

	ldl	$29, 64(sp)

	ldl	$26, 0(sp)
	.cfi_restore	26

	ldi	sp, 160(sp)
	.cfi_def_cfa	30, 0

	ret	$31,($26),1
	.cfi_endproc

	.end	crosscall1
