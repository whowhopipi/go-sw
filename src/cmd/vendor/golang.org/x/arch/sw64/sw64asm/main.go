package sw64asm

import (
	"encoding/binary"
	"fmt"
	"io"
)

// An Inst is a single instruction.
type Inst struct {
	Op   OP       // Opcode mnemonic
	Enc  uint32   // Raw encoding bits.
	Args []string // Instruction arguments, in native SW64 order.
}

func immName(v uint32) string {
	return fmt.Sprintf("$0x%x", v)
}

func memory(rno uint32, offset uint32) string {
	return fmt.Sprintf("$%d(%s)", offset, iRegName(rno))
}

func targetName(pc uint64, name string, base int64) string {
	if base == 0 {
		return fmt.Sprintf("%x <%s>", pc, name)
	} else {
		return fmt.Sprintf("%x <%s+0x%x>", pc, name, base)
	}
}

// Decode decodes the 4 bytes in src as a single instruction.
func Decode(src []byte) (inst Inst, err error) {
	if len(src) < 4 {
		return Inst{}, nil
	}
	v := binary.LittleEndian.Uint32(src)
	i := Inst{
		Op:   ParseOPtable(v),
		Enc:  v,
		Args: make([]string, 0, 4),
	}
	return i, nil
}

func GoSyntax(i Inst, pc uint64, symname func(uint64) (string, uint64), text io.ReaderAt) string {
	switch i.Op.Class {
	case OPC_SYSCALL:
		i.AddArg(immName(fetchBit(i.Enc, 0, 25)))
	case OPC_CONTROL:
		i.AddArg(i.RegNameAt(0)(fetchRa(i.Enc)))
		pc_t := uint64(fetchBit(i.Enc, 0, 21)) + 4 + pc
		name, base := symname(pc_t)
		i.AddArg(targetName(pc_t, name, int64(pc_t-base)))
	case OPC_MEMORY:
		i.AddArg(i.RegNameAt(0)(fetchRa(i.Enc)))
		i.AddArg(memory(fetchRb(i.Enc), fetchDisp(i.Enc)))
	case OPC_FUNC_MEMORY:
		i.AddArg(i.RegNameAt(0)(fetchRa(i.Enc)))
		i.AddArg(memory(fetchRb(i.Enc), fetchDisp(i.Enc)))
	case OPC_MISI_MEMORY:
		i.AddArg(i.RegNameAt(0)(fetchRa(i.Enc)))
	case OPC_ARITHMETIC:
		i.AddArg(i.RegNameAt(0)(fetchRa(i.Enc)))
		i.AddArg(i.RegNameAt(1)(fetchRb(i.Enc)))
		i.AddArg(i.RegNameAt(2)(fetchBit(i.Enc, 0, 5)))
	case OPC_ARITHMETIC_I:
		i.AddArg(i.RegNameAt(0)(fetchRa(i.Enc)))
		i.AddArg(immName(fetchBit(i.Enc, 13, 8)))
		i.AddArg(i.RegNameAt(0)(fetchBit(i.Enc, 0, 5)))
	case OPC_COMPLEX_ARITHMETIC:
		i.AddArg(i.RegNameAt(0)(fetchRa(i.Enc)))
		i.AddArg(i.RegNameAt(1)(fetchRb(i.Enc)))
		i.AddArg(i.RegNameAt(2)(fetchBit(i.Enc, 5, 5)))
		i.AddArg(i.RegNameAt(3)(fetchBit(i.Enc, 0, 5)))
	case OPC_COMPLEX_ARITHMETIC_I:
		i.AddArg(i.RegNameAt(0)(fetchRa(i.Enc)))
		i.AddArg(immName(fetchBit(i.Enc, 13, 8)))
		i.AddArg(i.RegNameAt(2)(fetchBit(i.Enc, 5, 5)))
		i.AddArg(i.RegNameAt(3)(fetchBit(i.Enc, 0, 5)))
	}

	switch len(i.Args) {
	case 0:
		return fmt.Sprintf("%7s", i.Op.Name)
	case 1:
		return fmt.Sprintf("%7s  %3s", i.Op.Name, i.Args[0])
	case 2:
		return fmt.Sprintf("%7s  %3s, %v", i.Op.Name, i.Args[0], i.Args[1])
	case 3:
		return fmt.Sprintf("%7s  %3s, %v, %v", i.Op.Name, i.Args[0], i.Args[1], i.Args[2])
	default:
		return fmt.Sprintf("%7s  %3s, %v, %v", i.Op.Name, "?", "?", "?")
	}
}
func (i *Inst) AddArg(arg string) {
	if arg == "" {
		return
	}
	i.Args = append(i.Args, arg)
}
