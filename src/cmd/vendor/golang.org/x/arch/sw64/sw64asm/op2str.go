package sw64asm

import (
	"fmt"
	"strings"
)

var classTabs = map[int]int{
	0x00: OPC_SYSCALL,
	0x01: OPC_MEMORY,
	0x02: OPC_MEMORY,
	0x03: OPC_MEMORY,
	0x04: OPC_CONTROL,
	0x05: OPC_CONTROL,
	0x06: OPC_MISI_MEMORY,
	0x08: OPC_FUNC_MEMORY,
	0x10: OPC_ARITHMETIC,
	0x11: OPC_ARITHMETIC,
	0x12: OPC_ARITHMETIC_I,
	0x13: OPC_ARITHMETIC_I,
	0x20: OPC_MEMORY,
	0x21: OPC_MEMORY,
	0x22: OPC_MEMORY,
	0x23: OPC_MEMORY,
	0x24: OPC_MEMORY,
	0x25: OPC_MEMORY,
	0x26: OPC_MEMORY,
	0x27: OPC_MEMORY,
	0x28: OPC_MEMORY,
	0x29: OPC_MEMORY,
	0x2A: OPC_MEMORY,
	0x2B: OPC_MEMORY,
	0x2C: OPC_MEMORY,
	0x2D: OPC_MEMORY,
	0x2E: OPC_MEMORY,
	0x2F: OPC_MEMORY,
	0x30: OPC_CONTROL,
	0x31: OPC_CONTROL,
	0x32: OPC_CONTROL,
	0x33: OPC_CONTROL,
	0x34: OPC_CONTROL,
	0x35: OPC_CONTROL,
	0x36: OPC_CONTROL,
	0x37: OPC_CONTROL,
	0x38: OPC_CONTROL,
	0x39: OPC_CONTROL,
	0x3A: OPC_CONTROL,
	0x3B: OPC_CONTROL,
	0x3C: OPC_CONTROL,
	0x3D: OPC_CONTROL,
	0x3e: OPC_MEMORY,
	0x3f: OPC_MEMORY,
	0x18: OPC_ARITHMETIC,
	0x19: OPC_COMPLEX_ARITHMETIC,
}

