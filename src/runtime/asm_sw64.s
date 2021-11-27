// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

// _rt0_sw64_lib is common startup code for most sw64 systems when
// using -buildmode=c-archive or -buildmode=c-shared. The linker will
// arrange to invoke this function as a global constructor (for
// c-archive) or when the shared library is loaded (for c-shared).
// We expect argc and argv to be passed in the usual C ABI registers
// a0 and a1.
TEXT _rt0_sw64_lib(SB), NOSPLIT, $0x50
	// 1. SAVE R16, R17
	STL	R16, _rt0_sw64_lib_argc<>(SB)
	STL	R17, _rt0_sw64_lib_argv<>(SB)
	
	// 2. save C ABI registers
	LDI	SP, $-64(SP)
	STL	R9, 0*8(SP)
	STL	R10, 1*8(SP)
	STL	R11, 2*8(SP)
	STL	R12, 3*8(SP)
	STL	R13, 4*8(SP)
	STL	R14, 5*8(SP)
	STL	R15, 6*8(SP)
	
	CALL	runtime·libpreinit(SB)
	// 3. Create a new thread to finish Go runtime initialization.
	LDL	R27, _cgo_sys_thread_create(SB)
	BEQ	R27, nocgo
	
	SYMADDR	R16, $_rt0_sw64_lib_go(SB)
	LDI	R17, $0
	CALL	R26, (R27)
	JMP	restore
nocgo:
	LDI	R16, $0x800000
	STL	R16, 0(SP)
	SYMADDR	R17, $_rt0_sw64_lib_go(SB)
	STL	R17, 8(SP)
	CALL	runtime·newosproc0(SB)
restore:
	LDL	R9, 0*8(SP)
	LDL	R10, 1*8(SP)
	LDL	R11, 2*8(SP)
	LDL	R12, 3*8(SP)
	LDL	R13, 4*8(SP)
	LDL	R14, 5*8(SP)
	LDL	R15, 6*8(SP)
	LDI	SP, $64(SP)
	RET

// _rt0_sw64_lib_go initializes the Go runtime.
// This is started in a separate thread by _rt0_sw64_lib.
TEXT _rt0_sw64_lib_go(SB),NOSPLIT,$0
	LDL	R16, _rt0_sw64_lib_argc<>(SB)
	LDL	R17, _rt0_sw64_lib_argv<>(SB)
	CALL	runtime·rt0_go(SB)
	RET

DATA	_rt0_sw64_lib_argc<>(SB)/8, $0
GLOBL	_rt0_sw64_lib_argc<>(SB),NOPTR, $8
DATA	_rt0_sw64_lib_argv<>(SB)/8, $0
GLOBL	_rt0_sw64_lib_argv<>(SB),NOPTR, $8


// _rt0_sw64 is common startup code for most sw64 systems when using
// internal linking. This is the entry point for the program from the
// kernel for an ordinary -buildmode=exe program. The stack holds the
// number of arguments and the C-style argv.
TEXT _rt0_sw64(SB), NOFRAME|NOSPLIT, $0
	LDL	R16, 0(SP) // argc
	LDI	R17, 8(SP) // argv
	JMP	runtime·rt0_go(SB)

TEXT main(SB), NOFRAME|NOSPLIT, $0
	SETFPEC1
	JMP	runtime·rt0_go(SB)

TEXT runtime·rt0_go(SB), NOSPLIT, $16
	// copy args
	STL	R16, $argc-16(SP)
	STL	R17, $argv-8(SP)
	
	SYMADDR	g, $runtime·g0(SB)
	LDI	R1, $-64*1024(SP)
	STL	R1, g_stackguard0(g)
	STL	R1, g_stackguard1(g)
	STL	R1, (g_stack+stack_lo)(g)
	STL	SP, (g_stack+stack_hi)(g)
	
	// if there is a _cgo_init, call it using the gcc ABI.
	LDL	R27, _cgo_init(SB)
	BEQ	R27, nocgo
	
	LDI	R16, g
	SYMADDR	R17, $setg_gcc<>(SB)
	LDI	R18, ZERO
	LDI	R19, ZERO
	CALL	R26, (R27)
	SETFPEC1
nocgo:
	SYMADDR	g, $runtime·g0(SB)
	SYMADDR	R0, $runtime·m0(SB)
	STL	g, m_g0(R0)
	STL	R0,  g_m(g)
	
	CALL	runtime·check(SB)

	// args are already prepared
	CALL	runtime·args(SB)
	CALL	runtime·osinit(SB)
	CALL	runtime·schedinit(SB)
	// create a new goroutine to start program
	SYMADDR	R1, $runtime·mainPC(SB)
	STL	R1, $fn-8(SP)
	STL	ZERO, $sz-16(SP)
	CALL	runtime·newproc(SB)
	CALL	runtime·mstart(SB)
	RET

