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

var pltHeaderSize int64
var internalLinkingCgo bool = false

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
	var targType sym.SymKind
	if targ != 0 {
		targType = ldr.SymType(targ)
	}

	const pcrel = 1
	switch r.Type() {
	default:
		if r.Type() >= objabi.ElfRelocOffset {
			ldr.Errorf(s, "unexpected relocation type %d (%s)", r.Type(), sym.RelocName(target.Arch, r.Type()))
			return false
		}

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_GPDISP):
		if targType == sym.SDYNIMPORT {
			ldr.Errorf(s, "unexpected relocation for dynamic symbol %s", ldr.SymName(targ))
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_SW64_GPDISP)
		su.SetRelocAdd(rIdx, r.Add()+4)
		return true

	// Handle relocations found in ELF object files.
	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_SREL16):
		if targType == sym.SDYNIMPORT {
			ldr.Errorf(s, "unexpected R_SW64_SREL16 relocation for dynamic symbol %s", ldr.SymName(targ))
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_PCREL)
		su.SetRelocAdd(rIdx, r.Add()+2)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_SREL32):
		if targType == sym.SDYNIMPORT {
			ldr.Errorf(s, "unexpected R_SW64_SREL16 relocation for dynamic symbol %s", ldr.SymName(targ))
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_PCREL)
		su.SetRelocAdd(rIdx, r.Add()+4)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_SREL64):
		if targType == sym.SDYNIMPORT {
			ldr.Errorf(s, "unexpected R_SW64_SREL64 relocation for dynamic symbol %s", ldr.SymName(targ))
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_PCREL)
		su.SetRelocAdd(rIdx, r.Add()+8)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_BRADDR):
		// TODO: do we need do something about it?
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_SW64_BRADDR)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_LITERAL_GOT):
		// TODO:as literal_got doesn't achieve completily we just ignore it
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_LITUSE),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_HINT):
		// TODO:this is not necessary for reloc ignore it
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_LITERAL):
		if targType == sym.SDYNIMPORT {
			// have symbol
			// literal need a plt slot in dynmic link
			addpltsym(target, ldr, syms, targ)
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_SW64_LITERAL)
		// if LITERAL contant a reg call,
		// it is already created a plt entry in .got,
		// so we change target to it.
		sv := ldr.SymGot(targ)
		splt := ldr.SymPlt(targ)
		if sv != -1 && splt != -1 {
			su.SetRelocSym(rIdx, syms.GOT)
			su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymGot(targ)))
			return true
		}
		// fall back to using GOT
		ld.AddGotSym(target, ldr, syms, targ, uint32(elf.R_SW64_GLOB_DAT))
		su.SetRelocSym(rIdx, syms.GOT)
		su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymGot(targ))+8)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_GPRELHIGH),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_GPRELLOW):
		if targType == sym.SDYNIMPORT {
			ldr.Errorf(s, "unexpected relocation for dynamic symbol %s", ldr.SymName(targ))
		}
		if targType == 0 || targType == sym.SXREF {
			ldr.Errorf(s, "unknown symbol %s", ldr.SymName(targ))
		}
		var rel objabi.RelocType
		if r.Type() == objabi.ElfRelocOffset+objabi.RelocType(elf.R_SW64_GPRELHIGH) {
			rel = objabi.R_SW64_GPRELHIGH
		} else {
			rel = objabi.R_SW64_GPRELLOW
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, rel)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_REFLONG),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_SW64_REFQUAD):
		if targType == sym.SDYNIMPORT {
			ldr.Errorf(s, "unexpected R_SW64_*ABS* relocation for dynamic symbol %s", ldr.SymName(targ))
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_ADDR)
		if target.IsPIE() && target.IsInternal() {
			// For internal linking PIE, this R_ADDR relocation cannot
			// be resolved statically. We need to generate a dynamic
			// relocation. Let the code below handle it.
			break
		}
		return true
	}

	// Reread the reloc to incorporate any changes in type above.
	relocs := ldr.Relocs(s)
	r = relocs.At(rIdx)

	switch r.Type() {
	case objabi.R_CALL,
		objabi.R_PCREL:
		if targType != sym.SDYNIMPORT {
			// nothing to do, the relocation will be laid out in reloc
			return true
		}
		if target.IsExternal() {
			// External linker will do this relocation.
			return true
		}
		// Internal linking.
		if r.Add() != 0 {
			ldr.Errorf(s, "PLT call with non-zero addend (%v)", r.Add())
		}
		// Build a PLT entry and change the relocation target to that entry.
		addpltsym(target, ldr, syms, targ)
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocSym(rIdx, syms.PLT)
		su.SetRelocAdd(rIdx, int64(ldr.SymPlt(targ)))
		return true

	case objabi.R_ADDR:
		if ldr.SymType(s) == sym.STEXT {
			// The code is asking for the address of an external
			// function. We provide it with the address of the
			// correspondent GOT symbol.
			ld.AddGotSym(target, ldr, syms, targ, uint32(elf.R_SW64_GLOB_DAT))
			su := ldr.MakeSymbolUpdater(s)
			su.SetRelocSym(rIdx, syms.GOT)
			su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymGot(targ)))
			return true
		}

		// Process dynamic relocations for the data sections.
		if target.IsPIE() && target.IsInternal() {
			// When internally linking, generate dynamic relocations
			// for all typical R_ADDR relocations. The exception
			// are those R_ADDR that are created as part of generating
			// the dynamic relocations and must be resolved statically.
			//
			// There are three phases relevant to understanding this:
			//
			//  dodata()  // we are here
			//  address() // symbol address assignment
			//  reloc()   // resolution of static R_ADDR relocs
			//
			// At this point symbol addresses have not been
			// assigned yet (as the final size of the .rela section
			// will affect the addresses), and so we cannot write
			// the Elf64_Rela.r_offset now. Instead we delay it
			// until after the 'address' phase of the linker is
			// complete. We do this via Addaddrplus, which creates
			// a new R_ADDR relocation which will be resolved in
			// the 'reloc' phase.
			//
			// These synthetic static R_ADDR relocs must be skipped
			// now, or else we will be caught in an infinite loop
			// of generating synthetic relocs for our synthetic
			// relocs.
			//
			// Furthermore, the rela sections contain dynamic
			// relocations with R_ADDR relocations on
			// Elf64_Rela.r_offset. This field should contain the
			// symbol offset as determined by reloc(), not the
			// final dynamically linked address as a dynamic
			// relocation would provide.
			switch ldr.SymName(s) {
			case ".dynsym", ".rela", ".rela.plt", ".got.plt", ".dynamic":
				return false
			}
		} else {
			// Either internally linking a static executable,
			// in which case we can resolve these relocations
			// statically in the 'reloc' phase, or externally
			// linking, in which case the relocation will be
			// prepared in the 'reloc' phase and passed to the
			// external linker in the 'asmb' phase.
			if ldr.SymType(s) != sym.SDATA && ldr.SymType(s) != sym.SRODATA {
				break
			}
		}

		// Generate R_SW64_RELATIVE relocations for best
		// efficiency in the dynamic linker.
		//
		// As noted above, symbol addresses have not been
		// assigned yet, so we can't generate the final reloc
		// entry yet. We ultimately want:
		//
		// r_offset = s + r.Off
		// r_info = R_SW64_RELATIVE
		// r_addend = targ + r.Add
		//
		// The dynamic linker will set *offset = base address +
		// addend.
		//
		// AddAddrPlus is used for r_offset and r_addend to
		// generate new R_ADDR relocations that will update
		// these fields in the 'reloc' phase.
		rela := ldr.MakeSymbolUpdater(syms.Rela)
		rela.AddAddrPlus(target.Arch, s, int64(r.Off()))
		if r.Siz() == 8 {
			rela.AddUint64(target.Arch, elf.R_INFO(0, uint32(elf.R_SW64_RELATIVE)))
		} else {
			ldr.Errorf(s, "unexpected relocation for dynamic symbol %s", ldr.SymName(targ))
		}
		rela.AddAddrPlus(target.Arch, targ, int64(r.Add()))
		// Not mark r done here. So we still apply it statically,
		// so in the file content we'll also have the right offset
		// to the relocation target. So it can be examined statically
		// (e.g. go version).
		return true
	}

	return false
}