//                op      fn   name
var nameTabs = map[int]map[int]string{
	0x00: {0x0: "SYSCALL/B", 0x1: "SYSCALL"},
	0x01: {0x0: "CALL"},
	0x02: {0x0: "RET"},
	0x03: {0x0: "JMP"},
	0x04: {0x0: "BR"},
	0x05: {0x0: "BSR"},
	0x06: {
		0x0000: "MEMB",
		0x0001: "IMEMB",
		0x1000: "RD_F",
		0x1020: "WR_F",
	},
	0x08: {
		0x0: "LLDW",
		0x1: "LLDL",
		0x8: "LSTW",
		0x9: "LSTL",
	},
	0x10: {
		0x00: "ADDW",
		0x01: "SUBW",
		0x02: "S4ADDW",
		0x03: "S4SUBW",
		0x04: "S8ADDW",
		0x05: "S8SUBW",
		0x08: "ADDL",
		0x09: "SUBL",
		0x0a: "S4ADDL",
		0x0b: "S4SUBL",
		0x0c: "S8ADDL",
		0x0d: "S8SUBL",
		0x10: "MULW",
		0x18: "MULL",
		0x19: "UMULH",
		0x28: "CMPEQ",
		0x29: "CMPLT",
		0x2a: "CMPLE",
		0x2b: "CMPULT",
		0x2c: "CMPULE",
		0x38: "AND",
		0x39: "BIC",
		0x3a: "BIS",
		0x3b: "ORNOT",
		0x3c: "XOR",
		0x3d: "EQV",
		0x40: "INSLB",
		0x41: "INSLH",
		0x42: "INSLW",
		0x43: "INSLL",
		0x44: "INSHB",
		0x45: "INSHH",
		0x46: "INSHW",
		0x47: "INSHL",
		0x48: "SLL",
		0x49: "SRL",
		0x4a: "SRA",
		0x50: "EXTLB",
		0x51: "EXTLH",
		0x52: "EXTLW",
		0x53: "EXTLL",
		0x54: "EXTHB",
		0x55: "EXTHH",
		0x56: "EXTHW",
		0x57: "EXTHL",
		0x58: "CTPOP",
		0x59: "CTLZ",
		0x5a: "CTTZ",
		0x60: "MASKLB",
		0x61: "MASKLH",
		0x62: "MASKLW",
		0x63: "MASKLL",
		0x64: "MASKHB",
		0x65: "MASKHH",
		0x66: "MASKHW",
		0x67: "MASKHL",
		0x68: "ZAP",
		0x69: "ZAPNOT",
		0x6a: "SEXTB",
		0x6b: "SEXTH",
		0x6c: "CMPGEB",
		0x70: "FIMOVS",
		0x78: "FIMOVD",
	},
	0x11: {
		0x0: "SELEQ",
		0x1: "SELGE",
		0x2: "SELGT",
		0x3: "SELLE",
		0x4: "SELLT",
		0x5: "SELNE",
		0x6: "SELLBC",
		0x7: "SELLBS",
	},
	0x12: {
		0x00: "ADDW",
		0x01: "SUBW",
		0x02: "S4ADDW",
		0x03: "S4SUBW",
		0x04: "S8ADDW",
		0x05: "S8SUBW",
		0x08: "ADDL",
		0x09: "SUBL",
		0x0a: "S4ADDL",
		0x0b: "S4SUBL",
		0x0c: "S8ADDL",
		0x0d: "S8SUBL",
		0x10: "MULW",
		0x18: "MULL",
		0x19: "UMULH",
		0x28: "CMPEQ",
		0x29: "CMPLT",
		0x2a: "CMPLE",
		0x2b: "CMPULT",
		0x2c: "CMPULE",
		0x38: "AND",
		0x39: "BIC",
		0x3a: "BIS",
		0x3b: "ORNOT",
		0x3c: "XOR",
		0x3d: "EQV",
		0x40: "INSLB",
		0x41: "INSLH",
		0x42: "INSLW",
		0x43: "INSLL",
		0x44: "INSHB",
		0x45: "INSHH",
		0x46: "INSHW",
		0x47: "INSHL",
		0x48: "SLL",
		0x49: "SRL",
		0x4a: "SRA",
		0x50: "EXTLB",
		0x51: "EXTLH",
		0x52: "EXTLW",
		0x53: "EXTLL",
		0x54: "EXTHB",
		0x55: "EXTHH",
		0x56: "EXTHW",
		0x57: "EXTHL",
		0x60: "MASKLB",
		0x61: "MASKLH",
		0x62: "MASKLW",
		0x63: "MASKLL",
		0x64: "MASKHB",
		0x65: "MASKHH",
		0x66: "MASKHW",
		0x67: "MASKHL",
		0x68: "ZAP",
		0x69: "ZAPNOT",
		0x6a: "SEXTB",
		0x6b: "SEXTH",
		0x6c: "CMPGEB",
	},
	0x13: {
		0x0: "SELEQ",
		0x1: "SELGE",
		0x2: "SELGT",
		0x3: "SELLE",
		0x4: "SELLT",
		0x5: "SELNE",
		0x6: "SELLBC",
		0x7: "SELLBS",
	},
	0x20: {0: "LDBU"},
	0x21: {0: "LDHU"},
	0x22: {0: "LDW"},
	0x23: {0: "LDL"},
	0x24: {0: "LDL_U"},
	0x25: {0: "PRI_LD"},
	0x26: {0: "FLDS"},
	0x27: {0: "FLDD"},
	0x28: {0: "STB"},
	0x29: {0: "STH"},
	0x2A: {0: "STW"},
	0x2B: {0: "STL"},
	0x2C: {0: "STL_U"},
	0x2D: {0: "PRI_ST"},
	0x2E: {0: "FSTS"},
	0x2F: {0: "FSTD"},
	0x30: {0: "BEQ"},
	0x31: {0: "BNE"},
	0x32: {0: "BLT"},
	0x33: {0: "BLE"},
	0x34: {0: "BGT"},
	0x35: {0: "BGE"},
	0x36: {0: "BLBC"},
	0x37: {0: "BLBS"},
	0x38: {0: "FBEQ"},
	0x39: {0: "FBNE"},
	0x3A: {0: "FBLT"},
	0x3B: {0: "FBLE"},
	0x3C: {0: "FBGT"},
	0x3D: {0: "FBGE"},
	0x3e: {0: "LDI"},
	0x3f: {0: "LDIH"},
	0x18: {
		0x00: "FADDS",
		0x01: "FADDD",
		0x02: "FSUBS",
		0x03: "FSUBD",
		0x04: "FMULS",
		0x05: "FMULD",
		0x06: "FDIVS",
		0x07: "FDIVD",
		0x08: "FSQRTS",
		0x09: "FSQRTD",
		0x10: "FCMPEQ",
		0x11: "FCMPLE",
		0x12: "FCMPLT",
		0x13: "FCMPUN",
		0x20: "FCVTSD",
		0x21: "FCVTDS",
		0x22: "FCVTDL_G",
		0x23: "FCVTDL_P",
		0x24: "FCVTDL_Z",
		0x25: "FCVTDL_N",
		0x27: "FCVTDL",
		0x28: "FCVTWL",
		0x29: "FCVTLW",
		0x2D: "FCVTLS",
		0x2F: "FCVTLD",
		0x30: "FCPYS",
		0x31: "FCPYSE",
		0x32: "FCPYSN",
		0x40: "IFMOVS",
		0x41: "IFMOVD",
		0x50: "RFPCR",
		0x51: "WFPCR",
		0x54: "SETFPEC0",
		0x55: "SETFPEC1",
		0x56: "SETFPEC2",
		0x57: "SETFPEC3",
	},
	0x19: {
		0x00: "FMAS",
		0x01: "FMAD",
		0x02: "FMSS",
		0x03: "FMSD",
		0x04: "FNMAS",
		0x05: "FNMAD",
		0x06: "FNMSS",
		0x07: "FNMSD",
	},
}

