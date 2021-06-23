#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

#define SYS_exit 1
#define SYS_read 3
#define SYS_write 4
#define SYS_close 6
#define SYS_brk 17
#define SYS_getxpid 20
#define SYS_kill 37
#define SYS_pipe 42
#define SYS_open 45
#define SYS_mmap 71
#define SYS_munmap 73
#define SYS_madvise 75
#define SYS_fcntl 92
#define SYS_sigaltstack 235
#define SYS_clone 312
#define SYS_sched_yield 334
#define SYS_rt_sigreturn 351
#define SYS_rt_sigaction 352
#define SYS_rt_sigprocmask 353
#define SYS_select 358
#define SYS_setitimer 362
#define SYS_mincore 375
#define SYS_gettid 378
#define SYS_tkill 381
#define SYS_futex 394
#define SYS_sched_getaffinity 396
#define SYS_exit_group 405
#define SYS_epoll_create 407
#define SYS_epoll_ctl 408
#define SYS_epoll_wait 409
#define SYS_epoll_pwait 474
#define SYS_epoll_create1 486
#define SYS_clock_gettime 420
#define SYS_tgkill 424
#define SYS_pipe2 488

#define SYSCALL(n) \
	LDI	R0, $n \
	SYS_CALL_B	$131

// func exit(code int32)
TEXT runtime·exit(SB), NOFRAME|NOSPLIT, $0-4
	LDW	R16, code+0(FP)
	SYSCALL(SYS_exit_group)
	RET

// func exitThread(wait *uint32)
TEXT runtime·exitThread(SB), NOFRAME|NOSPLIT, $0-8
	LDL	R1, wait+0(FP)
	MEMB
	STW	ZERO, (R1)
	MEMB

	LDI	R16, $0
	SYSCALL(SYS_exit)
	JMP	0(PC)
//zxw new change
// func write(fd uintptr, p unsafe.Pointer, n int32) int32
TEXT runtime·write1(SB), NOSPLIT, $0-28
	LDL	R16, fd+0(FP)
	LDL	R17, p+8(FP)
	LDW	R18, n+16(FP)
	SYSCALL(SYS_write)
	BEQ	R19, ok
	SUBL	ZERO, R0, R0
ok:
	STW	R0, ret+24(FP)
	RET

// func read(fd int32, p unsafe.Pointer, n int32) int32
TEXT runtime·read(SB),NOSPLIT,$0-28
	LDW	R16, fd+0(FP)
	LDL	R17, p+8(FP)
	LDW	R18, n+16(FP)
	SYSCALL(SYS_read)

	BEQ	R19, ok
	SUBL	ZERO, R0, R0
ok:
	STW	R0, ret+24(FP)
	RET

// func open(name *byte, mode, perm int32) int32
TEXT runtime·open(SB), NOSPLIT, $0-20
	LDL	R16, name+0(FP)
	LDW	R17, mode+8(FP)
	LDW	R18, perm+12(FP)
	SYSCALL(SYS_open)

	BEQ	R19, ok
	LDI	R0, -1
ok:
	STW	R0, ret+16(FP)
	RET

// func closefd(fd int32) int32
TEXT runtime·closefd(SB), NOSPLIT, $0-12
	LDW	R16, fd+0(FP)
	SYSCALL(SYS_close)

	BEQ	R19, ok
	LDI	R0, $-1
ok:
	STW	R0, ret+8(FP)
	RET

// func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) (p unsafe.Pointer, err int)
TEXT runtime·mmap(SB),NOSPLIT,$0-48
	LDL	R16, addr+0(FP)
	LDL	R17, n+8(FP)
	LDW	R18, prot+16(FP)
	LDW	R19, flags+20(FP)
	LDW	R20, fd+24(FP)
	LDW	R21, off+28(FP)

	SYSCALL(SYS_mmap)

	BEQ	R19, ok
	STL	ZERO, p+32(FP)
	STL	R0, err+40(FP)
	RET
ok:
	STL	R0, p+32(FP)
	STL	ZERO, err+40(FP)
	RET

// func munmap(addr unsafe.Pointer, n uintptr)
TEXT runtime·munmap(SB), NOSPLIT, $0-16
	LDL	R16, addr+0(FP)
	LDL	R17, n+8(FP)
	SYSCALL(SYS_munmap)
	RET

