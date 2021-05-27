// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw64

import (
	"cmd/compile/internal/gc"
	"cmd/compile/internal/logopt"
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/types"
	"cmd/internal/obj"
	"cmd/internal/obj/sw64"
	"math"
)

// isFPreg returns whether r is an FP register
func isFPreg(r int16) bool {
	return sw64.REG_F0 <= r && r <= sw64.REG_F31
}

func loadByType(t *types.Type, r int16) obj.As {
	if sw64.IsFReg(int(r)) {
		if t.Size() == 4 {
			return sw64.AFLDS
		}
		if t.Size() == 8 {
			return sw64.AFLDD
		}
		panic("bad load type")
	}
	switch t.Size() {
	case 1:
		return sw64.ALDBU
	case 2:
		return sw64.ALDHU
	case 4:
		return sw64.ALDW
	case 8:
		return sw64.ALDL
	}
	panic("bad load type")
}

func storeByType(t *types.Type, r int16) obj.As {
	if sw64.IsFReg(int(r)) {
		if t.Size() == 4 {
			return sw64.AFSTS
		}
		if t.Size() == 8 {
			return sw64.AFSTD
		}
		panic("bad store type")
	}

	switch t.Size() {
	case 1:
		return sw64.ASTB
	case 2:
		return sw64.ASTH
	case 4:
		return sw64.ASTW
	case 8:
		return sw64.ASTL
	}
	panic("bad store type")
}

func ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpLoadReg:
		r := v.Reg()
		p := s.Prog(loadByType(v.Type, r))

		p.From.Type = obj.TYPE_REG
		p.From.Reg = r

		gc.AddrAuto(&p.To, v.Args[0])
	case ssa.OpStoreReg:
		r := v.Args[0].Reg()
		p := s.Prog(storeByType(v.Type, r))

		p.From.Type = obj.TYPE_REG
		p.From.Reg = r

		gc.AddrAuto(&p.To, v)
	case ssa.OpSW64CALLstatic, ssa.OpSW64CALLclosure, ssa.OpSW64CALLinter:
		_, ok := v.Aux.(*obj.LSym)
		if !ok && v.Args[0].Reg() != sw64.REG_R27 {
			// TODO(snyh): remove this if we can restrict ACALL use R27 in ssa
			p := s.Prog(sw64.ALDI)
			p.From = obj.Addr{Type: obj.TYPE_REG, Reg: sw64.REG_R27}
			p.To = obj.Addr{Type: obj.TYPE_ADDR, Reg: v.Args[0].Reg()}
		}
		s.Call(v)
		//zxw add
	case ssa.OpSW64LoweredWB:
		p := s.Prog(obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = v.Aux.(*obj.LSym)
	case ssa.OpSW64FCVTSD,
		ssa.OpSW64FCVTDS,
		ssa.OpSW64IFMOVD,
		ssa.OpSW64FIMOVD,
		ssa.OpSW64FCVTDL,
		ssa.OpSW64FCVTLS,
		ssa.OpSW64FCVTLD,
		ssa.OpSW64FCVTLW,
		ssa.OpSW64FCVTWL,
		ssa.OpSW64FSQRTS,
		ssa.OpSW64FSQRTD,
		ssa.OpSW64CTLZ,
		ssa.OpSW64CTTZ,
		ssa.OpSW64CTPOP,
		ssa.OpSW64SEXTB,
		ssa.OpSW64SEXTH,
		ssa.OpSW64FCVTDL_Z,
		ssa.OpSW64FCVTDL_P,
		ssa.OpSW64FCVTDL_G,
		ssa.OpSW64FCVTDL_N:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpSW64SEXTBconst,
		ssa.OpSW64SEXTHconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpSW64FADDS,
		ssa.OpSW64FADDD,
		ssa.OpSW64FSUBD,
		ssa.OpSW64FSUBS,
		ssa.OpSW64FMULS,
		ssa.OpSW64FMULD,
		ssa.OpSW64FDIVS,
		ssa.OpSW64FDIVD,
		ssa.OpSW64FCPYS,
		ssa.OpSW64ADDV,
		ssa.OpSW64ADDW,
		ssa.OpSW64SUBV,
		ssa.OpSW64MULW,
		ssa.OpSW64MULL,
		ssa.OpSW64UMULH,
		ssa.OpSW64AND,
		ssa.OpSW64BIS,
		ssa.OpSW64XOR,
		ssa.OpSW64ORNOT,
		ssa.OpSW64SLL,
		ssa.OpSW64SRL,
		ssa.OpSW64SRA:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Args[1].Reg(),
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSW64FABS:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sw64.REG_F31
		p.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Args[0].Reg(),
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSW64CALLudiv:
		p := s.Prog(obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = gc.Udiv
	case ssa.OpSW64ADDVconst,
		ssa.OpSW64ADDWconst,
		ssa.OpSW64SUBVconst,
		ssa.OpSW64MULWconst,
		ssa.OpSW64MULLconst,
		ssa.OpSW64UMULHconst,
		ssa.OpSW64ANDconst,
		ssa.OpSW64BISconst,
		ssa.OpSW64XORconst,
		ssa.OpSW64ORNOTconst,
		ssa.OpSW64SLLconst,
		ssa.OpSW64SRLconst,
		ssa.OpSW64SRAconst,
		ssa.OpSW64CMPLEconst,
		ssa.OpSW64CMPLTconst,
		ssa.OpSW64CMPEQconst,
		ssa.OpSW64CMPULEconst,
		ssa.OpSW64CMPULTconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		ib := obj.Addr{
			Type:   obj.TYPE_CONST,
			Offset: v.AuxInt,
		}
		p.SetFrom3(ib)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpSW64NEGV:
		// SUB from REGZERO
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sw64.REGZERO
		p.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Args[0].Reg(),
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpSW64NEGF,
		ssa.OpSW64NEGD:
		//FCPYSN Fx,Fx,Fy
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Args[0].Reg(),
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpSW64MOVBstorezero,
		ssa.OpSW64MOVHstorezero,
		ssa.OpSW64MOVWstorezero,
		ssa.OpSW64MOVVstorezero:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sw64.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		gc.AddAux(&p.To, v)
	case ssa.OpSW64MOVBstore,
		ssa.OpSW64MOVHstore,
		ssa.OpSW64MOVWstore,
		ssa.OpSW64MOVVstore,
		ssa.OpSW64MOVFstore,
		ssa.OpSW64MOVDstore:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		gc.AddAux(&p.To, v)
	case ssa.OpSW64MOVBload,
		ssa.OpSW64MOVBUload,
		ssa.OpSW64MOVHload,
		ssa.OpSW64MOVHUload,
		ssa.OpSW64MOVWload,
		ssa.OpSW64MOVWUload,
		ssa.OpSW64MOVVload,
		ssa.OpSW64MOVFload,
		ssa.OpSW64MOVDload:
		p := s.Prog(v.Op.Asm())
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		gc.AddAux(&p.To, v)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Reg()

	case ssa.OpSW64SYMADDR:
		// arg0 + auxInt + aux.(*gc.Sym), arg0=SP/SB
		p := s.Prog(sw64.ASYMADDR)
		p.To.Type = obj.TYPE_ADDR
		p.To.Reg = v.Args[0].Reg()
		var wantreg string
		// LDL R, $sym+off(base)
		// the assembler expands it as the following:
		// - base is SP: add constant offset to SP
		//               when constant is large, tmp registe may be used
		// - base is SB: load external address with relocation
		switch v.Aux.(type) {
		default:
			v.Fatalf("aux is of unknown type %T", v.Aux)
		case *obj.LSym:
			wantreg = "SB"
			gc.AddAux(&p.To, v)
		case *gc.Node:
			p.As = sw64.ALDI
			wantreg = "SP"
			gc.AddAux(&p.To, v)
		case nil:
			p.As = sw64.ALDI
			// No sym, just LDL R, $off(SP)
			wantreg = "SP"
			p.To.Offset = v.AuxInt
		}
		if reg := v.Args[0].RegName(); reg != wantreg {
			v.Fatalf("bad reg %s for symbol type %T, want %s", reg, v.Aux, wantreg)
		}
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Reg()

	case ssa.OpCopy, ssa.OpSW64MOVVconvert:
		if v.Type.IsMemory() {
			return
		}
		x := int(v.Args[0].Reg())
		y := int(v.Reg())
		if x == y {
			return
		}
		var as obj.As
		switch {
		case sw64.IsFReg(x) && sw64.IsRReg(y) && v.Type.Size() == 4:
			as = sw64.AFIMOVS
		case sw64.IsFReg(x) && sw64.IsRReg(y) && v.Type.Size() == 8:
			as = sw64.AFIMOVD
		case sw64.IsRReg(x) && sw64.IsFReg(y) && v.Type.Size() == 4:
			as = sw64.AIFMOVS
		case sw64.IsRReg(x) && sw64.IsFReg(y) && v.Type.Size() == 8:
			as = sw64.AIFMOVD
		case sw64.IsRReg(x) && sw64.IsRReg(y):
			as = sw64.ALDI
			x, y = y, x
		case sw64.IsFReg(x) && sw64.IsFReg(y):
			as = sw64.AFCPYS
		default:
			v.Fatalf("not implement OpCopy with %v and %v. %s\n",
				x&63, y&63,
				v.LongString())
		}
		p := s.Prog(as)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = int16(x)
		if as == sw64.AFCPYS {
			p.SetFrom3(obj.Addr{
				Type: obj.TYPE_REG,
				Reg:  int16(x),
			})
		}
		p.To.Type = obj.TYPE_REG
		p.To.Reg = int16(y)
	case ssa.OpSW64MOVFconst,
		ssa.OpSW64MOVDconst:
		p := s.Prog(v.Op.Asm())
		p.To.Type = obj.TYPE_FCONST
		p.To.Val = math.Float64frombits(uint64(v.AuxInt))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Reg()
	case ssa.OpSW64CMPEQ, ssa.OpSW64CMPLE, ssa.OpSW64CMPLT,
		ssa.OpSW64CMPULE, ssa.OpSW64CMPULT, ssa.OpSW64FCMPUN,
		ssa.OpSW64FCMPEQ, ssa.OpSW64FCMPLE, ssa.OpSW64FCMPLT:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Args[1].Reg(),
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpSW64MOVVconst:
		r := v.Reg()
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = v.AuxInt
		if isFPreg(r) {
			// cannot move into FP registers, use TMP as intermediate
			p.From.Reg = sw64.REGTMP
			p = s.Prog(sw64.AIFMOVD)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = sw64.REGTMP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
		}

	case ssa.OpSW64LoweredNilCheck:
		// Issue a load which will fault if arg is nil.
		p := s.Prog(sw64.ALDBU)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		gc.AddAux(&p.To, v)

		p.From.Type = obj.TYPE_REG
		p.From.Reg = sw64.REGTMP
		//zxw new add
		if logopt.Enabled() {
			logopt.LogOpt(v.Pos, "nilcheck", "genssa", v.Block.Func.Name)
		}

		if gc.Debug_checknil != 0 && v.Pos.Line() > 1 { // v.Pos.Line()==1 in generated wrappers
			gc.Warnl(v.Pos, "generated nil check")
		}
	case ssa.OpSW64DUFFZERO:
		// runtime.duffzero expects start address - 8 in R1

		//SUBL arg0.REG, $8, R1
		p := s.Prog(sw64.ASUBL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.SetFrom3(obj.Addr{
			Type:   obj.TYPE_CONST,
			Offset: 8,
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sw64.REG_R1
		// DUFFZERO runtime.duffzero+v.AuxInt(SB)
		p = s.Prog(obj.ADUFFZERO)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = gc.Duffzero
		p.To.Offset = v.AuxInt
	case ssa.OpSW64LoweredZero:
		//SUBL   R1, $sz ,R1
		//STx    R31,sz(R1)
		//ADDL   R1, $sz, R1
		//CMPULT R1, Rarg1, REGTMP
		//BNE    REGTMP, -3(PC)
		var sz int64
		var mov obj.As
		switch {
		case v.AuxInt%8 == 0:
			sz = 8
			mov = sw64.ASTL
		case v.AuxInt%4 == 0:
			sz = 4
			mov = sw64.ASTW
		case v.AuxInt%2 == 0:
			sz = 2
			mov = sw64.ASTH
		default:
			sz = 1
			mov = sw64.ASTB
		}
		//SUBL R1, $sz ,R1
		p := s.Prog(sw64.ASUBL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sw64.REG_R1
		p.SetFrom3(obj.Addr{
			Type:   obj.TYPE_CONST,
			Offset: sz,
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sw64.REG_R1
		//STx R31, sz(R1)
		p2 := s.Prog(mov)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = sw64.REGZERO
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = sw64.REG_R1
		p2.To.Offset = sz
		//ADDL R1, $sz, R1
		p3 := s.Prog(sw64.AADDL)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = sw64.REG_R1
		p3.SetFrom3(obj.Addr{
			Type:   obj.TYPE_CONST,
			Offset: sz,
		})
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = sw64.REG_R1
		//CMPULT R1, Rarg1, REGTMP
		p4 := s.Prog(sw64.ACMPULT)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = sw64.REG_R1
		p4.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Args[1].Reg(),
		})
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = sw64.REGTMP
		//BNE RGETMP, -3(PC)
		p5 := s.Prog(sw64.ABNE)
		p5.From.Type = obj.TYPE_REG
		p5.From.Reg = sw64.REGTMP
		p5.To.Type = obj.TYPE_BRANCH
		gc.Patch(p5, p2)
	case ssa.OpSW64LoweredMove:
		//SUBL R1, $8, R1
		//LDx  Rtmp, 8(R1)
		//STx  Rtmp, (R2)
		//ADDL R1, $8, R1
		//ADDL R2, $8, R2
		//CMPULT R1, Rarg2, Rtmp
		//BNE Rtmp, -5(PC)
		var sz int64
		var ldx obj.As
		var stx obj.As
		switch {
		case v.AuxInt%8 == 0:
			sz = 8
			ldx = sw64.ALDL
			stx = sw64.ASTL
		case v.AuxInt%4 == 0:
			sz = 4
			ldx = sw64.ALDW
			stx = sw64.ASTW
		case v.AuxInt%2 == 0:
			sz = 2
			ldx = sw64.ALDHU
			stx = sw64.ASTH
		default:
			sz = 1
			ldx = sw64.ALDBU
			stx = sw64.ASTB
		}
		//SUBL R1, $sz, R1
		p := s.Prog(sw64.ASUBL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sw64.REG_R1
		p.SetFrom3(obj.Addr{
			Type:   obj.TYPE_CONST,
			Offset: sz,
		})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sw64.REG_R1
		//LDL Rtmp, sz(R1)
		p2 := s.Prog(ldx)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = sw64.REGTMP
		p2.To.Offset = sz
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = sw64.REG_R1
		//STL Rtmp, (R2)
		p3 := s.Prog(stx)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = sw64.REGTMP
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = sw64.REG_R2
		//ADDL R1, $sz, R1
		p4 := s.Prog(sw64.AADDL)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = sw64.REG_R1
		p4.SetFrom3(obj.Addr{
			Type:   obj.TYPE_CONST,
			Offset: sz,
		})
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = sw64.REG_R1
		//ADDL R2, $sz, R2
		p5 := s.Prog(sw64.AADDL)
		p5.From.Type = obj.TYPE_REG
		p5.From.Reg = sw64.REG_R2
		p5.SetFrom3(obj.Addr{
			Type:   obj.TYPE_CONST,
			Offset: sz,
		})
		p5.To.Type = obj.TYPE_REG
		p5.To.Reg = sw64.REG_R2
		//CMPULT R1, Rarg2, Rtmp
		p6 := s.Prog(sw64.ACMPULT)
		p6.From.Type = obj.TYPE_REG
		p6.From.Reg = sw64.REG_R1
		p6.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Args[2].Reg(),
		})
		p6.To.Type = obj.TYPE_REG
		p6.To.Reg = sw64.REGTMP
		//BNE Rtmp, -5(PC)
		p7 := s.Prog(sw64.ABNE)
		p7.From.Type = obj.TYPE_REG
		p7.From.Reg = sw64.REGTMP
		p7.To.Type = obj.TYPE_BRANCH
		gc.Patch(p7, p2)
	case ssa.OpSW64LoweredGetClosurePtr:
		// Closure pointer is R27 (sw64.REGCTXT).
		gc.CheckLoweredGetClosurePtr(v)
	case ssa.OpSW64LoweredGetCallerSP:
		// caller's SP is FixedFrameSize below the address of the first arg
		p := s.Prog(sw64.ALDI)
		p.To.Type = obj.TYPE_ADDR
		p.To.Offset = -gc.Ctxt.FixedFrameSize()
		p.To.Name = obj.NAME_PARAM

		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Reg()
	//zxw add
	case ssa.OpSW64LoweredGetCallerPC:
		p := s.Prog(obj.AGETCALLERPC)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Reg()
	case ssa.OpSW64FEqual,
		ssa.OpSW64FNotEqual:
		as := sw64.ACMPULT
		if v.Op == ssa.OpSW64FNotEqual {
			as = sw64.ACMPEQ
		}
		p := s.Prog(sw64.AFIMOVD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
		p1 := s.Prog(as)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = sw64.REGZERO
		p1.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Reg(),
		})
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = v.Reg()

		//zxw new add
	case ssa.OpSW64LoweredPanicBoundsA, ssa.OpSW64LoweredPanicBoundsB, ssa.OpSW64LoweredPanicBoundsC:
		p := s.Prog(obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = gc.BoundsCheckFunc[v.AuxInt]
		s.UseArgs(16) // space used in callee args area by assembly stubs

		//zxw new add
	case ssa.OpSW64LoweredAtomicLoad8, ssa.OpSW64LoweredAtomicLoad32, ssa.OpSW64LoweredAtomicLoad64:
		as := sw64.ALDL
		switch v.Op {
		case ssa.OpSW64LoweredAtomicLoad8:
			as = sw64.ALDBU
		case ssa.OpSW64LoweredAtomicLoad32:
			as = sw64.ALDW
		}
		s.Prog(sw64.AMEMB)
		p := s.Prog(as)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Reg0()
		s.Prog(sw64.AMEMB)
	case ssa.OpSW64LoweredAtomicStore8, ssa.OpSW64LoweredAtomicStore32, ssa.OpSW64LoweredAtomicStore64:
		as := sw64.ASTL
		switch v.Op {
		case ssa.OpSW64LoweredAtomicStore8:
			as = sw64.ASTB
		case ssa.OpSW64LoweredAtomicStore32:
			as = sw64.ASTW
		}
		s.Prog(sw64.AMEMB)
		p := s.Prog(as)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		s.Prog(sw64.AMEMB)
	case ssa.OpSW64LoweredAtomicAdd32, ssa.OpSW64LoweredAtomicAdd64:
		// MEMB
		// LLDx Rout, (Rarg0)
		// LDI  Rtmp, 1
		// WR_F Rtmp
		// ADDL Rarg1, Rout, Rtmp
		// LSTx Rtmp, (Rarg0)
		// RD_F Rtmp
		// BEQ  Rtmp, -6(PC)
		// ADDL Rarg1, Rout, Rout
		// MEMB
		var sz int64
		sz = 1
		lldx := sw64.ALLDL
		lstx := sw64.ALSTL
		if v.Op == ssa.OpSW64LoweredAtomicAdd32 {
			lldx = sw64.ALLDW
			lstx = sw64.ALSTW
		}
		s.Prog(sw64.AMEMB)
		// LLDx Rout, (Rarg0)
		p := s.Prog(lldx)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Reg0()
		// LDI  Rtmp, 1
		p1 := s.Prog(sw64.ALDI)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = sw64.REGTMP
		p1.To.Type = obj.TYPE_CONST
		p1.To.Offset = sz
		// WR_F Rtmp
		p2 := s.Prog(sw64.AWR_F)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = sw64.REGTMP
		// ADDL Rarg1, Rout, Rtmp
		p3 := s.Prog(sw64.AADDL)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = v.Args[1].Reg()
		p3.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Reg0(),
		})
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = sw64.REGTMP
		// LSTx Rtmp, (Rarg0)
		p4 := s.Prog(lstx)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = sw64.REGTMP
		p4.To.Type = obj.TYPE_MEM
		p4.To.Reg = v.Args[0].Reg()
		// RD_F Rtmp
		p5 := s.Prog(sw64.ARD_F)
		p5.From.Type = obj.TYPE_REG
		p5.From.Reg = sw64.REGTMP
		// BEQ  Rtmp, -6(PC)
		p6 := s.Prog(sw64.ABEQ)
		p6.From.Type = obj.TYPE_REG
		p6.From.Reg = sw64.REGTMP
		p6.To.Type = obj.TYPE_BRANCH
		gc.Patch(p6, p)
		// ADDL Rarg1, Rout, Rout
		p7 := s.Prog(sw64.AADDL)
		p7.From.Type = obj.TYPE_REG
		p7.From.Reg = v.Args[1].Reg()
		p7.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Reg0(),
		})
		p7.To.Type = obj.TYPE_REG
		p7.To.Reg = v.Reg0()
		s.Prog(sw64.AMEMB)
	case ssa.OpSW64LoweredAtomicExchange32, ssa.OpSW64LoweredAtomicExchange64:
		// MEMB
		// LLDx Rout, (Rarg0)
		// LDI  Rtmp, 1
		// WR_F Rtmp
		// BIS  Rarg1, R31, Rtmp
		// LSTx Rtmp, (Rarg0)
		// RD_F Rtmp
		// BEQ  Rtmp, -6(PC)
		// MEMB
		var sz int64
		sz = 1
		lldx := sw64.ALLDL
		lstx := sw64.ALSTL
		if v.Op == ssa.OpSW64LoweredAtomicExchange32 {
			lldx = sw64.ALLDW
			lstx = sw64.ALSTW
		}
		s.Prog(sw64.AMEMB)
		// LLDx Rout, (Rarg0)
		p := s.Prog(lldx)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Reg0()
		// LDI  Rtmp, 1
		p1 := s.Prog(sw64.ALDI)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = sw64.REGTMP
		p1.To.Type = obj.TYPE_CONST
		p1.To.Offset = sz
		// WR_F Rtmp
		p2 := s.Prog(sw64.AWR_F)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = sw64.REGTMP
		// BIS  Rarg1, R31, Rtmp
		p3 := s.Prog(sw64.ABIS)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = v.Args[1].Reg()
		p3.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  sw64.REGZERO,
		})
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = sw64.REGTMP
		// LSTx Rtmp, (Rarg0)
		p4 := s.Prog(lstx)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = sw64.REGTMP
		p4.To.Type = obj.TYPE_MEM
		p4.To.Reg = v.Args[0].Reg()
		// RD_F Rtmp
		p5 := s.Prog(sw64.ARD_F)
		p5.From.Type = obj.TYPE_REG
		p5.From.Reg = sw64.REGTMP
		// BEQ  Rtmp, -6(PC)
		p6 := s.Prog(sw64.ABEQ)
		p6.From.Type = obj.TYPE_REG
		p6.From.Reg = sw64.REGTMP
		p6.To.Type = obj.TYPE_BRANCH
		gc.Patch(p6, p)
		s.Prog(sw64.AMEMB)
	case ssa.OpSW64LoweredAtomicCas32, ssa.OpSW64LoweredAtomicCas64:
		// MEMB
		// LLDx  Rout, (Rarg0)
		// CMPEQ Rout, Rarg1, Rout
		// WR_F	 Rout
		// BIS   Rarg2, R31, Rtmp
		// LSTx  Rtmp, (Rarg0)
		// RD_F  Rtmp
		// BEQ   Rout, 2(PC)
		// BEQ   Rtmp, -7(PC)
		// MEMB
		lldx := sw64.ALLDL
		lstx := sw64.ALSTL
		if v.Op == ssa.OpSW64LoweredAtomicCas32 {
			lldx = sw64.ALLDW
			lstx = sw64.ALSTW
		}
		s.Prog(sw64.AMEMB)
		// LLDx Rout, (Rarg0)
		p := s.Prog(lldx)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Reg0()
		// CMPEQ Rout, Rarg1, Rout
		p1 := s.Prog(sw64.ACMPEQ)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = v.Reg0()
		p1.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  v.Args[1].Reg(),
		})
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = v.Reg0()
		// WR_F Rout
		p2 := s.Prog(sw64.AWR_F)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = v.Reg0()
		// BIS  Rarg2, R31, Rtmp
		p3 := s.Prog(sw64.ABIS)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = v.Args[2].Reg()
		p3.SetFrom3(obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  sw64.REGZERO,
		})
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = sw64.REGTMP
		// LSTx Rtmp, (Rarg0)
		p4 := s.Prog(lstx)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = sw64.REGTMP
		p4.To.Type = obj.TYPE_MEM
		p4.To.Reg = v.Args[0].Reg()
		// RD_F Rtmp
		p5 := s.Prog(sw64.ARD_F)
		p5.From.Type = obj.TYPE_REG
		p5.From.Reg = sw64.REGTMP
		// BEQ  Rout, 2(PC)
		p6 := s.Prog(sw64.ABEQ)
		p6.From.Type = obj.TYPE_REG
		p6.From.Reg = v.Reg0()
		p6.To.Type = obj.TYPE_BRANCH
		// BEQ  Rtmp, -7(PC)
		p7 := s.Prog(sw64.ABEQ)
		p7.From.Type = obj.TYPE_REG
		p7.From.Reg = sw64.REGTMP
		p7.To.Type = obj.TYPE_BRANCH
		gc.Patch(p7, p)
		p8 := s.Prog(sw64.AMEMB)
		gc.Patch(p6, p8)

	default:
		v.Fatalf("genValue not implemented: %s", v.LongString())
	}
}

