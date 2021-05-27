// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw64

import (
	"cmd/compile/internal/gc"
	"cmd/compile/internal/ssa"
	"cmd/internal/obj/sw64"
)

func Init(arch *gc.Arch) {
	arch.LinkArch = &sw64.LinkSW64
	arch.REGSP = sw64.REGSP
	arch.MAXWIDTH = 1 << 50

	arch.ZeroRange = zerorange
	//arch.ZeroAuto = zeroAuto //zxw new change
	arch.Ginsnop = ginsnop
	arch.Ginsnopdefer = ginsnop //zxw new add

	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
