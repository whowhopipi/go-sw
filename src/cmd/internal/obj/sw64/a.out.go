// cmd/9c/9.out.h from Vita Nuova.
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
)

//go:generate go run ../stringer_sw64.go -i $GOFILE -o anames.go -p sw64

/*
 * sw64
 */
const (
	NREG  = 32        /* number of general registers */
	NFREG = 32        /* number of floating point registers */
	NINST = 253 + 114 /* number of instructions */
)

const (
	REG_R0 = obj.RBaseSW64 + iota
	REG_R1
	REG_R2
	REG_R3
	REG_R4
	REG_R5
	REG_R6
	REG_R7
	REG_R8
	REG_R9
	REG_R10
	REG_R11
	REG_R12
	REG_R13
	REG_R14
	REG_R15
	REG_R16
	REG_R17
	REG_R18
	REG_R19
	REG_R20
	REG_R21
	REG_R22
	REG_R23
	REG_R24
	REG_R25
	REG_R26
	REG_R27
	REG_R28
	REG_R29
	REG_R30
	REG_R31

	REG_F0
	REG_F1
	REG_F2
	REG_F3
	REG_F4
	REG_F5
	REG_F6
	REG_F7
	REG_F8
	REG_F9
	REG_F10
	REG_F11
	REG_F12
	REG_F13
	REG_F14
	REG_F15
	REG_F16
	REG_F17
	REG_F18
	REG_F19
	REG_F20
	REG_F21
	REG_F22
	REG_F23
	REG_F24
	REG_F25
	REG_F26
	REG_F27
	REG_F28
	REG_F29
	REG_F30
	REG_F31
	REG_LAST = REG_F31 // the last defined register

	REGG    = REG_R15 // fp
	REGLINK = REG_R26 // ra
	REGCTXT = REG_R25 // t12/pv
	REGTMP  = REG_R28 // at
	REGSB   = REG_R29 // gp
	REGSP   = REG_R30 // sp
	REGZERO = REG_R31 // zero

)

//zxw add
var SW64DWARFRegisters = map[int16]int16{}
func init() {
        // f assigns dwarfregisters[from:to] = (base):(to-from+base)
        f := func(from, to, base int16) {
                for r := int16(from); r <= to; r++ {
                        SW64DWARFRegisters[r] = (r - from) + base
                }
        }
        f(REG_R0, REG_R31, 0)
        f(REG_F0, REG_F31, 32) // For 32-bit SW64, compiler only uses even numbered registers --  see cmd/compile/internal/ssa/gen/MIPSOps.go
}
const (
	BIG = 32766
)

const (
	C_NONE  = iota
	C_RREG  // integer register
	C_FREG  // float-point register
	C_REG   // any register
	C_SUCON // integer constant not greater than 0xff
	C_MCON  // integer constant [-0x8000,0x7fff]
	C_LCON  // integer constant without restriction
	C_FCON  // float-point constant without restriction
	C_LEXT  // external symbol
	C_ADDR  // all symbol
	C_SAUTO // (a.NAME == NAME_PARAM || a.NAME == NAME_AUTO) && offset is [-0x7fff, 0x7fff]
	C_LAUTO // a.NAME == NAME_PARAM || a.NAME == NAME_AUTO
	C_ZOREG // a.NAME == NAME_NONE && offset is 0
	C_SOREG // (a.NAME == NAME_NONE) && offset is [-0x7fff, 0x7fff]
	C_LOREG // a.NAME == NAME_NONE
	C_SBRA  // a.NAME == NAME_NONE && offset is 0
	C_GOK
	C_TLS_LE
	C_TLS_IE
	C_NCLASS /* must be the last */
)

