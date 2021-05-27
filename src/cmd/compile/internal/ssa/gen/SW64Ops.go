// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import "strings"

// Notes:
//  - Integer types live in the low portion of registers. Upper portions are junk.
//  - Boolean types use the low-order byte of a register. 0=false, 1=true.
//    Upper bytes are junk.
//  - *const instructions may use a constant larger than the instruction can encode.
//    In this case the assembler expands to multiple instructions and uses tmp
//    register (R28).

// Suffixes encode the bit width of various instructions.
// V (vlong)     = 64 bit
// WU (word)     = 32 bit unsigned
// W (word)      = 32 bit
// H (half word) = 16 bit
// HU            = 16 bit unsigned
// B (byte)      = 8 bit
// BU            = 8 bit unsigned
// F (float)     = 32 bit float
// D (double)    = 64 bit float

// Note: registers not used in regalloc are not included in this list,
// so that regmask stays within int64
// Be careful when hand coding regmasks.
var regNamesSW64 = []string{
	"R0",
	"R1",
	"R2",
	"R3",
	"R4",
	"R5",
	"R6",
	"R7",
	"R8",
	"R9",
	"R10",
	"R11",
	"R12",
	"R13",
	"R14",
	"g", //R15
	"R16",
	"R17",
	"R18",
	"R19",
	"R20",
	"R21",
	"R22",
	"R23",
	"R24",
	"R25", //REGCTXT
	"R26", //link register
	//"R27", //PV
	//"R28", REGTMP
	//"R29", REGSB
	"SP",  //R30
	"R31", // REGZERO

	"F0",
	"F1",
	"F2",
	"F3",
	"F4",
	"F5",
	"F6",
	"F7",
	"F8",
	"F9",
	"F10",
	"F11",
	"F12",
	"F13",
	"F14",
	"F15",
	"F16",
	"F17",
	"F18",
	"F19",
	"F20",
	"F21",
	"F22",
	"F23",
	"F24",
	"F25",
	"F26",
	"F27",
	"F28",
	"F29",
	"F30",
	"F31",

	// pseudo-registers
	"SB",
}