// func mincore(addr unsafe.Pointer, n uintptr, dst *byte) int32
TEXT runtime·mincore(SB),NOSPLIT,$0-28
	LDL	R16, addr+0(FP)
	LDL	R17, n+8(FP)
	LDL	R18, dst+16(FP)
	SYSCALL(SYS_mincore)
	SUBL	ZERO, R0, R0
	STW	R0, ret+24(FP)
	RET

// func sched_getaffinity(pid, len uintptr, buf *byte) int32
TEXT runtime·sched_getaffinity(SB),NOSPLIT,$0-28
	LDL	R16, pid+0(FP)
	LDL	R17, len+8(FP)
	LDL	R18, buf+16(FP)
	SYSCALL(SYS_sched_getaffinity)
	BEQ	R19, ok
	LDI	R0, -1
ok:
	STW	R0, ret+24(FP)
	RET

// func sbrk0() uintptr
TEXT runtime·sbrk0(SB), NOSPLIT, $0-8
	// Implemented as brk(NULL).
	LDI	R16, $0
	SYSCALL(SYS_brk)
	BEQ	R19, ok
	SUBL	ZERO, R0, R0 
ok:
	STL	R0, ret+0(FP)
	RET

TEXT runtime·osyield(SB), NOFRAME|NOSPLIT, $0
	SYSCALL(SYS_sched_yield)
	RET

//zxw new change
/*TEXT runtime·nanotime1(SB), NOSPLIT, $16-8
	LDI R16, $1 // CLOCK_MONOTONIC
	LDI R17, 8(SP)
	SYSCALL(SYS_clock_gettime)

	LDL R0, 8(SP) //sec
	LDL R1, 16(SP) //nsec

	LDI R2, $1000000000
	MULL R0, R2, R0
	ADDL R0, R1, R0
	STL R0, ret+0(FP)
	RET
*/
TEXT runtime·nanotime1(SB), NOSPLIT, $16-8
	LDI	R9, R30    // R9 is unchanged by C code
	LDI	R1, R30

	LDL	R10, g_m(g) // R10 = m

	// Set vdsoPC and vdsoSP for SIGPROF traceback.
	STL	R26, m_vdsoPC(R10)
	STL	R30, m_vdsoSP(R10)

	LDL	R16, m_curg(R10)
	LDI	R17, g
	CMPEQ	R16, R17, R7
	BEQ	R7, noswitch

	LDL	R16, m_g0(R10)
	LDL	R1, (g_sched+gobuf_sp)(R16)  // Set SP to g0 stack

noswitch:
	SUBL	R1, $16, R1
	AND	R1, $~15, R1  // Align for C code
	LDI	R30, R1

	LDI	R16, $1  // CLOCK_MONOTONIC
	LDI	R17, $0(R30)

	LDL	R27, runtime·vdsoClockgettimeSym(SB)
	BEQ	R27, fallback

	CALL	R26, (R27)

finish:
	LDL	R0, 0(R30)   // sec
	LDL	R17, 8(R30)  // nsec

	LDI	R30, R9      // restore SP
	STL	ZERO, m_vdsoSP(R10)  // clear vdsoSP

	// sec is in R3, nsec in R5
        // return nsec in R3
	LDI	R16, $1000000000
	MULL	R0, R16, R0
	ADDL	R0, R17, R0
	STL	R0, ret+0(FP)
	RET

fallback:
	SYSCALL(SYS_clock_gettime)
	JMP	finish

TEXT runtime·usleep(SB),NOSPLIT,$16-4
	LDW	R1, usec+0(FP)
	LDI	R0, $1000000
	CALL	runtime·udiv(SB)
	STL	R0, 8(SP)
	STL	R1, 16(SP)

	// select(0, 0, 0, 0, &tv)
	LDI	R16, $0
	LDI	R17, $0
	LDI	R18, $0
	LDI	R19, $0
	LDI	R20, $8(SP)
	SYSCALL(SYS_select)
	RET

