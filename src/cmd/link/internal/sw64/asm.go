// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw64

import (
	"cmd/internal/objabi"
	"cmd/link/internal/ld"
	"cmd/link/internal/loader"
	"cmd/link/internal/sym"
	"debug/elf"
	"log"
)

func gentext(ctxt *ld.Link, ldr *loader.Loader) {
	initfunc, addmoduledata := ld.PrepareAddmoduledata(ctxt)
	if initfunc == nil {
		return
	}

	o := func(op uint32) {
		initfunc.AddUint32(ctxt.Arch, op)
	}

	// 0000000000000000 <local.dso_init>:
	// 0: 00 00 bb ff ldih ldih$r29,0($r27)
	//  0: R_SW_64_GPDISP .text+0x4
	// 4: 00 00 bd fb ldi $r29,0($r29)
	o(0xffbb0000)
	rel0, _ := initfunc.AddRel(objabi.R_SW64_GPDISP)
	rel0.SetOff(0)
	rel0.SetSiz(4)
	rel0.SetAdd(4)
	//TODO: how to find a new way to find pc symbol?
	rel0.SetSym(ldr.LookupOrCreateSym("go.link.addmoduledata", 0))
	o(0xfbbd0000)

	// 8: 00 00 1d fc ldih  r0, 0(r29) <runtime.firstmoduledata>
	//  0: R_SW_64_GPRELHI local.moduledata
	// c: 00 00 00 f8 ldi r0, 0(r0)
	//  4: R_SW_64_GPRELLO  local.moduledata
	o(0xfc1d0000)
	rel, _ := initfunc.AddRel(objabi.R_SW64_GPRELHIGH)
	rel.SetOff(8)
	rel.SetSiz(4)
	rel.SetSym(ctxt.Moduledata)

	o(0xf8000000)
	rel1, _ := initfunc.AddRel(objabi.R_SW64_GPRELLOW)
	rel1.SetOff(12)
	rel1.SetSiz(4)
	rel1.SetSym(ctxt.Moduledata)

	// 10: 00 00 7d ff ldih r27, 0(r29) <runtime.addmoduledata>
	//  8: R_SW_64_GPRELHI runtime.addmoduledata
	// 14: 00 00 7b fb ldi r27, 0(r27) <runtime.addmoduledata>
	//  12: R_SW_64_GPRELLO  runtime.addmoduledata
	// 18: 00 00 fb 0f  jmp (r31),(r27),1
	o(0xff7d0000)
	rel2, _ := initfunc.AddRel(objabi.R_SW64_GPRELHIGH)
	rel2.SetOff(16)
	rel2.SetSiz(4)
	rel2.SetSym(addmoduledata)
	o(0xfb7b0000)
	rel3, _ := initfunc.AddRel(objabi.R_SW64_GPRELLOW)
	rel3.SetOff(20)
	rel3.SetSiz(4)
	rel3.SetSym(addmoduledata)
	o(0x0ffb0000)

	// filled nop
	o(0x43ff075f)
	o(0x43ff075f)
	o(0x1bff0080) //unreachable
}

func adddynrel(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s loader.Sym, r loader.Reloc, rIdx int) bool {
	//TODO
	targ := r.Sym()
	rType := r.Type()
	Name := ldr.SymName(targ)
	log.Println("rType is:", rType, "Name is :", Name)
	return true
}

func elfsetupplt(ctxt *ld.Link, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym) {
	if plt.Size() == 0 {
		// br r27, pc+4
		// identifying information
		plt.AddUint32(ctxt.Arch, 0x13600000)

		// nop
		plt.SetUint32(ctxt.Arch, plt.Size()-4, 0x43ff075f)

		// ldl r27, 12(r27)
		plt.SetUint32(ctxt.Arch, plt.Size()-4, 0x8f7b000c)

		// jmp r27, r27, 0
		plt.SetUint32(ctxt.Arch, plt.Size()-4, 0x0f7b0000)

		// check gotplt.size == 0
		if gotplt.Size() != 0 {
			ctxt.Errorf(gotplt.Sym(), "got.plt is not empty at the very beginning")
		}
		gotplt.AddAddrPlus(ctxt.Arch, dynamic, 0)

		// fill by ld.so
		gotplt.AddUint64(ctxt.Arch, 0)
		gotplt.AddUint64(ctxt.Arch, 0)
	}
	return
}

