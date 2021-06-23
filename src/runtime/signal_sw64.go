// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux
// +build sw64

package runtime

import (
	"runtime/internal/sys"
	"unsafe"
)

// snyh_TODO: FIX ALL OF THIS

func dumpregs(c *sigctxt) {
	println("r0\t", hex(c.r0()))
	println("r1\t", hex(c.r1()))
	println("r2\t", hex(c.r2()))
	println("r3\t", hex(c.r3()))
	println("r4\t", hex(c.r4()))
	println("r5\t", hex(c.r5()))
	println("r6\t", hex(c.r6()))
	println("r7\t", hex(c.r7()))
	println("r8\t", hex(c.r8()))
	println("r9\t", hex(c.r9()))
	println("r10\t", hex(c.r10()))
	println("r11\t", hex(c.r11()))
	println("r12\t", hex(c.r12()))
	println("r13\t", hex(c.r13()))
	println("r14\t", hex(c.r14()))
	println("g\t", hex(c.g()))
	println("r16\t", hex(c.r16()))
	println("r17\t", hex(c.r17()))
	println("r18\t", hex(c.r18()))
	println("r19\t", hex(c.r19()))
	println("r20\t", hex(c.r20()))
	println("r21\t", hex(c.r21()))
	println("r22\t", hex(c.r22()))
	println("r23\t", hex(c.r23()))
	println("r24\t", hex(c.r24()))
	println("r25\t", hex(c.r25()))
	println("ra\t", hex(c.link()))
	println("ctxt\t", hex(c.regctxt()))
	println("rtmp\t", hex(c.rtmp()))
	println("r29\t", hex(c.r29()))
	println("sp\t", hex(c.sp()))
}

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) sigpc() uintptr { return uintptr(c.pc()) }

func (c *sigctxt) sigsp() uintptr { return uintptr(c.sp()) }
func (c *sigctxt) siglr() uintptr { return uintptr(c.link()) }
func (c *sigctxt) fault() uintptr { return uintptr(c.sigaddr()) }

// preparePanic sets up the stack to look like a call to sigpanic.
func (c *sigctxt) preparePanic(sig uint32, gp *g) {
	// We arrange link, and pc to pretend the panicking
	// function calls sigpanic directly.
	// Always save LINK to stack so that panics in leaf
	// functions are correctly handled. This smashes
	// the stack frame but we're not going back there
	// anyway.
	sp := c.sp() - sys.PtrSize
	c.set_sp(sp)
	*(*uint64)(unsafe.Pointer(uintptr(sp))) = c.link()

	pc := gp.sigpc

	//zxw change
	if shouldPushSigpanic(gp, pc, uintptr(c.link())) {
		// Make it look the like faulting PC called sigpanic.
		c.set_link(uint64(pc))
	}
	// In case we are panicking from external C code
	sigpanicPC := uint64(funcPC(sigpanic))
	c.set_g(uint64(uintptr(unsafe.Pointer(gp))))
	c.set_pc(sigpanicPC)
}

func (c *sigctxt) pushCall(targetPC, resumePC uintptr) {
	// Push the LR to stack, as we'll clobber it in order to
	// push the call. The function being pushed is responsible
	// for restoring the LR and setting the SP back.
	// This extra slot is known to gentraceback.
	sp := c.sp() - 8
	c.set_sp(sp)
	*(*uint64)(unsafe.Pointer(uintptr(sp))) = c.link()
	// Set up PC and LR to pretend the function being signaled
	// calls targetPC at the faulting PC.
	c.set_link(uint64(resumePC))
	c.set_pc(uint64(targetPC))
}