DATA	runtime·mainPC+0(SB)/8,$runtime·main(SB)
GLOBL	runtime·mainPC(SB),RODATA,$8

// This is not obey by Go ABI, see SW64Ops.go
//           R1 R0         R0 R1
// func udiv(n, d uint64) (q, r uint64) {
TEXT runtime·udiv(SB), NOSPLIT, $0
	SUBL	SP, $80, SP

	LDI	R5, ZERO
	STL	R1, $capn+16(SP)

	LDI 	R3, R1
	LDI	R2, $0x8000000000000000
	CMPULE	R2, R3, R3
	BEQ	R3, l80
	STL 	R2, $capn+16(SP)
l80:
   	LDI	R3, ZERO
	STL	ZERO, $i+8(SP)
l96:
	LDL	R4, $capn+16(SP)

	CMPULT	R0, R4, R2
	BEQ	R2, l184
	LDI	R2, R0

	SLL	R2, $1, R2
	LDI	R0, R2
	LDI	R2, $i+8(SP)
	LDL	R2, (R2)

	ADDL	R2, $1, R2

	STL	R2, $i+8(SP)

	BR	ZERO, l96
l184:
	LDL	R2, $i+8(SP)
	CMPLE	R3, R2, R2
	BEQ	R2, l384
l208:
	SLL	R5, $1, R5
	CMPULE	R0, R1, R2
	BEQ	R2, l316
	SUBL	R1, R0, R1
	BIS	R5, $1, R5
l316:
	LDI	R2, R0
	SRL	R2, $1, R2
	LDI	R0, R2
	BR	ZERO, l340
l340:
	LDL	R2, $i+8(SP)
	SUBL	R2, $1, R2
	STL	R2, $i+8(SP)
	BR	ZERO, l184
l384:
	LDI	R0, R5
	ADDL	SP, $80, SP
	RET

TEXT runtime·breakpoint(SB), NOFRAME|NOSPLIT, $0
	SYS_CALL_B	$0x80
	RET

// void mcall(fn func(*g))
// Switch to m->g0's stack, call fn(g).
// Fn must never return. It should gogo(&g->sched)
// to keep running g.
TEXT runtime·mcall(SB), NOFRAME|NOSPLIT, $0-8
	// Save caller state in g->sched
	STL	SP, (g_sched+gobuf_sp)(g)
	STL	R26, (g_sched+gobuf_pc)(g)
	STL	ZERO, (g_sched+gobuf_lr)(g)
	STL	g, (g_sched+gobuf_g)(g)

	// Switch to m->g0 & its stack, call fn.
	LDI	R1, g
	LDL	R3, g_m(g)
	LDL	g, m_g0(R3)
	CALL	runtime·save_g(SB)
	CMPEQ	g, R1, R7
	BEQ	R7, ok
	CALL	runtime·badmcall(SB)
ok:
	LDL	REGCTXT, fn+0(FP)
	LDL	R27, 0(REGCTXT)
	LDL	SP, (g_sched+gobuf_sp)(g)
	
	SUBL	SP, $16, SP
	STL	R1, 8(SP)
	STL	ZERO, 0(SP)
	JMP	(R27)
	ADDL	SP, $16, SP
	CALL	runtime·badmcall2(SB)

// func systemstack(fn func())
TEXT runtime·systemstack(SB), NOSPLIT, $0-8
	LDL	R1, fn+0(FP)
	LDI	REGCTXT, R1
	
	LDL	R2, g_m(g)
	
	LDL	R3, m_gsignal(R2)
	CMPEQ	g, R3, R3
	BNE	R3, noswitch
	
	LDL	R3, m_g0(R2) // save g0 in R3
	CMPEQ	g, R3, R4
	BNE	R4, noswitch
	
	LDL	R4, m_curg(R2)
	CMPEQ	g, R4, R4
	BNE	R4, switch
	
	CALL	runtime·badsystemstack(SB)
	CALL	runtime·abort(SB)
	RET