func elfreloc1(ctxt *ld.Link, out *ld.OutBuf, ldr *loader.Loader, s loader.Sym, r loader.ExtReloc, ri int, sectoff int64) bool {
	out.Write64(uint64(sectoff))

	elfsym := ld.ElfSymForReloc(ctxt, r.Xsym)
	switch r.Type {
	case objabi.R_SW64_HINT:
		out.Write64(uint64(elf.R_SW64_HINT) | uint64(elfsym)<<32)
	case objabi.R_SW64_GPDISP:
		out.Write64(uint64(elf.R_SW64_GPDISP) | uint64(elfsym)<<32)
	case objabi.R_SW64_BRADDR:
		out.Write64(uint64(elf.R_SW64_BRADDR) | uint64(elfsym)<<32)
	case objabi.R_SW64_GPRELHIGH:
		out.Write64(uint64(elf.R_SW64_GPRELHIGH) | uint64(elfsym)<<32)
	case objabi.R_SW64_GPRELLOW:
		out.Write64(uint64(elf.R_SW64_GPRELLOW) | uint64(elfsym)<<32)
	case objabi.R_SW64_LITERAL_GOT:
		out.Write64(uint64(elf.R_SW64_LITERAL_GOT) | uint64(elfsym)<<32)
	case objabi.R_SW64_LITERAL:
		out.Write64(uint64(elf.R_SW64_LITERAL) | uint64(elfsym)<<32)
	case objabi.R_SW64_TPRELHI:
		out.Write64(uint64(39) | uint64(elfsym)<<32)
	case objabi.R_SW64_TPRELLO:
		out.Write64(uint64(40) | uint64(elfsym)<<32)
	case objabi.R_SW64_GOTTPREL:
		out.Write64(uint64(37) | uint64(elfsym)<<32)
	case objabi.R_ADDR, objabi.R_DWARFSECREF:
		switch r.Size {
		case 4:
			out.Write64(uint64(elf.R_SW64_REFLONG) | uint64(elfsym)<<32)
		case 8:
			out.Write64(uint64(elf.R_SW64_REFQUAD) | uint64(elfsym)<<32)
		default:
			return false
		}
	case objabi.R_CALL, objabi.R_CALLIND:
		return true
	default:
		return false
	}
	out.Write64(uint64(r.Xadd))
	return true
}

//zxw new change
func archreloc(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, r loader.Reloc, s loader.Sym, val int64) (o int64, nExtReloc int, ok bool) {
	if target.IsExternal() {
		switch r.Type() {
		case
			objabi.R_CALL, objabi.R_CALLIND,
			objabi.R_SW64_GPDISP,
			objabi.R_SW64_BRADDR,
			objabi.R_SW64_HINT,
			objabi.R_SW64_GOTTPREL,
			objabi.R_SW64_GPRELHIGH, objabi.R_SW64_GPRELLOW,
			objabi.R_SW64_LITERAL_GOT, objabi.R_SW64_LITERAL,
			objabi.R_SW64_TPRELHI, objabi.R_SW64_TPRELLO:
			return val, 1, true
		default:
			return val, 0, false
		}
	}

	const isOk = true
	const noExtReloc = 0
	rs := r.Sym()
	rs = ldr.ResolveABIAlias(rs)
	switch r.Type() {
	case objabi.R_CALL, objabi.R_CALLIND:
		return val, noExtReloc, isOk
	case objabi.R_SW64_GPDISP:
		pc := ldr.SymValue(rs) + int64(r.Off())
		hi, lo := gpdispAddr(pc)

		base := uint32(val) & 0xffff0000
		if base != uint32(val) {
			log.Fatalf("The R_SW64_GPDISP %v has been broken in %v.", r, s)
		}
		val = int64((uint32(hi) << 16) + uint32(uint16(lo)))
		return val, noExtReloc, isOk
	case objabi.R_SW64_TPRELHI, objabi.R_SW64_TPRELLO:
		hi, lo := splitSymAddr(ldr, r, 16)
		base := uint32(val) & 0xffff0000
		if base != uint32(val) {
			log.Fatalf("The R_SW64_TPRELLO/HI %v has been broken in %v.", r, s)
		}
		if r.Type() == objabi.R_SW64_TPRELHI {
			val = int64(base + uint32(uint16(hi)))
		} else {
			val = int64(base + uint32(uint16(lo)))
		}
		return val, noExtReloc, isOk

	case objabi.R_SW64_GPRELLOW, objabi.R_SW64_GPRELHIGH:
		hi, lo := splitGPRelAddr(ldr, r)
		base := uint32(val) & 0xffff0000
		if base != uint32(val) {
			log.Fatalf("The R_SW64_GPRELxx %v has been broken in %v.", r, s)
		}
		if r.Type() == objabi.R_SW64_GPRELHIGH {
			val = int64(base + uint32(uint16(hi)))
		} else {
			val = int64(base + uint32(uint16(lo)))
		}
		return val, noExtReloc, isOk

	case objabi.R_SW64_BRADDR:
		off := (ldr.SymValue(rs)+r.Add()-(ldr.SymValue(s)+int64(r.Off())))/4 - 1
		mask := (int64(1) << 21) - 1
		disp := (int64(1) << 20) - 1
		if off > disp || off < -disp {
			log.Fatalf("BRADDR from %s to %s is too long %v\n",
				s, r.Sym(), off)
		}
		off &= mask
		val = off + val
		return val, noExtReloc, isOk
	case objabi.R_SW64_HINT:
		off := (ldr.SymValue(rs)+r.Add()-(ldr.SymValue(s)+int64(r.Off())))/4 - 1
		mask := (int64(1) << 16) - 1
		if int64(int16(off)) != off {
			return val, noExtReloc, isOk
		}
		off &= mask
		val = off + val
		return val, noExtReloc, isOk
	case objabi.R_SW64_LITERAL_GOT, objabi.R_SW64_LITERAL:
		log.Fatalf("literal reloc shouldn't process in internal link mode")
	}
	return val, 0, false
}