//zxw new change
/*
//func walltime() (sec int64, nsec int32)
TEXT runtime·walltime1(SB), NOSPLIT, $16-12
    LDI R16, $0  // CLOCK_REALTIME
    LDI R17, 8(SP) //struct timespec*
    SYSCALL(SYS_clock_gettime)
    LDL R0, 8(SP)
    LDL R1, 16(SP)
    STL R0, sec+0(FP)
    STW R1, nsec+8(FP)
    RET
*/
TEXT runtime·walltime1(SB), NOSPLIT, $16-12
	LDI	R9, R30    // R9 is unchanged by C code
	LDI	R1, R30

	LDL	R10, g_m(g) // R10 = m

	// Set vdsoPC and vdsoSP for SIGPROF traceback.
	STL	R26, m_vdsoPC(R10)
	STL	R30, m_vdsoSP(R10)

	LDL	R16, m_curg(R10)
	LDI	R17, g
	CMPEQ	R16, R17, R7
	BEQ	R7, noswitch

	LDL	R16, m_g0(R10)
	LDL	R1, (g_sched+gobuf_sp)(R16)  // Set SP to g0 stack

noswitch:
	SUBL	R1, $16, R1
	AND	R1, $~15, R1  // Align for C code
	LDI	R30, R1

	LDI	R16, $0  // CLOCK_MONOTONIC
	LDI	R17, $0(R30)

	LDL	R27, runtime·vdsoClockgettimeSym(SB)
	BEQ	R27, fallback

	CALL	R26, (R27)

finish:
	LDL	R0, 0(R30)   // sec
	LDL	R17, 8(R30)  // nsec

	LDI	R30, R9      // restore SP
	STL	ZERO, m_vdsoSP(R10)  // clear vdsoSP

	STL	R0, sec+0(FP)
	STW	R17, nsec+8(FP)
	RET

fallback:
	SYSCALL(SYS_clock_gettime)
	JMP	finish

// func madvise(addr unsafe.Pointer, n uintptr, flags int32)
TEXT runtime·madvise(SB),NOSPLIT,$0-28
	LDL	R16, addr+0(FP)
	LDL	R17, n+8(FP)
	LDW	R18, flags+16(FP)
	SYSCALL(SYS_madvise)
	STW	R0,ret+24(FP)
	// ignore failure - maybe pages are locked
	RET

// func rtsigprocmask(how int32, new, old *sigset, size int32)
TEXT runtime·rtsigprocmask(SB), NOFRAME|NOSPLIT,$0-28
	LDW	R16, how+0(FP) //there has 4byte pad space
	LDL	R17, new+8(FP)
	LDL	R18, old+16(FP)
	LDW	R19, size+24(FP)
	SYSCALL(SYS_rt_sigprocmask)
	BEQ	R19, ok
	LDL	ZERO, 0xf1(ZERO) //crash
	LDI	R0, -1
ok:
	RET

TEXT glibc_sigreturn<>(SB), NOFRAME|NOSPLIT, $0
	LDI	R16, SP
	SYSCALL(SYS_rt_sigreturn)

// func rt_sigaction(sig uintptr, new, old *sigactiont, size uintptr) int32
TEXT runtime·rt_sigaction(SB), NOFRAME|NOSPLIT, $0-36
	LDL	R16, sig+0(FP)
	LDL	R17, new+8(FP)
	LDL	R18, old+16(FP)
	LDL	R19, size+24(FP)
	SYMADDR	R20, glibc_sigreturn<>(SB)
	
	SYSCALL(SYS_rt_sigaction)
	
	BEQ	R19, ok
	STW	R19, ret+32(FP)
	RET
ok:
	STW	ZERO, ret+32(FP)
	RET

TEXT runtime·cgoSigtramp(SB), NOFRAME|NOSPLIT, $0
	JMP	runtime·sigtramp(SB)

// void (*sa_sigaction)(int, siginfo_t *, void *);
TEXT runtime·sigtramp(SB), NOSPLIT, $64
	// initialize REGSB = PC&0xffffffff00000000
	// BGEZAL	R0, 1(PC)
	// SRLV	$32, R31, RSB
	// SLLV	$32, RSB

	// this might be called in external code context,
	// where g is not set.
	LDBU	R1, runtime·iscgo(SB)
	BEQ	R1, ok
	CALL	runtime·load_g(SB)
ok:
	STW	R16, 8(SP) // sig
	STL	R17, 16(SP) // info
	STL	R18, 24(SP) // ctx
	// func sigtrampgo(sig uint32, info *siginfo, ctx unsafe.Pointer)
	CALL	runtime·sigtrampgo(SB)
	LDL	R17, 16(SP)
	STL	R17, 8(SP)
	CALL	sw_sigreturn(SB)
	RET