switch:
	// save our state in g->sched. Pretend to
	// be systemstack_switch if the G stack is scanned.
	SYMADDR	R4, $runtime·systemstack_switch(SB)
	ADDL	R4, $8, R4 // get past prologue
	STL	R4, (g_sched+gobuf_pc)(g)
	STL	SP, (g_sched+gobuf_sp)(g)
	STL	ZERO, (g_sched+gobuf_lr)(g)
	STL	g, (g_sched+gobuf_g)(g)
	
	// switch to g0
	LDI	g, R3
	CALL	runtime·save_g(SB)

	LDL	R1, (g_sched+gobuf_sp)(g)
	// make it look like mstart called systemstack on g0, to stop traceback
	SUBL	R1, $8, R1
	SYMADDR	R2, $runtime·mstart(SB)
	STL	R2, 0(R1)
	LDI	SP, R1
	
	LDL	R27, 0(REGCTXT)
	CALL	(R27)
	
	// switch back to g
	LDL	R1, g_m(g)
	LDL	g, m_curg(R1)
	CALL	runtime·save_g(SB)
	LDL	SP, (g_sched+gobuf_sp)(g)
	STL	ZERO, (g_sched+gobuf_sp)(g)
	RET
noswitch:
	// already on m stack, just call directly
	// Using a tail call here cleans up tracebacks since we won't stop
	// at an intermediate systemstack.
	LDL	R27, 0(REGCTXT)
	LDL	R26, 0(SP)
	ADDL	SP, $8, SP
	JMP	(R27)

TEXT runtime·systemstack_switch(SB), NOSPLIT, $0-0
	UNDEF
	CALL	(R27)	// make sure this function is not leaf
	RET

TEXT runtime·asminit(SB), NOFRAME|NOSPLIT, $0-0
	RET

/*
 * support for morestack
 */

// Called during function prolog when more stack is needed.
//
// The traceback routines see morestack on a g0 as being
// the top of a stack (for example, morestack calling newstack
// calling the scheduler calling newm calling gc), so we must
// record an argument size. For that purpose, it has no arguments.
//
// R2 is return address; R3 is PC
TEXT runtime·morestack(SB), NOFRAME|NOSPLIT, $0-0
	// Cannot grow scheduler stack (m->g0).
	LDL	R7, g_m(g)
	LDL	R8, m_g0(R7)
	CMPEQ	g, R8, R1
	BEQ	R1, 3(PC) // if g != g.m.g0
	CALL	runtime·badmorestackg0(SB)
	CALL	runtime·abort(SB)

	// Cannot grow signal stack (m->gsignal).
	LDL	R8, m_gsignal(R7)
	CMPEQ	g, R8, R1
	BEQ	R1, 3(PC) // if g != g.m.gsignal
	CALL	runtime·badmorestackgsignal(SB)
	CALL	runtime·abort(SB)

	// Called from f.
	// Set g->sched to context in f.
	STL	SP, (g_sched+gobuf_sp)(g)
	STL	R2, (g_sched+gobuf_pc)(g)
	STL	R3, (g_sched+gobuf_lr)(g)
	STL	REGCTXT, (g_sched+gobuf_ctxt)(g)

	// Called from f.
	// Set m->morebuf to f's caller.
	STL	R3, (m_morebuf+gobuf_pc)(R7)	// f's caller's PC
	STL	SP, (m_morebuf+gobuf_sp)(R7)	// f's caller's SP
	STL	g, (m_morebuf+gobuf_g)(R7)

	// Call newstack on m->g0's stack.
	LDL	g, m_g0(R7)
	CALL	runtime·save_g(SB)
	LDL	SP, (g_sched+gobuf_sp)(g)
	// Create a stack frame on g0 to call newstack.
	STL	ZERO, -8(SP)
	SUBL	SP, $8, SP
	CALL	runtime·newstack(SB)

	// Not reached, but make sure the return PC from the call to newstack
	// is still in this function, and not the beginning of the next.
	UNDEF


TEXT runtime·morestack_noctxt(SB), NOFRAME|NOSPLIT, $0-0
	LDI	REGCTXT, ZERO
	JMP	runtime·morestack(SB)


TEXT runtime·publicationBarrier(SB), NOFRAME|NOSPLIT, $0
	MEMB
	RET

// void setg(G*); set g. for use by needm.
TEXT runtime·setg(SB), NOSPLIT, $0-8
	LDL	g, gg+0(FP)
	// This only happens if iscgo, so jump straight to save_g
	CALL	runtime·save_g(SB)
	RET


TEXT runtime·return0(SB), NOSPLIT, $0
	LDI	R0, $0
	RET


// The top-most function running on a goroutine
TEXT runtime·goexit(SB), NOFRAME|NOSPLIT|TOPFRAME, $0-0
	LDI	ZERO, $0
	JMP	runtime·goexit1(SB)
	LDI	ZERO, $0

// This is called from .init_array and follows the platform, not Go, ABI
TEXT runtime·addmoduledata(SB),NOSPLIT,$0-0
  LDL  R1, runtime·lastmoduledatap(SB)
  STL  R0, moduledata_next(R1) // local.moduledata passed to R0
  STL  R0, runtime·lastmoduledatap(SB)
  RET