const (
	ASYS_CALL = obj.ABaseSW64 + obj.A_ARCHSPECIFIC + iota
	ASYS_CALL_B
	AEXCB
	ABR //控制指令, JMP,RET,CALL定义为obj.AJMP,obj.ARET,obj.CALL
	ABSR
	ABEQ
	ABNE
	ABLT
	ABLE
	ABGT
	ABGE
	ABLBC
	ABLBS
	AFBEQ
	AFBNE
	AFBLT
	AFBLE
	AFBGT
	AFBGE
	AMOVB
	AMOVBU
	AMOVH
	AMOVHU
	AMOVW
	AMOVWU
	AMOVV
	AMOVF
	AMOVD
	ALDBU //存储装载指令
	ALDHU
	ALDW
	ALDL
	ALDL_U
	ASTB
	ASTH
	ASTW
	ASTL
	ASTL_U
	AFLDS
	AFLDD
	AFSTS
	AFSTD
	ALDI //装入立即数
	ALDF //装入浮点立即数
	ALDIH
	AADDW //整数算数运算
	ASUBW
	AS4ADDW
	AS4SUBW
	AS8ADDW
	AS8SUBW
	AADDL
	ASUBL
	AS4ADDL
	AS4SUBL
	AS8ADDL
	AS8SUBL
	AMULW
	AMULL
	AUMULH
	ACTPOP
	ACTLZ
	ACTTZ
	AZAP
	AZAPNOT
	ASEXTB
	ASEXTH
	AFADDS //浮点数运算
	AFADDD
	AFSUBS
	AFSUBD
	AFMULS
	AFMULD
	AFMAS
	AFMAD
	AFMSS
	AFMSD
	AFNMAS
	AFNMAD
	AFNMSS
	AFNMSD
	AFSELEQ
	AFSELNE
	AFSELLT
	AFSELLE
	AFSELGT
	AFSELGE
	AFDIVS
	AFDIVD
	AFCPYS
	AFCVTSD
	AFCVTDS
	AFCVTDL
	AFCVTLS
	AFCVTLD
	AFCVTLW
	AFCVTWL
	AFSQRTS
	AFSQRTD
	AFCVTDL_Z
	AFCVTDL_P
	AFCVTDL_G
	AFCVTDL_N
	AFCMPLE
	AFCMPLT
	AFCMPEQ
	AFCMPUN
	AFCPYSN
	AFCPYSE
	AIFMOVS
	AIFMOVD
	AFIMOVS
	AFIMOVD
	ARFPCR
	AWFPCR
	ASETFPEC0
	ASETFPEC1
	ASETFPEC2
	ASETFPEC3
	AMEMB //杂项指令
	AIMEMB
	ARTC
	ARCID
	AHALT
	ARD_F
	AWR_F
	ARTID
	ALLDW
	ALLDL
	ALDW_INC
	ALDL_INC
	ALDW_DEC
	ALDL_DEC
	ALDW_SET
	ALDL_SET
	ALSTW
	ALSTL
	ALDW_NC
	ALDL_NC
	ALDD_NC
	ASTW_NC
	ASTL_NC
	ASTD_NC
	ACMPEQ
	ACMPLT
	ACMPLE
	ACMPULT
	ACMPULE
	AAND
	ABIC
	ABIS
	AORNOT
	AXOR
	AEQV
	AINSLB
	AINSLH
	AINSLW
	AINSLL
	AINSHB
	AINSHH
	AINSHW
	AINSHL
	ASLL
	ASRL
	ASRA
	AEXTLB
	AEXTLH
	AEXTLW
	AEXTLL
	AEXTHB
	AEXTHH
	AEXTHW
	AEXTHL
	AMASKLB
	AMASKLH
	AMASKLW
	AMASKLL
	AMASKHB
	AMASKHH
	AMASKHW
	AMASKHL
	ACMPGEB
	ASELEQ
	ASELGE
	ASELGT
	ASELLE
	ASELLT
	ASELNE
	ASELLBC
	ASELLBS
	APRI_ST
	APRI_LD
	ASYMADDR
	AWORD
	ANOOP
	ALDGP
	ALAST
	// aliases
	AJMP  = obj.AJMP
	ACALL = obj.ACALL
	ARET  = obj.ARET
)

//zxw add
const (
       /* mark flags */
       FOLL    = 1 << 0
       LABEL   = 1 << 1
       LEAF    = 1 << 2
       SYNC    = 1 << 3
       BRANCH  = 1 << 4
       LOAD    = 1 << 5
       FCMP    = 1 << 6
       NOSCHED = 1 << 7

       NSCHED = 20
)
