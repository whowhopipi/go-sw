// cmd/9l/optab.c, cmd/9l/asmout.c from Vita Nuova.
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2008 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2008 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sw64

import (
	"cmd/internal/obj"
	"cmd/internal/objabi"
	"fmt"
	"log"
	"math"
	"sort"
)

// ctxt77 holds state while assembling a single function.
// Each function gets a fresh ctxt77.
// This allows for multiple functions to be safely concurrently assembled.
type ctxt77 struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	autosize   int32
	instoffset int64
	pc         int64
}

const (
	funcAlign = 16
)

type Optab struct {
	as    obj.As
	a1    uint8
	a2    uint8
	a3    uint8
	a4    uint8
	type_ int8
	size  int8
	param int16
	flag  uint8
}

const (
	// Optab.flag
	NOTUSETMP = 1 << iota // p expands to multiple instructions, but does NOT use REGTMP
)

var optab = []Optab{
	{obj.ANOP, C_NONE, C_NONE, C_NONE, C_NONE, 0, 0, 0, 0},
	{obj.AFUNCDATA, C_SUCON, C_NONE, C_NONE, C_LEXT, 0, 0, 0, 0},
	{obj.APCDATA, C_SUCON, C_NONE, C_NONE, C_LCON, 0, 0, 0, 0},
	{obj.ARET, C_RREG, C_NONE, C_NONE, C_SOREG, 2, 4, 0, 0},
	//zxw add
	{obj.ARET, C_NONE, C_NONE, C_NONE, C_LEXT, 24, 12, REG_R29, NOTUSETMP},
	{AMEMB, C_NONE, C_NONE, C_NONE, C_NONE, 4, 4, 0, 0},
	{ARD_F, C_RREG, C_NONE, C_NONE, C_NONE, 4, 4, 0, 0},
	{ASYS_CALL, C_SUCON, C_NONE, C_NONE, C_NONE, 1, 4, 0, 0},
	{AEXCB, C_RREG, C_NONE, C_NONE, C_SAUTO, 2, 4, 0, 0},

	{ARFPCR, C_FREG, C_NONE, C_NONE, C_NONE, 6, 4, 0, 0},
	{ASETFPEC0, C_NONE, C_NONE, C_NONE, C_NONE, 6, 4, 0, 0},

	{AADDW, C_RREG, C_RREG, C_NONE, C_RREG, 6, 4, 0, 0},
	{AADDW, C_RREG, C_SUCON, C_NONE, C_RREG, 7, 4, 0, 0},
	{AADDW, C_RREG, C_MCON, C_NONE, C_RREG, 14, 8, 0, 0},
	{AADDW, C_RREG, C_LCON, C_NONE, C_RREG, 16, 16, 0, 0},

	{ACTPOP, C_RREG, C_RREG, C_NONE, C_RREG, 6, 4, 0, 0},
	{AIFMOVS, C_RREG, C_NONE, C_NONE, C_FREG, 6, 4, 0, 0},
	{AFIMOVS, C_FREG, C_NONE, C_NONE, C_RREG, 6, 4, 0, 0},
	{AFADDS, C_FREG, C_FREG, C_NONE, C_FREG, 6, 4, 0, 0},
	{AFSQRTS, C_RREG, C_FREG, C_NONE, C_FREG, 6, 4, 0, 0},

	{ALDBU, C_RREG, C_NONE, C_NONE, C_SAUTO, 2, 4, REG_R30, 0},
	{ALDBU, C_RREG, C_NONE, C_NONE, C_LAUTO, 17, 20, REG_R30, 0},
	{ALDBU, C_RREG, C_NONE, C_NONE, C_SOREG, 2, 4, 0, 0},
	{ALDBU, C_RREG, C_NONE, C_NONE, C_LOREG, 17, 20, 0, 0},
	{ALDBU, C_RREG, C_NONE, C_NONE, C_LEXT, 12, 8, REG_R29, NOTUSETMP},
	{ALDBU, C_RREG, C_NONE, C_NONE, C_TLS_LE, 22, 12, 0, NOTUSETMP},
	{ALDBU, C_RREG, C_NONE, C_NONE, C_TLS_IE, 23, 16, 0, 0},
	{ASTB, C_RREG, C_NONE, C_NONE, C_SAUTO, 2, 4, REG_R30, 0},
	{ASTB, C_RREG, C_NONE, C_NONE, C_LAUTO, 17, 20, REG_R30, 0},
	{ASTB, C_RREG, C_NONE, C_NONE, C_SOREG, 2, 4, 0, 0},
	{ASTB, C_RREG, C_NONE, C_NONE, C_LOREG, 17, 20, 0, 0},
	{ASTB, C_RREG, C_NONE, C_NONE, C_LEXT, 13, 8, REG_R29, 0},
	{ASTB, C_RREG, C_NONE, C_NONE, C_TLS_LE, 22, 12, 0, NOTUSETMP},
	{ASTB, C_RREG, C_NONE, C_NONE, C_TLS_IE, 23, 16, 0, 0},

	{AFLDD, C_FREG, C_NONE, C_NONE, C_SAUTO, 2, 4, REG_R30, 0},
	{AFLDD, C_FREG, C_NONE, C_NONE, C_LAUTO, 17, 20, REG_R30, 0},
	{AFLDD, C_FREG, C_NONE, C_NONE, C_SOREG, 2, 4, 0, 0},
	{AFLDD, C_FREG, C_NONE, C_NONE, C_LOREG, 17, 20, 0, 0},
	{AFLDD, C_FREG, C_NONE, C_NONE, C_LEXT, 15, 12, REG_R29, 0},

	{AFSTD, C_FREG, C_NONE, C_NONE, C_SAUTO, 2, 4, REG_R30, 0},
	{AFSTD, C_FREG, C_NONE, C_NONE, C_LAUTO, 17, 20, REG_R30, 0},
	{AFSTD, C_FREG, C_NONE, C_NONE, C_SOREG, 2, 4, 0, 0},
	{AFSTD, C_FREG, C_NONE, C_NONE, C_LOREG, 17, 20, 0, 0},
	{AFSTD, C_FREG, C_NONE, C_NONE, C_LEXT, 13, 8, REG_R29, 0},

	{ALDI, C_RREG, C_NONE, C_NONE, C_MCON, 2, 4, 0, 0},
	{ALDI, C_RREG, C_NONE, C_NONE, C_SAUTO, 2, 4, REG_R30, 0},
	{ALDI, C_RREG, C_NONE, C_NONE, C_LAUTO, 17, 20, REG_R30, 0},
	{ALDI, C_RREG, C_NONE, C_NONE, C_RREG, 2, 4, 0, 0},
	{ALDI, C_RREG, C_NONE, C_NONE, C_SOREG, 2, 4, 0, 0},
	{ALDI, C_RREG, C_NONE, C_NONE, C_LOREG, 17, 20, 0, 0},

	{ASYMADDR, C_RREG, C_NONE, C_NONE, C_LEXT, 12, 8, REG_R29, NOTUSETMP},
	{ALLDW, C_RREG, C_NONE, C_NONE, C_SOREG, 5, 4, 0, 0},
	{ABR, C_RREG, C_NONE, C_NONE, C_SBRA, 3, 4, 0, 0},
	{ABR, C_RREG, C_NONE, C_NONE, C_LEXT, 10, 4, REG_R29, 0},
	{AFBGE, C_FREG, C_NONE, C_NONE, C_SBRA, 3, 4, 0, 0},
	{AFBGE, C_FREG, C_NONE, C_NONE, C_LEXT, 10, 4, REG_R29, 0},
	{obj.AJMP, C_RREG, C_NONE, C_NONE, C_LEXT, 9, 12, REG_R29, NOTUSETMP},
	{obj.AJMP, C_RREG, C_NONE, C_NONE, C_SOREG, 21, 4, 0, 0},

	{AFMAS, C_FREG, C_FREG, C_FREG, C_FREG, 8, 4, 0, 0},
	{ASELEQ, C_RREG, C_RREG, C_RREG, C_RREG, 8, 4, 0, 0},

	{ALDGP, C_RREG, C_NONE, C_NONE, C_RREG, 20, 16, 0, 0},
	{ANOOP, C_NONE, C_NONE, C_NONE, C_NONE, 11, 4, 0, 0},
	{obj.AUNDEF, C_NONE, C_NONE, C_NONE, C_NONE, 18, 4, 0, 0},
	{AWORD, C_LCON, C_NONE, C_NONE, C_NONE, 19, 4, 0, 0},
	{obj.AXXX, C_NONE, C_NONE, C_NONE, C_NONE, 0, 4, 0, 0},
}