// void gogo(Gobuf*)
// restore state from Gobuf; longjmp
TEXT runtime·gogo(SB), NOSPLIT, $0-8
	LDL	R3, buf+0(FP)
	LDL	g, gobuf_g(R3)
	CALL	runtime·save_g(SB)

	LDL	SP, gobuf_sp(R3)
	LDL	R26, gobuf_lr(R3)
	LDL	R0, gobuf_ret(R3)
	LDL	REGCTXT, gobuf_ctxt(R3)

	STL	ZERO, gobuf_sp(R3)
	STL	ZERO, gobuf_ret(R3)
	STL	ZERO, gobuf_lr(R3)
	STL	ZERO, gobuf_ctxt(R3)

	LDL	R27, gobuf_pc(R3)
	JMP	(R27)
	RET

TEXT runtime·procyield(SB), NOFRAME|NOSPLIT, $0-0
	RET

#define DISPATCH(NAME, MAXSIZE) \
	LDI	R0, $MAXSIZE \
	CMPULT	R0, R19, R0 \
	BNE	R0, 4(PC) \
	SYMADDR	R27, $NAME(SB) \
	JMP	(R27) \
	RET

TEXT reflect·call(SB), NOFRAME|NOSPLIT, $0-0
	JMP	·reflectcall(SB)

// func reflectcall(argtype *_type, fn, arg unsafe.Pointer, argsize uint32, retoffset uint32)
TEXT ·reflectcall(SB), NOFRAME|NOSPLIT, $0-32
	LDW	R19, argsize+24(FP) // R19 used in DISPACH macro
	DISPATCH(runtime·call32, 32)
	DISPATCH(runtime·call64, 64)
	DISPATCH(runtime·call128, 128)
	DISPATCH(runtime·call256, 256)
	DISPATCH(runtime·call512, 512)
	DISPATCH(runtime·call1024, 1024)
	DISPATCH(runtime·call2048, 2048)
	DISPATCH(runtime·call4096, 4096)
	DISPATCH(runtime·call8192, 8192)
	DISPATCH(runtime·call16384, 16384)
	DISPATCH(runtime·call32768, 32768)
	DISPATCH(runtime·call65536, 65536)
	DISPATCH(runtime·call131072, 131072)
	DISPATCH(runtime·call262144, 262144)
	DISPATCH(runtime·call524288, 524288)
	DISPATCH(runtime·call1048576, 1048576)
	DISPATCH(runtime·call2097152, 2097152)
	DISPATCH(runtime·call4194304, 4194304)
	DISPATCH(runtime·call8388608, 8388608)
	DISPATCH(runtime·call16777216, 16777216)
	DISPATCH(runtime·call33554432, 33554432)
	DISPATCH(runtime·call67108864, 67108864)
	DISPATCH(runtime·call134217728, 134217728)
	DISPATCH(runtime·call268435456, 268435456)
	DISPATCH(runtime·call536870912, 536870912)
	DISPATCH(runtime·call1073741824, 1073741824)
	
	CALL	runtime·badreflectcall(SB)
	RET

#define CALLFN(NAME,MAXSIZE) \
TEXT NAME(SB), WRAPPER, $MAXSIZE-24 \
	NO_LOCAL_POINTERS;            \
	/* copy arguments to stack */        \
	LDL	R1, argtype+16(FP) \
	LDW	R2, argsize+24(FP) \
	LDI	R3, SP \
	ADDL	R3, $8, R3 \
	ADDL	R3, R2, R2 \
loop: \
	CMPEQ	R3, R2, R0 \
	BNE	R0, ok \
	LDBU	R4, (R1) \
	STB	R4, (R3) \
	ADDL	R1, $1, R1 \
	ADDL	R3, $1, R3 \
	JMP	loop \
ok: \
	/* call function */  \
	LDL	REGCTXT, fn+8(FP) \
	LDL	R27, (REGCTXT) \
	PCDATA	$PCDATA_StackMapIndex, $0; \
	CALL	(R27) \
	/* copy return values back */        \
	LDL	R5, argtype+0(FP) \
	LDL	R1, arg+16(FP) \
	LDW	R2, argsize+24(FP) \
	LDW	R4, retoffset+28(FP) \
	ADDL	SP, $8, R3 \
	ADDL	R4, R3, R3 \
	ADDL	R4, R1, R1 \
	SUBL	R2, R4, R2 \
	CALL	callRet<>(SB) \
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $32-0
	STL	R5, 8(SP) // argtype
	STL	R1, 16(SP) // dst
	STL	R3, 24(SP) // src
	STL R2, 32(SP) // size
	// reflectcallmove(typ *_type, dst, src unsafe.Pointer, size uintptr)
	CALL	runtime·reflectcallmove(SB)
	RET

