package sw64asm

import (
	"fmt"
)

func ParseOPtable(raw uint32) OP {
	op := fetchOpcode(raw)
	class, ok := classTabs[int(op)]
	if !ok {
		return OP{
			OPcode: int(op),
			Name:   fmt.Sprintf("?0x%x", op),
		}
	}
	fn := fetchFncode(raw, class)
	name, ok := nameTabs[int(op)][int(fn)]
	if !ok {
		name = fmt.Sprintf("?0x%x:0x%x?", op, fn)
	}
	return OP{
		OPcode:  int(op),
		Fncode:  int(fn),
		Name:    name,
		Class:   class,
		IsFloat: name[1] == 'F',
	}
}

func fetchBit(v uint32, begin, width uint32) uint32 {
	t := v >> begin
	mask := (uint32(1) << width) - 1
	return t & mask
}

func fetchOpcode(v uint32) uint32 { return fetchBit(v, 26, 6) }
func fetchRa(v uint32) uint32     { return fetchBit(v, 21, 5) }
func fetchRb(v uint32) uint32     { return fetchBit(v, 16, 5) }
func fetchDisp(v uint32) uint32   { return fetchBit(v, 0, 16) }

func fetchFncode(v uint32, c int) uint32 {
	switch c {
	case OPC_SYSCALL:
		return fetchBit(v, 25, 1)
	case OPC_ARITHMETIC,
		OPC_ARITHMETIC_I:
		return fetchBit(v, 5, 8)
	case OPC_COMPLEX_ARITHMETIC,
		OPC_COMPLEX_ARITHMETIC_I:
		return fetchBit(v, 10, 6)
	case OPC_MISI_MEMORY:
		return fetchBit(v, 0, 16)
	case OPC_FUNC_MEMORY:
		return fetchBit(v, 12, 4)
	case OPC_CONTROL,
		OPC_MEMORY:
		return 0
	default:
		panic(fmt.Sprintf("Invalid instruction for %v %v", v, c))
	}
}

const (
	OPC_NULL    = iota // 类别        |31  26|25 21|20 16|15       5|4   0|
	OPC_SYSCALL        //系统调用指令, |Opcode|         Function           |
	OPC_CONTROL        //转移控制指令, |Opcode| Ra  |       disp           |
	OPC_MEMORY         //存储器指令,   |Opcode| Ra  | Rb  |  disp          |

	OPC_MISI_MEMORY // 杂项指令       |Opcode| Ra  | Rb  |  Function       |

	OPC_FUNC_MEMORY //带功能域的存     |Opcode| Ra  | Rb  | Function | disp |

	OPC_ARITHMETIC   //简单运算指令, |Opcode| Ra  | Rb  | Function | RC  |
	OPC_ARITHMETIC_I //简单运算指令, |Opcode| Ra  | Ib    |Function| RC  |
	//                             |31  26|25 21|20   13|12     5|4   0|

	OPC_COMPLEX_ARITHMETIC   // 浮点复合运算指令格式 |Opcode|  Fa  |  Fb  |Function | Fc  |Fd  |
	OPC_COMPLEX_ARITHMETIC_I // 浮点复合运算指令格式 |Opcode|  Fa  |  Fb  |Function | Ib  |Fd  |
	//                                            |31  26|25  21|20  16|15     10|9   5|4  0|
)

type OP struct {
	OPcode  int
	Fncode  int
	Name    string
	Class   int
	IsFloat bool
}