var oprange [ALAST & obj.AMask][]Optab

var xcmp [C_NCLASS][C_NCLASS]bool

func span77(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	if ctxt.Retpoline {
		ctxt.Diag("-spectre=ret not supported on mips")
		ctxt.Retpoline = false // don't keep printing
	}
	p := cursym.Func().Text
	if p == nil || p.Link == nil { // handle external functions and ELF section symbols
		return
	}

	c := ctxt77{ctxt: ctxt, newprog: newprog, cursym: cursym, autosize: int32(p.To.Offset + ctxt.FixedFrameSize())}

	if oprange[AADDL&obj.AMask] == nil {
		c.ctxt.Diag("sw64 ops not initialized, call sw64.buildop first")

	}

	pc := int64(0)
	p.Pc = pc

	var m int
	var o *Optab
	for p = p.Link; p != nil; p = p.Link {
		p.Pc = pc
		o = c.oplook(p)
		m = int(o.size)
		if m == 0 {
			if p.As != obj.ANOP && p.As != obj.AFUNCDATA && p.As != obj.APCDATA {
				c.ctxt.Diag("zero-width instruction\n%v", p)
			}
			continue
		}

		if (p.As == ALSTW || p.As == ALSTL) && p.Pc&0x7 != 0 {
			m += 4
		}
		pc += int64(m)
	}

	c.cursym.Size = pc

	/*
	 * if any procedure is large enough to
	 * generate a large SBRA branch, then
	 * generate extra passes putting branches
	 * around jmps to fix. this is rare.
	 */
	/*	bflag := 1

		var otxt int64
		var q *obj.Prog
		for bflag != 0 {
			bflag = 0
			pc = 0
			for p = c.cursym.Func().Text.Link; p != nil; p = p.Link {
				p.Pc = pc
				o = c.oplook(p)
				// very large conditional branches
				if (o.type_ == 3 || o.type_ == 10) && p.To.Target() != nil {
					otxt = p.To.Target().Pc - pc
					if otxt < -(1<<17)+10 || otxt >= (1<<17)-10 {
						q = c.newprog()
						q.Link = p.Link
						p.Link = q
						q.As = ABR
						q.To.Type = obj.TYPE_BRANCH
						q.To.SetTarget(p.To.Target())
					 	p.To.SetTarget(q)
						q = c.newprog()
						q.Link = p.Link
						p.Link = q
						q.As = ABR
						q.To.Type = obj.TYPE_BRANCH
						q.To.SetTarget(q.Link.Link)
						c.addnop(p.Link)
						c.addnop(p)
						bflag = 1
					}
				}
				m = int(o.size)
				if m == 0 {
					if p.As != obj.ANOP && p.As != obj.AFUNCDATA && p.As != obj.APCDATA {
						c.ctxt.Diag("zero-width instruction\n%v", p)
					}
					continue
				}

				pc += int64(m)
			}

			c.cursym.Size = pc
		}

		pc += -pc & (funcAlign - 1)
		c.cursym.Size = pc
	*/
	/*
	 * lay out the code, emitting code and data relocations.
	 */
	c.cursym.Grow(c.cursym.Size)
	bp := c.cursym.P
	var i int32
	var out [5]uint32
	for p := c.cursym.Func().Text.Link; p != nil; p = p.Link {
		c.pc = p.Pc
		o = c.oplook(p)
		if int(o.size) > 4*len(out) {
			log.Fatalf("out array in span77 is too small, need at least %d for %v", o.size/4, p)
		}
		c.asmout(p, o, out[:])
		n := o.size
		if (p.As == ALSTW || p.As == ALSTL) && p.Pc&0x7 != 0 {
			n += 4
		}
		for i = 0; i < int32(n/4); i++ {
			c.ctxt.Arch.ByteOrder.PutUint32(bp, out[i])
			bp = bp[4:]
		}
	}

	//zxw new add
	// Mark nonpreemptible instruction sequences.
	// We use REGTMP as a scratch register during call injection,
	// so instruction sequences that use REGTMP are unsafe to
	// preempt asynchronously.
	obj.MarkUnsafePoints(c.ctxt, c.cursym.Func().Text, c.newprog, c.isUnsafePoint, c.isRestartable)

	verifyLock(c)
}

// isUnsafePoint returns whether p is an unsafe point.
func (c *ctxt77) isUnsafePoint(p *obj.Prog) bool {
	// If p explicitly uses REGTMP, it's unsafe to preempt, because the
	// preemption sequence clobbers REGTMP.
	return p.From.Reg == REGTMP || p.To.Reg == REGTMP || p.Reg == REGTMP
}

// isRestartable returns whether p is a multi-instruction sequence that,
// if preempted, can be restarted.
func (c *ctxt77) isRestartable(p *obj.Prog) bool {
	if c.isUnsafePoint(p) {
		return false
	}
	// If p is a multi-instruction sequence with uses REGTMP inserted by
	// the assembler in order to materialize a large constant/offset, we
	// can restart p (at the start of the instruction sequence), recompute
	// the content of REGTMP, upon async preemption. Currently, all cases
	// of assembler-inserted REGTMP fall into this category.
	// If p doesn't use REGTMP, it can be simply preempted, so we don't
	// mark it.
	o := c.oplook(p)
	return o.size > 4 && o.flag&NOTUSETMP == 0
}

func verifyLock(c ctxt77) {
	for p := c.cursym.Func().Text.Link; p != nil; p = p.Link {
		// 判断锁装入指令与写锁标志指令是否成对出现
		if p.As == ALLDW || p.As == ALLDL {
			for q := p; q != nil; q = q.Link {
				if q.As != AWR_F {
					if q.As == ALSTW || q.As == ALSTL {
						c.ctxt.Diag("missing WR_F after %s\n", p.As)
						return
					} else {
						if (q.Pc + 4) == c.cursym.Size {
							c.ctxt.Diag("Atomic operation is not logical!\n")
							return
						}
						continue
					}
				} else {
					break
				}
			}
		}
		if p.As == ALSTW || p.As == ALSTL {
			// 判断锁存储指令与读锁标志指令是否成对出现
			if p.Link.As != ARD_F {
				c.ctxt.Diag("missing RD_F after %s\n", p.As)
				return
			}
		}
	}

}

func IsReg(r int) bool {
	return REG_R0 <= r && r <= REG_F31
}
func IsFReg(r int) bool {
	return REG_F0 <= r && r <= REG_F31
}
func IsRReg(r int) bool {
	return REG_R0 <= r && r <= REG_R31
}
func isint16(v int64) bool {
	return int64(int16(v)) == v
}
func isuint8(v uint64) bool {
	return uint64(uint8(v)) == v
}

func isint32(v int64) bool {
	return int64(int32(v)) == v
}

func isuint32(v uint64) bool {
	return uint64(uint32(v)) == v
}