CALLFN(·call16, 16)
CALLFN(·call32, 32)
CALLFN(·call64, 64)
CALLFN(·call128, 128)
CALLFN(·call256, 256)
CALLFN(·call512, 512)
CALLFN(·call1024, 1024)
CALLFN(·call2048, 2048)
CALLFN(·call4096, 4096)
CALLFN(·call8192, 8192)
CALLFN(·call16384, 16384)
CALLFN(·call32768, 32768)
CALLFN(·call65536, 65536)
CALLFN(·call131072, 131072)
CALLFN(·call262144, 262144)
CALLFN(·call524288, 524288)
CALLFN(·call1048576, 1048576)
CALLFN(·call2097152, 2097152)
CALLFN(·call4194304, 4194304)
CALLFN(·call8388608, 8388608)
CALLFN(·call16777216, 16777216)
CALLFN(·call33554432, 33554432)
CALLFN(·call67108864, 67108864)
CALLFN(·call134217728, 134217728)
CALLFN(·call268435456, 268435456)
CALLFN(·call536870912, 536870912)
CALLFN(·call1073741824, 1073741824)


// void jmpdefer(fv, sp);
// called from deferreturn.
// 1. grab stored LR for caller
// 2. sub 3 instructions, e.g. 12 bytes to get back to deferreturn
// 3. JMP to fn
TEXT runtime·jmpdefer(SB), NOFRAME|NOSPLIT, $0-16
	LDL	R26, 0(SP)
	SUBL	R26, $12, R26
	
	LDL	REGCTXT, fv+0(FP)
	
	LDL	SP, argp+8(FP)
	SUBL	SP, $8, SP
	LDI	ZERO, ZERO // prevent scheduling
	
	LDL	R27, 0(REGCTXT)
	JMP	(R27)
	RET

TEXT ·checkASM(SB),NOSPLIT,$0-1
	LDI	R1, $1
	STB	R1, ret+0(FP)
	RET


TEXT runtime·abort(SB), NOFRAME|NOSPLIT, $0-0
	LDW	ZERO,(ZERO)
	UNDEF

// Save state of caller into g->sched. Smashes R1.
TEXT gosave<>(SB), NOFRAME|NOSPLIT,$0
	STL	R26, (g_sched+gobuf_pc)(g)
	STL	SP, (g_sched+gobuf_sp)(g)
	STL	ZERO, (g_sched+gobuf_lr)(g)
	STL	ZERO, (g_sched+gobuf_ret)(g)
	// Assert ctxt is zero. See func save.
	LDL	R1, (g_sched+gobuf_ctxt)(g)
	BEQ	R1, 2(PC)
	CALL	runtime·badctxt(SB)
	RET

// func asmcgocall(fn, arg unsafe.Pointer) int32
// Call fn(arg) on the scheduler stack,
// aligned appropriately for the gcc ABI.
// See cgocall.go for more details.
TEXT ·asmcgocall(SB), NOSPLIT, $8-20
	NO_LOCAL_POINTERS
	LDL	R27, fn+0(FP)
	LDL	R16, arg+8(FP)
	
	STL	R27, 8(SP)
	
	LDI	R3, SP
	LDI	R2, g

	// Figure out if we need to switch to m->g0 stack.
	// We get called to create new OS threads too, and those
	// come in on the m->g0 stack already.
	LDL	R5, g_m(g)
	LDL	R6, m_g0(R5)
	CMPEQ	R6, g, R0
	
	BNE	R0, g0
	
	CALL	gosave<>(SB)
	LDI	g, R6
	CALL	runtime·save_g(SB)
	
	// restore fn value from old SP
	LDL	R27, 8(SP)
	LDL	SP, (g_sched+gobuf_sp)(g)
	// Now on a scheduling stack (a pthread-created stack).
g0:
	// Save room for two of our pointers.
	LDI	SP, $-16(SP)
	STL	R2, 0(SP)
	LDL	R2, (g_stack+stack_hi)(R2)
	SUBL	R2, R3, R2
	STL	R2, 8(SP)
	CALL	R26, (R27)
	// Restore g, stack pointer. R2 is return value.
	LDL	g, 0(SP)
	CALL	runtime·save_g(SB)
	LDL	R5, (g_stack+stack_hi)(g)
	LDL	R6, 8(SP)
	SUBL	R5, R6, R5
	LDI	SP, R5
	
	STW	R0, ret+16(FP)
	RET

