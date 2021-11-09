// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw64

import (
	"cmd/internal/obj"
	"cmd/internal/objabi"
	"cmd/internal/sys"
)

func (c *ctxt77) stacksplit(p *obj.Prog, framesize int32) *obj.Prog {
	//还有问题，不要开启。
	//目前是在src/runtime/stack.go中将StackMin调整到2M，避免stack不够的情况
	//	return p

	//	if framesize == 0 {
	//		return p
	//	}
	p = insertLDGP(p, c.newprog, REGZERO)

	// LDL R1, g_stackguard(g)
	p = obj.Appendp(p, c.newprog)
	begin := p

	p.As = ALDL
	p.From.Type = obj.TYPE_REG
	p.From.Reg = REG_R1

	p.To.Type = obj.TYPE_MEM
	p.To.Reg = REGG
	p.To.Offset = 2 * int64(c.ctxt.Arch.PtrSize) // G.stackguard0
	if c.cursym.CFunc() {
		p.To.Offset = 3 * int64(c.ctxt.Arch.PtrSize) // G.stackguard1
	}

	//zxw new add
	// Mark the stack bound check and morestack call async nonpreemptible.
	// If we get preempted here, when resumed the preemption request is
	// cleared, but we'll still call morestack, which will double the stack
	// unnecessarily. See issue #35470.
	p = c.ctxt.StartUnsafePoint(p, c.newprog)

	var q *obj.Prog
	if framesize <= objabi.StackSmall {
		// small stack: SP < stackguard
		// CMPULE SP, stackguard, R1
		p = obj.Appendp(p, c.newprog)

		p.As = ACMPULE
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REGSP
		p.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  REG_R1,
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R1

	} else if framesize <= objabi.StackBig {
		// large stack: SP-framesize < stackguard-StackSmall
		// SUBL SP, $(framesize-StackSmall), R2
		// CMPULE R2, stackgurad, R1
		p = obj.Appendp(p, c.newprog)

		p.As = ASUBL
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REGSP

		p.SetFrom3(obj.Addr{
			Type:   obj.TYPE_CONST,
			Offset: (int64(framesize) - objabi.StackSmall),
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = ACMPULE
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R2
		p.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  REG_R1,
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R1
	} else {
		// Such a large stack we need to protect against wraparound.
		// If SP is close to zero:
		//	SP-stackguard+StackGuard <= framesize + (StackGuard-StackSmall)
		// The +StackGuard on both sides is required to keep the left side positive:
		// SP is allowed to be slightly below stackguard. See stack.h.
		//
		// Preemption sets stackguard to StackPreempt, a very large value.
		// That breaks the math above, so we have to check for that explicitly.
		//	// stackguard is R1

		//	LDI R2, $StackPreempt
		p = obj.Appendp(p, c.newprog)
		p.As = ALDI
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = objabi.StackPreempt
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R2

		// CMPEQ R1, R2, R0
		// BNE R0, label-of-call-to-morestack
		p = obj.Appendp(p, c.newprog)
		p.As = ACMPEQ
		p.From.Type, p.From.Reg = obj.TYPE_REG, REG_R1
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: REG_R2})
		p.To.Type, p.To.Reg = obj.TYPE_REG, REG_R0
		p = obj.Appendp(p, c.newprog)
		q = p
		p.As = ABNE
		p.From.Type, p.From.Reg = obj.TYPE_REG, REG_R0
		p.To.Type = obj.TYPE_BRANCH

		//	ADDL SP, $StackGuard, R2
		p = obj.Appendp(p, c.newprog)
		p.As = AADDL
		p.From.Type, p.From.Reg = obj.TYPE_REG, REGSP
		p.SetFrom3(obj.Addr{Type: obj.TYPE_CONST, Offset: int64(objabi.StackGuard)}) //zxw new change
		p.To.Type, p.To.Reg = obj.TYPE_REG, REG_R2

		//	SUBL R2, R1, R2
		p = obj.Appendp(p, c.newprog)
		p.As = ASUBL
		p.From.Type, p.From.Reg = obj.TYPE_REG, REG_R2
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: REG_R1})
		p.To.Type, p.To.Reg = obj.TYPE_REG, REG_R2

		//	LDI R1, $(framesize+(StackGuard-StackSmall))
		v := int64(framesize) + int64(objabi.StackGuard) - objabi.StackSmall //zxw new change
		hi := int16(v >> 16)
		lo := int16(v & 0xffff)
		if lo < 0 {
			hi = hi + 1
			lo = int16(v - int64(hi)<<16)
		}
		p = obj.Appendp(p, c.newprog)
		p.As = ALDIH
		p.From.Type, p.From.Reg = obj.TYPE_REG, REG_R1
		p.To.Type, p.To.Offset = obj.TYPE_CONST, int64(hi)
		p = obj.Appendp(p, c.newprog)
		p.As = ALDI
		p.From.Type, p.From.Reg = obj.TYPE_REG, REG_R1
		p.To.Type, p.To.Offset, p.To.Reg = obj.TYPE_ADDR, int64(lo), REG_R1

		// CMPULE R2, R1, R1
		p = obj.Appendp(p, c.newprog)
		p.As = ACMPULE
		p.From.Type, p.From.Reg = obj.TYPE_REG, REG_R2
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: REG_R1})
		p.To.Type, p.To.Reg = obj.TYPE_REG, REG_R1
	}

	// q1: BEQ R1, done
	p = obj.Appendp(p, c.newprog)
	q1 := p

	p.As = ABEQ
	p.From.Type = obj.TYPE_REG
	p.From.Reg = REG_R1
	p.To.Type = obj.TYPE_BRANCH

	// LDI R3, LINK
	p = obj.Appendp(p, c.newprog)
	p.As = ALDI
	p.From.Type = obj.TYPE_REG
	p.From.Reg = REG_R3
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REGLINK
	if q != nil {
		q.To.Val = p
	}

	//zxw new change
	p = c.ctxt.EmitEntryStackMap(c.cursym, p, c.newprog)

	// JMP runtime.morestack(SB)
	p = obj.Appendp(p, c.newprog)

	p.As = AJMP
	p.From.Type = obj.TYPE_REG
	//Don't use RA, otherwise the GP value will wrong because when
	//the morestack return the RA is not the next instruction address.
	p.From.Reg = REG_R2
	p.To.Type = obj.TYPE_BRANCH
	p.To.Name = obj.NAME_EXTERN

	if c.cursym.CFunc() {
		p.To.Sym = c.ctxt.Lookup("runtime.morestackc")
	} else if !c.cursym.Func().Text.From.Sym.NeedCtxt() {
		p.To.Sym = c.ctxt.Lookup("runtime.morestack_noctxt")
	} else {
		p.To.Sym = c.ctxt.Lookup("runtime.morestack")
	}

	//zxw new add
	p = c.ctxt.EndUnsafePoint(p, c.newprog, -1)

	// JMP	start
	p = obj.Appendp(p, c.newprog)
	p.As = ABR
	p.From.Type, p.From.Reg = obj.TYPE_REG, REGZERO
	p.To.Type, p.To.Val = obj.TYPE_BRANCH, begin

	// placeholder for q1's jump target
	p = obj.Appendp(p, c.newprog)
	p.As = obj.ANOP // zero-width place holder
	q1.To.Val = p

	return p
}