func (c *ctxt77) aclass(a *obj.Addr) int {
	if a.Sym != nil { // use relocation
		if a.Sym.Type == objabi.STLSBSS {
			if c.ctxt.Flag_shared {
				return C_TLS_IE
			} else {
				return C_TLS_LE
			}
		}
	}
	switch a.Type {
	case obj.TYPE_NONE:
		return C_NONE

	case obj.TYPE_REG:
		if IsRReg(int(a.Reg)) {
			return C_RREG
		}

		if IsFReg(int(a.Reg)) {
			return C_FREG
		}

	case obj.TYPE_CONST:
		if isuint8(uint64(a.Offset)) {
			return C_SUCON
		}
		if isint16(a.Offset) {
			return C_MCON
		}
		return C_LCON

	case obj.TYPE_FCONST:
		return C_FCON

	case obj.TYPE_MEM,
		obj.TYPE_ADDR:
		switch a.Name {
		case obj.NAME_EXTERN, obj.NAME_STATIC:
			if a.Sym.Type == objabi.STLSBSS {
				c.ctxt.Diag("taking address of TLS variable is not supported")
			}
			return C_LEXT
		case obj.NAME_PARAM, obj.NAME_AUTO:
			if isint16(a.Offset) {
				return C_SAUTO
			}
			if a.Offset > math.MaxInt16 || -a.Offset > math.MaxInt16 {
				return C_LAUTO
			}
		case obj.NAME_NONE:
			if isint16(a.Offset) {
				return C_SOREG
			}
			if a.Offset > math.MaxInt16 || -a.Offset > math.MaxInt16 {
				return C_LOREG
			}
		}
		return C_ADDR

	case obj.TYPE_BRANCH:
		switch a.Name {
		case obj.NAME_EXTERN, obj.NAME_STATIC:
			return C_LEXT
		}
		return C_SBRA
	case obj.TYPE_INDIR:
		switch a.Name {
		case obj.NAME_EXTERN, obj.NAME_STATIC:
			return C_LEXT
		}
	}
	return C_GOK
}

func prasm(p *obj.Prog) {
	fmt.Printf("%v\n", p)
}

func (c *ctxt77) oplook(p *obj.Prog) *Optab {
	if oprange[AADDL&obj.AMask] == nil {
		c.ctxt.Diag("sw64 ops not initialized, call sw64.buildop first")
	}
	a1 := int(p.Optab)
	if a1 != 0 {
		return &optab[a1-1]
	}
	a1 = int(p.From.Class)
	if a1 == 0 {
		a1 = c.aclass(&p.From) + 1
		p.From.Class = int8(a1)
	}

	a1--

	a3 := C_NONE
	if len(p.RestArgs) == 2 {
		a3 = int(p.RestArgs[1].Addr.Class)
		if a3 == 0 {
			a3 = c.aclass(&p.RestArgs[1].Addr) + 1
			p.GetFrom3().Class = int8(a3)
		}
		a3--

	}

	a2 := C_NONE
	if p.Reg != 0 {
		a2 = C_RREG
		if IsFReg(int(p.Reg)) {
			a2 = C_FREG
		}
	}
	if (p.Reg == 0) && (len(p.RestArgs) != 0) {
		a2 = int(p.RestArgs[0].Addr.Class)
		if a2 == 0 {
			a2 = c.aclass(&p.RestArgs[0].Addr) + 1
			p.RestArgs[0].Addr.Class = int8(a2)
		}
		a2--
	}
	a4 := int(p.To.Class)
	if a4 == 0 {
		a4 = c.aclass(&p.To) + 1
		p.To.Class = int8(a4)
	}

	a4--

	ops := oprange[p.As&obj.AMask]
	c1 := &xcmp[a1]
	c2 := &xcmp[a2]
	c3 := &xcmp[a3]
	c4 := &xcmp[a4]
	for i := range ops {
		op := &ops[i]
		if c1[op.a1] && c2[op.a2] && c3[op.a3] && c4[op.a4] {
			p.Optab = uint16(cap(optab) - cap(ops) + i + 1)
			return op
		}
	}

	c.ctxt.Diag("illegal combination %v %v %v %v %v", p.As, DRconv(a1), DRconv(a2), DRconv(a3), DRconv(a4))
	prasm(p)
	if ops == nil {
		ops = optab
	}
	return &ops[0]
}

func cmp(a int, b int) bool {
	if a == b {
		return true
	}
	switch a {
	case C_LCON:
		if b == C_SUCON || b == C_MCON {
			return true
		}

	case C_MCON:
		if b == C_SUCON {
			return true
		}

	case C_LAUTO:
		if b == C_SAUTO {
			return true
		}

	case C_LOREG:
		if b == C_ZOREG || b == C_SOREG {
			return true
		}

	case C_SOREG:
		if b == C_ZOREG {
			return true
		}

	}

	return false
}

type ocmp []Optab

func (x ocmp) Len() int {
	return len(x)
}

func (x ocmp) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x ocmp) Less(i, j int) bool {
	p1 := &x[i]
	p2 := &x[j]
	n := int(p1.as) - int(p2.as)
	if n != 0 {
		return n < 0
	}
	n = int(p1.a1) - int(p2.a1)
	if n != 0 {
		return n < 0
	}
	n = int(p1.a2) - int(p2.a2)
	if n != 0 {
		return n < 0
	}
	n = int(p1.a3) - int(p2.a3)
	if n != 0 {
		return n < 0
	}
	n = int(p1.a4) - int(p2.a4)
	if n != 0 {
		return n < 0
	}
	return false
}

func opset(a, b0 obj.As) {
	oprange[a&obj.AMask] = oprange[b0]
}