func init() {
	// Make map from reg names to reg integers.
	if len(regNamesSW64) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNamesSW64 {
		num[name] = i
	}
	buildReg := func(s string) regMask {
		m := regMask(0)
		for _, r := range strings.Split(s, " ") {
			if n, ok := num[r]; ok {
				m |= regMask(1) << uint(n)
				continue
			}
			panic("register " + r + " not found")
		}
		return m
	}

	// Common individual register masks
	var (
		gp         = buildReg("R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25")
		gpg        = gp | buildReg("g")
		gpsp       = gp | buildReg("SP")
		gpspg      = gpg | buildReg("SP")
		gpspsbg    = gpspg | buildReg("SB")
		fp         = buildReg("F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30")
		fp_src     = buildReg("F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20")
		fp_dst     = buildReg("F21 F22 F23 F24 F25 F26 F27 F28 F29 F30")
		callerSave = gp | fp | buildReg("g") // runtime.setg (and anything calling it) may clobber g
		//zxw new add
		r1 = buildReg("R1")
		r2 = buildReg("R2")
		r3 = buildReg("R3")
		r4 = buildReg("R4")
	)
	// Common regInfo
	var (
		gp01   = regInfo{inputs: nil, outputs: []regMask{gp}}
		gp11   = regInfo{inputs: []regMask{gpg}, outputs: []regMask{gp}}
		gp11sp = regInfo{inputs: []regMask{gpspg}, outputs: []regMask{gp}}
		gp21   = regInfo{inputs: []regMask{gpg, gpg}, outputs: []regMask{gp}}
		//gp2hilo  = regInfo{inputs: []regMask{gpg, gpg}, outputs: []regMask{hi, lo}}
		gpload   = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{gp}}
		gpstore  = regInfo{inputs: []regMask{gpspsbg, gpg}}
		gpstore0 = regInfo{inputs: []regMask{gpspsbg}}
		gpxchg   = regInfo{inputs: []regMask{gpspsbg, gpg}, outputs: []regMask{gp}}
		gpcas    = regInfo{inputs: []regMask{gpspsbg, gpg, gpg}, outputs: []regMask{gp}}
		fp01     = regInfo{inputs: nil, outputs: []regMask{fp}}
		fp11     = regInfo{inputs: []regMask{fp}, outputs: []regMask{fp}}
		ffp11    = regInfo{inputs: []regMask{fp_src}, outputs: []regMask{fp_dst}}
		//fp1flags  = regInfo{inputs: []regMask{fp}}
		//fpgp      = regInfo{inputs: []regMask{fp}, outputs: []regMask{gp}}
		//gpfp      = regInfo{inputs: []regMask{gp}, outputs: []regMask{fp}}
		fp21  = regInfo{inputs: []regMask{fp, fp}, outputs: []regMask{fp}}
		ffp21 = regInfo{inputs: []regMask{fp_src, fp_src}, outputs: []regMask{fp_dst}}
		//fp31  = regInfo{inputs: []regMask{fp, fp, fp}, outputs: []regMask{fp}}
		// fp2flags  = regInfo{inputs: []regMask{fp, fp}}
		fpload    = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{fp}}
		fpstore   = regInfo{inputs: []regMask{gpspsbg, fp}}
		readflags = regInfo{inputs: []regMask{fp}, outputs: []regMask{gp}}
		ifp11     = regInfo{inputs: []regMask{gp}, outputs: []regMask{fp}}
		fip11     = regInfo{inputs: []regMask{fp}, outputs: []regMask{gp}}
	)
	ops := []opData{
		// binary ops
		{name: "ADDV", argLength: 2, reg: gp21, asm: "ADDL", commutative: true},
		{name: "ADDVconst", argLength: 1, reg: gp11sp, asm: "ADDL", aux: "Int64"},
		{name: "ADDW", argLength: 2, reg: gp21, asm: "ADDW", commutative: true},
		{name: "ADDWconst", argLength: 1, reg: gp11sp, asm: "ADDW", aux: "Int32"},
		{name: "SUBV", argLength: 2, reg: gp21, asm: "SUBL"},
		{name: "SUBVconst", argLength: 1, reg: gp11sp, asm: "SUBL", aux: "Int64"},
		{name: "MULW", argLength: 2, reg: gp21, asm: "MULW", commutative: true},
		{name: "MULWconst", argLength: 1, reg: gp11sp, asm: "MULW", aux: "Int64"},
		{name: "MULL", argLength: 2, reg: gp21, asm: "MULL", commutative: true},
		{name: "MULLconst", argLength: 1, reg: gp11sp, asm: "MULL", aux: "Int64"},
		{name: "UMULH", argLength: 2, reg: gp21, asm: "UMULH", commutative: true},
		{name: "UMULHconst", argLength: 1, reg: gp11sp, asm: "UMULH", aux: "Int64"},
		// udiv runtime call for soft division
		// output0 = arg0/arg1, output1 = arg0%arg1
		// see ../../../../../runtime/vlop_arm.s
		{
			name:      "CALLudiv",
			argLength: 2,
			reg: regInfo{
				inputs:   []regMask{buildReg("R1"), buildReg("R0")},
				outputs:  []regMask{buildReg("R0"), buildReg("R1")},
				clobbers: buildReg("R2 R3 R4 R5 R26"),
			},
			clobberFlags: true,
			typ:          "(UInt64,UInt64)",
			call:         false,
		},

		{name: "FADDD", argLength: 2, reg: ffp21, asm: "FADDD", commutative: true},
		{name: "FADDS", argLength: 2, reg: ffp21, asm: "FADDS", commutative: true},
		{name: "FSUBD", argLength: 2, reg: ffp21, asm: "FSUBD"},
		{name: "FSUBS", argLength: 2, reg: ffp21, asm: "FSUBS"},
		{name: "FMULD", argLength: 2, reg: ffp21, asm: "FMULD", commutative: true},
		{name: "FMULS", argLength: 2, reg: ffp21, asm: "FMULS", commutative: true},
		{name: "FDIVD", argLength: 2, reg: ffp21, asm: "FDIVD"},
		{name: "FDIVS", argLength: 2, reg: ffp21, asm: "FDIVS"},
		{name: "FCPYS", argLength: 2, reg: fp21, asm: "FCPYS"},
		{name: "FABS", argLength: 1, reg: fp11, asm: "FCPYS"},
		{name: "IFMOVD", argLength: 1, reg: ifp11, asm: "IFMOVD", typ: "Float64"},
		{name: "FIMOVD", argLength: 1, reg: fip11, asm: "FIMOVD", typ: "Int64"},
		{name: "FCVTSD", argLength: 1, reg: ffp11, asm: "FCVTSD", typ: "Float64"},
		{name: "FCVTDS", argLength: 1, reg: ffp11, asm: "FCVTDS"},
		{name: "FCVTDL", argLength: 1, reg: fp11, asm: "FCVTDL"},
		{name: "FCVTLS", argLength: 1, reg: fp11, asm: "FCVTLS"},
		{name: "FCVTLD", argLength: 1, reg: fp11, asm: "FCVTLD"},
		{name: "FCVTLW", argLength: 1, reg: fp11, asm: "FCVTLW"},
		{name: "FCVTWL", argLength: 1, typ: "Float64", reg: fp11, asm: "FCVTWL"},

		{name: "FSQRTS", argLength: 1, reg: fp11, asm: "FSQRTS"},
		{name: "FSQRTD", argLength: 1, reg: ffp11, asm: "FSQRTD"},
		{name: "FCVTDL_Z", argLength: 1, reg: fp11, typ: "Float64", asm: "FCVTDL_Z"},
		{name: "FCVTDL_P", argLength: 1, reg: fp11, asm: "FCVTDL_P"},
		{name: "FCVTDL_G", argLength: 1, reg: fp11, asm: "FCVTDL_G"},
		{name: "FCVTDL_N", argLength: 1, reg: fp11, asm: "FCVTDL_N"},

		{name: "AND", argLength: 2, reg: gp21, asm: "AND", commutative: true},
		{name: "ANDconst", argLength: 1, reg: gp11, asm: "AND", aux: "Int64"},
		{name: "BIS", argLength: 2, reg: gp21, asm: "BIS", commutative: true},
		{name: "BISconst", argLength: 1, reg: gp11, asm: "BIS", aux: "Int64"},
		{name: "XOR", argLength: 2, reg: gp21, asm: "XOR", commutative: true},
		{name: "XORconst", argLength: 1, reg: gp11, asm: "XOR", aux: "Int64"},

		{name: "ORNOT", argLength: 2, reg: gp21, asm: "ORNOT"},
		{name: "ORNOTconst", argLength: 1, reg: gp11, asm: "ORNOT", aux: "Int64"},
		{name: "NEGV", argLength: 1, reg: gp11, asm: "SUBL"},   // -arg0
		{name: "NEGF", argLength: 1, reg: fp11, asm: "FCPYSN"}, // -arg0, float32
		{name: "NEGD", argLength: 1, reg: fp11, asm: "FCPYSN"}, // -arg0, float64

		// shifts
		{name: "SLL", argLength: 2, reg: gp21, asm: "SLL"},                                   // arg0 << arg1, shift amount is mod 64
		{name: "SLLconst", argLength: 1, reg: gp11, asm: "SLL", aux: "Int64"},                // arg0 << auxInt
		{name: "SRL", argLength: 2, reg: gp21, asm: "SRL"},                                   // arg0 >> arg1, unsigned, shift amount is mod 64
		{name: "SRLconst", argLength: 1, reg: gp11, asm: "SRL", typ: "UInt64", aux: "Int64"}, // arg0 >> auxInt, unsigned
		{name: "SRA", argLength: 2, reg: gp21, asm: "SRA"},                                   // arg0 >> arg1, signed, shift amount is mod 64
		{name: "SRAconst", argLength: 1, reg: gp11, asm: "SRA", aux: "Int64"},                // arg0 >> auxInt, signed

		{name: "CTTZ", argLength: 1, reg: gp11, asm: "CTTZ", typ: "UInt64"},
		{name: "CTLZ", argLength: 1, reg: gp11, asm: "CTLZ"},
		{name: "CTPOP", argLength: 1, reg: gp11, asm: "CTPOP"},
		{name: "SEXTB", argLength: 1, reg: gp11, asm: "SEXTB"},
		{name: "SEXTBconst", argLength: 0, reg: gp01, aux: "Int64", asm: "SEXTB"},
		{name: "SEXTH", argLength: 1, reg: gp11, asm: "SEXTH"},
		{name: "SEXTHconst", argLength: 0, reg: gp01, aux: "Int64", asm: "SEXTH"},

		// arg0 + auxInt + aux.(*gc.Sym), arg0=SP/SB
		{name: "SYMADDR", argLength: 1, reg: regInfo{inputs: []regMask{buildReg("SP") | buildReg("SB")}, outputs: []regMask{gp}}, aux: "SymOff", asm: "SYMADDR", rematerializeable: true, symEffect: "Addr"},

		{name: "MOVVconst", argLength: 0, reg: gp01, aux: "Int64", asm: "LDI", typ: "UInt64", rematerializeable: true},    // auxint
		{name: "MOVFconst", argLength: 0, reg: fp01, aux: "Float32", asm: "LDF", typ: "Float32", rematerializeable: true}, // auxint as 32-bit float
		{name: "MOVDconst", argLength: 0, reg: fp01, aux: "Float64", asm: "LDI", typ: "Float64", rematerializeable: true}, // auxint as 64-bit float

		{name: "MOVBstorezero", argLength: 2, reg: gpstore0, asm: "STB", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"}, // store zero byte to arg0+aux.  arg1=mem
		{name: "MOVHstorezero", argLength: 2, reg: gpstore0, asm: "STH", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"}, // store zero 2 bytes to ...
		{name: "MOVWstorezero", argLength: 2, reg: gpstore0, asm: "STW", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"}, // store zero 4 bytes to ...
		{name: "MOVVstorezero", argLength: 2, reg: gpstore0, asm: "STL", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"}, // store zero 8 bytes to ...

		{name: "MOVBload", argLength: 2, reg: gpload, aux: "SymOff", asm: "LDBU", typ: "Int8", faultOnNilArg0: true, symEffect: "Read"},    // load from arg0 + auxInt + aux.  arg1=mem.
		{name: "MOVBUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "LDBU", typ: "UInt8", faultOnNilArg0: true, symEffect: "Read"},  // load from arg0 + auxInt + aux.  arg1=mem.
		{name: "MOVHload", argLength: 2, reg: gpload, aux: "SymOff", asm: "LDHU", typ: "Int16", faultOnNilArg0: true, symEffect: "Read"},   // load from arg0 + auxInt + aux.  arg1=mem.
		{name: "MOVHUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "LDHU", typ: "UInt16", faultOnNilArg0: true, symEffect: "Read"}, // load from arg0 + auxInt + aux.  arg1=mem.
		{name: "MOVWload", argLength: 2, reg: gpload, aux: "SymOff", asm: "LDW", typ: "Int32", faultOnNilArg0: true, symEffect: "Read"},    // load from arg0 + auxInt + aux.  arg1=mem.
		{name: "MOVWUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "LDW", typ: "UInt32", faultOnNilArg0: true, symEffect: "Read"},  // load from arg0 + auxInt + aux.  arg1=mem.
		{name: "MOVVload", argLength: 2, reg: gpload, aux: "SymOff", asm: "LDL", typ: "UInt64", faultOnNilArg0: true, symEffect: "Read"},   // load from arg0 + auxInt + aux.  arg1=mem.
		{name: "MOVFload", argLength: 2, reg: fpload, aux: "SymOff", asm: "FLDS", typ: "Float32", faultOnNilArg0: true, symEffect: "Read"}, // load from arg0 + auxInt + aux.  arg1=mem.
		{name: "MOVDload", argLength: 2, reg: fpload, aux: "SymOff", asm: "FLDD", typ: "Float64", faultOnNilArg0: true, symEffect: "Read"}, // load from arg0 + auxInt + aux.  arg1=mem.

		{name: "MOVBstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "STB", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},  // store 1 byte of arg1 to arg0 + auxInt + aux.  arg2=mem.
		{name: "MOVHstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "STH", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},  // store 2 bytes of arg1 to arg0 + auxInt + aux.  arg2=mem.
		{name: "MOVWstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "STW", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},  // store 4 bytes of arg1 to arg0 + auxInt + aux.  arg2=mem.
		{name: "MOVVstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "STL", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},  // store 8 bytes of arg1 to arg0 + auxInt + aux.  arg2=mem.
		{name: "MOVFstore", argLength: 3, reg: fpstore, aux: "SymOff", asm: "FSTS", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"}, // store 4 bytes of arg1 to arg0 + auxInt + aux.  arg2=mem.
		{name: "MOVDstore", argLength: 3, reg: fpstore, aux: "SymOff", asm: "FSTD", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"}, // store 8 bytes of arg1 to arg0 + auxInt + aux.  arg2=mem.

		{name: "CMPULE", argLength: 2, reg: gp21, asm: "CMPULE", typ: "Bool"}, // 1 if arg
		{name: "CMPULEconst", argLength: 1, typ: "Bool", reg: gp11, asm: "CMPULE", aux: "Int64"},
		{name: "CMPULT", argLength: 2, reg: gp21, asm: "CMPULT", typ: "Bool"}, // 1 if arg
		{name: "CMPULTconst", argLength: 1, typ: "Bool", reg: gp11, asm: "CMPULT", aux: "Int64"},
		{name: "CMPLE", argLength: 2, reg: gp21, asm: "CMPLE", typ: "Bool"}, // 1 if arg
		{name: "CMPLEconst", argLength: 1, typ: "Bool", reg: gp11, asm: "CMPLE", aux: "Int64"},
		{name: "CMPLT", argLength: 2, reg: gp21, asm: "CMPLT", typ: "Bool"}, // 1 if arg
		{name: "CMPLTconst", argLength: 1, typ: "Bool", reg: gp11, asm: "CMPLT", aux: "Int64"},
		{name: "CMPEQ", argLength: 2, reg: gp21, asm: "CMPEQ", typ: "Bool"},
		{name: "CMPEQconst", argLength: 1, typ: "Bool", reg: gp11, asm: "CMPEQ", aux: "Int64"},

		{name: "FCMPLE", argLength: 2, reg: ffp21, asm: "FCMPLE", typ: "Float64"}, // 1 if arg
		{name: "FCMPLT", argLength: 2, reg: ffp21, asm: "FCMPLT", typ: "Float64"}, // 1 if arg
		{name: "FCMPEQ", argLength: 2, reg: ffp21, asm: "FCMPEQ", typ: "Float64"},
		{name: "FCMPUN", argLength: 2, reg: ffp21, asm: "FCMPUN", typ: "Float64"},

		{name: "CALLstatic", argLength: 1, reg: regInfo{clobbers: callerSave}, aux: "SymOff", clobberFlags: true, call: true, symEffect: "None"},                           // call static function aux.(*obj.LSym).  arg0=mem, auxint=argsize, returns mem
		{name: "CALLclosure", argLength: 3, reg: regInfo{inputs: []regMask{gpsp, buildReg("R25"), 0}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true}, // call function via closure.  arg0=codeptr, arg1=closure, arg2=mem, auxint=argsize, returns mem
		{name: "CALLinter", argLength: 2, reg: regInfo{inputs: []regMask{gp}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},                         // call fn by pointer.  arg0=codeptr, arg1=mem, auxint=argsize, returns mem

		// duffzero
		{
			name:      "DUFFZERO",
			aux:       "Int64",
			argLength: 2,
			reg: regInfo{
				inputs:   []regMask{gp},
				clobbers: buildReg("R1"),
			},
			faultOnNilArg0: true,
		},

		//LoweredZero
		{
			name:      "LoweredZero",
			aux:       "Int64",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("R1"), gp},
				clobbers: buildReg("R1"),
			},
			clobberFlags:   true,
			faultOnNilArg0: true,
		},

		//LoweredMove
		{
			name:      "LoweredMove",
			aux:       "Int64",
			argLength: 4,
			reg: regInfo{
				inputs:   []regMask{buildReg("R2"), buildReg("R1"), gp},
				clobbers: buildReg("R1 R2"),
			},
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		// pseudo-ops
		{name: "LoweredNilCheck", argLength: 2, reg: regInfo{inputs: []regMask{gpg}}, nilCheck: true, faultOnNilArg0: true}, // panic if arg0 is nil.  arg1=mem.

		//zxw change
		// Scheduler ensures LoweredGetClosurePtr occurs only in entry block,
		// and sorts it to the very beginning of the block to prevent other
		// use of R25 (sw64.REGCTXT, the closure pointer)
		{name: "LoweredGetClosurePtr", reg: regInfo{outputs: []regMask{buildReg("R25")}}, zeroWidth: true},

		// LoweredGetCallerSP returns the SP of the caller of the current function.
		{name: "LoweredGetCallerSP", reg: gp01, rematerializeable: true},

		//zxw add
		// LoweredGetCallerPC evaluates to the PC to which its "caller" will return.
		// I.e., if f calls g "calls" getcallerpc,
		// the result should be the PC within f that g will return to.
		// See runtime/stubs.go for a more detailed discussion.
		{name: "LoweredGetCallerPC", reg: gp01, rematerializeable: true},

		// MOVDconvert converts between pointers and integers.
		// We have a special op for this so as to not confuse GC
		// (particularly stack maps).  It takes a memory arg so it
		// gets correctly ordered with respect to GC safepoints.
		// arg0=ptr/int arg1=mem, output=int/ptr
		{name: "MOVVconvert", argLength: 2, reg: gp11, asm: "LDI"},

		{name: "FEqual", argLength: 1, reg: readflags},
		{name: "FNotEqual", argLength: 1, reg: readflags},

		//zxw add
		// LoweredWB invokes runtime.gcWriteBarrier. arg0=destptr, arg1=srcptr, arg2=mem, aux=runtime.gcWriteBarrier
		// It saves all GP registers if necessary,
		// but clobbers R26 (LR) because it's a call
		// and R28 (REGTMP).
		{name: "LoweredWB", argLength: 3, reg: regInfo{inputs: []regMask{buildReg("R13"), buildReg("R14")}, clobbers: (callerSave &^ gpg) | buildReg("R26")}, clobberFlags: true, aux: "Sym", symEffect: "None"},

		//zxw new add
		// There are three of these functions so that they can have three different register inputs.
		// When we check 0 <= c <= cap (A), then 0 <= b <= c (B), then 0 <= a <= b (C), we want the
		// default registers to match so we don't need to copy registers around unnecessarily.
		{name: "LoweredPanicBoundsA", argLength: 3, aux: "Int64", reg: regInfo{inputs: []regMask{r3, r4}}, typ: "Mem"}, // arg0=idx, arg1=len, arg2=mem, returns memory. AuxInt contains report code (see PanicBounds in genericOps.go).
		{name: "LoweredPanicBoundsB", argLength: 3, aux: "Int64", reg: regInfo{inputs: []regMask{r2, r3}}, typ: "Mem"}, // arg0=idx, arg1=len, arg2=mem, returns memory. AuxInt contains report code (see PanicBounds in genericOps.go).
		{name: "LoweredPanicBoundsC", argLength: 3, aux: "Int64", reg: regInfo{inputs: []regMask{r1, r2}}, typ: "Mem"}, // arg0=idx, arg1=len, arg2=mem, returns memory. AuxInt contains report code (see PanicBounds in genericOps.go).

		//zxw new add
		// atomic loads.
		// load from arg0. arg1=mem.
		// returns <value,memory> so they can be properly ordered with other loads.
		{name: "LoweredAtomicLoad8", argLength: 2, reg: gpload, faultOnNilArg0: true},
		{name: "LoweredAtomicLoad32", argLength: 2, reg: gpload, faultOnNilArg0: true},
		{name: "LoweredAtomicLoad64", argLength: 2, reg: gpload, faultOnNilArg0: true},

		// atomic stores.
		// store arg1 to arg0. arg2=mem. returns memory.
		{name: "LoweredAtomicStore8", argLength: 3, reg: gpstore, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicStore32", argLength: 3, reg: gpstore, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicStore64", argLength: 3, reg: gpstore, faultOnNilArg0: true, hasSideEffects: true},

		// atomic exchange.
		// store arg1 to arg0. arg2=mem. returns <old content of *arg0, memory>.
		// MEMB
		// LLDx Rout, (Rarg0)
		// LDI	Rtmp, 1
		// WR_F	Rtmp
		// BIS  Rarg1, R31, Rtmp
		// LSTx Rtmp, (Rarg0)
		// RD_F	Rtmp
		// BEQ  Rtmp, -6(PC)
		// MEMB
		{name: "LoweredAtomicExchange32", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true, unsafePoint: true},
		{name: "LoweredAtomicExchange64", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true, unsafePoint: true},

		// atomic add.
		// *arg0 += arg1. arg2=mem. returns <new content of *arg0, memory>.
		// MEMB
		// LLDx	Rout, (Rarg0)
		// LDI	Rtmp, 1
		// WR_F	Rtmp
		// ADDL Rarg1, Rout, Rtmp
		// LSTx Rtmp, (Rarg0)
		// RD_F	Rtmp
		// BEQ  Rtmp, -6(PC)
		// ADDL Rarg1, Rout, Rout
		// MEMB
		{name: "LoweredAtomicAdd32", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true, unsafePoint: true},
		{name: "LoweredAtomicAdd64", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true, unsafePoint: true},

		// atomic compare and swap.
		// arg0 = pointer, arg1 = old value, arg2 = new value, arg3 = memory.
		// if *arg0 == arg1 {
		//   *arg0 = arg2
		//   return (true, memory)
		// } else {
		//   return (false, memory)
		// }
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
		{name: "LoweredAtomicCas32", argLength: 4, reg: gpcas, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true, unsafePoint: true},
		{name: "LoweredAtomicCas64", argLength: 4, reg: gpcas, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true, unsafePoint: true},
	}

	//zxw  new change
	blocks := []blockData{
		{name: "NE", controls: 1},
		{name: "EQ", controls: 1},
		{name: "LT", controls: 1},
		{name: "LE", controls: 1},
		{name: "GT", controls: 1},
		{name: "GE", controls: 1},
		{name: "FNE", controls: 1},
		{name: "FEQ", controls: 1},
		{name: "FLE", controls: 1},
		{name: "FLT", controls: 1},
		{name: "FGE", controls: 1},
		{name: "FGT", controls: 1},
	}

	archs = append(archs, arch{
		name:            "SW64",
		pkg:             "cmd/internal/obj/sw64",
		genfile:         "../../sw64/ssa.go",
		ops:             ops,
		blocks:          blocks,
		regnames:        regNamesSW64,
		gpregmask:       gp,
		fpregmask:       fp,
		framepointerreg: -1, // not used
		linkreg:         int8(num["R26"]),
	})
}