func splitAddr(addr int64) (hi int16, lo int16) {
	hi = int16(addr >> 16)
	lo = int16(addr & 0xffff)
	if lo < 0 {
		hi = hi + 1
		lo = int16(addr - int64(hi)<<16)
	}
	return
}

func gpAddr() int64 {
	return 0x7fffffff
}

func gpdispAddr(pc int64) (hi int16, lo int16) {
	addr := gpAddr() - pc
	hi, lo = splitAddr(addr)
	if int64(hi)<<16+int64(lo) != addr {
		log.Fatalf("PC 0x%x is out of range when build GP displacement\n", pc)
	}
	return
}

func splitGPRelAddr(ldr *loader.Loader, r loader.Reloc) (hi int16, lo int16) {
	rs := ldr.ResolveABIAlias(r.Sym())
	addr := ldr.SymValue(rs) + r.Add() - gpAddr()
	hi, lo = splitAddr(addr)
	if int64(hi)<<16+int64(lo) != addr {
		log.Fatalf("Symbol %q is out of range when split GP relative address\n", r.Sym())
	}
	return
}

// splitSymaddr split address of s to two 16 signed bit
func splitSymAddr(ldr *loader.Loader, r loader.Reloc, off int64) (hi int16, lo int16) {
	rs := ldr.ResolveABIAlias(r.Sym())
	addr := ldr.SymValue(rs) + r.Add() + off
	hi, lo = splitAddr(addr)
	if int64(hi)<<16+int64(lo) != addr {
		log.Fatalf("Symbol %q is out of range when split symbol address\n",
			r.Sym())
	}
	return
}

func archrelocvariant(*ld.Target, *loader.Loader, loader.Reloc, sym.RelocVariant, loader.Sym, int64) int64 {
	log.Fatalf("unexpected relocation variant")
	return -1
}

func extreloc(target *ld.Target, ldr *loader.Loader, r loader.Reloc, s loader.Sym) (loader.ExtReloc, bool) {
	switch r.Type() {
	case objabi.R_CALL, objabi.R_CALLIND,
		objabi.R_SW64_GPDISP,
		objabi.R_SW64_BRADDR,
		objabi.R_SW64_HINT,
		objabi.R_SW64_GOTTPREL,
		objabi.R_SW64_LITERAL_GOT, objabi.R_SW64_LITERAL,
		objabi.R_SW64_TPRELHI, objabi.R_SW64_TPRELLO:
		return ld.ExtrelocViaOuterSym(ldr, r, s), true
	case objabi.R_SW64_GPRELHIGH, objabi.R_SW64_GPRELLOW:
		return ld.ExtrelocViaOuterSym(ldr, r, s), true
	}
	return loader.ExtReloc{}, false
}