func buildop(ctxt *obj.Link) {
	if oprange[AADDL&obj.AMask] != nil {
		// Already initialized; stop now.
		// This happens in the cmd/asm tests,
		// each of which re-initializes the arch.
		return
	}

	var n int

	for i := 0; i < C_NCLASS; i++ {
		for n = 0; n < C_NCLASS; n++ {
			if cmp(n, i) {
				xcmp[i][n] = true
			}
		}
	}
	for n = 0; optab[n].as != obj.AXXX; n++ {
	}
	sort.Sort(ocmp(optab[:n]))
	for i := 0; i < n; i++ {
		r := optab[i].as
		r0 := r & obj.AMask
		start := i
		for optab[i].as == r {
			i++
		}
		oprange[r0] = optab[start:i]
		i--

		switch r {
		default:
			ctxt.Diag("unknown op in build: %v", r)
			ctxt.DiagFlush()
			log.Fatalf("bad code")
		case ASYS_CALL:
			opset(ASYS_CALL_B, r0)
		case AEXCB:
			opset(APRI_LD, r0)
			opset(APRI_ST, r0)
		case AADDW:
			opset(AADDL, r0)
			opset(ASUBL, r0)
			opset(ASUBW, r0)
			opset(AMULW, r0)
			opset(AMULL, r0)
			opset(AUMULH, r0)
			opset(ACMPEQ, r0)
			opset(ACMPLT, r0)
			opset(ACMPLE, r0)
			opset(ACMPULT, r0)
			opset(ACMPULE, r0)
			opset(AAND, r0)
			opset(ABIS, r0)
			opset(AXOR, r0)
			opset(AORNOT, r0)
			opset(ASLL, r0)
			opset(ASRL, r0)
			opset(ASRA, r0)
			opset(ASEXTB, r0)
			opset(ASEXTH, r0)
			opset(AS4ADDW, r0)
			opset(AS4SUBW, r0)
			opset(AS8ADDW, r0)
			opset(AS8SUBW, r0)
			opset(AS4ADDL, r0)
			opset(AS4SUBL, r0)
			opset(AS8ADDL, r0)
			opset(AS8SUBL, r0)
			opset(ABIC, r0)
			opset(AEQV, r0)
			opset(AINSLB, r0)
			opset(AINSLH, r0)
			opset(AINSLW, r0)
			opset(AINSLL, r0)
			opset(AINSHB, r0)
			opset(AINSHH, r0)
			opset(AINSHW, r0)
			opset(AINSHL, r0)
			opset(AEXTLB, r0)
			opset(AEXTLH, r0)
			opset(AEXTLW, r0)
			opset(AEXTLL, r0)
			opset(AEXTHB, r0)
			opset(AEXTHH, r0)
			opset(AEXTHW, r0)
			opset(AEXTHL, r0)
			opset(AMASKLB, r0)
			opset(AMASKLH, r0)
			opset(AMASKLW, r0)
			opset(AMASKLL, r0)
			opset(AMASKHB, r0)
			opset(AMASKHH, r0)
			opset(AMASKHW, r0)
			opset(AMASKHL, r0)
			opset(ACMPGEB, r0)
			opset(AZAP, r0)
			opset(AZAPNOT, r0)
		case ABR:
			opset(ABSR, r0)
			opset(ABEQ, r0)
			opset(ABGE, r0)
			opset(ABGT, r0)
			opset(ABLE, r0)
			opset(ABLT, r0)
			opset(ABNE, r0)
			opset(ABLBC, r0)
			opset(ABLBS, r0)
		case AFBGE:
			opset(AFBGT, r0)
			opset(AFBEQ, r0)
			opset(AFBLE, r0)
			opset(AFBLT, r0)
			opset(AFBNE, r0)
		case AMEMB:
			opset(AIMEMB, r0)
		case ARD_F:
			opset(AWR_F, r0)
		case ALLDW:
			opset(ALLDL, r0)
			opset(ALSTW, r0)
			opset(ALSTL, r0)
		case AFMAS:
			opset(AFMAD, r0)
			opset(AFMSS, r0)
			opset(AFMSD, r0)
			opset(AFNMAS, r0)
			opset(AFNMAD, r0)
			opset(AFNMSS, r0)
			opset(AFNMSD, r0)
			opset(AFSELEQ, r0)
			opset(AFSELNE, r0)
			opset(AFSELLT, r0)
			opset(AFSELLE, r0)
			opset(AFSELGT, r0)
			opset(AFSELGE, r0)
		case ASELEQ:
			opset(ASELNE, r0)
			opset(ASELGE, r0)
			opset(ASELGT, r0)
			opset(ASELLE, r0)
			opset(ASELLT, r0)
			opset(ASELLBC, r0)
			opset(ASELLBS, r0)
		case ALDBU:
			opset(ALDHU, r0)
			opset(ALDW, r0)
			opset(ALDL, r0)
			opset(ALDL_U, r0)
		case ASTB:
			opset(ASTH, r0)
			opset(ASTW, r0)
			opset(ASTL, r0)
			opset(ASTL_U, r0)
		case ACTPOP:
			opset(ACTLZ, r0)
			opset(ACTTZ, r0)
		case AFIMOVS:
			opset(AFIMOVD, r0)
		case AIFMOVS:
			opset(AIFMOVD, r0)
		case AFADDS:
			opset(AFADDD, r0)
			opset(AFSUBS, r0)
			opset(AFSUBD, r0)
			opset(AFMULS, r0)
			opset(AFMULD, r0)
			opset(AFDIVS, r0)
			opset(AFDIVD, r0)
			opset(AFCMPEQ, r0)
			opset(AFCMPLE, r0)
			opset(AFCMPLT, r0)
			opset(AFCMPUN, r0)
			opset(AFCPYS, r0)
			opset(AFCPYSE, r0)
			opset(AFCPYSN, r0)
		case AFSQRTS:
			opset(AFSQRTD, r0)
			opset(AFCVTSD, r0)
			opset(AFCVTDS, r0)
			opset(AFCVTDL_G, r0)
			opset(AFCVTDL_P, r0)
			opset(AFCVTDL_Z, r0)
			opset(AFCVTDL_N, r0)
			opset(AFCVTDL, r0)
			opset(AFCVTWL, r0)
			opset(AFCVTLW, r0)
			opset(AFCVTLS, r0)
			opset(AFCVTLD, r0)
		case ARFPCR:
			opset(AWFPCR, r0)
		case ASETFPEC0:
			opset(ASETFPEC1, r0)
			opset(ASETFPEC2, r0)
			opset(ASETFPEC3, r0)
		case AFLDD:
			opset(AFLDS, r0)
		case AFSTD:
			opset(AFSTS, r0)
		case ALDI:
			opset(ALDIH, r0)
		case obj.AJMP:
			opset(obj.ACALL, r0)
		case ASYMADDR,
			ALDGP,
			obj.ANOP,
			obj.AFUNCDATA,
			obj.APCDATA,
			obj.AUNDEF,
			obj.ARET,
			ANOOP,
			AWORD:
			break
		}

	}

}

//操作码
func SP(x uint32, y uint32) uint32 {
	return x<<30 | y<<26
}

//带功能码的指令
func FMM(x uint32, y uint32) uint32 {
	return SP(0, 6) | x<<8 | y<<0
}

func FA(x uint32, y uint32) uint32 {
	return SP(1, 0) | x<<9 | y<<5
}

func FAF(x uint32, y uint32) uint32 {
	return SP(1, 8) | x<<9 | y<<5
}

func FAI(x uint32, y uint32) uint32 {
	return SP(1, 2) | x<<9 | y<<5
}

func FCA(x uint32, y uint32) uint32 {
	return SP(1, 9) | x<<14 | y<<10
}

//系统调用指令
func OP_SYSCALL(op uint32, fn uint32) uint32 {
	return op | fn
}

//转移控制指令
func OP_CONTROL(op uint32, ra int16, disp int32) uint32 {
	return op | uint32(ra)<<21 | uint32(disp)&(1<<21-1)
}

//存储器指令
func OP_MEMORY(op uint32, ra int16, rb int16, disp int16) uint32 {
	return op | uint32(ra)<<21 | uint32(rb)<<16 | uint32(uint16(disp))
}

//杂项指令
func OP_MISI_MEMORY(op uint32, ra int16, rb int16) uint32 {
	return op | uint32(ra)<<21 | uint32(rb)<<16
}

//带功能域的存储器指令
func OP_FUNC_MEMORY(op uint32, ra int16, rb int16, disp int16) uint32 {
	return op | uint32(ra)<<21 | uint32(rb)<<16 | uint32(uint16(disp))
}

//简单运算指令寄存器格式
func OP_ARITHMETIC(op uint32, ra int16, rb int16, rc int16) uint32 {
	return op | uint32(ra)<<21 | uint32(rb)<<16 | uint32(rc)
}

//简单运算指令立即数格式
func OP_ARITHMETIC_I(op uint32, ra int16, ib int16, rc int16) uint32 {
	return op | uint32(ra)<<21 | uint32(ib)<<13 | uint32(rc)
}

//浮点复合运算指令
func OP_COMPLEX_ARITHMETIC(op uint32, ra int16, rb int16, rc int16, rd int16) uint32 {
	return op | uint32(ra)<<21 | uint32(rb)<<16 | uint32(rc)<<5 | uint32(rd)
}