// func sw_sigreturn(rt_sigframe)
TEXT sw_sigreturn(SB), NOSPLIT, $0-0
	LDL	R16, $frame+0(FP)
	SYSCALL(SYS_rt_sigreturn)
	RET

// func sigfwd(fn uintptr, sig uint32, info *siginfo, ctx unsafe.Pointer)
TEXT runtime·sigfwd(SB), NOSPLIT, $0-32
	LDW	R16, sig+8(FP)
	LDL	R17, info+16(FP)
	LDL	R18, ctx+24(FP)
	LDL	R27, fn+0(FP)
	CALL	(R27)
	RET

// func futex(addr unsafe.Pointer, op int32, val uint32, ts, addr2 unsafe.Pointer, val3 uint32) int32
TEXT runtime·futex(SB), NOFRAME|NOSPLIT, $0-44
	LDL	R16, addr+0(FP)
	LDW	R17, op+8(FP)
	LDW	R18, val+12(FP)
	LDL	R19, ts+16(FP)
	LDL	R20, addr2+24(FP)
	LDW	R21, val3+32(FP)
	SYSCALL(SYS_futex)
	BEQ	R19, ok
	SUBL	ZERO, R0, R0 
ok:
	STW	R0, ret+40(FP)
	RET

// func sigaltstack(new, old *stackt)
TEXT runtime·sigaltstack(SB), NOFRAME|NOSPLIT, $0-16
	LDL	R16, new+0(FP)
	LDL	R17, old+8(FP)
	SYSCALL(SYS_sigaltstack)
	BEQ	R19, ok
	LDL	R19, (ZERO) // crash
ok:
	RET

// func gettid() uint32
TEXT runtime·gettid(SB), NOFRAME|NOSPLIT, $0-4
	SYSCALL(SYS_gettid)
	STW	R0, ret+0(FP)
	RET

// func clone(flags int32, stk, mp, gp, fn unsafe.Pointer) int32
TEXT runtime·clone(SB), NOSPLIT, $0-44
	LDW	R16, flags+0(FP)
	LDL	R17, stk+8(FP)

	// Copy mp, gp, fn off parent stack for use by child.
	// Careful: Linux system call clobbers ???.
	LDL	R1, mp+16(FP)
	LDL	R2, gp+24(FP)
	LDL	R3, fn+32(FP)

	STL	R1, -8(R17) // mp
	STL	R2, -16(R17) // gp
	STL	R3, -24(R17) // fn

	LDI	R1, $77
	STL	R1, -32(R17) // guard
	SYSCALL(SYS_clone)
	BEQ	R0, child
	STW	R0, ret+40(FP)  // on parent
	RET
child:
	LDL	R4, -32(SP)     // on child
	LDI	R5, $77
	CMPEQ	R4, R5, R5
	BNE	R5, guard_ok
	LDL	R5, 0(ZERO)
guard_ok:
	// set up new stack
	LDL	R27, -24(SP) // fn
	LDL	R2, -16(SP) // g
	LDL	R1, -8(SP)  // m
	BEQ	R2, nog
	BEQ	R1, nog
	
	SYSCALL(SYS_gettid)
	STL	R0, m_procid(R1)

	STL	R1, g_m(R2)
	LDI	g, R2
	//CALL	runtime·stackcheck(SB)
nog:
	// Call fn
	CALL	(R27)

	// It shouldn't return.	 If it does, exit that thread.
	LDI	R16, $111
	SYSCALL(SYS_exit)
	RET

// func raise(sig uint32)
TEXT runtime·raise(SB), NOFRAME|NOSPLIT, $0
	SYSCALL(SYS_getxpid)
        LDI     R9, R0
        SYSCALL(SYS_gettid)
        LDI     R17, R0 // arg 2 tid
        LDI     R16, R9 // arg 1 pid
        LDW     R18, sig+0(FP) // arg 3
        SYSCALL(SYS_tgkill)
        RET

// func raiseproc(sig uint32)
TEXT runtime·raiseproc(SB), NOFRAME|NOSPLIT, $0
	SYSCALL(SYS_getxpid)
	LDI	R16, R0
	LDW	R17, sig+0(FP)
	SYSCALL(SYS_kill)
	RET