// void setg_gcc(G*); set g in C TLS.
// Must obey the gcc calling convention.
TEXT setg_gcc<>(SB), NOSPLIT, $0
	LDI	g, R16
	CALL	runtime·save_g(SB)
	RET

// Called from cgo wrappers, this function returns g->m->curg.stack.hi.
// Must obey the gcc calling convention.
TEXT _cgo_topofstack(SB), NOSPLIT, $8
	// g (R15)  might be clobbered by load_g. They
	// are callee-save in the gcc calling convention, so save them.
	STL	g, savedG-8(SP)

	CALL	runtime·load_g(SB)

	LDL	R1, g_m(g)
	LDL	R1, m_curg(R1)
	LDL	R0, (g_stack+stack_hi)(R1) // return value in R0
	
	LDL	g, savedG-8(SP)
	RET


// cgocallback_gofunc(FuncVal*, void *frame, uintptr framesize, uintptr ctxt)
// See cgocall.go for more details.
TEXT ·cgocallback(SB), NOSPLIT, $24-24
	NO_LOCAL_POINTERS

	// Load m and g from thread-local storage.
	LDBU	R1, runtime·iscgo(SB)
	BEQ	R1, nocgo
	CALL	runtime·load_g(SB)
nocgo:

	// If g is nil, Go did not create the current thread.
	// Call needm to obtain one for temporary use.
	// In this case, we're running on the thread stack, so there's
	// lots of space, but the linker doesn't know. Hide the call from
	// the linker analysis by using an indirect call.
	BEQ	g, needm
	LDL	R3, g_m(g)
	STL	R3, savedm-8(SP)
	JMP	havem

needm:
	STL	g, savedm-8(SP) // g is zero, so is m.
	SYMADDR	R27, $runtime·needm(SB)
	CALL	R26, (R27)

	// Set m->sched.sp = SP, so that if a panic happens
	// during the function we are about to execute, it will
	// have a valid SP to run on the g0 stack.
	// The next few lines (after the havem label)
	// will save this SP onto the stack and then write
	// the same SP back to m->sched.sp. That seems redundant,
	// but if an unrecovered panic happens, unwindm will
	// restore the g->sched.sp from the stack location
	// and then systemstack will try to use it. If we don't set it here,
	// that restored SP will be uninitialized (typically 0) and
	// will not be usable.
	LDL	R3, g_m(g)
	LDL	R1, m_g0(R3)
	STL	SP, (g_sched+gobuf_sp)(R1)

havem:
	// Now there's a valid m, and we're running on its m->g0.
	// Save current m->g0->sched.sp on stack and then set it to SP.
	// Save current sp in m->g0->sched.sp in preparation for
	// switch back to m->curg stack.
	// NOTE: unwindm knows that the saved g->sched.sp is at 8(R29) aka savedsp-16(SP).

	LDL	R1, m_g0(R3)
	LDL	R2, (g_sched+gobuf_sp)(R1)
	STL	R2, savedsp-24(SP)
	STL	SP,  (g_sched+gobuf_sp)(R1)

	// Switch to m->curg stack and call runtime.cgocallbackg.
	// Because we are taking over the execution of m->curg
	// but *not* resuming what had been running, we need to
	// save that information (m->curg->sched) so we can restore it.
	// We can restore m->curg->sched.sp easily, because calling
	// runtime.cgocallbackg leaves SP unchanged upon return.
	// To save m->curg->sched.pc, we push it onto the stack.
	// This has the added benefit that it looks to the traceback
	// routine like cgocallbackg is going to return to that
	// PC (because the frame we allocate below has the same
	// size as cgocallback_gofunc's frame declared above)
	// so that the traceback will seamlessly trace back into
	// the earlier calls.
	//
	// In the new goroutine, -8(SP) is unused (where SP refers to
	// m->curg's SP while we're setting it up, before we've adjusted it).
	LDL	g, m_curg(R3)
	CALL	runtime·save_g(SB)
	LDL	R2, (g_sched+gobuf_sp)(g)
	LDL	R27, (g_sched+gobuf_pc)(g)
	STL	R27, -32(R2) // save LR 
  	//  Gather our arguments into registers.
  	LDL 	R16, fn+0(FP)
  	LDL	R17, frame+8(FP)
  	LDL 	R18, ctxt+16(FP)
  	LDI 	SP, $-32(R2) //switch stack
	STL 	R16, 8(SP)
	STL 	R17, 16(SP)
	STL 	R18, 24(SP)
	CALL	runtime·cgocallbackg(SB)

	// Restore g->sched (== m->curg->sched) from saved values.
	LDL	R27, 0(SP)
	STL	R27, (g_sched+gobuf_pc)(g)
	LDI	R2, $32(SP)
	STL	R2, (g_sched+gobuf_sp)(g)

	// Switch back to m->g0's stack and restore m->g0->sched.sp.
	// (Unlike m->curg, the g0 goroutine never uses sched.pc,
	// so we do not have to restore it.)
	LDL	R3, g_m(g)
	LDL	g, m_g0(R3)
	CALL	runtime·save_g(SB)
	LDL	SP, (g_sched+gobuf_sp)(g)
	LDL	R2, savedsp-24(SP)
	STL	R2, (g_sched+gobuf_sp)(g)
	
	// If the m on entry was nil, we called needm above to borrow an m
	// for the duration of the call. Since the call is over, return it with dropm.
	LDL	R3, savedm-8(SP)
	BNE	R3, droppedm
	SYMADDR	R27, $runtime·dropm(SB)
	CALL	R26, (R27)