func (c *ctxt77) asmout(p *obj.Prog, o *Optab, out []uint32) {
	o1 := uint32(0)
	o2 := uint32(0)
	o3 := uint32(0)
	o4 := uint32(0)
	o5 := uint32(0)

	switch o.type_ {
	default:
		c.ctxt.Diag("unknown type %d %v", o.type_)
		prasm(p)

	case 0: /* buildNull */
		break
	case 1: /*OPC_SYSCALL*/
		o1 = OP_SYSCALL(c.oprrr(p.As), uint32(p.From.Offset))
	case 2: /*OPC_MEMORY*/
		disp := p.To.Offset

		if (disp > 0 && disp&0xffff != disp) ||
			(disp < 0 && -disp&0xffff != -disp) {
			c.ctxt.Diag("Memory instruction with immediate(0x%x) bigger than 0xffff at\n%v\n", disp, p)
		}

		o1 = OP_MEMORY(c.oprrr(p.As), getRegister(p.From.Reg), getRegister(p.To.Reg), int16(disp))
	case 3: /*OPC_CONTROL*/
		if p.To.Type == obj.TYPE_BRANCH {
			target := p.To.Val.(*obj.Prog)
			offset := (target.Pc - p.Pc - 4)
			//offset := (p.To.Target().Pc - p.Pc - 4)
			o1 = OP_CONTROL(c.oprrr(p.As), getRegister(p.From.Reg), int32(offset)/4)
		} else {
			o1 = OP_CONTROL(c.oprrr(p.As), getRegister(p.From.Reg), int32(p.To.Offset)/4)
		}
	case 4: /*OPC_MISI_MEMORY*/
		o1 = OP_MISI_MEMORY(c.oprrr(p.As), getRegister(p.From.Reg), getRegister(p.To.Reg))
	case 5: /*OPC_FUNC_MEMORY*/
		disp := p.To.Offset

		if (disp >= 0x1000) || (disp <= -0x1000) {
			c.ctxt.Diag("Memory instruction with immediate(0x%x) bigger than 0x1000 at\n%v\n", disp, p)
		}
		if (p.As == ALSTW || p.As == ALSTL) && p.Pc&0x7 != 0 {
			o1 = OP_MEMORY(c.oprrr(ALDIH), REGZERO, REGSP, 0)
			o2 = OP_FUNC_MEMORY(c.oprrr(p.As), getRegister(p.From.Reg), getRegister(p.To.Reg), int16(disp))
		} else {
			o1 = OP_FUNC_MEMORY(c.oprrr(p.As), getRegister(p.From.Reg), getRegister(p.To.Reg), int16(disp))
		}
	case 6: /*OPC_ARITHMETIC*/
		o1 = OP_ARITHMETIC(c.oprrr(p.As), getRegister(p.From.Reg), getReg2(p), getRegister(p.To.Reg))
	case 7: /*OPC_ARITHMETIC_I*/
		ib := p.RestArgs[0].Offset

		if ib > 255 {
			c.ctxt.Diag("Arithmetic instruction with immediate(0x%x) bigger than 0xff at\n%v\n", ib, p)
		}

		o1 = OP_ARITHMETIC_I(c.opirr(p.As), getRegister(p.From.Reg), int16(ib), getRegister(p.To.Reg))
	case 8: /*OPC_COMPLEX_ARITHMETIC*/
		o1 = OP_COMPLEX_ARITHMETIC(c.oprrr(p.As), getRegister(p.From.Reg), getReg2(p), getFrom3Reg(p), getRegister(p.To.Reg))

	case 9: /*buildCall,obj.AJMP, obj.ACALL*/
		sym := p.To.Sym
		r := getRegister(p.From.Reg)
		pv := REG_R27 & int16(31)
		o1 = OP_MEMORY(c.oprrr(ALDIH), pv, REGSB&31, 0) // ldih R27, sym(GP)
		addrel(c.cursym, int32(p.Pc), p.To.Offset, sym, objabi.R_SW64_GPRELHIGH)
		o2 = OP_MEMORY(c.oprrr(ALDI), pv, pv, 0) // ldi  R27, sym_lo(R27)
		addrel(c.cursym, int32(p.Pc+4), p.To.Offset, sym, objabi.R_SW64_GPRELLOW)
		o3 = OP_MEMORY(c.oprrr(p.As), r, pv, 0) //CALL/JMP R, (R27)
		addrel(c.cursym, int32(p.Pc+8), p.To.Offset, sym, objabi.R_SW64_HINT)
	case 10: /*buildCall,ABSR, ABR,ABEQ, ABGE, ABGT, ABLE, ABLT,ABNE, ABLBC, ABLBS, AFBEQ, AFBGE, AFBGT, AFBLE, AFBLT, AFBNE*/
		sym := p.To.Sym
		addrel(c.cursym, int32(p.Pc), p.To.Offset, sym, objabi.R_SW64_BRADDR)

		if p.To.Type == obj.TYPE_BRANCH {
			target := p.To.Val.(*obj.Prog)
			offset := (target.Pc - p.Pc - 4)
			//offset := (p.To.Target().Pc - p.Pc - 4)
			o1 = OP_CONTROL(c.oprrr(p.As), getRegister(p.From.Reg), int32(offset)/4)
		} else {
			o1 = OP_CONTROL(c.oprrr(p.As), getRegister(p.From.Reg), int32(p.To.Offset)/4)
		}
	case 11: /*buildNoop*/
		p.As = ALDIH
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REGSP
		disp := p.To.Offset

		if (disp > 0 && disp&0xffff != disp) ||
			(disp < 0 && -disp&0xffff != -disp) {
			c.ctxt.Diag("Memory instruction with immediate(0x%x) bigger than 0xffff at\n%v\n", disp, p)
		}

		o1 = OP_MEMORY(c.oprrr(p.As), getRegister(p.From.Reg), getRegister(p.To.Reg), int16(disp))
	case 12: /*buildLoad,ALDL, ALDW, ALDHU, ALDBU, ASYMADDR*/
		sym := p.To.Sym
		r := getRegister(p.From.Reg)
		as := p.As
		if as == ASYMADDR {
			as = ALDI
		}
		o1 = OP_MEMORY(c.oprrr(ALDIH), r, REGSB&31, 0) // ldih r, sym(gp)
		addrel(c.cursym, int32(p.Pc), p.To.Offset, sym, objabi.R_SW64_GPRELHIGH)
		o2 = OP_MEMORY(c.oprrr(as), r, r, 0) // ldx  r, sym_lo(r)
		addrel(c.cursym, int32(p.Pc+4), p.To.Offset, sym, objabi.R_SW64_GPRELLOW)
	case 13: /*buildStore*/
		sym := p.To.Sym

		if sym == nil {
			c.ctxt.Diag("buildStore not support %v", p)
		}

		// STx Rn, sym(SB)
		r := getRegister(p.From.Reg)
		o1 = OP_MEMORY(c.oprrr(ALDIH), REGTMP&31, REGSB&31, 0) // ldih TMP, sym(gp)
		addrel(c.cursym, int32(p.Pc), p.To.Offset, sym, objabi.R_SW64_GPRELHIGH)
		o2 = OP_MEMORY(c.oprrr(p.As), r, REGTMP&31, 0) // STx  Rn, sym_lo(TMP)
		addrel(c.cursym, int32(p.Pc+4), p.To.Offset, sym, objabi.R_SW64_GPRELLOW)
	case 14:
		//LDI AT, $const
		ib := p.RestArgs[0].Offset
		q := c.newprog()
		q.Pc = p.Pc
		q.As = ALDI
		q.From.Type = obj.TYPE_REG
		q.From.Reg = REGTMP
		q.To.Type = obj.TYPE_CONST
		q.To.Offset = ib
		disp := q.To.Offset

		if (disp > 0 && disp&0xffff != disp) ||
			(disp < 0 && -disp&0xffff != -disp) {
			c.ctxt.Diag("Memory instruction with immediate(0x%x) bigger than 0xffff at\n%v\n", disp, q)
		}

		o1 = OP_MEMORY(c.oprrr(q.As), getRegister(q.From.Reg), getRegister(q.To.Reg), int16(disp))
		// XXX Rn1, AT, Rn2
		p.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  REGTMP,
		})
		o2 = OP_ARITHMETIC(c.oprrr(p.As), getRegister(p.From.Reg), getReg2(p), getRegister(p.To.Reg))
	case 15: /*buildLoad,AFLDD, AFLDS*/
		sym := p.To.Sym
		o1 = OP_MEMORY(c.oprrr(ALDIH), REGTMP&31, REGSB&31, 0) // LDIH AT, sym(gp)
		addrel(c.cursym, int32(p.Pc), p.To.Offset, sym, objabi.R_SW64_GPRELHIGH)
		o2 = OP_MEMORY(c.oprrr(ALDI), REGTMP&31, REGTMP&31, 0) // LDI  AT, sym_lo(AT)
		addrel(c.cursym, int32(p.Pc+4), p.To.Offset, sym, objabi.R_SW64_GPRELLOW)
		o3 = OP_MEMORY(c.oprrr(p.As), getRegister(p.From.Reg), REGTMP&31, 0) // FLDX, Fn, (AT)
	case 16:
		//SYMADDR AT, $int64.ib
		ib := p.RestArgs[0].Offset
		q := c.newprog()
		q.Pc = p.Pc
		q.As = ASYMADDR
		q.From.Type = obj.TYPE_REG
		q.From.Reg = REGTMP
		q.To.Type = obj.TYPE_MEM
		q.To.Sym = c.ctxt.Int64Sym(ib)
		q.To.Name = obj.NAME_EXTERN
		sym := q.To.Sym
		r := getRegister(q.From.Reg)
		as := q.As
		if as == ASYMADDR {
			as = ALDI
		}
		o1 = OP_MEMORY(c.oprrr(ALDIH), r, REGSB&31, 0) // ldih r, sym(gp)
		addrel(c.cursym, int32(q.Pc), q.To.Offset, sym, objabi.R_SW64_GPRELHIGH)
		o2 = OP_MEMORY(c.oprrr(as), r, r, 0) // ldx  r, sym_lo(r)
		addrel(c.cursym, int32(q.Pc+4), q.To.Offset, sym, objabi.R_SW64_GPRELLOW)
		// LDL AT, (AT)
		q2 := c.newprog()
		q2.As = ALDL
		q2.Pc = p.Pc + 8
		q2.From.Type = obj.TYPE_REG
		q2.From.Reg = REGTMP
		q2.To.Type = obj.TYPE_MEM
		q2.To.Reg = REGTMP
		disp := q2.To.Offset

		if (disp > 0 && disp&0xffff != disp) || (disp < 0 && -disp&0xffff != -disp) {
			c.ctxt.Diag("Memory instruction with immediate(0x%x) bigger than 0xffff at\n%v\n", disp, q2)
		}

		o3 = OP_MEMORY(c.oprrr(q2.As), getRegister(q2.From.Reg), getRegister(q2.To.Reg), int16(disp))
		// XXX Rn1, AT, Rn2
		p.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  REGTMP,
		})
		o4 = OP_ARITHMETIC(c.oprrr(p.As), getRegister(p.From.Reg), getReg2(p), getRegister(p.To.Reg))

	case 17:
		// LDx/STx   Ra, off(Rb)
		//  convert to --->
		// ADDL      Rb, $off, REGTMP  #len 16
		// LDx/STx   Ra, (REGTMP)       #len 4
		rb := p.To.Reg
		off := p.To.Offset

		// ADDL Rb, $off, REGTMP
		addl := c.newprog()
		addl.Pc = p.Pc
		addl.As = AADDL
		addl.From = obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  rb,
		}
		addl.SetFrom3(obj.Addr{
			Type:   obj.TYPE_CONST,
			Offset: off,
		})
		addl.To = obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  REGTMP,
		}

		//SYMADDR AT, $int64.ib
		ib := addl.RestArgs[0].Offset
		q := c.newprog()
		q.Pc = addl.Pc
		q.As = ASYMADDR
		q.From.Type = obj.TYPE_REG
		q.From.Reg = REGTMP
		q.To.Type = obj.TYPE_MEM
		q.To.Sym = c.ctxt.Int64Sym(ib)
		q.To.Name = obj.NAME_EXTERN
		sym := q.To.Sym
		r := getRegister(q.From.Reg)
		as := q.As
		if as == ASYMADDR {
			as = ALDI
		}
		o1 = OP_MEMORY(c.oprrr(ALDIH), r, REGSB&31, 0) // ldih r, sym(gp)
		addrel(c.cursym, int32(q.Pc), q.To.Offset, sym, objabi.R_SW64_GPRELHIGH)
		o2 = OP_MEMORY(c.oprrr(as), r, r, 0) // ldx  r, sym_lo(r)
		addrel(c.cursym, int32(q.Pc+4), q.To.Offset, sym, objabi.R_SW64_GPRELLOW)
		// LDL AT, (AT)
		q2 := c.newprog()
		q2.As = ALDL
		q2.Pc = addl.Pc + 8
		q2.From.Type = obj.TYPE_REG
		q2.From.Reg = REGTMP
		q2.To.Type = obj.TYPE_MEM
		q2.To.Reg = REGTMP
		disp := q2.To.Offset

		if (disp > 0 && disp&0xffff != disp) || (disp < 0 && -disp&0xffff != -disp) {
			c.ctxt.Diag("Memory instruction with immediate(0x%x) bigger than 0xffff at\n%v\n", disp, q2)
		}

		o3 = OP_MEMORY(c.oprrr(q2.As), getRegister(q2.From.Reg), getRegister(q2.To.Reg), int16(disp))
		// XXX Rn1, AT, Rn2
		addl.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  REGTMP,
		})
		o4 = OP_ARITHMETIC(c.oprrr(addl.As), getRegister(addl.From.Reg), getReg2(addl), getRegister(addl.To.Reg))
		// LDx Ra, (REGTMP)
		p.To = obj.Addr{
			Type: obj.TYPE_MEM,
			Reg:  REGTMP,
		}

		//disp := p.To.Offset
		if (disp > 0 && disp&0xffff != disp) ||
			(disp < 0 && -disp&0xffff != -disp) {
			c.ctxt.Diag("Memory instruction with immediate(0x%x) bigger than 0xffff at\n%v\n", disp, p)
		}

		o5 = OP_MEMORY(c.oprrr(p.As), getRegister(p.From.Reg), getRegister(p.To.Reg), int16(p.To.Offset))

	case 18: /* buildTrap */
		o1 = 80
	case 19: /* buildWORD */
		o1 = uint32(p.From.Offset)
	case 20: /*buildLDGP*/
		dummysym := c.cursym

		gp := getRegister(p.From.Reg)
		r := getRegister(p.To.Reg)

		if p.To.Reg == REGZERO {
			// If the ATEXT+0 isn't the noop, then
			// runtime.goexit will be invoked at o3 which is ldgp
			// without setup $at. see the logical in
			// runtime/proc.go' func newproc1 about
			// newg.sched.pc = funcPC(goexit) + sys.PCQuantum
			o1 = OP_MEMORY(c.oprrr(ALDIH), gp, gp, 0) // only for alignment

			//insert a "BR AT, 0" to find the correct PC value
			o2 = OP_CONTROL(c.oprrr(ABR), REGTMP, 0)
			o3 = OP_MEMORY(c.oprrr(ALDIH), gp, REGTMP&31, 0)
			addrel(c.cursym, int32(p.Pc+8), 4, dummysym, objabi.R_SW64_GPDISP)
			o4 = OP_MEMORY(c.oprrr(ALDI), gp, gp, 0) //ldi  $gp, m($gp)
		} else {
			if p.To.Reg == REG_R27 && p.Pc != 0 {
				p.To.Offset = p.Pc
			}
			o1 = OP_MEMORY(c.oprrr(ALDIH), gp, r, 0) //ldih $gp, n($r)
			addrel(c.cursym, int32(p.Pc), 4, dummysym, objabi.R_SW64_GPDISP)
			o2 = OP_MEMORY(c.oprrr(ALDI), gp, gp, 0) //ldi  $gp, m($gp)

			o3 = OP_MEMORY(c.oprrr(ALDI), gp, gp, int16(p.To.Offset))
			o4 = OP_MEMORY(c.oprrr(ALDIH), gp, gp, 0) // only for alignment
		}
	case 21: /*buildCallByReg*/
		offset := p.To.Offset
		if (p.As != obj.AJMP && p.As != obj.ACALL) || p.To.Sym != nil || !isint16(offset) {
			c.ctxt.Diag("buildCallByReg doesn't support %v", p)
			return
		}
		ra := getRegister(p.From.Reg)
		pv := getRegister(p.To.Reg)
		o1 = OP_MEMORY(c.oprrr(p.As), ra, pv, int16(offset)) //CALL/JMP RA, (Rn
	case 22: /*LDx Reg,(TLS_LE)(SB)*/
		if p.To.Sym == nil || p.From.Type != obj.TYPE_REG || !isint16(p.To.Offset) {
			c.ctxt.Diag("%v is not support TLS_LE", p)
		}

		offset := int16(p.To.Offset)
		rn := int16(p.From.Reg & 31)
		r0 := int16(REG_R0 & 31)
		sym := p.To.Sym
		// sys_call 0x9e
		o1 = OP_SYSCALL(c.oprrr(ASYS_CALL), 0x9e)
		// LDIH R0, sym(R0) !tprelhi
		o2 = OP_MEMORY(c.oprrr(ALDIH), r0, r0, 0)
		addrel(c.cursym, int32(p.Pc+4), 0, sym, objabi.R_SW64_TPRELHI)
		// STx/LDx Rn, sym(R0) !tprello
		o3 = OP_MEMORY(c.oprrr(p.As), rn, r0, offset)
		addrel(c.cursym, int32(p.Pc+8), 0, sym, objabi.R_SW64_TPRELLO)

	case 23: /*LDx Reg,(TLS_IE)(SB)*/
		if p.To.Sym == nil || p.From.Type != obj.TYPE_REG || !isint16(p.To.Offset) {
			c.ctxt.Diag("%v is not support TLS_IE", p)
		}
		offset := int16(p.To.Offset)
		rn := int16(p.From.Reg & 31)
		r0 := int16(REG_R0 & 31)
		rtmp := int16(REGTMP & 31)
		sym := p.To.Sym
		// sys_call 0x9e
		o1 = OP_SYSCALL(c.oprrr(ASYS_CALL), 0x9e)

		//LDL RTMP, offset($sym) !gottprel
		o2 = OP_MEMORY(c.oprrr(ALDL), rtmp, REGSB&31, offset)
		addrel(c.cursym, int32(p.Pc+4), 0, sym, objabi.R_SW64_GOTTPREL)

		//ADDL R0, RTMP, RTMP
		o3 = OP_MEMORY(c.oprrr(AADDL), r0, rtmp, rtmp)
		//STx/LDx Rn, 0(RTMP)
		o4 = OP_MEMORY(c.oprrr(p.As), rn, rtmp, 0)
	//zxw add
	case 24: /*retjmp*/
		sym := p.To.Sym
		//r := getRegister(p.From.Reg)
		pv := REG_R27 & int16(31)
		o1 = OP_MEMORY(c.oprrr(ALDIH), pv, REGSB&31, 0) // ldih R27, sym(GP)
		addrel(c.cursym, int32(p.Pc), p.To.Offset, sym, objabi.R_SW64_GPRELHIGH)
		o2 = OP_MEMORY(c.oprrr(ALDI), pv, pv, 0) // ldi  R27, sym_lo(R27)
		addrel(c.cursym, int32(p.Pc+4), p.To.Offset, sym, objabi.R_SW64_GPRELLOW)
		o3 = OP_MEMORY(c.oprrr(obj.AJMP), REGLINK&31, pv, 0) //CALL/JMP R, (R27)
		addrel(c.cursym, int32(p.Pc+8), p.To.Offset, sym, objabi.R_SW64_HINT)
	}
	out[0] = o1
	out[1] = o2
	out[2] = o3
	out[3] = o4
	out[4] = o5
	return
}

