// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw64

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/sw64"
)

func zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {
	base := p.Ctxt.FixedFrameSize() + off
	for i := int64(0); i < cnt; i += int64(gc.Widthptr) {
		// STL ZERO, FixedFrameSize+off+i(SP)
		p = pp.Appendpp(p, sw64.ASTL, obj.TYPE_REG, sw64.REGZERO, 0, obj.TYPE_MEM, sw64.REGSP, base+i)
	}

	return p
}

func zeroAuto(pp *gc.Progs, n *gc.Node) {
	// Note: this code must not clobber any registers.
	sym := n.Sym.Linksym()
	size := n.Type.Size()
	for i := int64(0); i < size; i += 8 {
		p := pp.Prog(sw64.ASTL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sw64.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = sw64.REGSP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

//zxw new change
func ginsnop(pp *gc.Progs) *obj.Prog {
	p := pp.Prog(sw64.ALDI)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = sw64.REG_R31

	p.To.Type = obj.TYPE_CONST
	p.To.Offset = 0
	return p
}
