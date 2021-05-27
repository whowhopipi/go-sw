// +build sw64
// +build linux

package runtime

const (
	_EINTR  = 0x4
	_EAGAIN = 0x23
	_ENOMEM = 0xc
	_ENOSYS = 0x59

	_PROT_NONE  = 0x0
	_PROT_READ  = 0x1
	_PROT_WRITE = 0x2
	_PROT_EXEC  = 0x4

	_MAP_ANON    = 0x10
	_MAP_PRIVATE = 0x2
	_MAP_FIXED   = 0x100

	_MADV_DONTNEED   = 0x6
	_MADV_FREE       = 0x8
	_MADV_HUGEPAGE   = 0xe
	_MADV_NOHUGEPAGE = 0xf

	_SA_ONSTACK = 0x1
	_SA_RESTART = 0x2
	_SA_SIGINFO = 0x40

	_SIGHUP    = 0x1
	_SIGINT    = 0x2
	_SIGQUIT   = 0x3
	_SIGILL    = 0x4
	_SIGTRAP   = 0x5
	_SIGABRT   = 0x6
	_SIGEMT    = 0x7
	_SIGFPE    = 0x8
	_SIGKILL   = 0x9
	_SIGBUS    = 0xa
	_SIGSEGV   = 0xb
	_SIGSYS    = 0xc
	_SIGPIPE   = 0xd
	_SIGALRM   = 0xe
	_SIGTERM   = 0xf
	_SIGURG    = 0x10
	_SIGSTOP   = 0x11
	_SIGTSTP   = 0x12
	_SIGCONT   = 0x13
	_SIGCHLD   = 0x14
	_SIGTTIN   = 0x15
	_SIGTTOU   = 0x16
	_SIGIO     = 0x17
	_SIGXCPU   = 0x18
	_SIGXFSZ   = 0x19
	_SIGVTALRM = 0x1a
	_SIGPROF   = 0x1b
	_SIGWINCH  = 0x1c
	_SIGINFO   = 0x1d
	_SIGUSR1   = 0x1e
	_SIGUSR2   = 0x1f

	_FPE_INTDIV = 0x1
	_FPE_INTOVF = 0x2
	_FPE_FLTDIV = 0x3
	_FPE_FLTOVF = 0x4
	_FPE_FLTUND = 0x5
	_FPE_FLTRES = 0x6
	_FPE_FLTINV = 0x7
	_FPE_FLTSUB = 0x8

	_BUS_ADRALN = 0x1
	_BUS_ADRERR = 0x2
	_BUS_OBJERR = 0x3

	_SEGV_MAPERR = 0x1
	_SEGV_ACCERR = 0x2

	_ITIMER_REAL    = 0x0
	_ITIMER_VIRTUAL = 0x1
	_ITIMER_PROF    = 0x2

	_EPOLLIN       = 0x1
	_EPOLLOUT      = 0x4
	_EPOLLERR      = 0x8
	_EPOLLHUP      = 0x10
	_EPOLLRDHUP    = 0x2000
	_EPOLLET       = 0x80000000
	_EPOLL_CLOEXEC = 0x200000
	_EPOLL_CTL_ADD = 0x1
	_EPOLL_CTL_DEL = 0x2
	_EPOLL_CTL_MOD = 0x3
)

type usigset struct {
	__val [16]uint64
}

type timespec struct {
	tv_sec  int64
	tv_nsec int64
}

//go:nosplit
func (ts *timespec) setNsec(ns int64) {
	ts.tv_sec = ns / 1e9
	ts.tv_nsec = ns % 1e9
}

/*
func (ts *timespec) set_sec(x int64) {
	ts.tv_sec = x
}

func (ts *timespec) set_nsec(x int32) {
	ts.tv_nsec = int64(x)
}
*/
type timeval struct {
	tv_sec  int64
	tv_usec int64
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = int64(x)
}

type sigactiont struct {
	sa_handler uintptr
	sa_flags   int32
	sa_mask    [1]uint64

	sa_restorer uintptr //noused
}

type siginfo struct {
	si_signo   int32
	si_errno   int32
	si_code    int32
	_pad_cgo_0 [4]byte

	si_addr     uint64
	_X_sifields [128 - 8*3]byte
}

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

type epollevent struct {
	events    uint32
	pad_cgo_0 [4]byte
	data      uint64
}

const (
	_O_RDONLY    = 0x0
	_O_NONBLOCK  = 0x4
	_O_CLOEXEC   = 0x200000
	_SA_RESTORER = 0 //don't set this
)

type stackt struct {
	ss_sp      *byte
	ss_flags   int32
	_pad_cgo_0 [4]byte
	ss_size    uintptr
}

type sigcontext struct {
	sc_onstack uint64
	sc_mask    uint64
	sc_pc      uint64
	sc_ps      uint64
	sc_regs    [32]uint64
	sc_ownedfp int64
	sc_fpregs  [128]int64 // if not defined CONFIG_SW_SMID

	sc_fpcr            uint64
	sc_fp_control      uint64
	sc_reserved1       uint64
	sc_reserved2       uint64
	sc_ssize           uint64
	sc_sbase           *int8
	sc_traparg_a0      uint64
	sc_traparg_a1      uint64
	sc_traparg_a2      uint64
	sc_fp_trap_pc      uint64
	sc_fp_trigger_sum  uint64
	sc_fp_trigger_inst uint64
}

type ucontext struct {
	uc_flags         uint64
	uc_link          *ucontext
	__uc_osf_sigmask uint64
	uc_stack         stackt
	uc_mcontext      sigcontext
	uc_sigmask       usigset
}