func (c *ctxt77) oprrr(a obj.As) uint32 {
	switch a {
	case ASYS_CALL_B:
		return SP(0, 0) | (0 << 25)
	case ASYS_CALL:
		return SP(0, 0) | (1 << 25)
	case obj.AUNDEF:
		return 0
	case AEXCB:
		return SP(1, 0)
	case ACALL:
		return SP(0, 1)
	case AJMP:
		return SP(0, 3)
	case ARET:
		return SP(0, 2)
	case ABR:
		return SP(0, 4)
	case ABSR:
		return SP(0, 5)
	case AMEMB:
		return FMM(0, 0)
	case AIMEMB:
		return FMM(0, 1)
	case ARD_F:
		return FMM(16, 0)
	case AWR_F:
		return FMM(16, 32)
	case ALLDW:
		return SP(0, 8) | (0 << 12)
	case ALLDL:
		return SP(0, 8) | (1 << 12)
	case ALSTW:
		return SP(0, 8) | (8 << 12)
	case ALSTL:
		return SP(0, 8) | (9 << 12)
	case AADDW:
		return FA(0, 0)
	case ASUBW:
		return FA(0, 1)
	case AS4ADDW:
		return FA(0, 2)
	case AS4SUBW:
		return FA(0, 3)
	case AS8ADDW:
		return FA(0, 4)
	case AS8SUBW:
		return FA(0, 5)
	case AADDL:
		return FA(0, 8)
	case ASUBL:
		return FA(0, 9)
	case AS4ADDL:
		return FA(0, 10)
	case AS4SUBL:
		return FA(0, 11)
	case AS8ADDL:
		return FA(0, 12)
	case AS8SUBL:
		return FA(0, 13)
	case AMULW:
		return FA(1, 0)
	case AMULL:
		return FA(1, 8)
	case AUMULH:
		return FA(1, 9)
	case ACMPEQ:
		return FA(2, 8)
	case ACMPLT:
		return FA(2, 9)
	case ACMPLE:
		return FA(2, 10)
	case ACMPULT:
		return FA(2, 11)
	case ACMPULE:
		return FA(2, 12)
	case AAND:
		return FA(3, 8)
	case ABIC:
		return FA(3, 9)
	case ABIS:
		return FA(3, 10)
	case AORNOT:
		return FA(3, 11)
	case AXOR:
		return FA(3, 12)
	case AEQV:
		return FA(3, 13)
	case AINSLB:
		return FA(4, 0)
	case AINSLH:
		return FA(4, 1)
	case AINSLW:
		return FA(4, 2)
	case AINSLL:
		return FA(4, 3)
	case AINSHB:
		return FA(4, 4)
	case AINSHH:
		return FA(4, 5)
	case AINSHW:
		return FA(4, 6)
	case AINSHL:
		return FA(4, 7)
	case ASLL:
		return FA(4, 8)
	case ASRL:
		return FA(4, 9)
	case ASRA:
		return FA(4, 10)
	case AEXTLB:
		return FA(5, 0)
	case AEXTLH:
		return FA(5, 1)
	case AEXTLW:
		return FA(5, 2)
	case AEXTLL:
		return FA(5, 3)
	case AEXTHB:
		return FA(5, 4)
	case AEXTHH:
		return FA(5, 5)
	case AEXTHW:
		return FA(5, 6)
	case AEXTHL:
		return FA(5, 7)
	case ACTPOP:
		return FA(5, 8)
	case ACTLZ:
		return FA(5, 9)
	case ACTTZ:
		return FA(5, 10)
	case AMASKLB:
		return FA(6, 0)
	case AMASKLH:
		return FA(6, 1)
	case AMASKLW:
		return FA(6, 2)
	case AMASKLL:
		return FA(6, 3)
	case AMASKHB:
		return FA(6, 4)
	case AMASKHH:
		return FA(6, 5)
	case AMASKHW:
		return FA(6, 6)
	case AMASKHL:
		return FA(6, 7)
	case AZAP:
		return FA(6, 8)
	case AZAPNOT:
		return FA(6, 9)
	case ASEXTB:
		return FA(6, 10)
	case ASEXTH:
		return FA(6, 11)
	case ACMPGEB:
		return FA(6, 12)
	case AFIMOVS:
		return FA(7, 0)
	case AFIMOVD:
		return FA(7, 8)
	case ASELEQ:
		return SP(1, 1) | (0 << 5)
	case ASELGE:
		return SP(1, 1) | (1 << 5)
	case ASELGT:
		return SP(1, 1) | (2 << 5)
	case ASELLE:
		return SP(1, 1) | (3 << 5)
	case ASELLT:
		return SP(1, 1) | (4 << 5)
	case ASELNE:
		return SP(1, 1) | (5 << 5)
	case ASELLBC:
		return SP(1, 1) | (6 << 5)
	case ASELLBS:
		return SP(1, 1) | (7 << 5)
	case APRI_LD:
		return SP(2, 5)
	case ABEQ:
		return SP(3, 0)
	case ABNE:
		return SP(3, 1)
	case ABLT:
		return SP(3, 2)
	case ABLE:
		return SP(3, 3)
	case ABGT:
		return SP(3, 4)
	case ABGE:
		return SP(3, 5)
	case ABLBC:
		return SP(3, 6)
	case ABLBS:
		return SP(3, 7)
	case AFBEQ:
		return SP(3, 8)
	case AFBNE:
		return SP(3, 9)
	case AFBLT:
		return SP(3, 10)
	case AFBLE:
		return SP(3, 11)
	case AFBGT:
		return SP(3, 12)
	case AFBGE:
		return SP(3, 13)
	case AFADDS:
		return FAF(0, 0)
	case AFADDD:
		return FAF(0, 1)
	case AFSUBS:
		return FAF(0, 2)
	case AFSUBD:
		return FAF(0, 3)
	case AFMULS:
		return FAF(0, 4)
	case AFMULD:
		return FAF(0, 5)
	case AFDIVS:
		return FAF(0, 6)
	case AFDIVD:
		return FAF(0, 7)
	case AFSQRTS:
		return FAF(0, 8)
	case AFSQRTD:
		return FAF(0, 9)
	case AFCMPEQ:
		return FAF(1, 0)
	case AFCMPLE:
		return FAF(1, 1)
	case AFCMPLT:
		return FAF(1, 2)
	case AFCMPUN:
		return FAF(1, 3)
	case AFCVTSD:
		return FAF(2, 0)
	case AFCVTDS:
		return FAF(2, 1)
	case AFCVTDL_G:
		return FAF(2, 2)
	case AFCVTDL_P:
		return FAF(2, 3)
	case AFCVTDL_Z:
		return FAF(2, 4)
	case AFCVTDL_N:
		return FAF(2, 5)
	case AFCVTDL:
		return FAF(2, 7)
	case AFCVTWL:
		return FAF(2, 8)
	case AFCVTLW:
		return FAF(2, 9)
	case AFCVTLS:
		return FAF(2, 13)
	case AFCVTLD:
		return FAF(2, 15)
	case AFCPYS:
		return FAF(3, 0)
	case AFCPYSE:
		return FAF(3, 1)
	case AFCPYSN:
		return FAF(3, 2)
	case AIFMOVS:
		return FAF(4, 0)
	case AIFMOVD:
		return FAF(4, 1)
	case ARFPCR:
		return FAF(5, 0)
	case AWFPCR:
		return FAF(5, 1)
	case ASETFPEC0:
		return FAF(5, 4)
	case ASETFPEC1:
		return FAF(5, 5)
	case ASETFPEC2:
		return FAF(5, 6)
	case ASETFPEC3:
		return FAF(5, 7)
	case AFMAS:
		return FCA(0, 0)
	case AFMAD:
		return FCA(0, 1)
	case AFMSS:
		return FCA(0, 2)
	case AFMSD:
		return FCA(0, 3)
	case AFNMAS:
		return FCA(0, 4)
	case AFNMAD:
		return FCA(0, 5)
	case AFNMSS:
		return FCA(0, 6)
	case AFNMSD:
		return FCA(0, 7)
	case AFSELEQ:
		return FCA(1, 0)
	case AFSELNE:
		return FCA(1, 1)
	case AFSELLT:
		return FCA(1, 2)
	case AFSELLE:
		return FCA(1, 3)
	case AFSELGT:
		return FCA(1, 4)
	case AFSELGE:
		return FCA(1, 5)
	case ALDBU:
		return SP(2, 0)
	case ALDHU:
		return SP(2, 1)
	case ALDW:
		return SP(2, 2)
	case ALDL:
		return SP(2, 3)
	case ALDL_U:
		return SP(2, 4)
	case AFLDS:
		return SP(2, 6)
	case AFLDD:
		return SP(2, 7)
	case ASTB:
		return SP(2, 8)
	case ASTH:
		return SP(2, 9)
	case ASTW:
		return SP(2, 10)
	case ASTL:
		return SP(2, 11)
	case ASTL_U:
		return SP(2, 12)
	case APRI_ST:
		return SP(2, 13)
	case AFSTS:
		return SP(2, 14)
	case AFSTD:
		return SP(2, 15)
	case ALDI:
		return SP(3, 14)
	case ALDIH:
		return SP(3, 15)
	}

	if a < 0 {
		c.ctxt.Diag("bad orr opcode -%v", -a)
	} else {
		c.ctxt.Diag("bad orr opcode %v", a)
	}
	return 0
}