droppedm:
	// Done!
	RET


// gcWriteBarrier performs a heap pointer write and informs the GC.
//
// gcWriteBarrier does NOT follow the Go ABI. It takes two arguments:
// - R13 is the destination of the write
// - R14 is the value being written at R13.
// It clobbers R28 (the linker temp register).
// The act of CALLing gcWriteBarrier will clobber R26 (LR).
// It does not clobber any other general-purpose registers,
// but may clobber others (e.g., floating point registers).
TEXT runtime·gcWriteBarrier(SB),NOSPLIT,$208
	// Save the registers clobbered by the fast path.
	STL	R1, 200(SP)
	STL	R2, 208(SP)
	LDL	R1, g_m(g)
	LDL	R1, m_p(R1)
	LDL	R2, (p_wbBuf+wbBuf_next)(R1)
	// Increment wbBuf.next position.
	ADDL	R2, $16, R2
	STL	R2, (p_wbBuf+wbBuf_next)(R1)
	LDL	R1, (p_wbBuf+wbBuf_end)(R1)
	LDI	R28, R1 	// R28 is linker temp register
	// Record the write.
	STL	R14, -16(R2)	// Record value
	LDL	R1, (R13)	// TODO: This turns bad writes into bad reads.
	STL	R1, -8(R2)	// Record *slot
	// Is the buffer full?
	CMPEQ	R2, R28, R28
	BNE	R28, flush
ret:
	LDL	R1, 200(SP)
	LDL	R2, 208(SP)
	// Do the write.
	STL	R14, (R13)
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	STL	R13, 8(SP)	// Also first argument to wbBufFlush
	STL	R14, 16(SP)	// Also second argument to wbBufFlush
	STL	R0, 24(SP)
	// R1 already saved
	// R2 already saved
	STL	R3, 32(SP)
	STL	R4, 40(SP)
	STL	R5, 48(SP)
	STL	R6, 56(SP)
	STL	R7, 64(SP)
	STL	R8, 72(SP)
	STL	R9, 80(SP)
	STL	R10, 88(SP)
	STL	R11, 96(SP)
	STL	R12, 104(SP)
	// R13 already saved
	// R14 already saved.
	// R15 is g.
	STL	R16, 112(SP)
	STL	R17, 120(SP)
	STL	R18, 128(SP)
	STL	R19, 136(SP)
	STL	R20, 144(SP)
	STL	R21, 152(SP)
	STL	R22, 160(SP)
	STL	R23, 168(SP)
	STL	R24, 176(SP)
	STL	R25, 184(SP)
	// R26 is link register.
	STL	R27, 192(SP)
	// R28 is tmp register.
	// R29 is SB.
	// R30 is SP.
	// R31 is ZERO.

	// This takes arguments R13 and R14.
	CALL	runtime·wbBufFlush(SB)

	LDL     R13, 8(SP)     
        LDL     R14, 16(SP)    
        LDL     R0, 24(SP)
        LDL     R3, 32(SP)
        LDL     R4, 40(SP)
        LDL     R5, 48(SP)
        LDL     R6, 56(SP)
        LDL     R7, 64(SP)
        LDL     R8, 72(SP)
        LDL     R9, 80(SP)
        LDL     R10, 88(SP)
        LDL     R11, 96(SP)
        LDL     R12, 104(SP)
        LDL     R16, 112(SP)
        LDL     R17, 120(SP)
        LDL     R18, 128(SP)
        LDL     R19, 136(SP)
        LDL     R20, 144(SP)
        LDL     R21, 152(SP)
        LDL     R22, 160(SP)
        LDL     R23, 168(SP)
        LDL     R24, 176(SP)
        LDL     R25, 184(SP)
        LDL     R27, 192(SP)
	LDL	R1, 200(SP)
	LDL	R2, 208(SP)
	// Do the write.
	STL	R14, (R13)
	RET