// int32 runtime·epollcreate(int32 flags);
TEXT runtime·epollcreate(SB), NOFRAME|NOSPLIT, $0
	LDW	R16, size+0(FP)
	SYSCALL(SYS_epoll_create)
	BEQ	R19, ok
	SUBL	ZERO, R0, R0 
ok:
	STW	R0, ret+8(FP)
	RET

// int32 runtime·epollcreate1(int32 flags);
TEXT runtime·epollcreate1(SB), NOFRAME|NOSPLIT, $0
	LDW	R16, flags+0(FP)
	SYSCALL(SYS_epoll_create1)
	BEQ	R19, ok
	SUBL	ZERO, R0, R0 
ok:
	STW	R0, ret+8(FP)
	RET

// func epollctl(epfd, op, fd int32, ev *epollEvent) int
TEXT runtime·epollctl(SB), NOFRAME|NOSPLIT, $0
	LDW	R16, epfd+0(FP)
	LDW	R17, op+4(FP)
	LDW	R18, fd+8(FP)
	LDL	R19, ev+16(FP)
	SYSCALL(SYS_epoll_ctl)
	BEQ	R19, ok
	SUBL	ZERO, R0, R0 
ok:
	STL	R0, ret+24(FP)
	RET

// int32 runtime·epollwait(int32 epfd, EpollEvent *ev, int32 nev, int32 timeout);
TEXT runtime·epollwait(SB), NOSPLIT|NOFRAME, $0
	LDW	R16, epfd+0(FP)
	LDL	R17, ev+8(FP)
	LDW	R18, nev+16(FP)
	LDW	R19, timeout+20(FP)
	LDI	R20, ZERO
	SYSCALL(SYS_epoll_pwait)
	BEQ	R19, ok
	SUBL	ZERO, R0, R0 
ok:
	STW	R0, ret+24(FP)
	RET

// void runtime·closeonexec(int32 fd);
TEXT runtime·closeonexec(SB), NOFRAME|NOSPLIT, $0
	LDW	R16, fd+0(FP)
	LDI	R17, $2 // F_SETFD
	LDI	R18, $1 // FD_CLOEXEC
	SYSCALL(SYS_fcntl)
	RET

// func setitimer(mode int32, new, old *itimerval)
TEXT runtime·setitimer(SB), NOFRAME|NOSPLIT, $0-24
	LDW	R16, mode+0(FP)
	LDL	R17, new+8(FP)
	LDL	R18, old+16(FP)
	SYSCALL(SYS_setitimer)
	RET

//zxw new add
// func pipe() (r, w int32, errno int32)
TEXT runtime·pipe(SB),NOSPLIT|NOFRAME,$0-12
	LDI	R16, $r+0(FP)
	LDI	R17, ZERO
	SYSCALL(SYS_pipe2)
	BEQ	R19, ok
	SUBL	ZERO, R0, R0 
ok:
	STW	R0, errno+8(FP)
	RET

// func pipe2(flags int32) (r, w int32, errno int32)
TEXT runtime·pipe2(SB),NOSPLIT|NOFRAME,$0-20
	LDI	R16, $r+8(FP)
	LDW	R17, flags+0(FP)
	SYSCALL(SYS_pipe2)
	BEQ	R19, ok
	SUBL	ZERO, R0, R0 
ok:
	STW	R0, errno+16(FP)
	RET

// func runtime·setNonblock(int32 fd)
TEXT runtime·setNonblock(SB),NOSPLIT|NOFRAME,$0-4
	LDW	R16, fd+0(FP) // fd
	LDI	R17, $3	// F_GETFL
	LDI	R18, $0
	SYSCALL(SYS_fcntl)
	LDI	R18, $0x4 // O_NONBLOCK
	BIS	R18, R0, R18
	LDW	R16, fd+0(FP) // fd
	LDI	R17, $4	// F_SETFL
	SYSCALL(SYS_fcntl)
	RET

TEXT ·getpid(SB),NOSPLIT|NOFRAME,$0-8
	SYSCALL(SYS_getxpid)
	STL	R0, ret+0(FP)
	RET

TEXT ·tgkill(SB),NOSPLIT|NOFRAME,$0-24
	LDL	R16, tgid+0(FP)
	LDL	R17, tid+8(FP)
	LDL	R18, sig+16(FP)
	SYSCALL(SYS_tgkill)
	RET