func (c *ctxt77) opirr(a obj.As) uint32 {
	switch a {
	case AADDW:
		return FAI(0, 0)
	case ASUBW:
		return FAI(0, 1)
	case AS4ADDW:
		return FAI(0, 2)
	case AS4SUBW:
		return FAI(0, 3)
	case AS8ADDW:
		return FAI(0, 4)
	case AS8SUBW:
		return FAI(0, 5)
	case AADDL:
		return FAI(0, 8)
	case ASUBL:
		return FAI(0, 9)
	case AS4ADDL:
		return FAI(0, 10)
	case AS4SUBL:
		return FAI(0, 11)
	case AS8ADDL:
		return FAI(0, 12)
	case AS8SUBL:
		return FAI(0, 13)
	case AMULW:
		return FAI(1, 0)
	case AMULL:
		return FAI(1, 8)
	case AUMULH:
		return FAI(1, 9)
	case ACMPEQ:
		return FAI(2, 8)
	case ACMPLT:
		return FAI(2, 9)
	case ACMPLE:
		return FAI(2, 10)
	case ACMPULT:
		return FAI(2, 11)
	case ACMPULE:
		return FAI(2, 12)
	case AAND:
		return FAI(3, 8)
	case ABIC:
		return FAI(3, 9)
	case ABIS:
		return FAI(3, 10)
	case AORNOT:
		return FAI(3, 11)
	case AXOR:
		return FAI(3, 12)
	case AEQV:
		return FAI(3, 13)
	case AINSLB:
		return FAI(4, 0)
	case AINSLH:
		return FAI(4, 1)
	case AINSLW:
		return FAI(4, 2)
	case AINSLL:
		return FAI(4, 3)
	case AINSHB:
		return FAI(4, 4)
	case AINSHH:
		return FAI(4, 5)
	case AINSHW:
		return FAI(4, 6)
	case AINSHL:
		return FAI(4, 7)
	case ASLL:
		return FAI(4, 8)
	case ASRL:
		return FAI(4, 9)
	case ASRA:
		return FAI(4, 10)
	case AEXTLB:
		return FAI(5, 0)
	case AEXTLH:
		return FAI(5, 1)
	case AEXTLW:
		return FAI(5, 2)
	case AEXTLL:
		return FAI(5, 3)
	case AEXTHB:
		return FAI(5, 4)
	case AEXTHH:
		return FAI(5, 5)
	case AEXTHW:
		return FAI(5, 6)
	case AEXTHL:
		return FAI(5, 7)
	case AMASKLB:
		return FAI(6, 0)
	case AMASKLH:
		return FAI(6, 1)
	case AMASKLW:
		return FAI(6, 2)
	case AMASKLL:
		return FAI(6, 3)
	case AMASKHB:
		return FAI(6, 4)
	case AMASKHH:
		return FAI(6, 5)
	case AMASKHW:
		return FAI(6, 6)
	case AMASKHL:
		return FAI(6, 7)
	case AZAP:
		return FAI(6, 8)
	case AZAPNOT:
		return FAI(6, 9)
	case ASEXTB:
		return FAI(6, 10)
	case ASEXTH:
		return FAI(6, 11)
	case ACMPGEB:
		return FAI(6, 12)
	case ASELEQ:
		return SP(1, 3) | (0 << 5)
	case ASELGE:
		return SP(1, 3) | (1 << 5)
	case ASELGT:
		return SP(1, 3) | (2 << 5)
	case ASELLE:
		return SP(1, 3) | (3 << 5)
	case ASELLT:
		return SP(1, 3) | (4 << 5)
	case ASELNE:
		return SP(1, 3) | (5 << 5)
	case ASELLBC:
		return SP(1, 3) | (6 << 5)
	case ASELLBS:
		return SP(1, 3) | (7 << 5)
	}
	if a < 0 {
		c.ctxt.Diag("bad irr opcode -%v", -a)
	} else {
		c.ctxt.Diag("bad irr opcode %v", a)
	}
	return 0
}

func getRegister(v int16) int16 {
	if v == 0 {
		return REGZERO & 31
	}
	return v & 31
}
func getReg2(p *obj.Prog) int16 {
	reg := p.Reg
	if reg == 0 && len(p.RestArgs) != 0 {
		reg = p.RestArgs[0].Reg
	}
	return reg & 31
}

func getFrom3Reg(p *obj.Prog) int16 {
	reg := p.RestArgs[1].Reg
	if reg == 0 {
		fmt.Errorf("illegal From3 reg: %v", reg)
	}
	return reg & 31
}

func addrel(cursym *obj.LSym, off int32, add int64, target *obj.LSym, rt objabi.RelocType) {
	r := obj.Addrel(cursym)
	r.Add = add
	r.Siz = 4
	r.Off = off
	r.Sym = target
	r.Type = rt
}
