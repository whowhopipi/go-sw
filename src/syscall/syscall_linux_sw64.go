// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux
// +build sw64

package syscall

const (
	//generate by handle
	//_snyh_TODO: this should be generate by improving build script
	FIOCLEX    = 0x20006601
	FIONCLEX   = 0x20006602
	FIOASYNC   = 0xffffffff8004667d
	FIONBIO    = 0xffffffff8004667e
	FIONREAD   = 0x4004667f
	TIOCINQ    = 0x4004667f
	FIOQSIZE   = 0x40086680
	TIOCGETP   = 0x40067408
	TIOCSETP   = 0xffffffff80067409
	TIOCSETN   = 0xffffffff8006740a
	TIOCSETC   = 0xffffffff80067411
	TIOCGETC   = 0x40067412
	TCSETSW    = 0xffffffff802c7415
	TCSETSF    = 0xffffffff802c7416
	TCGETA     = 0x40127417
	TCSETA     = 0xffffffff80127418
	TCSETAW    = 0xffffffff80127419
	TCSETAF    = 0xffffffff8012741c
	TIOCSWINSZ = 0xffffffff80087467
	TIOCGWINSZ = 0x40087468
	TIOCGLTC   = 0x40067474
	TIOCSLTC   = 0xffffffff80067475
)

const (
	//ALL OF THIS constants are WORKAROUND, and should be removing
	_SYS_dup       = SYS_DUP2
	_SYS_setgroups = SYS_SETGROUPS
	SYS_UMOUNT2    = SYS_UMOUNT

	SYS_GETPID  = SYS_GETXPID
	SYS_GETPPID = SYS_GETXPID
)

//sysnb getxpid() (pid int, ppid int)
// TODO(snyh):  correct handle Getppid and Getpid
// currently manually remove the implements of Getpid and Getppid
// in zsyscall_linux_sw64.go
func Getpid() (pid int)   { pid, _ = getxpid(); return }
func Getppid() (ppid int) { _, ppid = getxpid(); return }

// TODO(snyh):  correct handle Utime
func Utime(path string, buf *Utimbuf) error {
	tv := [2]Timeval{
		{Sec: buf.Actime},
		{Sec: buf.Modtime},
	}
	return utimes(path, &tv)
}

//sys	Fstat64(fd int, st *Stat_t) (err error)
//sys	Lstat64(path string, st *Stat_t) (err error)
//sys	Stat64(path string, st *Stat_t) (err error)
func Fstat(fd int, st *Stat_t) (err error)      { return Fstat64(fd, st) }
func Lstat(path string, st *Stat_t) (err error) { return Lstat64(path, st) }
func Stat(path string, st *Stat_t) (err error)  { return Stat64(path, st) }

//sys getxuid() (uid int, euid int)
func Getuid() (uid int)   { uid, _ = getxuid(); return }
func Geteuid() (euid int) { _, euid = getxuid(); return }

//sys getxgid() (gid int, egid int)
func Getgid() (gid int)   { gid, _ = getxgid(); return }
func Getegid() (egid int) { _, egid = getxgid(); return }

//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	Fstatfs(fd int, buf *Statfs_t) (err error)
//sys	Dup2(oldfd int, newfd int) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error)
//sysnb	InotifyInit() (fd int, err error)
//sys	Lchown(path string, uid int, gid int) (err error)
//sys	Listen(s int, n int) (err error)
//sys	Pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sys	Pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sys	Setfsgid(gid int) (err error)
//sys	Setfsuid(uid int) (err error)
//sysnb	Setregid(rgid int, egid int) (err error)
//sysnb	Setresgid(rgid int, egid int, sgid int) (err error)
//sysnb	Setresuid(ruid int, euid int, suid int) (err error)
//sysnb	Setrlimit(resource int, rlim *Rlimit) (err error)
//sysnb	Setreuid(ruid int, euid int) (err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)

//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	Truncate(path string, length int64) (err error)
//sys	accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error)
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sysnb	setgroups(n int, list *_Gid_t) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sysnb	socket(domain int, typ int, proto int) (fd int, err error)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)

type sigset_t struct {
	X__val [16]uint64
}

//sys	pselect(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timespec, sigmask *sigset_t) (n int, err error) = SYS_PSELECT6

func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) {
	ts := Timespec{Sec: timeout.Sec, Nsec: timeout.Usec * 1000}
	return pselect(nfd, r, w, e, &ts, nil)
}

//sysnb	Gettimeofday(tv *Timeval) (err error)

func Time(t *Time_t) (tt Time_t, err error) {
	var tv Timeval
	err = Gettimeofday(&tv)
	if err != nil {
		return 0, err
	}
	if t != nil {
		*t = Time_t(tv.Sec)
	}
	return Time_t(tv.Sec), nil
}

func setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: sec, Nsec: nsec}
}

func setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: sec, Usec: usec}
}

func Pipe(p []int) (err error) {
	if len(p) != 2 {
		return EINVAL
	}
	var pp [2]_C_int
	err = pipe2(&pp, 0)
	p[0] = int(pp[0])
	p[1] = int(pp[1])
	return
}

//sysnb pipe2(p *[2]_C_int, flags int) (err error)

func Pipe2(p []int, flags int) (err error) {
	if len(p) != 2 {
		return EINVAL
	}
	var pp [2]_C_int
	err = pipe2(&pp, flags)
	p[0] = int(pp[0])
	p[1] = int(pp[1])
	return
}

func Ioperm(from int, num int, on int) (err error) {
	return ENOSYS
}

func Iopl(level int) (err error) {
	return ENOSYS
}

// func (r *PtraceRegs) PC() uint64 { return r.Regs[64] }

// func (r *PtraceRegs) SetPC(pc uint64) { r.Regs[64] = pc }

func (iov *Iovec) SetLen(length int) {
	iov.Len = uint64(length)
}

func (msghdr *Msghdr) SetControllen(length int) {
	msghdr.Controllen = uint64(length)
}

func (cmsg *Cmsghdr) SetLen(length int) {
	cmsg.Len = uint64(length)
}

func rawVforkSyscall(trap, a1 uintptr) (r1 uintptr, err Errno) {
	panic("not implemented")
}

//zxw new add
//sysnb	EpollCreate(size int) (fd int, err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	Ustat(dev int, ubuf *Ustat_t) (err error)
//sys	futimesat(dirfd int, path string, times *[2]Timeval) (err error)
//sys	utimes(path string, times *[2]Timeval) (err error)
//sys   fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sys   EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