func (c *ctxt77) addnop(p *obj.Prog) {
	q := c.newprog()
	q.As = ANOOP
	q.Pos = p.Pos
	q.Link = p.Link
	p.Link = q
}

func insertLDGP(q *obj.Prog, newprog obj.ProgAlloc, rip int16) *obj.Prog {
	q = obj.Appendp(q, newprog)
	q.As = ALDGP
	q.From = obj.Addr{Type: obj.TYPE_REG, Reg: REG_R29}
	q.To = obj.Addr{Type: obj.TYPE_REG, Reg: rip}
	return q
}

func preprocess(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	// 1. handle ARET and ATEXT to setup function prologue and epilogue
	// 2. adjust Prog.Spadj if the prog change hardward SP

	if cursym.Func().Text == nil || cursym.Func().Text.Link == nil {
		return
	}

	c := ctxt77{ctxt: ctxt, newprog: newprog, cursym: cursym}
	p := c.cursym.Func().Text
	textstksiz := p.To.Offset

	//zxw add
	if textstksiz == -ctxt.FixedFrameSize() {
		// Historical way to mark NOFRAME.
		p.From.Sym.Set(obj.AttrNoFrame, true)
		textstksiz = 0
	}
	if textstksiz < 0 {
		c.ctxt.Diag("negative frame size %d - did you mean NOFRAME?", textstksiz)
	}
	if p.From.Sym.NoFrame() {
		if textstksiz != 0 {
			c.ctxt.Diag("NOFRAME functions must have a frame size of 0, not %d", textstksiz)
		}
	}

	c.cursym.Func().Args = p.To.Val.(int32)
	c.cursym.Func().Locals = int32(textstksiz)

	if textstksiz < 0 {
		ctxt.Diag("Use NOFRAME attribute instead of a negative frame size.\n%v\n", p)
		return
	}

	//zxw add
	/*
	 * find leaf subroutines
	 * strip NOPs
	 * expand RET
	 */

	for p := c.cursym.Func().Text; p != nil; p = p.Link {
		switch p.As {
		case obj.ATEXT:
			p.Mark |= LEAF

		case obj.ARET:
			break

		case ACALL,
			obj.ADUFFZERO,
			obj.ADUFFCOPY:
			c.cursym.Func().Text.Mark &^= LEAF
			fallthrough

		case AJMP,
			ABR,
			ABSR,
			ABNE,
			ABEQ,
			ABLT,
			ABLE,
			ABGT,
			ABGE,
			ABLBC,
			ABLBS,
			AFBNE,
			AFBEQ,
			AFBLT,
			AFBLE,
			AFBGT,
			AFBGE:
			q1 := p.To.Target()

			if q1 != nil {
				for q1.As == obj.ANOP {
					q1 = q1.Link
					//p.Pcond = q1
					p.To.Val = q1 //zxw new change
				}
			}

			break
		}

	}

	autosize := int32(0)
	var retjmp *obj.LSym
	for p := cursym.Func().Text; p != nil; p = p.Link {
		o := p.As
		switch o {
		case obj.ATEXT:
			autosize = int32(textstksiz)
			q := p
			if p.Mark&LEAF != 0 && autosize == 0 {
				// A leaf function with no locals has no frame.
				p.From.Sym.Set(obj.AttrNoFrame, true)
			}

			if !p.From.Sym.NoFrame() {
				// If there is a stack frame at all, it includes
				// space to save the LR.
				autosize += int32(c.ctxt.FixedFrameSize())
			}
			if cursym.Attribute.NoFrame() {
				if textstksiz != 0 {
					ctxt.Diag("NoFrame symbol %v with non zero frame size(%d)", cursym, textstksiz)
				}
				autosize = 0
			}

			//convert virtual SP to hardware SP
			//see also: https://bj.git.sndu.cn/xiabin/go-sw64/wikis/stack-frame-layout
			for p := cursym.Func().Text; p != nil; p = p.Link {
				switch p.To.Name {
				case obj.NAME_PARAM:
					p.To.Type = obj.TYPE_ADDR
					p.To.Reg = REGSP
					p.To.Offset += int64(autosize) + c.ctxt.FixedFrameSize()
					p.To.Sym = nil
				case obj.NAME_AUTO:
					p.To.Type = obj.TYPE_ADDR
					p.To.Reg = REGSP
					p.To.Offset += int64(autosize)
					p.To.Sym = nil
				}
			}

			if autosize == 0 && c.cursym.Func().Text.Mark&LEAF == 0 {
				if c.cursym.Func().Text.From.Sym.NoSplit() {
					if ctxt.Debugvlog {
						ctxt.Logf("save suppressed in: %s\n", c.cursym.Name)
					}

					c.cursym.Func().Text.Mark |= LEAF
				}
			}

			//p.To.Offset = int64(autosize) - ctxt.FixedFrameSize()

			if c.cursym.Func().Text.Mark&LEAF != 0 {
				c.cursym.Set(obj.AttrLeaf, true)
				if p.From.Sym.NoFrame() {
					//break
				}
			}

			if !p.From.Sym.NoSplit() {
				q = c.stacksplit(q, autosize) // emit split check
			}

			if autosize != 0 {
				//zxw new add
				q = c.ctxt.StartUnsafePoint(q, c.newprog)

				// LDI SP, $-autosize(SP)
				q = obj.Appendp(q, newprog)
				q.As = ALDI
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REGSP
				q.To = obj.Addr{
					Type:   obj.TYPE_ADDR,
					Reg:    REGSP,
					Offset: int64(-autosize),
				}
				// STL RA, $0(SP)
				q = obj.Appendp(q, newprog)
				q.Spadj = autosize
				q.As = ASTL
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REGLINK
				q.To.Type = obj.TYPE_MEM
				q.To.Reg = REGSP

				//zxw new add
				q = c.ctxt.EndUnsafePoint(q, c.newprog, -1)
			}
			// snyh_TODO: #131
			q = insertLDGP(q, newprog, REGZERO)

			if c.cursym.Func().Text.From.Sym.Wrapper() {
				// if(g->panic != nil && g->panic->argp == FP)
				//    g->panic->argp = bottom-of-frame

				// LDL R1, g_panic(g)
				q = obj.Appendp(q, newprog)
				q.As = ALDL
				q.From = obj.Addr{Type: obj.TYPE_REG, Reg: REG_R1}
				q.To = obj.Addr{
					Type:   obj.TYPE_MEM,
					Reg:    REGG,
					Offset: 4 * int64(c.ctxt.Arch.PtrSize), // G.panic
				}

				// BEQ R1, end
				q = obj.Appendp(q, newprog)
				q.As = ABEQ
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R1
				q.To.Type = obj.TYPE_BRANCH
				p1 := q

				// LDL R2, panic_argp(R1)
				q = obj.Appendp(q, newprog)
				q.As = ALDL
				q.From = obj.Addr{Type: obj.TYPE_REG, Reg: REG_R2}
				q.To = obj.Addr{
					Type:   obj.TYPE_MEM,
					Reg:    REG_R1,
					Offset: 0, // Panic.argp
				}

				// ADDL SP, $autosize+FixedFrameSize, R3
				q = obj.Appendp(q, newprog)
				q.As = AADDL
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REGSP
				q.SetFrom3(obj.Addr{
					Type:   obj.TYPE_CONST,
					Offset: int64(autosize) + ctxt.FixedFrameSize(),
				})
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R3

				// CMPEQ R2, R3, R0
				q = obj.Appendp(q, newprog)
				q.As = ACMPEQ
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R2
				q.SetFrom3(obj.Addr{
					Type: obj.TYPE_REG,
					Reg:  REG_R3,
				})
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R0
				// BEQ R0, end
				q = obj.Appendp(q, newprog)
				q.As = ABEQ
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R0
				q.To.Type = obj.TYPE_BRANCH
				p2 := q

				// ADD SP, $FIXED_FRAME, R2
				q = obj.Appendp(q, newprog)
				q.As = AADDL
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REGSP
				q.SetFrom3(obj.Addr{
					Type:   obj.TYPE_CONST,
					Offset: ctxt.FixedFrameSize(),
				})
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R2

				// STL R2, panic_argp(R1)
				q = obj.Appendp(q, newprog)
				q.As = ASTL
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R2
				q.To.Type = obj.TYPE_MEM
				q.To.Reg = REG_R1
				q.To.Offset = 0 // Panic.argp

				// end:
				q = obj.Appendp(q, newprog)
				q.As = obj.ANOP
				p1.To.Val = q
				p2.To.Val = q
			}
			p = q

		case obj.ARET:
			retjmp = p.To.Sym
			if autosize != 0 && retjmp == nil {
				tmp := newprog()
				*tmp = *p

				// We can't use q = obj.Appendp(q, newprog),
				// because the RET may be a jump target, and Appendp
				// will change the pointer value of q
				// LDL RA, $0(SP)
				q := p
				q.As = ALDL
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REGLINK
				q.To.Type = obj.TYPE_MEM
				q.To.Reg = REGSP

				//LDI SP, $autosize(SP)
				q = obj.Appendp(q, newprog)
				q.As = ALDI
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REGSP
				q.To = obj.Addr{
					Type:   obj.TYPE_ADDR,
					Reg:    REGSP,
					Offset: int64(autosize),
				}
				q.Spadj = -autosize

				q = obj.Appendp(q, newprog)
				*q = *tmp

				//insert a noop for restore spadj
				q = obj.Appendp(q, newprog)
				q.As = obj.ANOP
				q.Spadj = autosize
				p = q

			}
			//zxw add
			if retjmp != nil {
				if autosize == 0 {
					tmp := newprog()
					*tmp = *p

					// We can't use q = obj.Appendp(q, newprog),
					// because the RET may be a jump target, and Appendp
					// will change the pointer value of q
					q := p
					q.As = ALDI
					q.From.Type = obj.TYPE_REG
					q.From.Reg = REGSP
					q.To = obj.Addr{
						Type:   obj.TYPE_ADDR,
						Reg:    REGSP,
						Offset: int64(autosize) - c.ctxt.FixedFrameSize(),
					}
					// STL RA, $0(SP)
					q = obj.Appendp(q, newprog)
					q.Spadj = autosize
					q.As = ASTL
					q.From.Type = obj.TYPE_REG
					q.From.Reg = REGLINK
					q.To.Type = obj.TYPE_MEM
					q.To.Reg = REGSP

					q = obj.Appendp(q, newprog)
					*q = *tmp
					//insert a noop for restore spadj
					q = obj.Appendp(q, newprog)
					q.As = obj.ANOP
					q.Spadj = autosize
					p = q
				}

				tmp1 := newprog()
				*tmp1 = *p

				q1 := p
				// snyh_TODO: #131
				q1 = insertLDGP(q1, newprog, REGZERO)

				//insert a noop for restore spadj
				q1 = obj.Appendp(q1, newprog)
				q1.As = obj.ANOP
				q1.Spadj = autosize

				// LDL RA, $0(SP)
				q1 = obj.Appendp(q1, newprog)
				q1.As = ALDL
				q1.From.Type = obj.TYPE_REG
				q1.From.Reg = REGLINK
				q1.To.Type = obj.TYPE_MEM
				q1.To.Reg = REGSP

				if autosize != 0 {
					//LDI SP, $autosize(SP)
					q1 = obj.Appendp(q1, newprog)
					q1.As = ALDI
					q1.From.Type = obj.TYPE_REG
					q1.From.Reg = REGSP
					q1.To = obj.Addr{
						Type:   obj.TYPE_ADDR,
						Reg:    REGSP,
						Offset: int64(autosize),
					}
					q1.Spadj = -autosize
				} else {
					//LDI SP, $8(SP)
					q1 = obj.Appendp(q1, newprog)
					q1.As = ALDI
					q1.From.Type = obj.TYPE_REG
					q1.From.Reg = REGSP
					q1.To = obj.Addr{
						Type:   obj.TYPE_ADDR,
						Reg:    REGSP,
						Offset: int64(autosize) + c.ctxt.FixedFrameSize(),
					}
					q1.Spadj = -autosize - int32(c.ctxt.FixedFrameSize())
				}

				//RET ZERO,RA
				q1 = obj.Appendp(q1, newprog)
				q1.As = ARET
				q1.From.Type = obj.TYPE_REG
				q1.From.Reg = REGZERO
				q1.To.Type = obj.TYPE_MEM
				q1.To.Reg = REGLINK

				p = q1
			}
			//zxw add
		case obj.AGETCALLERPC:
			if cursym.Leaf() {
				q := p
				/* LDI Rd, LR */
				q.As = ALDI
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REGLINK
				p = q
			} else {
				q := p
				/* LDL Rd, (RSP) */
				q.As = ALDL
				q.To.Type = obj.TYPE_MEM
				q.To.Reg = REGSP
				p = q
			}
		}
	}
	for p := cursym.Func().Text; p != nil; p = p.Link {
		switch p.As {
		case obj.ACALL, obj.AJMP:
			p = insertLDGP(p, newprog, REGZERO)
		}
	}

}

func progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	c := ctxt77{ctxt: ctxt, newprog: newprog}

	p.From.Class = 0
	p.To.Class = 0
	switch p.As {
	case ACALL:
		if p.From.Type == obj.TYPE_NONE {
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REGLINK
		}
	case AJMP:
		if p.To.Type == obj.TYPE_BRANCH {
			p.As = ABR
		}
		if p.From.Type == obj.TYPE_NONE {
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REGZERO
		}
	case ARET:
		if p.To.Type == obj.TYPE_NONE && p.From.Type == obj.TYPE_NONE {
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REGZERO
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = REGLINK
		}
	case ABR,
		ABSR:
		if p.From.Type == obj.TYPE_NONE {
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REGZERO
		}
	case ALDF:
		if p.To.Type == obj.TYPE_FCONST {
			f32 := float32(p.To.Val.(float64))
			p.As = AFLDS
			p.To.Type = obj.TYPE_MEM
			p.To.Sym = ctxt.Float32Sym(f32)
			p.To.Name = obj.NAME_EXTERN
			p.To.Offset = 0
		}
	case ALDI:
		if p.To.Type == obj.TYPE_FCONST {
			f64 := float64(p.To.Val.(float64))
			p.As = AFLDD
			p.To.Type = obj.TYPE_MEM
			p.To.Sym = ctxt.Float64Sym(f64)
			p.To.Name = obj.NAME_EXTERN
			p.To.Offset = 0
		} else if p.To.Type == obj.TYPE_CONST && !isint16(p.To.Offset) {
			p.As = ALDL
			p.To.Type = obj.TYPE_MEM
			p.To.Sym = ctxt.Int64Sym(p.To.Offset)
			p.To.Name = obj.NAME_EXTERN
			p.To.Offset = 0
		}
	/*case AIFMOVS,
	AIFMOVD,
	AFIMOVS,
	AFIMOVD:
	p.SetFrom3(obj.Addr{
		Type: obj.TYPE_REG,
		Reg:  REGZERO,
	})*/
	case AFCVTSD,
		AFCVTDS,
		AFCVTDL,
		AFCVTLS,
		AFCVTLD,
		AFCVTLW,
		AFCVTWL,
		AFSQRTS,
		AFSQRTD,
		ACTLZ,
		ACTTZ,
		ACTPOP,
		ASEXTB,
		ASEXTH,
		AFCVTDL_Z,
		AFCVTDL_P,
		AFCVTDL_G,
		AFCVTDL_N:
		p.SetFrom3(p.From)
		p.From.Reg = REGZERO
		p.From.Type = obj.TYPE_REG
	}

	// use in plugin mode or share mode
	if c.ctxt.Flag_dynlink {
		c.rewriteToUseGot(p)
	}
}