var __iRegName = make(map[uint32]string)

func iRegName(v uint32) string {
	if len(__iRegName) == 0 {
		for i := uint32(0); i < 32; i++ {
			switch i {
			case 15:
				__iRegName[i] = fmt.Sprintf("G")
			case 25:
				__iRegName[i] = fmt.Sprintf("CTXT")
			case 26:
				__iRegName[i] = fmt.Sprintf("RA")
			case 28:
				__iRegName[i] = fmt.Sprintf("TMP")
			case 29:
				__iRegName[i] = fmt.Sprintf("GP")
			case 30:
				__iRegName[i] = fmt.Sprintf("SP")
			case 31:
				__iRegName[i] = fmt.Sprintf("ZERO")
			default:
				__iRegName[i] = fmt.Sprintf("R%d", i)
			}
		}
	}
	return __iRegName[v]
}

var __fRegName = make(map[uint32]string)

func fRegName(v uint32) string {
	if len(__fRegName) == 0 {
		for i := uint32(0); i < 32; i++ {
			__fRegName[i] = fmt.Sprintf("F%d", i)
		}
	}
	return __fRegName[v]
}
func nullName(_ uint32) string { return "" }

func (i Inst) RegNameAt(pos int) func(uint32) string {
	name := i.Op.Name

	switch {
	case name == "MEMB",
		name == "IMEMB",
		strings.HasPrefix(name, "SETFPEC"):
		return nullName
	case strings.HasPrefix(name, "FIMOV"):
		switch pos {
		case 0:
			return fRegName
		case 1:
			return nullName
		case 2:
			return iRegName
		}
	case strings.HasPrefix(name, "IFMOV"):
		switch pos {
		case 0:
			return iRegName
		case 1:
			return nullName
		case 2:
			return fRegName
		}
	default:
		if name[0] == 'F' {
			return fRegName
		}
	}

	return iRegName
}