//zxw new add
// Note: these functions use a special calling convention to save generated code space.
// Arguments are passed in registers, but the space for those arguments are allocated
// in the caller's stack frame. These stubs write the args into that stack space and
// then tail call to the corresponding runtime handler.
// The tail call makes these stubs disappear in backtraces.
TEXT runtime·panicIndex(SB),NOSPLIT,$0-16
	STL	R1, x+0(FP)
	STL	R2, y+8(FP)
	JMP	runtime·goPanicIndex(SB)
TEXT runtime·panicIndexU(SB),NOSPLIT,$0-16
	STL	R1, x+0(FP)
	STL	R2, y+8(FP)
	JMP	runtime·goPanicIndexU(SB)
TEXT runtime·panicSliceAlen(SB),NOSPLIT,$0-16
	STL	R2, x+0(FP)
	STL	R3, y+8(FP)
	JMP	runtime·goPanicSliceAlen(SB)
TEXT runtime·panicSliceAlenU(SB),NOSPLIT,$0-16
	STL	R2, x+0(FP)
	STL	R3, y+8(FP)
	JMP	runtime·goPanicSliceAlenU(SB)
TEXT runtime·panicSliceAcap(SB),NOSPLIT,$0-16
	STL	R2, x+0(FP)
	STL	R3, y+8(FP)
	JMP	runtime·goPanicSliceAcap(SB)
TEXT runtime·panicSliceAcapU(SB),NOSPLIT,$0-16
	STL	R2, x+0(FP)
	STL	R3, y+8(FP)
	JMP	runtime·goPanicSliceAcapU(SB)
TEXT runtime·panicSliceB(SB),NOSPLIT,$0-16
	STL	R1, x+0(FP)
	STL	R2, y+8(FP)
	JMP	runtime·goPanicSliceB(SB)
TEXT runtime·panicSliceBU(SB),NOSPLIT,$0-16
	STL	R1, x+0(FP)
	STL	R2, y+8(FP)
	JMP	runtime·goPanicSliceBU(SB)
TEXT runtime·panicSlice3Alen(SB),NOSPLIT,$0-16
	STL	R3, x+0(FP)
	STL	R4, y+8(FP)
	JMP	runtime·goPanicSlice3Alen(SB)
TEXT runtime·panicSlice3AlenU(SB),NOSPLIT,$0-16
	STL	R3, x+0(FP)
	STL	R4, y+8(FP)
	JMP	runtime·goPanicSlice3AlenU(SB)
TEXT runtime·panicSlice3Acap(SB),NOSPLIT,$0-16
	STL	R3, x+0(FP)
	STL	R4, y+8(FP)
	JMP	runtime·goPanicSlice3Acap(SB)
TEXT runtime·panicSlice3AcapU(SB),NOSPLIT,$0-16
	STL	R3, x+0(FP)
	STL	R4, y+8(FP)
	JMP	runtime·goPanicSlice3AcapU(SB)
TEXT runtime·panicSlice3B(SB),NOSPLIT,$0-16
	STL	R2, x+0(FP)
	STL	R3, y+8(FP)
	JMP	runtime·goPanicSlice3B(SB)
TEXT runtime·panicSlice3BU(SB),NOSPLIT,$0-16
	STL	R2, x+0(FP)
	STL	R3, y+8(FP)
	JMP	runtime·goPanicSlice3BU(SB)
TEXT runtime·panicSlice3C(SB),NOSPLIT,$0-16
	STL	R1, x+0(FP)
	STL	R2, y+8(FP)
	JMP	runtime·goPanicSlice3C(SB)
TEXT runtime·panicSlice3CU(SB),NOSPLIT,$0-16
	STL	R1, x+0(FP)
	STL	R2, y+8(FP)
	JMP	runtime·goPanicSlice3CU(SB)

TEXT runtime·memhash(SB),NOSPLIT|NOFRAME,$0-32
	JMP	runtime·memhashFallback(SB)
TEXT runtime·strhash(SB),NOSPLIT|NOFRAME,$0-24
	JMP	runtime·strhashFallback(SB)
TEXT runtime·memhash32(SB),NOSPLIT|NOFRAME,$0-24
	JMP	runtime·memhash32Fallback(SB)
TEXT runtime·memhash64(SB),NOSPLIT|NOFRAME,$0-24
	JMP	runtime·memhash64Fallback(SB)