// Rewrite p, if necessary, to access global data via the global offset table.
func (c *ctxt77) rewriteToUseGot(p *obj.Prog) {
	if p.As == obj.ADUFFCOPY || p.As == obj.ADUFFZERO {
		//     ADUFFxxx $offset
		// becomes
		//     MOVD runtime.duffxxx@GOT, REGTMP
		//     LDI $offset, REGTMP
		//     LDI LR, REGTMP
		//     CALL LR
		var sym *obj.LSym
		if p.As == obj.ADUFFZERO {
			sym = c.ctxt.Lookup("runtime.duffzero")
		} else {
			sym = c.ctxt.Lookup("runtime.duffcopy")
		}
		offset := p.To.Offset
		p.As = AMOVD
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_GOTREF
		p.From.Sym = sym
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REGTMP
		p.To.Name = obj.NAME_NONE
		p.To.Offset = 0
		p.To.Sym = nil

		p1 := obj.Appendp(p, c.newprog)
		p1.As = ALDI
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = REGTMP
		p1.SetFrom3(obj.Addr{Type: obj.TYPE_CONST, Offset: offset})
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = REGTMP

		p2 := obj.Appendp(p, c.newprog)
		p2.As = ALDI
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = REG_R27
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = REGTMP

		p3 := obj.Appendp(p1, c.newprog)
		p3.As = obj.ACALL
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = REG_R27
	}

	// We only care about global data: NAME_EXTERN means a global
	// symbol in the Go sense, and p.Sym.Local is true for a few
	// internally defined symbols.
	if p.To.Type == obj.TYPE_ADDR && p.To.Name == obj.NAME_EXTERN && !p.To.Sym.Local() {
		// SYMADDR $sym, Rx becomes SYMADDR sym@GOT, Rx
		// SYMADDR $sym+<off>, Rx becomes SYMADDR sym@GOT, Rx; LDI <off>, Rx
		if p.As != ASYMADDR {
			c.ctxt.Diag("do not know how to handle symbol not symaddr in %v with -dynlink", p)
		}
		if p.From.Type != obj.TYPE_REG {
			c.ctxt.Diag("do not know how to handle symbol address insn to non-register in %v with -dynlink", p)
		}
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_GOTREF
		if p.From.Offset != 0 {
			q := obj.Appendp(p, c.newprog)
			q.As = ALDI
			q.From = p.To
			p.SetFrom3(obj.Addr{Type: obj.TYPE_CONST, Offset: p.From.Offset}) //zxw new change
			q.To = p.To
		}
	}
	//	if p.GetFrom3() != nil && p.GetFrom3().Name == obj.NAME_EXTERN {
	//		c.ctxt.Diag("don't know how to handle %v with -dynlink", p)
	//	}

	var source *obj.Addr
	// LDx Ry, sym becomes LDL REGTMP, sym@GOT; LDx Ry, (REGTMP)
	// STx Ry, sym becomes LDL REGTMP, sym@GOT; STx Ry, (REGTMP)
	// An addition may be inserted between the two MOVs if there is an offset.
	if p.From.Name == obj.NAME_EXTERN && !p.From.Sym.Local() {
		if p.To.Name == obj.NAME_EXTERN && !p.To.Sym.Local() {
			c.ctxt.Diag("cannot handle NAME_EXTERN on both sides in %v with -dynlink", p)
		}
		source = &p.From
	} else if p.To.Name == obj.NAME_EXTERN && !p.To.Sym.Local() {
		source = &p.To
	} else {
		return
	}
	if p.As == obj.ATEXT || p.As == obj.AFUNCDATA || p.As == obj.ACALL || p.As == obj.ARET || p.As == obj.AJMP {
		return
	}
	if source.Sym.Type == objabi.STLSBSS {
		return
	}
	if source.Type != obj.TYPE_MEM {
		c.ctxt.Diag("don't know how to handle %v with -dynlink", p)
	}
	p1 := obj.Appendp(p, c.newprog)
	p2 := obj.Appendp(p1, c.newprog)
	p1.As = AMOVD
	p1.From.Type = obj.TYPE_MEM
	p1.From.Sym = source.Sym
	p1.From.Name = obj.NAME_GOTREF
	p1.To.Type = obj.TYPE_REG
	p1.To.Reg = REGTMP

	var final obj.As
	if p.As == ASYMADDR {
		final = ALDI
	} else {
		final = p.As
	}
	p2.As = final
	p2.From = p.From
	p2.To = p.To
	if p.From.Name == obj.NAME_EXTERN {
		p2.From.Reg = REGTMP
		p2.From.Name = obj.NAME_NONE
		p2.From.Sym = nil
	} else if p.To.Name == obj.NAME_EXTERN {
		p2.To.Reg = REGTMP
		p2.To.Name = obj.NAME_NONE
		p2.To.Sym = nil
	} else {
		return
	}
	obj.Nopout(p)
}

var LinkSW64 = obj.LinkArch{
	Arch:           sys.ArchSW64,
	Init:           buildop,
	Preprocess:     preprocess,
	Assemble:       span77,
	Progedit:       progedit,
	DWARFRegisters: SW64DWARFRegisters,
}

//单参数指令，且参数为目的数 ==>(Prog.To = a[0])
//不在这个列表的则参数为源数 ==>(Prog.From = a[0])
var unaryDst = map[obj.As]bool{}