func elfsetupplt(ctxt *ld.Link, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym) {
	internalLinkingCgo = true
	if plt.Size() == 0 {
		// TODO: this is a little different with gnu pltHeader

		// rel offset
		// subl  r27,r28,r25
		plt.AddUint32(ctxt.Arch, 0x437c0139)

		// set .got to r28
		// ldi r28, 17(r28)
		plt.AddUint32(ctxt.Arch, 0xfb9d8001)

		// s4subl r25,r25,r25
		plt.AddUint32(ctxt.Arch, 0x43390179)

		// ldi r28, 32767(r28)
		plt.AddUint32(ctxt.Arch, 0xfb9cffef)

		// load resolver to jump target
		// ldl r27,0(r28)
		plt.AddUint32(ctxt.Arch, 0x8f7c0000)

		// addl r25,r25,r25
		plt.AddUint32(ctxt.Arch, 0x43390119)

		// load map info to reg r28
		// ldl r28,8(r28)
		plt.AddUint32(ctxt.Arch, 0x8f9c0008)
		// jmp $r31,($r27),1f27a0
		plt.AddUint32(ctxt.Arch, 0x0ffb0000)

		// br r28, .plt
		plt.AddUint32(ctxt.Arch, 0x139ffff7)

		// check gotplt.size == 0
		if gotplt.Size() != 0 {
			ctxt.Errorf(gotplt.Sym(), "got.plt is not empty at the very beginning")
		}
		pltHeaderSize = plt.Size()

		// will be fill in ld.so
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
		base := uint32(val) & 0xffff0000
		var hi, lo int16
		if internalLinkingCgo {
			hi, lo = gpdispAddrDyn(pc, ldr, syms)
		} else {
			hi, lo = gpdispAddr(pc)
		}

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

	case objabi.R_SW64_GOTTPREL:
		base := uint32(val) & 0xffff0000
		v := ldr.SymAddr(rs) + int64(2*target.Arch.PtrSize) + r.Add()
		if base != uint32(val) {
			log.Fatalf("The R_SW64_GOTTPREL %v has been broken in %v.", r, s)
		}
		val := int64(base + uint32(uint16(v)))
		return val, noExtReloc, isOk

	case objabi.R_SW64_GPRELLOW, objabi.R_SW64_GPRELHIGH:
		var hi, lo int16
		if internalLinkingCgo {
			hi, lo = splitGPRelAddrDyn(ldr, r, syms)
		} else {
			hi, lo = splitGPRelAddr(ldr, r)
		}
		base := uint32(val) & 0xffff0000
		if base != uint32(val) {
			log.Fatalf("The R_SW64_GPRELxx %v has been broken in %v.", ldr.SymName(s), r, val, base)
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
		base := uint32(val) & 0xffff0000
		//TODO: As we process off set in adddynrel, we try doing this
		off := int64(r.Add()) - 8 - 32784
		if off > 32768 {
			log.Fatalf("off is too big, we can't handle it")
		}
		hi, lo := splitAddr(off)
		if r.Type() == objabi.R_SW64_LITERAL_GOT {
			val = int64(base + uint32(uint16(hi)))
		} else {
			val = int64(base + uint32(uint16(lo)))
		}
		return val, noExtReloc, isOk
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

func gpDynmic(ldr *loader.Loader, syms *ld.ArchSyms) int64 {
	return ldr.SymValue(syms.GOT) + 32784
}

func gpdispAddrDyn(pc int64, ldr *loader.Loader, syms *ld.ArchSyms) (hi int16, lo int16) {
	addr := gpDynmic(ldr, syms) - pc
	hi, lo = splitAddr(addr)
	if int64(hi)<<16+int64(lo) != addr {
		log.Fatalf("PC 0x%x is out of range when build GP displacement\n", pc)
	}
	return
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

func splitGPRelAddrDyn(ldr *loader.Loader, r loader.Reloc, syms *ld.ArchSyms) (hi int16, lo int16) {
	rs := ldr.ResolveABIAlias(r.Sym())
	addr := ldr.SymValue(rs) + r.Add() - gpDynmic(ldr, syms)
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

func addpltsym(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s loader.Sym) {
	if ldr.SymPlt(s) >= 0 {
		return
	}

	ld.Adddynsym(ldr, target, syms, s)

	if target.IsElf() {
		plt := ldr.MakeSymbolUpdater(syms.PLT)
		got := ldr.MakeSymbolUpdater(syms.GOT)
		rela := ldr.MakeSymbolUpdater(syms.RelaPLT)
		if plt.Size() == 0 {
			panic("plt is not set up")
		}

		// br $R28, pltHeaderEnd
		disp := plt.Size() - pltHeaderSize + 8
		br := 0x13e00000 | (-int32(disp/4) & 0x1fffff)
		plt.AddUint32(target.Arch, uint32(br))

		// got
		// create got entry for extern jump
		got.AddAddrPlus(target.Arch, plt.Sym(), plt.Size()-4)
		ldr.SetGot(s, int32(got.Size()))

		// rela
		rela.AddAddrPlus(target.Arch, got.Sym(), got.Size()-8)
		sDynid := ldr.SymDynid(s)

		rela.AddUint64(target.Arch, elf.R_INFO(uint32(sDynid), uint32(elf.R_SW64_JMP_SLOT)))
		rela.AddUint64(target.Arch, 0)

		ldr.SetPlt(s, int32(plt.Size())-16)
	} else {
		ldr.Errorf(s, "addpltsym: unsupported binary format")
	}
}
