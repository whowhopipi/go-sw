// Code generated from gen/SW64.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "cmd/compile/internal/types"

func rewriteValueSW64(v *Value) bool {
	switch v.Op {
	case OpAbs:
		v.Op = OpSW64FABS
		return true
	case OpAdd16:
		v.Op = OpSW64ADDV
		return true
	case OpAdd32:
		v.Op = OpSW64ADDV
		return true
	case OpAdd32F:
		v.Op = OpSW64FADDS
		return true
	case OpAdd64:
		v.Op = OpSW64ADDV
		return true
	case OpAdd64F:
		v.Op = OpSW64FADDD
		return true
	case OpAdd8:
		v.Op = OpSW64ADDV
		return true
	case OpAddPtr:
		v.Op = OpSW64ADDV
		return true
	case OpAddr:
		return rewriteValueSW64_OpAddr(v)
	case OpAnd16:
		v.Op = OpSW64AND
		return true
	case OpAnd32:
		v.Op = OpSW64AND
		return true
	case OpAnd64:
		v.Op = OpSW64AND
		return true
	case OpAnd8:
		v.Op = OpSW64AND
		return true
	case OpAndB:
		v.Op = OpSW64AND
		return true
	case OpAtomicAdd32:
		v.Op = OpSW64LoweredAtomicAdd32
		return true
	case OpAtomicAdd64:
		v.Op = OpSW64LoweredAtomicAdd64
		return true
	case OpAtomicCompareAndSwap32:
		v.Op = OpSW64LoweredAtomicCas32
		return true
	case OpAtomicCompareAndSwap64:
		v.Op = OpSW64LoweredAtomicCas64
		return true
	case OpAtomicExchange32:
		v.Op = OpSW64LoweredAtomicExchange32
		return true
	case OpAtomicExchange64:
		v.Op = OpSW64LoweredAtomicExchange64
		return true
	case OpAtomicLoad32:
		v.Op = OpSW64LoweredAtomicLoad32
		return true
	case OpAtomicLoad64:
		v.Op = OpSW64LoweredAtomicLoad64
		return true
	case OpAtomicLoad8:
		v.Op = OpSW64LoweredAtomicLoad8
		return true
	case OpAtomicLoadPtr:
		v.Op = OpSW64LoweredAtomicLoad64
		return true
	case OpAtomicStore32:
		v.Op = OpSW64LoweredAtomicStore32
		return true
	case OpAtomicStore64:
		v.Op = OpSW64LoweredAtomicStore64
		return true
	case OpAtomicStore8:
		v.Op = OpSW64LoweredAtomicStore8
		return true
	case OpAtomicStorePtrNoWB:
		v.Op = OpSW64LoweredAtomicStore64
		return true
	case OpAvg64u:
		return rewriteValueSW64_OpAvg64u(v)
	case OpCeil:
		v.Op = OpSW64FCVTDL_P
		return true
	case OpClosureCall:
		return rewriteValueSW64_OpClosureCall(v)
	case OpCom16:
		return rewriteValueSW64_OpCom16(v)
	case OpCom32:
		return rewriteValueSW64_OpCom32(v)
	case OpCom64:
		return rewriteValueSW64_OpCom64(v)
	case OpCom8:
		return rewriteValueSW64_OpCom8(v)
	case OpConst16:
		return rewriteValueSW64_OpConst16(v)
	case OpConst32:
		return rewriteValueSW64_OpConst32(v)
	case OpConst32F:
		return rewriteValueSW64_OpConst32F(v)
	case OpConst64:
		return rewriteValueSW64_OpConst64(v)
	case OpConst64F:
		return rewriteValueSW64_OpConst64F(v)
	case OpConst8:
		return rewriteValueSW64_OpConst8(v)
	case OpConstBool:
		return rewriteValueSW64_OpConstBool(v)
	case OpConstNil:
		return rewriteValueSW64_OpConstNil(v)
	case OpConvert:
		v.Op = OpSW64MOVVconvert
		return true
	case OpCopysign:
		return rewriteValueSW64_OpCopysign(v)
	case OpCtz32:
		return rewriteValueSW64_OpCtz32(v)
	case OpCtz64:
		v.Op = OpSW64CTTZ
		return true
	case OpCvt32Fto32:
		return rewriteValueSW64_OpCvt32Fto32(v)
	case OpCvt32Fto64:
		return rewriteValueSW64_OpCvt32Fto64(v)
	case OpCvt32Fto64F:
		v.Op = OpSW64FCVTSD
		return true
	case OpCvt32to32F:
		return rewriteValueSW64_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValueSW64_OpCvt32to64F(v)
	case OpCvt64Fto32:
		return rewriteValueSW64_OpCvt64Fto32(v)
	case OpCvt64Fto32F:
		v.Op = OpSW64FCVTDS
		return true
	case OpCvt64Fto64:
		return rewriteValueSW64_OpCvt64Fto64(v)
	case OpCvt64to32F:
		v.Op = OpSW64FCVTLS
		return true
	case OpCvt64to64F:
		v.Op = OpSW64FCVTLD
		return true
	case OpDiv16:
		return rewriteValueSW64_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueSW64_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueSW64_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpSW64FDIVS
		return true
	case OpDiv32u:
		return rewriteValueSW64_OpDiv32u(v)
	case OpDiv64:
		return rewriteValueSW64_OpDiv64(v)
	case OpDiv64F:
		v.Op = OpSW64FDIVD
		return true
	case OpDiv64u:
		return rewriteValueSW64_OpDiv64u(v)
	case OpDiv8:
		return rewriteValueSW64_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueSW64_OpDiv8u(v)
	case OpEq16:
		return rewriteValueSW64_OpEq16(v)
	case OpEq32:
		return rewriteValueSW64_OpEq32(v)
	case OpEq32F:
		return rewriteValueSW64_OpEq32F(v)
	case OpEq64:
		v.Op = OpSW64CMPEQ
		return true
	case OpEq64F:
		return rewriteValueSW64_OpEq64F(v)
	case OpEq8:
		return rewriteValueSW64_OpEq8(v)
	case OpEqB:
		return rewriteValueSW64_OpEqB(v)
	case OpEqPtr:
		v.Op = OpSW64CMPEQ
		return true
	case OpFloor:
		v.Op = OpSW64FCVTDL_N
		return true
	case OpGetCallerPC:
		v.Op = OpSW64LoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpSW64LoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpSW64LoweredGetClosurePtr
		return true
	case OpHmul32:
		return rewriteValueSW64_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueSW64_OpHmul32u(v)
	case OpHmul64:
		return rewriteValueSW64_OpHmul64(v)
	case OpHmul64u:
		v.Op = OpSW64UMULH
		return true
	case OpInterCall:
		return rewriteValueSW64_OpInterCall(v)
	case OpIsInBounds:
		v.Op = OpSW64CMPULT
		return true
	case OpIsNonNil:
		return rewriteValueSW64_OpIsNonNil(v)
	case OpIsSliceInBounds:
		v.Op = OpSW64CMPULE
		return true
	case OpLeq16:
		return rewriteValueSW64_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueSW64_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueSW64_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueSW64_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueSW64_OpLeq32U(v)
	case OpLeq64:
		v.Op = OpSW64CMPLE
		return true
	case OpLeq64F:
		return rewriteValueSW64_OpLeq64F(v)
	case OpLeq64U:
		v.Op = OpSW64CMPULE
		return true
	case OpLeq8:
		return rewriteValueSW64_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueSW64_OpLeq8U(v)
	case OpLess16:
		return rewriteValueSW64_OpLess16(v)
	case OpLess16U:
		return rewriteValueSW64_OpLess16U(v)
	case OpLess32:
		return rewriteValueSW64_OpLess32(v)
	case OpLess32F:
		return rewriteValueSW64_OpLess32F(v)
	case OpLess32U:
		return rewriteValueSW64_OpLess32U(v)
	case OpLess64:
		v.Op = OpSW64CMPLT
		return true
	case OpLess64F:
		return rewriteValueSW64_OpLess64F(v)
	case OpLess64U:
		v.Op = OpSW64CMPULT
		return true
	case OpLess8:
		return rewriteValueSW64_OpLess8(v)
	case OpLess8U:
		return rewriteValueSW64_OpLess8U(v)
	case OpLoad:
		return rewriteValueSW64_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueSW64_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueSW64_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueSW64_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueSW64_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueSW64_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueSW64_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueSW64_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueSW64_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueSW64_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValueSW64_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValueSW64_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValueSW64_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValueSW64_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValueSW64_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueSW64_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueSW64_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueSW64_OpLsh8x8(v)
	case OpMod16:
		return rewriteValueSW64_OpMod16(v)
	case OpMod16u:
		return rewriteValueSW64_OpMod16u(v)
	case OpMod32:
		return rewriteValueSW64_OpMod32(v)
	case OpMod32u:
		return rewriteValueSW64_OpMod32u(v)
	case OpMod64:
		return rewriteValueSW64_OpMod64(v)
	case OpMod64u:
		return rewriteValueSW64_OpMod64u(v)
	case OpMod8:
		return rewriteValueSW64_OpMod8(v)
	case OpMod8u:
		return rewriteValueSW64_OpMod8u(v)
	case OpMove:
		return rewriteValueSW64_OpMove(v)
	case OpMul16:
		v.Op = OpSW64MULW
		return true
	case OpMul32:
		v.Op = OpSW64MULW
		return true
	case OpMul32F:
		v.Op = OpSW64FMULS
		return true
	case OpMul64:
		v.Op = OpSW64MULL
		return true
	case OpMul64F:
		v.Op = OpSW64FMULD
		return true
	case OpMul8:
		v.Op = OpSW64MULW
		return true
	case OpNeg16:
		v.Op = OpSW64NEGV
		return true
	case OpNeg32:
		v.Op = OpSW64NEGV
		return true
	case OpNeg32F:
		v.Op = OpSW64NEGF
		return true
	case OpNeg64:
		v.Op = OpSW64NEGV
		return true
	case OpNeg64F:
		v.Op = OpSW64NEGD
		return true
	case OpNeg8:
		v.Op = OpSW64NEGV
		return true
	case OpNeq16:
		return rewriteValueSW64_OpNeq16(v)
	case OpNeq32:
		return rewriteValueSW64_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueSW64_OpNeq32F(v)
	case OpNeq64:
		return rewriteValueSW64_OpNeq64(v)
	case OpNeq64F:
		return rewriteValueSW64_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueSW64_OpNeq8(v)
	case OpNeqB:
		v.Op = OpSW64XOR
		return true
	case OpNeqPtr:
		return rewriteValueSW64_OpNeqPtr(v)
	case OpNilCheck:
		v.Op = OpSW64LoweredNilCheck
		return true
	case OpNot:
		return rewriteValueSW64_OpNot(v)
	case OpOffPtr:
		return rewriteValueSW64_OpOffPtr(v)
	case OpOr16:
		v.Op = OpSW64BIS
		return true
	case OpOr32:
		v.Op = OpSW64BIS
		return true
	case OpOr64:
		v.Op = OpSW64BIS
		return true
	case OpOr8:
		v.Op = OpSW64BIS
		return true
	case OpOrB:
		v.Op = OpSW64BIS
		return true
	case OpPanicBounds:
		return rewriteValueSW64_OpPanicBounds(v)
	case OpPopCount16:
		return rewriteValueSW64_OpPopCount16(v)
	case OpPopCount32:
		return rewriteValueSW64_OpPopCount32(v)
	case OpPopCount64:
		v.Op = OpSW64CTPOP
		return true
	case OpPopCount8:
		return rewriteValueSW64_OpPopCount8(v)
	case OpRotateLeft16:
		return rewriteValueSW64_OpRotateLeft16(v)
	case OpRotateLeft32:
		return rewriteValueSW64_OpRotateLeft32(v)
	case OpRotateLeft64:
		return rewriteValueSW64_OpRotateLeft64(v)
	case OpRotateLeft8:
		return rewriteValueSW64_OpRotateLeft8(v)
	case OpRound:
		v.Op = OpSW64FCVTDL_G
		return true
	case OpRound32F:
		v.Op = OpCopy
		return true
	case OpRound64F:
		v.Op = OpCopy
		return true
	case OpRsh16Ux16:
		return rewriteValueSW64_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueSW64_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueSW64_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueSW64_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueSW64_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueSW64_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueSW64_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueSW64_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueSW64_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueSW64_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueSW64_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueSW64_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueSW64_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueSW64_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueSW64_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueSW64_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValueSW64_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValueSW64_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValueSW64_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValueSW64_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValueSW64_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValueSW64_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValueSW64_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValueSW64_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValueSW64_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueSW64_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueSW64_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueSW64_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueSW64_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueSW64_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueSW64_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueSW64_OpRsh8x8(v)
	case OpSW64ADDV:
		return rewriteValueSW64_OpSW64ADDV(v)
	case OpSW64ADDVconst:
		return rewriteValueSW64_OpSW64ADDVconst(v)
	case OpSW64ADDWconst:
		return rewriteValueSW64_OpSW64ADDWconst(v)
	case OpSW64AND:
		return rewriteValueSW64_OpSW64AND(v)
	case OpSW64ANDconst:
		return rewriteValueSW64_OpSW64ANDconst(v)
	case OpSW64BIS:
		return rewriteValueSW64_OpSW64BIS(v)
	case OpSW64BISconst:
		return rewriteValueSW64_OpSW64BISconst(v)
	case OpSW64CMPEQ:
		return rewriteValueSW64_OpSW64CMPEQ(v)
	case OpSW64CMPLE:
		return rewriteValueSW64_OpSW64CMPLE(v)
	case OpSW64CMPLT:
		return rewriteValueSW64_OpSW64CMPLT(v)
	case OpSW64CMPULE:
		return rewriteValueSW64_OpSW64CMPULE(v)
	case OpSW64CMPULT:
		return rewriteValueSW64_OpSW64CMPULT(v)
	case OpSW64MOVBUload:
		return rewriteValueSW64_OpSW64MOVBUload(v)
	case OpSW64MOVBload:
		return rewriteValueSW64_OpSW64MOVBload(v)
	case OpSW64MOVBstore:
		return rewriteValueSW64_OpSW64MOVBstore(v)
	case OpSW64MOVBstorezero:
		return rewriteValueSW64_OpSW64MOVBstorezero(v)
	case OpSW64MOVDload:
		return rewriteValueSW64_OpSW64MOVDload(v)
	case OpSW64MOVDstore:
		return rewriteValueSW64_OpSW64MOVDstore(v)
	case OpSW64MOVFload:
		return rewriteValueSW64_OpSW64MOVFload(v)
	case OpSW64MOVFstore:
		return rewriteValueSW64_OpSW64MOVFstore(v)
	case OpSW64MOVHUload:
		return rewriteValueSW64_OpSW64MOVHUload(v)
	case OpSW64MOVHload:
		return rewriteValueSW64_OpSW64MOVHload(v)
	case OpSW64MOVHstore:
		return rewriteValueSW64_OpSW64MOVHstore(v)
	case OpSW64MOVHstorezero:
		return rewriteValueSW64_OpSW64MOVHstorezero(v)
	case OpSW64MOVVload:
		return rewriteValueSW64_OpSW64MOVVload(v)
	case OpSW64MOVVstore:
		return rewriteValueSW64_OpSW64MOVVstore(v)
	case OpSW64MOVVstorezero:
		return rewriteValueSW64_OpSW64MOVVstorezero(v)
	case OpSW64MOVWUload:
		return rewriteValueSW64_OpSW64MOVWUload(v)
	case OpSW64MOVWload:
		return rewriteValueSW64_OpSW64MOVWload(v)
	case OpSW64MOVWstore:
		return rewriteValueSW64_OpSW64MOVWstore(v)
	case OpSW64MOVWstorezero:
		return rewriteValueSW64_OpSW64MOVWstorezero(v)
	case OpSW64MULL:
		return rewriteValueSW64_OpSW64MULL(v)
	case OpSW64MULW:
		return rewriteValueSW64_OpSW64MULW(v)
	case OpSW64NEGV:
		return rewriteValueSW64_OpSW64NEGV(v)
	case OpSW64ORNOT:
		return rewriteValueSW64_OpSW64ORNOT(v)
	case OpSW64SEXTB:
		return rewriteValueSW64_OpSW64SEXTB(v)
	case OpSW64SEXTH:
		return rewriteValueSW64_OpSW64SEXTH(v)
	case OpSW64SLLconst:
		return rewriteValueSW64_OpSW64SLLconst(v)
	case OpSW64SRAconst:
		return rewriteValueSW64_OpSW64SRAconst(v)
	case OpSW64SRLconst:
		return rewriteValueSW64_OpSW64SRLconst(v)
	case OpSW64SUBV:
		return rewriteValueSW64_OpSW64SUBV(v)
	case OpSW64SUBVconst:
		return rewriteValueSW64_OpSW64SUBVconst(v)
	case OpSW64XOR:
		return rewriteValueSW64_OpSW64XOR(v)
	case OpSW64XORconst:
		return rewriteValueSW64_OpSW64XORconst(v)
	case OpSignExt16to32:
		v.Op = OpSW64SEXTH
		return true
	case OpSignExt16to64:
		v.Op = OpSW64SEXTH
		return true
	case OpSignExt32to64:
		return rewriteValueSW64_OpSignExt32to64(v)
	case OpSignExt8to16:
		v.Op = OpSW64SEXTB
		return true
	case OpSignExt8to32:
		v.Op = OpSW64SEXTB
		return true
	case OpSignExt8to64:
		v.Op = OpSW64SEXTB
		return true
	case OpSignmask:
		return rewriteValueSW64_OpSignmask(v)
	case OpSlicemask:
		return rewriteValueSW64_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpSW64FSQRTD
		return true
	case OpStaticCall:
		return rewriteValueSW64_OpStaticCall(v)
	case OpStore:
		return rewriteValueSW64_OpStore(v)
	case OpSub16:
		v.Op = OpSW64SUBV
		return true
	case OpSub32:
		v.Op = OpSW64SUBV
		return true
	case OpSub32F:
		v.Op = OpSW64FSUBS
		return true
	case OpSub64:
		v.Op = OpSW64SUBV
		return true
	case OpSub64F:
		v.Op = OpSW64FSUBD
		return true
	case OpSub8:
		v.Op = OpSW64SUBV
		return true
	case OpSubPtr:
		v.Op = OpSW64SUBV
		return true
	case OpTrunc:
		v.Op = OpSW64FCVTDL_Z
		return true
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpTrunc64to16:
		v.Op = OpCopy
		return true
	case OpTrunc64to32:
		v.Op = OpCopy
		return true
	case OpTrunc64to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpSW64LoweredWB
		return true
	case OpXor16:
		v.Op = OpSW64XOR
		return true
	case OpXor32:
		v.Op = OpSW64XOR
		return true
	case OpXor64:
		v.Op = OpSW64XOR
		return true
	case OpXor8:
		v.Op = OpSW64XOR
		return true
	case OpZero:
		return rewriteValueSW64_OpZero(v)
	case OpZeroExt16to32:
		return rewriteValueSW64_OpZeroExt16to32(v)
	case OpZeroExt16to64:
		return rewriteValueSW64_OpZeroExt16to64(v)
	case OpZeroExt32to64:
		return rewriteValueSW64_OpZeroExt32to64(v)
	case OpZeroExt8to16:
		return rewriteValueSW64_OpZeroExt8to16(v)
	case OpZeroExt8to32:
		return rewriteValueSW64_OpZeroExt8to32(v)
	case OpZeroExt8to64:
		return rewriteValueSW64_OpZeroExt8to64(v)
	}
	return false
}
func rewriteValueSW64_OpAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Addr {sym} base)
	// result: (SYMADDR {sym} base)
	for {
		sym := auxToSym(v.Aux)
		base := v_0
		v.reset(OpSW64SYMADDR)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
}
func rewriteValueSW64_OpAvg64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg64u <t> x y)
	// result: (ADDV (SRLconst <t> (SUBV <t> x y) [1]) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64ADDV)
		v0 := b.NewValue0(v.Pos, OpSW64SRLconst, t)
		v0.AuxInt = int64ToAuxInt(1)
		v1 := b.NewValue0(v.Pos, OpSW64SUBV, t)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueSW64_OpClosureCall(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ClosureCall [argwid] entry closure mem)
	// result: (CALLclosure [argwid] entry closure mem)
	for {
		argwid := auxIntToInt32(v.AuxInt)
		entry := v_0
		closure := v_1
		mem := v_2
		v.reset(OpSW64CALLclosure)
		v.AuxInt = int64ToAuxInt(argwid)
		v.AddArg3(entry, closure, mem)
		return true
	}
}
func rewriteValueSW64_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com16 x)
	// result: (ORNOT (MOVVconst [0]) x)
	for {
		x := v_0
		v.reset(OpSW64ORNOT)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueSW64_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com32 x)
	// result: (ORNOT (MOVVconst [0]) x)
	for {
		x := v_0
		v.reset(OpSW64ORNOT)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueSW64_OpCom64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com64 x)
	// result: (ORNOT (MOVVconst [0]) x)
	for {
		x := v_0
		v.reset(OpSW64ORNOT)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueSW64_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com8 x)
	// result: (ORNOT (MOVVconst [0]) x)
	for {
		x := v_0
		v.reset(OpSW64ORNOT)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueSW64_OpConst16(v *Value) bool {
	// match: (Const16 [val])
	// result: (MOVVconst [int64(val)])
	for {
		val := auxIntToInt16(v.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSW64_OpConst32(v *Value) bool {
	// match: (Const32 [val])
	// result: (MOVVconst [int64(val)])
	for {
		val := auxIntToInt32(v.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSW64_OpConst32F(v *Value) bool {
	// match: (Const32F [val])
	// result: (MOVFconst [int64(val)])
	for {
		val := auxIntToFloat32(v.AuxInt)
		v.reset(OpSW64MOVFconst)
		v.AuxInt = float32ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSW64_OpConst64(v *Value) bool {
	// match: (Const64 [val])
	// result: (MOVVconst [int64(val)])
	for {
		val := auxIntToInt64(v.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSW64_OpConst64F(v *Value) bool {
	// match: (Const64F [val])
	// result: (MOVDconst [int64(val)])
	for {
		val := auxIntToFloat64(v.AuxInt)
		v.reset(OpSW64MOVDconst)
		v.AuxInt = float64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSW64_OpConst8(v *Value) bool {
	// match: (Const8 [val])
	// result: (MOVVconst [int64(val)])
	for {
		val := auxIntToInt8(v.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSW64_OpConstBool(v *Value) bool {
	// match: (ConstBool [b])
	// result: (MOVVconst [b])
	for {
		b := auxIntToBool(v.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(b)
		return true
	}
}
func rewriteValueSW64_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVVconst [0])
	for {
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
}
func rewriteValueSW64_OpCopysign(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Copysign x y)
	// result: (FCPYS y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64FCPYS)
		v.AddArg2(y, x)
		return true
	}
}
func rewriteValueSW64_OpCtz32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz32 <t> x)
	// result: (SUBV (CTTZ (ZeroExt32to64 x)) (MULLconst <t> (CMPEQ (ZeroExt32to64 x) (MOVVconst [0]) ) [32]))
	for {
		t := v.Type
		x := v_0
		v.reset(OpSW64SUBV)
		v0 := b.NewValue0(v.Pos, OpSW64CTTZ, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSW64MULLconst, t)
		v2.AuxInt = int64ToAuxInt(32)
		v3 := b.NewValue0(v.Pos, OpSW64CMPEQ, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(0)
		v3.AddArg2(v1, v4)
		v2.AddArg(v3)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueSW64_OpCvt32Fto32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32Fto32 x)
	// result: (FCVTLW (FCVTDL_Z (FCVTSD x)))
	for {
		x := v_0
		v.reset(OpSW64FCVTLW)
		v0 := b.NewValue0(v.Pos, OpSW64FCVTDL_Z, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpSW64FCVTSD, typ.Float64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpCvt32Fto64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32Fto64 x)
	// result: (FCVTDL_Z (FCVTSD x))
	for {
		x := v_0
		v.reset(OpSW64FCVTDL_Z)
		v0 := b.NewValue0(v.Pos, OpSW64FCVTSD, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpCvt32to32F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to32F x)
	// result: (FCVTLS (IFMOVD x))
	for {
		x := v_0
		v.reset(OpSW64FCVTLS)
		v0 := b.NewValue0(v.Pos, OpSW64IFMOVD, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpCvt32to64F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to64F x)
	// result: (FCVTLD (IFMOVD (SignExt32to64 x)))
	for {
		x := v_0
		v.reset(OpSW64FCVTLD)
		v0 := b.NewValue0(v.Pos, OpSW64IFMOVD, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpCvt64Fto32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt64Fto32 x)
	// result: (FCVTLW (FCVTDL_Z x))
	for {
		x := v_0
		v.reset(OpSW64FCVTLW)
		v0 := b.NewValue0(v.Pos, OpSW64FCVTDL_Z, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpCvt64Fto64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt64Fto64 x)
	// result: (FIMOVD (FCVTDL_Z x))
	for {
		x := v_0
		v.reset(OpSW64FIMOVD)
		v0 := b.NewValue0(v.Pos, OpSW64FCVTDL_Z, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (Div64 (SignExt16to64 x) (SignExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv64)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (Div64u (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv64u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 x y)
	// result: (Div64 (SignExt32to64 x) (SignExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv64)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (Div64u (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv64u)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpDiv64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div64 x y)
	// result: (SUBV (XOR <typ.UInt64> (Select0 <typ.UInt64> (CALLudiv (SUBV <typ.UInt64> (XOR x <typ.UInt64> (Signmask x)) (Signmask x)) (SUBV <typ.UInt64> (XOR y <typ.UInt64> (Signmask y)) (Signmask y)))) (Signmask (XOR <typ.UInt64> x y))) (Signmask (XOR <typ.UInt64> x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64SUBV)
		v0 := b.NewValue0(v.Pos, OpSW64XOR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpSelect0, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSW64CALLudiv, types.NewTuple(typ.UInt64, typ.UInt64))
		v3 := b.NewValue0(v.Pos, OpSW64SUBV, typ.UInt64)
		v4 := b.NewValue0(v.Pos, OpSW64XOR, typ.UInt64)
		v5 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v5.AddArg(x)
		v4.AddArg2(x, v5)
		v3.AddArg2(v4, v5)
		v6 := b.NewValue0(v.Pos, OpSW64SUBV, typ.UInt64)
		v7 := b.NewValue0(v.Pos, OpSW64XOR, typ.UInt64)
		v8 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v8.AddArg(y)
		v7.AddArg2(y, v8)
		v6.AddArg2(v7, v8)
		v2.AddArg2(v3, v6)
		v1.AddArg(v2)
		v9 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v10 := b.NewValue0(v.Pos, OpSW64XOR, typ.UInt64)
		v10.AddArg2(x, y)
		v9.AddArg(v10)
		v0.AddArg2(v1, v9)
		v.AddArg2(v0, v9)
		return true
	}
}
func rewriteValueSW64_OpDiv64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div64u x y)
	// result: (Select0 <typ.UInt64> (CALLudiv x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpSW64CALLudiv, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (Div64 (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv64)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (Div64u (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv64u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (CMPEQ (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPEQ)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32 x y)
	// result: (CMPEQ (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPEQ)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32F x y)
	// result: (FEqual (FCMPEQ x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64FEqual)
		v0 := b.NewValue0(v.Pos, OpSW64FCMPEQ, typ.Float64)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq64F x y)
	// result: (FEqual (FCMPEQ x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64FEqual)
		v0 := b.NewValue0(v.Pos, OpSW64FCMPEQ, typ.Float64)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (CMPEQ (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPEQ)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (XOR (MOVVconst [1]) (XOR <typ.Bool> x y) )
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64XOR)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(1)
		v1 := b.NewValue0(v.Pos, OpSW64XOR, typ.Bool)
		v1.AddArg2(x, y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpHmul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32 x y)
	// result: (SRAconst (MULL <typ.Int64> (SignExt32to64 x) (SignExt32to64 y)) [32])
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(32)
		v0 := b.NewValue0(v.Pos, OpSW64MULL, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpHmul32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32u x y)
	// result: (SRLconst (MULL <typ.UInt64> (ZeroExt32to64 x) (ZeroExt32to64 y)) [32])
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64SRLconst)
		v.AuxInt = int64ToAuxInt(32)
		v0 := b.NewValue0(v.Pos, OpSW64MULL, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpHmul64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul64 x y)
	// result: ( SUBV <typ.Int64> (SUBV <typ.Int64> (UMULH <typ.Int64> x y) (MULL <typ.Int64> (SRLconst x [63]) y)) (MULL <typ.Int64> (SRLconst y [63]) x) )
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64SUBV)
		v.Type = typ.Int64
		v0 := b.NewValue0(v.Pos, OpSW64SUBV, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpSW64UMULH, typ.Int64)
		v1.AddArg2(x, y)
		v2 := b.NewValue0(v.Pos, OpSW64MULL, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpSW64SRLconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(63)
		v3.AddArg(x)
		v2.AddArg2(v3, y)
		v0.AddArg2(v1, v2)
		v4 := b.NewValue0(v.Pos, OpSW64MULL, typ.Int64)
		v5 := b.NewValue0(v.Pos, OpSW64SRLconst, typ.UInt64)
		v5.AuxInt = int64ToAuxInt(63)
		v5.AddArg(y)
		v4.AddArg2(v5, x)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpInterCall(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (InterCall [argwid] entry mem)
	// result: (CALLinter [argwid] entry mem)
	for {
		argwid := auxIntToInt32(v.AuxInt)
		entry := v_0
		mem := v_1
		v.reset(OpSW64CALLinter)
		v.AuxInt = int64ToAuxInt(argwid)
		v.AddArg2(entry, mem)
		return true
	}
}
func rewriteValueSW64_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsNonNil ptr)
	// result: (CMPULT (MOVVconst [0]) ptr)
	for {
		ptr := v_0
		v.reset(OpSW64CMPULT)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, ptr)
		return true
	}
}
func rewriteValueSW64_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (CMPLE (SignExt16to64 x) (SignExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPLE)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (CMPULE (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPULE)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32 x y)
	// result: (CMPLE (SignExt32to64 x) (SignExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPLE)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32F x y)
	// result: (FEqual (FCMPLE x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64FEqual)
		v0 := b.NewValue0(v.Pos, OpSW64FCMPLE, typ.Float64)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32U x y)
	// result: (CMPULE (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPULE)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq64F x y)
	// result: (FEqual (FCMPLE x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64FEqual)
		v0 := b.NewValue0(v.Pos, OpSW64FCMPLE, typ.Float64)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (CMPLE (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPLE)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (CMPULE (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPULE)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (CMPLT (SignExt16to64 x) (SignExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPLT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (CMPULT (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPULT)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32 x y)
	// result: (CMPLT (SignExt32to64 x) (SignExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPLT)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32F x y)
	// result: (FEqual (FCMPLT x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64FEqual)
		v0 := b.NewValue0(v.Pos, OpSW64FCMPLT, typ.Float64)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32U x y)
	// result: (CMPULT (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPULT)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less64F x y)
	// result: (FEqual (FCMPLT x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64FEqual)
		v0 := b.NewValue0(v.Pos, OpSW64FCMPLT, typ.Float64)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (CMPLT (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPLT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (CMPULT (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64CMPULT)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpSW64MOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && isSigned(t))
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpSW64MOVBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && !isSigned(t))
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpSW64MOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && isSigned(t))
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpSW64MOVHload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && !isSigned(t))
	// result: (MOVHUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpSW64MOVHUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && isSigned(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpSW64MOVWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && !isSigned(t))
	// result: (MOVWUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpSW64MOVWUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVVload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpSW64MOVVload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (MOVFload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpSW64MOVFload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpSW64MOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpLocalAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LocalAddr {sym} base _)
	// result: (SYMADDR {sym} base)
	for {
		sym := auxToSym(v.Aux)
		base := v_0
		v.reset(OpSW64SYMADDR)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
}
func rewriteValueSW64_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SLLconst (ZeroExt16to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpSW64SLLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Lsh16x64 x (MOVVconst [c]))
	// cond: uint64(c) < 16
	// result: (SLLconst (ZeroExt16to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpSW64SLLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Lsh16x64 <t> x y)
	// result: (AND (NEGV <t> (CMPULT y (MOVVconst <typ.UInt64> [64]))) (SLL <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v2.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(y, v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSW64_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SLLconst (ZeroExt32to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpSW64SLLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Lsh32x64 x (MOVVconst [c]))
	// cond: uint64(c) < 32
	// result: (SLLconst (ZeroExt32to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpSW64SLLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Lsh32x64 <t> x y)
	// result: (AND (NEGV <t> (CMPULT y (MOVVconst <typ.UInt64> [64]))) (SLL <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v2.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(y, v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSW64_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x16 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x32 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x64 x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SLLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSW64SLLconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x64 x (MOVVconst [c]))
	// cond: uint64(c) < 64
	// result: (SLLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSW64SLLconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x64 <t> x y)
	// result: (AND (NEGV <t> (CMPULT y (MOVVconst <typ.UInt64> [64]))) (SLL <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v2.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(y, v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSW64_OpLsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x8 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SLLconst (ZeroExt8to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpSW64SLLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Lsh8x64 x (MOVVconst [c]))
	// cond: uint64(c) < 8
	// result: (SLLconst (ZeroExt8to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpSW64SLLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Lsh8x64 <t> x y)
	// result: (AND (NEGV <t> (CMPULT y (MOVVconst <typ.UInt64> [64]))) (SLL <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v2.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(y, v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSW64_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [64]))) (SLL <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SLL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (Mod64 (SignExt16to64 x) (SignExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Mod64u (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 x y)
	// result: (Mod64 (SignExt32to64 x) (SignExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (Mod64u (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64u)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpMod64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64 x y)
	// result: (SUBV (XOR <typ.UInt64> (Select1 <typ.UInt64> (CALLudiv (SUBV <typ.UInt64> (XOR <typ.UInt64> x (Signmask x)) (Signmask x)) (SUBV <typ.UInt64> (XOR <typ.UInt64> y (Signmask y)) (Signmask y)))) (Signmask x)) (Signmask x))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64SUBV)
		v0 := b.NewValue0(v.Pos, OpSW64XOR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpSelect1, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSW64CALLudiv, types.NewTuple(typ.UInt64, typ.UInt64))
		v3 := b.NewValue0(v.Pos, OpSW64SUBV, typ.UInt64)
		v4 := b.NewValue0(v.Pos, OpSW64XOR, typ.UInt64)
		v5 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v5.AddArg(x)
		v4.AddArg2(x, v5)
		v3.AddArg2(v4, v5)
		v6 := b.NewValue0(v.Pos, OpSW64SUBV, typ.UInt64)
		v7 := b.NewValue0(v.Pos, OpSW64XOR, typ.UInt64)
		v8 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v8.AddArg(y)
		v7.AddArg2(y, v8)
		v6.AddArg2(v7, v8)
		v2.AddArg2(v3, v6)
		v1.AddArg(v2)
		v0.AddArg2(v1, v5)
		v.AddArg2(v0, v5)
		return true
	}
}
func rewriteValueSW64_OpMod64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64u x y)
	// result: (Select1 <typ.UInt64> (CALLudiv x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpSW64CALLudiv, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Mod64 (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Mod64u (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpSW64MOVBstore)
		v0 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v0 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVBstore [1] dst (MOVBload [1] src mem) (MOVBstore dst (MOVBload src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpSW64MOVBstore)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v0.AuxInt = int32ToAuxInt(1)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v0 := b.NewValue0(v.Pos, OpSW64MOVWload, typ.Int32)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v0.AuxInt = int32ToAuxInt(2)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVBstore [3] dst (MOVBload [3] src mem) (MOVBstore [2] dst (MOVBload [2] src mem) (MOVBstore [1] dst (MOVBload [1] src mem) (MOVBstore dst (MOVBload src mem) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpSW64MOVBstore)
		v.AuxInt = int32ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v0.AuxInt = int32ToAuxInt(3)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v2.AuxInt = int32ToAuxInt(2)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(1)
		v4 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v4.AuxInt = int32ToAuxInt(1)
		v4.AddArg2(src, mem)
		v5 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v6.AddArg2(src, mem)
		v5.AddArg3(dst, v6, mem)
		v3.AddArg3(dst, v4, v5)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore dst (MOVVload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVload, typ.UInt64)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpSW64MOVWload, typ.Int32)
		v0.AuxInt = int32ToAuxInt(4)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVWstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpSW64MOVWload, typ.Int32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [6] dst (MOVHload [6] src mem) (MOVHstore [4] dst (MOVHload [4] src mem) (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AuxInt = int32ToAuxInt(6)
		v0 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v0.AuxInt = int32ToAuxInt(6)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(4)
		v2 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v2.AuxInt = int32ToAuxInt(4)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(2)
		v4 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v4.AuxInt = int32ToAuxInt(2)
		v4.AddArg2(src, mem)
		v5 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v6.AddArg2(src, mem)
		v5.AddArg3(dst, v6, mem)
		v3.AddArg3(dst, v4, v5)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// result: (MOVBstore [2] dst (MOVBload [2] src mem) (MOVBstore [1] dst (MOVBload [1] src mem) (MOVBstore dst (MOVBload src mem) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpSW64MOVBstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v0.AuxInt = int32ToAuxInt(2)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(1)
		v2 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v2.AuxInt = int32ToAuxInt(1)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpSW64MOVBload, typ.Int8)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [6] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [4] dst (MOVHload [4] src mem) (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 6 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v0.AuxInt = int32ToAuxInt(4)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v2.AuxInt = int32ToAuxInt(2)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpSW64MOVHload, typ.Int16)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [12] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [8] dst (MOVWload [8] src mem) (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 12 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v.AuxInt = int32ToAuxInt(8)
		v0 := b.NewValue0(v.Pos, OpSW64MOVWload, typ.Int32)
		v0.AuxInt = int32ToAuxInt(8)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVWstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(4)
		v2 := b.NewValue0(v.Pos, OpSW64MOVWload, typ.Int32)
		v2.AuxInt = int32ToAuxInt(4)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpSW64MOVWstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpSW64MOVWload, typ.Int32)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [16] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore [8] dst (MOVVload [8] src mem) (MOVVstore dst (MOVVload src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 16 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v.AuxInt = int32ToAuxInt(8)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVload, typ.UInt64)
		v0.AuxInt = int32ToAuxInt(8)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVVstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVload, typ.UInt64)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [24] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore [16] dst (MOVVload [16] src mem) (MOVVstore [8] dst (MOVVload [8] src mem) (MOVVstore dst (MOVVload src mem) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 24 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v.AuxInt = int32ToAuxInt(16)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVload, typ.UInt64)
		v0.AuxInt = int32ToAuxInt(16)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpSW64MOVVstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(8)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVload, typ.UInt64)
		v2.AuxInt = int32ToAuxInt(8)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVload, typ.UInt64)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: s > 24 || t.(*types.Type).Alignment()%8 != 0
	// result: (LoweredMove [t.(*types.Type).Alignment()] dst src (ADDVconst <src.Type> src [s-moveSize(t.(*types.Type).Alignment(), config)]) mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 24 || t.(*types.Type).Alignment()%8 != 0) {
			break
		}
		v.reset(OpSW64LoweredMove)
		v.AuxInt = int64ToAuxInt(t.(*types.Type).Alignment())
		v0 := b.NewValue0(v.Pos, OpSW64ADDVconst, src.Type)
		v0.AuxInt = int64ToAuxInt(s - moveSize(t.(*types.Type).Alignment(), config))
		v0.AddArg(src)
		v.AddArg4(dst, src, v0, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (Not (CMPEQ (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpSW64CMPEQ, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32 x y)
	// result: (Not (CMPEQ (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpSW64CMPEQ, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32F x y)
	// result: (FNotEqual (FCMPEQ x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64FNotEqual)
		v0 := b.NewValue0(v.Pos, OpSW64FCMPEQ, typ.Float64)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpNeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq64 x y)
	// result: (Not (CMPEQ x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpSW64CMPEQ, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq64F x y)
	// result: (FNotEqual (FCMPEQ x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSW64FNotEqual)
		v0 := b.NewValue0(v.Pos, OpSW64FCMPEQ, typ.Float64)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (Not (CMPEQ (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpSW64CMPEQ, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqPtr x y)
	// result: (Not (CMPEQ x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpSW64CMPEQ, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not x)
	// result: (XORconst x [1])
	for {
		x := v_0
		v.reset(OpSW64XORconst)
		v.AuxInt = int64ToAuxInt(1)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSW64_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (OffPtr [off] ptr:(SP))
	// result: (SYMADDR [off] ptr)
	for {
		off := auxIntToInt64(v.AuxInt)
		ptr := v_0
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpSW64SYMADDR)
		v.AuxInt = int32ToAuxInt(off)
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADDVconst [off] ptr)
	for {
		off := auxIntToInt64(v.AuxInt)
		ptr := v_0
		v.reset(OpSW64ADDVconst)
		v.AuxInt = int64ToAuxInt(off)
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueSW64_OpPanicBounds(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicBoundsA [kind] x y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpSW64LoweredPanicBoundsA)
		v.AuxInt = int64ToAuxInt(kind)
		v.AddArg3(x, y, mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicBoundsB [kind] x y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpSW64LoweredPanicBoundsB)
		v.AuxInt = int64ToAuxInt(kind)
		v.AddArg3(x, y, mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicBoundsC [kind] x y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpSW64LoweredPanicBoundsC)
		v.AuxInt = int64ToAuxInt(kind)
		v.AddArg3(x, y, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpPopCount16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount16 x)
	// result: (CTPOP (ZeroExt16to64 x))
	for {
		x := v_0
		v.reset(OpSW64CTPOP)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpPopCount32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount32 x)
	// result: (CTPOP (ZeroExt32to64 x))
	for {
		x := v_0
		v.reset(OpSW64CTPOP)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpPopCount8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount8 x)
	// result: (CTPOP (ZeroExt8to64 x))
	for {
		x := v_0
		v.reset(OpSW64CTPOP)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVVconst [c]))
	// result: (Or16 (Lsh16x64 <t> x (MOVVconst [c&15])) (Rsh16Ux64 <t> x (MOVVconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x64, t)
		v1 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v1.AuxInt = int64ToAuxInt(c & 15)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux64, t)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(-c & 15)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueSW64_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft32 <t> x (MOVVconst [c]))
	// result: (Or32 (Lsh32x64 <t> x (MOVVconst [c&31])) (Rsh32Ux64 <t> x (MOVVconst [-c&31])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpLsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v1.AuxInt = int64ToAuxInt(c & 31)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh32Ux64, t)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(-c & 31)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueSW64_OpRotateLeft64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft64 <t> x (MOVVconst [c]))
	// result: (Or64 (Lsh64x64 <t> x (MOVVconst [c&63])) (Rsh64Ux64 <t> x (MOVVconst [-c&63])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpLsh64x64, t)
		v1 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v1.AuxInt = int64ToAuxInt(c & 63)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh64Ux64, t)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(-c & 63)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueSW64_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVVconst [c]))
	// result: (Or8 (Lsh8x64 <t> x (MOVVconst [c&7])) (Rsh8Ux64 <t> x (MOVVconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x64, t)
		v1 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v1.AuxInt = int64ToAuxInt(c & 7)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux64, t)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(-c & 7)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueSW64_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg2(v5, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt16to64 x) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg2(v5, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRLconst (ZeroExt16to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpSW64SRLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 x (MOVVconst [c]))
	// cond: uint64(c) < 16
	// result: (SRLconst (ZeroExt16to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpSW64SRLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 <t> x y)
	// result: (AND (NEGV <t> (CMPULT y (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt16to64 x) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v2.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(y, v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg2(v4, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSW64_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt16to64 x) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg2(v5, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 <t> x y)
	// result: (SRA (SignExt16to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt16to64 y))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v3.AddArg2(v4, v5)
		v2.AddArg(v3)
		v1.AddArg2(v2, v5)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 <t> x y)
	// result: (SRA (SignExt16to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt32to64 y))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v3.AddArg2(v4, v5)
		v2.AddArg(v3)
		v1.AddArg2(v2, v5)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRAconst (SignExt16to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (SRAconst (SignExt16to64 x) [63])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (MOVVconst [c]))
	// cond: uint64(c) >= 16
	// result: (SRAconst (SignExt16to64 x) [63])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (MOVVconst [c]))
	// cond: uint64(c) < 16
	// result: (SRAconst (SignExt16to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 <t> x y)
	// result: (SRA (SignExt16to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) y)) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v3.AddArg2(v4, y)
		v2.AddArg(v3)
		v1.AddArg2(v2, y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 <t> x y)
	// result: (SRA (SignExt16to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt8to64 y))) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v3.AddArg2(v4, v5)
		v2.AddArg(v3)
		v1.AddArg2(v2, v5)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt32to64 x) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg2(v5, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg2(v5, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRLconst (ZeroExt32to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpSW64SRLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32Ux64 x (MOVVconst [c]))
	// cond: uint64(c) < 32
	// result: (SRLconst (ZeroExt32to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpSW64SRLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32Ux64 <t> x y)
	// result: (AND (NEGV <t> (CMPULT y (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt32to64 x) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v2.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(y, v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg2(v4, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSW64_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt32to64 x) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg2(v5, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 <t> x y)
	// result: (SRA (SignExt32to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt16to64 y))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v3.AddArg2(v4, v5)
		v2.AddArg(v3)
		v1.AddArg2(v2, v5)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 <t> x y)
	// result: (SRA (SignExt32to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt32to64 y))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v3.AddArg2(v4, v5)
		v2.AddArg(v3)
		v1.AddArg2(v2, v5)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRAconst (SignExt32to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (SRAconst (SignExt32to64 x) [63])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 x (MOVVconst [c]))
	// cond: uint64(c) >= 32
	// result: (SRAconst (SignExt32to64 x) [63])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 x (MOVVconst [c]))
	// cond: uint64(c) < 32
	// result: (SRAconst (SignExt32to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 <t> x y)
	// result: (SRA (SignExt32to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) y)) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v3.AddArg2(v4, y)
		v2.AddArg(v3)
		v1.AddArg2(v2, y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 <t> x y)
	// result: (SRA (SignExt32to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt8to64 y))) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v3.AddArg2(v4, v5)
		v2.AddArg(v3)
		v1.AddArg2(v2, v5)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh64Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux16 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh64Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux32 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux64 x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSW64SRLconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux64 x (MOVVconst [c]))
	// cond: uint64(c) < 64
	// result: (SRLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSW64SRLconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux64 <t> x y)
	// result: (AND (NEGV <t> (CMPULT y (MOVVconst <typ.UInt64> [64]))) (SRL <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v2.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(y, v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSW64_OpRsh64Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux8 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v4.AddArg2(x, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x16 <t> x y)
	// result: (SRA x (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt16to64 y))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v1 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(63)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg2(v3, v4)
		v1.AddArg(v2)
		v0.AddArg2(v1, v4)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSW64_OpRsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x32 <t> x y)
	// result: (SRA x (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt32to64 y))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v1 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(63)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg2(v3, v4)
		v1.AddArg(v2)
		v0.AddArg2(v1, v4)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSW64_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x64 x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRAconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRAconst x [63])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 x (MOVVconst [c]))
	// cond: uint64(c) >= 64
	// result: (SRAconst x [63])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 x (MOVVconst [c]))
	// cond: uint64(c) < 64
	// result: (SRAconst x [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 <t> x y)
	// result: (SRA x (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) y)) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v1 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(63)
		v2.AddArg2(v3, y)
		v1.AddArg(v2)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSW64_OpRsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x8 <t> x y)
	// result: (SRA x (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt8to64 y))) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v1 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(63)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg2(v3, v4)
		v1.AddArg(v2)
		v0.AddArg2(v1, v4)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSW64_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt8to64 x) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg2(v5, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt8to64 x) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg2(v5, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRLconst (ZeroExt8to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpSW64SRLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 x (MOVVconst [c]))
	// cond: uint64(c) < 8
	// result: (SRLconst (ZeroExt8to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpSW64SRLconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 <t> x y)
	// result: (AND (NEGV <t> (CMPULT y (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt8to64 x) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v2.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(y, v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg2(v4, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSW64_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 <t> x y)
	// result: (AND (NEGV <t> (CMPULT (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [64]))) (SRL <t> (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64AND)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v1.AddArg2(v2, v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpSW64SRL, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg2(v5, v2)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSW64_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 <t> x y)
	// result: (SRA (SignExt8to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt16to64 y))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v3.AddArg2(v4, v5)
		v2.AddArg(v3)
		v1.AddArg2(v2, v5)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 <t> x y)
	// result: (SRA (SignExt8to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt32to64 y))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v3.AddArg2(v4, v5)
		v2.AddArg(v3)
		v1.AddArg2(v2, v5)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRAconst (SignExt8to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (SRAconst (SignExt8to64 x) [63])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (MOVVconst [c]))
	// cond: uint64(c) >= 8
	// result: (SRAconst (SignExt8to64 x) [63])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (MOVVconst [c]))
	// cond: uint64(c) < 8
	// result: (SRAconst (SignExt8to64 x) [c])
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 <t> x y)
	// result: (SRA (SignExt8to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) y)) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v3.AddArg2(v4, y)
		v2.AddArg(v3)
		v1.AddArg2(v2, y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 <t> x y)
	// result: (SRA (SignExt8to64 x) (BIS <t> (NEGV <t> (CMPULT (MOVVconst <typ.UInt64> [63]) (ZeroExt8to64 y))) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSW64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSW64BIS, t)
		v2 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpSW64CMPULT, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v4.AuxInt = int64ToAuxInt(63)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v3.AddArg2(v4, v5)
		v2.AddArg(v3)
		v1.AddArg2(v2, v5)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSW64_OpSW64ADDV(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDV x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (ADDVconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpSW64MOVVconst {
				continue
			}
			c := auxIntToInt64(v_1.AuxInt)
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpSW64ADDVconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDV x (NEGV y))
	// result: (SUBV x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpSW64NEGV {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpSW64SUBV)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueSW64_OpSW64ADDVconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDVconst [0] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ADDVconst [c] (MOVVconst [d]))
	// result: (MOVVconst [c+d])
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(c + d)
		return true
	}
	// match: (ADDVconst [c] (ADDVconst [d] x))
	// cond: is32Bit(c+d)
	// result: (ADDVconst [c+d] x)
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpSW64ADDVconst)
		v.AuxInt = int64ToAuxInt(c + d)
		v.AddArg(x)
		return true
	}
	// match: (ADDVconst [c] (SUBVconst [d] x))
	// cond: is32Bit(c-d)
	// result: (ADDVconst [c-d] x)
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64SUBVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		x := v_0.Args[0]
		if !(is32Bit(c - d)) {
			break
		}
		v.reset(OpSW64ADDVconst)
		v.AuxInt = int64ToAuxInt(c - d)
		v.AddArg(x)
		return true
	}
	// match: (ADDVconst [off1] (SYMADDR [off2] {sym} ptr))
	// result: (SYMADDR [off1+off2] {sym} ptr)
	for {
		off1 := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		v.reset(OpSW64SYMADDR)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg(ptr)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64ADDWconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDWconst (MOVVconst [c]) [0])
	// result: (MOVVconst [int64(int32(c))])
	for {
		if auxIntToInt32(v.AuxInt) != 0 || v_0.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(int32(c)))
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64AND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AND x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpSW64MOVVconst {
				continue
			}
			c := auxIntToInt64(v_1.AuxInt)
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpSW64ANDconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64ANDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst (MOVVconst [c]) [255])
	// result: (MOVVconst [int64(uint8(c))])
	for {
		if auxIntToInt64(v.AuxInt) != 255 || v_0.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(uint8(c)))
		return true
	}
	// match: (ANDconst (MOVVconst [c]) [65535])
	// result: (MOVVconst [int64(uint16(c))])
	for {
		if auxIntToInt64(v.AuxInt) != 65535 || v_0.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(uint16(c)))
		return true
	}
	// match: (ANDconst (MOVVconst [c]) [0xffffffff])
	// result: (MOVVconst [int64(uint32(c))])
	for {
		if auxIntToInt64(v.AuxInt) != 0xffffffff || v_0.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(uint32(c)))
		return true
	}
	// match: (ANDconst [0] _)
	// result: (MOVVconst [0])
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (ANDconst [-1] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != -1 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ANDconst [c] (MOVVconst [d]))
	// result: (MOVVconst [c&d])
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(c & d)
		return true
	}
	// match: (ANDconst [c] (ANDconst [d] x))
	// result: (ANDconst [c&d] x)
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64ANDconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpSW64ANDconst)
		v.AuxInt = int64ToAuxInt(c & d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64BIS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BIS x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (BISconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpSW64MOVVconst {
				continue
			}
			c := auxIntToInt64(v_1.AuxInt)
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpSW64BISconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (BIS x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64BISconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (BISconst [0] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (BISconst [-1] _)
	// result: (MOVVconst [-1])
	for {
		if auxIntToInt64(v.AuxInt) != -1 {
			break
		}
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(-1)
		return true
	}
	// match: (BISconst [c] (MOVVconst [d]))
	// result: (MOVVconst [c|d])
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(c | d)
		return true
	}
	// match: (BISconst [c] (BISconst [d] x))
	// cond: is32Bit(c|d)
	// result: (BISconst [c|d] x)
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64BISconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		x := v_0.Args[0]
		if !(is32Bit(c | d)) {
			break
		}
		v.reset(OpSW64BISconst)
		v.AuxInt = int64ToAuxInt(c | d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64CMPEQ(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMPEQ x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (CMPEQconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpSW64CMPEQconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64CMPLE(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMPLE x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (CMPLEconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpSW64CMPLEconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64CMPLT(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMPLT x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (CMPLTconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpSW64CMPLTconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64CMPULE(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMPULE x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (CMPULEconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpSW64CMPULEconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64CMPULT(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMPULT x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (CMPULTconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpSW64CMPULTconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVBUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVBUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVBUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUload [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVBUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVBload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVBload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVBload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVBload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVBstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVBstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVBstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (SYMADDR [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVBstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVVconst [0]) mem)
	// result: (MOVBstorezero [off] {sym} ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpSW64MOVVconst || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.reset(OpSW64MOVBstorezero)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVBstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVBstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVBstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVBstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVDload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVDload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVDload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVDstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVDstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVDstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (SYMADDR [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVDstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVFload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVFload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVFload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVFload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVFload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVFload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVFload [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVFload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVFload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVFstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVFstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVFstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVFstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVFstore [off1] {sym1} (SYMADDR [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVFstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVFstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVHUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVHUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVHUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUload [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVHUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVHload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVHload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVHload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVHload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVHstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (SYMADDR [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVVconst [0]) mem)
	// result: (MOVHstorezero [off] {sym} ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpSW64MOVVconst || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.reset(OpSW64MOVHstorezero)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVHstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVHstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVHstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVHstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVVload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVVload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVVload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVVload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVVload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVVload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVVload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVVload [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVVload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVVload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVVstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVVstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVVstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVVstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVVstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVVstore [off1] {sym1} (SYMADDR [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVVstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVVstore [off] {sym} ptr (MOVVconst [0]) mem)
	// result: (MOVVstorezero [off] {sym} ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpSW64MOVVconst || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.reset(OpSW64MOVVstorezero)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVVstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVVstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVVstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVVstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVVstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVVstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVVstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVVstorezero [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVVstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVVstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVWUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVWUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVWUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWUload [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVWUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVWload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVWload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVWload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVWload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVWstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (SYMADDR [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVVconst [0]) mem)
	// result: (MOVWstorezero [off] {sym} ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpSW64MOVVconst || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.reset(OpSW64MOVWstorezero)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MOVWstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVWstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpSW64MOVWstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym1} (SYMADDR [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpSW64SYMADDR {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpSW64MOVWstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64MULL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MULL x (MOVVconst [c]))
	// result: (MULLconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpSW64MOVVconst {
				continue
			}
			c := auxIntToInt64(v_1.AuxInt)
			v.reset(OpSW64MULLconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueSW64_OpSW64MULW(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MULW x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (MULWconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpSW64MOVVconst {
				continue
			}
			c := auxIntToInt64(v_1.AuxInt)
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpSW64MULWconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueSW64_OpSW64NEGV(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEGV (MOVVconst [c]))
	// result: (MOVVconst [-c])
	for {
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(-c)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64ORNOT(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ORNOT x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (ORNOTconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpSW64ORNOTconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64SEXTB(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SEXTB (MOVVconst [c]))
	// result: (MOVVconst [int64(int8(c))])
	for {
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(int8(c)))
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64SEXTH(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SEXTH (MOVVconst [c]))
	// result: (MOVVconst [int64(int16(c))])
	for {
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(int16(c)))
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64SLLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SLLconst [c] (MOVVconst [d]))
	// result: (MOVVconst [int64(d)<<uint64(c)])
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(d) << uint64(c))
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64SRAconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRAconst [c] (MOVVconst [d]))
	// result: (MOVVconst [int64(d)>>uint64(c)])
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(d) >> uint64(c))
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64SRLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRLconst [c] (MOVVconst [d]))
	// result: (MOVVconst [int64(uint64(d)>>uint64(c))])
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(int64(uint64(d) >> uint64(c)))
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64SUBV(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBV x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (SUBVconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpSW64MOVVconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpSW64SUBVconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (SUBV x x)
	// result: (MOVVconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (SUBV (MOVVconst [0]) x)
	// result: (NEGV x)
	for {
		if v_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0.AuxInt) != 0 {
			break
		}
		x := v_1
		v.reset(OpSW64NEGV)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64SUBVconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SUBVconst [0] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (SUBVconst [c] (MOVVconst [d]))
	// result: (MOVVconst [d-c])
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(d - c)
		return true
	}
	// match: (SUBVconst [c] (SUBVconst [d] x))
	// cond: is32Bit(-c-d)
	// result: (ADDVconst [-c-d] x)
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64SUBVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		x := v_0.Args[0]
		if !(is32Bit(-c - d)) {
			break
		}
		v.reset(OpSW64ADDVconst)
		v.AuxInt = int64ToAuxInt(-c - d)
		v.AddArg(x)
		return true
	}
	// match: (SUBVconst [c] (ADDVconst [d] x))
	// cond: is32Bit(-c+d)
	// result: (ADDVconst [-c+d] x)
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64ADDVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		x := v_0.Args[0]
		if !(is32Bit(-c + d)) {
			break
		}
		v.reset(OpSW64ADDVconst)
		v.AuxInt = int64ToAuxInt(-c + d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64XOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XOR x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpSW64MOVVconst {
				continue
			}
			c := auxIntToInt64(v_1.AuxInt)
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpSW64XORconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR x x)
	// result: (MOVVconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueSW64_OpSW64XORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (XORconst [0] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (XORconst [c] (MOVVconst [d]))
	// result: (MOVVconst [c^d])
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64MOVVconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSW64MOVVconst)
		v.AuxInt = int64ToAuxInt(c ^ d)
		return true
	}
	// match: (XORconst [c] (XORconst [d] x))
	// cond: is32Bit(c^d)
	// result: (XORconst [c^d] x)
	for {
		c := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSW64XORconst {
			break
		}
		d := auxIntToInt64(v_0.AuxInt)
		x := v_0.Args[0]
		if !(is32Bit(c ^ d)) {
			break
		}
		v.reset(OpSW64XORconst)
		v.AuxInt = int64ToAuxInt(c ^ d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSW64_OpSignExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt32to64 x)
	// result: (ADDWconst x [0])
	for {
		x := v_0
		v.reset(OpSW64ADDWconst)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSW64_OpSignmask(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Signmask x)
	// result: (SRAconst x [63])
	for {
		x := v_0
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSW64_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRAconst (NEGV <t> x) [63])
	for {
		t := v.Type
		x := v_0
		v.reset(OpSW64SRAconst)
		v.AuxInt = int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, OpSW64NEGV, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSW64_OpStaticCall(v *Value) bool {
	v_0 := v.Args[0]
	// match: (StaticCall [argwid] {target} mem)
	// result: (CALLstatic [argwid] {target} mem)
	for {
		argwid := auxIntToInt32(v.AuxInt)
		target := auxToCall(v.Aux)
		mem := v_0
		v.reset(OpSW64CALLstatic)
		v.AuxInt = int32ToAuxInt(argwid)
		v.Aux = symToAux(target)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueSW64_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpSW64MOVBstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && !is64BitFloat(val.Type)
	// result: (MOVVstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && !is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (MOVFstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpSW64MOVFstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (MOVDstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpSW64MOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_1
		v.copyOf(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// result: (MOVBstore ptr (MOVVconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpSW64MOVBstore)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVVconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// result: (MOVBstore [1] ptr (MOVVconst [0]) (MOVBstore [0] ptr (MOVVconst [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpSW64MOVBstore)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(ptr, v0, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVVconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] ptr (MOVVconst [0]) (MOVHstore [0] ptr (MOVVconst [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(ptr, v0, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [4] ptr mem)
	// result: (MOVBstore [3] ptr (MOVVconst [0]) (MOVBstore [2] ptr (MOVVconst [0]) (MOVBstore [1] ptr (MOVVconst [0]) (MOVBstore [0] ptr (MOVVconst [0]) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpSW64MOVBstore)
		v.AuxInt = int32ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(1)
		v3 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(0)
		v3.AddArg3(ptr, v0, mem)
		v2.AddArg3(ptr, v0, v3)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore ptr (MOVVconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [4] ptr (MOVVconst [0]) (MOVWstore [0] ptr (MOVVconst [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVWstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(ptr, v0, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [6] ptr (MOVVconst [0]) (MOVHstore [4] ptr (MOVVconst [0]) (MOVHstore [2] ptr (MOVVconst [0]) (MOVHstore [0] ptr (MOVVconst [0]) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AuxInt = int32ToAuxInt(6)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(4)
		v2 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(2)
		v3 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(0)
		v3.AddArg3(ptr, v0, mem)
		v2.AddArg3(ptr, v0, v3)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [3] ptr mem)
	// result: (MOVBstore [2] ptr (MOVVconst [0]) (MOVBstore [1] ptr (MOVVconst [0]) (MOVBstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpSW64MOVBstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(1)
		v2 := b.NewValue0(v.Pos, OpSW64MOVBstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(0)
		v2.AddArg3(ptr, v0, mem)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [6] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [4] ptr (MOVVconst [0]) (MOVHstore [2] ptr (MOVVconst [0]) (MOVHstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 6 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpSW64MOVHstore)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpSW64MOVHstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(0)
		v2.AddArg3(ptr, v0, mem)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [12] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [8] ptr (MOVVconst [0]) (MOVWstore [4] ptr (MOVVconst [0]) (MOVWstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 12 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpSW64MOVWstore)
		v.AuxInt = int32ToAuxInt(8)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVWstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(4)
		v2 := b.NewValue0(v.Pos, OpSW64MOVWstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(0)
		v2.AddArg3(ptr, v0, mem)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [16] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore [8] ptr (MOVVconst [0]) (MOVVstore [0] ptr (MOVVconst [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 16 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v.AuxInt = int32ToAuxInt(8)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVVstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(ptr, v0, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [24] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore [16] ptr (MOVVconst [0]) (MOVVstore [8] ptr (MOVVconst [0]) (MOVVstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 24 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpSW64MOVVstore)
		v.AuxInt = int32ToAuxInt(16)
		v0 := b.NewValue0(v.Pos, OpSW64MOVVconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpSW64MOVVstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(8)
		v2 := b.NewValue0(v.Pos, OpSW64MOVVstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(0)
		v2.AddArg3(ptr, v0, mem)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: s%8 == 0 && s > 24 && s <= 8*128 && t.(*types.Type).Alignment()%8 == 0 && !config.noDuffDevice
	// result: (DUFFZERO [8 * (128 - int64(s/8))] ptr mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(s%8 == 0 && s > 24 && s <= 8*128 && t.(*types.Type).Alignment()%8 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpSW64DUFFZERO)
		v.AuxInt = int64ToAuxInt(8 * (128 - int64(s/8)))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: (s > 8*128 || config.noDuffDevice) || t.(*types.Type).Alignment()%8 != 0
	// result: (LoweredZero [t.(*types.Type).Alignment()] ptr (ADDVconst <ptr.Type> ptr [s-moveSize(t.(*types.Type).Alignment(), config)]) mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !((s > 8*128 || config.noDuffDevice) || t.(*types.Type).Alignment()%8 != 0) {
			break
		}
		v.reset(OpSW64LoweredZero)
		v.AuxInt = int64ToAuxInt(t.(*types.Type).Alignment())
		v0 := b.NewValue0(v.Pos, OpSW64ADDVconst, ptr.Type)
		v0.AuxInt = int64ToAuxInt(s - moveSize(t.(*types.Type).Alignment(), config))
		v0.AddArg(ptr)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	return false
}
func rewriteValueSW64_OpZeroExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt16to32 x)
	// result: (ANDconst x [65535])
	for {
		x := v_0
		v.reset(OpSW64ANDconst)
		v.AuxInt = int64ToAuxInt(65535)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSW64_OpZeroExt16to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt16to64 x)
	// result: (ANDconst x [65535])
	for {
		x := v_0
		v.reset(OpSW64ANDconst)
		v.AuxInt = int64ToAuxInt(65535)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSW64_OpZeroExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt32to64 x)
	// result: (ANDconst x [0xffffffff])
	for {
		x := v_0
		v.reset(OpSW64ANDconst)
		v.AuxInt = int64ToAuxInt(0xffffffff)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSW64_OpZeroExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt8to16 x)
	// result: (ANDconst x [255])
	for {
		x := v_0
		v.reset(OpSW64ANDconst)
		v.AuxInt = int64ToAuxInt(255)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSW64_OpZeroExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt8to32 x)
	// result: (ANDconst x [255])
	for {
		x := v_0
		v.reset(OpSW64ANDconst)
		v.AuxInt = int64ToAuxInt(255)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSW64_OpZeroExt8to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt8to64 x)
	// result: (ANDconst x [255])
	for {
		x := v_0
		v.reset(OpSW64ANDconst)
		v.AuxInt = int64ToAuxInt(255)
		v.AddArg(x)
		return true
	}
}
func rewriteBlockSW64(b *Block) bool {
	switch b.Kind {
	case BlockSW64EQ:
		// match: (EQ (XORconst cmp:(CMPEQ _ _) [1]) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpSW64XORconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpSW64CMPEQ {
				break
			}
			b.resetWithControl(BlockSW64NE, cmp)
			return true
		}
		// match: (EQ (XORconst cmp:(CMPEQconst _) [1]) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpSW64XORconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpSW64CMPEQconst {
				break
			}
			b.resetWithControl(BlockSW64NE, cmp)
			return true
		}
		// match: (EQ (CMPEQconst x [0]) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpSW64CMPEQconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64NE, x)
			return true
		}
		// match: (EQ (CMPLTconst x [0]) yes no)
		// result: (GE x yes no)
		for b.Controls[0].Op == OpSW64CMPLTconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64GE, x)
			return true
		}
		// match: (EQ (CMPLEconst x [0]) yes no)
		// result: (GT x yes no)
		for b.Controls[0].Op == OpSW64CMPLEconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64GT, x)
			return true
		}
		// match: (EQ (CMPLTconst x [1]) yes no)
		// result: (GT x yes no)
		for b.Controls[0].Op == OpSW64CMPLTconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 1 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64GT, x)
			return true
		}
		// match: (EQ (CMPULTconst x [1]) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpSW64CMPULTconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 1 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64NE, x)
			return true
		}
		// match: (EQ (CMPULEconst x [0]) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpSW64CMPULEconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64NE, x)
			return true
		}
		// match: (EQ (CMPEQ x (MOVVconst [0])) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpSW64CMPEQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64NE, x)
			return true
		}
		// match: (EQ (CMPLT x (MOVVconst [0])) yes no)
		// result: (GE x yes no)
		for b.Controls[0].Op == OpSW64CMPLT {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64GE, x)
			return true
		}
		// match: (EQ (CMPLT x (MOVVconst [1])) yes no)
		// result: (GT x yes no)
		for b.Controls[0].Op == OpSW64CMPLT {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 1 {
				break
			}
			b.resetWithControl(BlockSW64GT, x)
			return true
		}
		// match: (EQ (CMPLE x (MOVVconst [0])) yes no)
		// result: (GT x yes no)
		for b.Controls[0].Op == OpSW64CMPLE {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64GT, x)
			return true
		}
		// match: (EQ (CMPULT x (MOVVconst [1])) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpSW64CMPULT {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 1 {
				break
			}
			b.resetWithControl(BlockSW64NE, x)
			return true
		}
		// match: (EQ (CMPULE x (MOVVconst [0])) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpSW64CMPULE {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64NE, x)
			return true
		}
		// match: (EQ (CMPEQ (MOVVconst [0]) x) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpSW64CMPEQ {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64NE, x)
			return true
		}
		// match: (EQ (CMPLT (MOVVconst [0]) x) yes no)
		// result: (LE x yes no)
		for b.Controls[0].Op == OpSW64CMPLT {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64LE, x)
			return true
		}
		// match: (EQ (CMPLE (MOVVconst [0]) x) yes no)
		// result: (LT x yes no)
		for b.Controls[0].Op == OpSW64CMPLE {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64LT, x)
			return true
		}
		// match: (EQ (CMPLE (MOVVconst [1]) x) yes no)
		// result: (LE x yes no)
		for b.Controls[0].Op == OpSW64CMPLE {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 1 {
				break
			}
			b.resetWithControl(BlockSW64LE, x)
			return true
		}
		// match: (EQ (CMPULT (MOVVconst [0]) x) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpSW64CMPULT {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64EQ, x)
			return true
		}
		// match: (EQ (CMPULE (MOVVconst [1]) x) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpSW64CMPULE {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 1 {
				break
			}
			b.resetWithControl(BlockSW64EQ, x)
			return true
		}
		// match: (EQ (FNotEqual cmp) yes no)
		// result: (FNE cmp yes no)
		for b.Controls[0].Op == OpSW64FNotEqual {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockSW64FNE, cmp)
			return true
		}
		// match: (EQ (FEqual cmp) yes no)
		// result: (FEQ cmp yes no)
		for b.Controls[0].Op == OpSW64FEqual {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockSW64FEQ, cmp)
			return true
		}
	case BlockIf:
		// match: (If cond yes no)
		// result: (NE cond yes no)
		for {
			cond := b.Controls[0]
			b.resetWithControl(BlockSW64NE, cond)
			return true
		}
	case BlockSW64NE:
		// match: (NE (XORconst cmp:(CMPEQ _ _) [1]) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpSW64XORconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpSW64CMPEQ {
				break
			}
			b.resetWithControl(BlockSW64EQ, cmp)
			return true
		}
		// match: (NE (XORconst cmp:(CMPEQconst _) [1]) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpSW64XORconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpSW64CMPEQconst {
				break
			}
			b.resetWithControl(BlockSW64EQ, cmp)
			return true
		}
		// match: (NE (CMPEQconst x [0]) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpSW64CMPEQconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64EQ, x)
			return true
		}
		// match: (NE (CMPLTconst x [0]) yes no)
		// result: (LT x yes no)
		for b.Controls[0].Op == OpSW64CMPLTconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64LT, x)
			return true
		}
		// match: (NE (CMPLEconst x [0]) yes no)
		// result: (LE x yes no)
		for b.Controls[0].Op == OpSW64CMPLEconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64LE, x)
			return true
		}
		// match: (NE (CMPLTconst x [1]) yes no)
		// result: (LE x yes no)
		for b.Controls[0].Op == OpSW64CMPLTconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 1 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64LE, x)
			return true
		}
		// match: (NE (CMPULTconst x [1]) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpSW64CMPULTconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 1 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64EQ, x)
			return true
		}
		// match: (NE (CMPULEconst x [0]) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpSW64CMPULEconst {
			v_0 := b.Controls[0]
			if auxIntToInt64(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockSW64EQ, x)
			return true
		}
		// match: (NE (CMPEQ x (MOVVconst [0])) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpSW64CMPEQ {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64EQ, x)
			return true
		}
		// match: (NE (CMPLT x (MOVVconst [0])) yes no)
		// result: (LT x yes no)
		for b.Controls[0].Op == OpSW64CMPLT {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64LT, x)
			return true
		}
		// match: (NE (CMPLT x (MOVVconst [1])) yes no)
		// result: (LE x yes no)
		for b.Controls[0].Op == OpSW64CMPLT {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 1 {
				break
			}
			b.resetWithControl(BlockSW64LE, x)
			return true
		}
		// match: (NE (CMPLE x (MOVVconst [0])) yes no)
		// result: (LE x yes no)
		for b.Controls[0].Op == OpSW64CMPLE {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64LE, x)
			return true
		}
		// match: (NE (CMPULT x (MOVVconst [1])) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpSW64CMPULT {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 1 {
				break
			}
			b.resetWithControl(BlockSW64EQ, x)
			return true
		}
		// match: (NE (CMPULE x (MOVVconst [0])) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpSW64CMPULE {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpSW64MOVVconst || auxIntToInt64(v_0_1.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64EQ, x)
			return true
		}
		// match: (NE (CMPEQ (MOVVconst [0]) x) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpSW64CMPEQ {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64EQ, x)
			return true
		}
		// match: (NE (CMPLT (MOVVconst [0]) x) yes no)
		// result: (GT x yes no)
		for b.Controls[0].Op == OpSW64CMPLT {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64GT, x)
			return true
		}
		// match: (NE (CMPLE (MOVVconst [0]) x) yes no)
		// result: (GE x yes no)
		for b.Controls[0].Op == OpSW64CMPLE {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64GE, x)
			return true
		}
		// match: (NE (CMPLE (MOVVconst [1]) x) yes no)
		// result: (GT x yes no)
		for b.Controls[0].Op == OpSW64CMPLE {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 1 {
				break
			}
			b.resetWithControl(BlockSW64GT, x)
			return true
		}
		// match: (NE (CMPULT (MOVVconst [0]) x) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpSW64CMPULT {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 0 {
				break
			}
			b.resetWithControl(BlockSW64NE, x)
			return true
		}
		// match: (NE (CMPULE (MOVVconst [1]) x) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpSW64CMPULE {
			v_0 := b.Controls[0]
			x := v_0.Args[1]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpSW64MOVVconst || auxIntToInt64(v_0_0.AuxInt) != 1 {
				break
			}
			b.resetWithControl(BlockSW64NE, x)
			return true
		}
		// match: (NE (FNotEqual cmp) yes no)
		// result: (FEQ cmp yes no)
		for b.Controls[0].Op == OpSW64FNotEqual {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockSW64FEQ, cmp)
			return true
		}
		// match: (NE (FEqual cmp) yes no)
		// result: (FNE cmp yes no)
		for b.Controls[0].Op == OpSW64FEqual {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockSW64FNE, cmp)
			return true
		}
	}
	return false
}