var blockJump = map[ssa.BlockKind]struct {
	asm, invasm obj.As
}{
	ssa.BlockSW64NE:  {sw64.ABNE, sw64.ABEQ},
	ssa.BlockSW64EQ:  {sw64.ABEQ, sw64.ABNE},
	ssa.BlockSW64LT:  {sw64.ABLT, sw64.ABGE},
	ssa.BlockSW64LE:  {sw64.ABLE, sw64.ABGT},
	ssa.BlockSW64GT:  {sw64.ABGT, sw64.ABLE},
	ssa.BlockSW64GE:  {sw64.ABGE, sw64.ABLT},
	ssa.BlockSW64FNE: {sw64.AFBNE, sw64.AFBEQ},
	ssa.BlockSW64FEQ: {sw64.AFBEQ, sw64.AFBNE},
	ssa.BlockSW64FLT: {sw64.AFBLT, sw64.AFBGE},
	ssa.BlockSW64FLE: {sw64.AFBLE, sw64.AFBGT},
	ssa.BlockSW64FGT: {sw64.AFBGT, sw64.AFBLE},
	ssa.BlockSW64FGE: {sw64.AFBGE, sw64.AFBLT},
}

func ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockPlain:
		if b.Succs[0].Block() != next {
			p := s.Prog(obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}

	case ssa.BlockDefer:
		// defer returns in R0:
		// 0 if we should continue executing
		// 1 if we should jump to deferreturn call
		//
		// see also runtime/asm_sw64.s:return0
		p := s.Prog(sw64.ABNE)
		p.From = obj.Addr{
			Type: obj.TYPE_REG,
			Reg:  sw64.REG_R0,
		}
		p.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[1].Block()})

		if b.Succs[0].Block() != next {
			p := s.Prog(obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}

	case ssa.BlockExit:
		//zxw new change
		//s.Prog(obj.AUNDEF) // tell plive.go that we never reach here

	case ssa.BlockRet:
		s.Prog(obj.ARET)

	case ssa.BlockRetJmp:
		p := s.Prog(obj.AJMP)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)
		s.Prog(obj.ARET)
	case ssa.BlockSW64NE,
		ssa.BlockSW64EQ,
		ssa.BlockSW64LT,
		ssa.BlockSW64LE,
		ssa.BlockSW64GT,
		ssa.BlockSW64GE,
		ssa.BlockSW64FNE,
		ssa.BlockSW64FEQ,
		ssa.BlockSW64FLT,
		ssa.BlockSW64FLE,
		ssa.BlockSW64FGT,
		ssa.BlockSW64FGE:
		jmp := blockJump[b.Kind]
		var p *obj.Prog
		switch next {
		case b.Succs[0].Block():
			p = s.Prog(jmp.invasm)
			p.To.Type = obj.TYPE_BRANCH
			p.From.Type = obj.TYPE_REG
			p.From.Reg = b.Controls[0].Reg() //zxw new change
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[1].Block()})
		case b.Succs[1].Block():
			p = s.Prog(jmp.asm)
			p.To.Type = obj.TYPE_BRANCH
			p.From.Type = obj.TYPE_REG
			p.From.Reg = b.Controls[0].Reg()
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		default:
			//zxw change
			if b.Likely != ssa.BranchUnlikely {
				p = s.Prog(jmp.asm)
				p.To.Type = obj.TYPE_BRANCH
				p.From.Type = obj.TYPE_REG
				p.From.Reg = b.Controls[0].Reg()
				s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
				q := s.Prog(obj.AJMP)
				q.To.Type = obj.TYPE_BRANCH
				s.Branches = append(s.Branches, gc.Branch{P: q, B: b.Succs[1].Block()})
			} else {
				p = s.Prog(jmp.invasm)
				p.To.Type = obj.TYPE_BRANCH
				p.From.Type = obj.TYPE_REG
				p.From.Reg = b.Controls[0].Reg()
				s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[1].Block()})
				q := s.Prog(obj.AJMP)
				q.To.Type = obj.TYPE_BRANCH
				s.Branches = append(s.Branches, gc.Branch{P: q, B: b.Succs[0].Block()})
			}
		}

	default:
		b.Fatalf("branch not implemented: %s", b.LongString())
	}
}
