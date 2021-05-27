// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sw64

import (
	"cmd/internal/objabi"
	"cmd/link/internal/ld"
	"cmd/link/internal/sym"
	"debug/elf"
	"fmt"
	"log"
)

func gentext(ctxt *ld.Link) {
	if !ctxt.DynlinkingGo() {
		return
	}
	addmoduledata := ctxt.Syms.Lookup("runtime.addmoduledata", 0)
	if addmoduledata.Type == sym.STEXT {
		// we're linking a module containing the runtime -> no need for
		// an init function
		return
	}
	panic("Not implement")
}

func adddynrel(ctxt *ld.Link, s *sym.Symbol, r *sym.Reloc) bool {
	log.Fatalf("adddynrel not implemented for %v with %v\n", s, r)
	return false
}

func elfsetupplt(ctxt *ld.Link) {
	// TODO(aram)
	return
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	ctxt.Out.Write64(uint64(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	switch r.Type {
	case objabi.R_SW64_HINT:
		ctxt.Out.Write64(uint64(elf.R_SW64_HINT) | uint64(elfsym)<<32)
	case objabi.R_SW64_GPDISP:
		ctxt.Out.Write64(uint64(elf.R_SW64_GPDISP) | uint64(elfsym)<<32)
	case objabi.R_SW64_BRADDR:
		ctxt.Out.Write64(uint64(elf.R_SW64_BRADDR) | uint64(elfsym)<<32)
	case objabi.R_SW64_GPRELHIGH:
		ctxt.Out.Write64(uint64(elf.R_SW64_GPRELHIGH) | uint64(elfsym)<<32)
	case objabi.R_SW64_GPRELLOW:
		ctxt.Out.Write64(uint64(elf.R_SW64_GPRELLOW) | uint64(elfsym)<<32)
	case objabi.R_SW64_TPRELHI:
		ctxt.Out.Write64(uint64(39) | uint64(elfsym)<<32)
	case objabi.R_SW64_TPRELLO:
		ctxt.Out.Write64(uint64(40) | uint64(elfsym)<<32)
	case objabi.R_SW64_GOTTPREL:
		ctxt.Out.Write64(uint64(37) | uint64(elfsym)<<32)
	case objabi.R_ADDR:
		switch r.Siz {
		case 4:
			ctxt.Out.Write64(uint64(elf.R_SW64_REFLONG) | uint64(elfsym)<<32)
		case 8:
			ctxt.Out.Write64(uint64(elf.R_SW64_REFQUAD) | uint64(elfsym)<<32)
		default:
			return false
		}
	case objabi.R_CALL, objabi.R_CALLIND:
		return true
	default:
		return false
	}
	ctxt.Out.Write64(uint64(r.Xadd))
	return true
}

//zxw new change
func archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val int64) (int64, bool) {
	if ctxt.LinkMode == ld.LinkExternal {
		switch r.Type {
		case
			objabi.R_CALL, objabi.R_CALLIND,
			objabi.R_SW64_GPDISP,
			objabi.R_SW64_BRADDR,
			objabi.R_SW64_HINT,
			objabi.R_SW64_GOTTPREL,
			objabi.R_SW64_TPRELHI, objabi.R_SW64_TPRELLO:
			r.Done = false
			r.Xadd = r.Add
			r.Xsym = r.Sym
			return val, true
		case objabi.R_SW64_GPRELHIGH, objabi.R_SW64_GPRELLOW:
			r.Done = false
			rs := r.Sym
			r.Xadd = r.Add
			for rs.Outer != nil {
				r.Xadd += ld.Symaddr(rs) - ld.Symaddr(rs.Outer)
				rs = rs.Outer
			}

			if rs.Type != sym.SHOSTOBJ && rs.Type != sym.SDYNIMPORT && rs.Type != sym.SUNDEFEXT && rs.Sect == nil {
				ld.Errorf(s, "missing section for %s", rs.Name)
			}
			r.Xsym = rs
			return val, true
		default:
			return val, false
		}
	}

	switch r.Type {
	case objabi.R_CALL, objabi.R_CALLIND:
		return val, true
	case objabi.R_SW64_GPDISP:
		pc := ld.Symaddr(s) + int64(r.Off)
		hi, lo := gpdispAddr(pc)

		_val := int64(ctxt.Arch.ByteOrder.Uint32(s.P[int64(r.Off):]))
		if _val != val {
			panic("Internal SW64 bug")
		}
		base := uint32(val) & 0xffff0000
		if base != uint32(val) {
			log.Fatalf("The R_SW64_GPDISP %v has been broken in %v.", r, s)
		}
		val = int64(base + uint32(uint16(hi)))

		nextPosition := int64(r.Off) + r.Add
		base2 := (ctxt.Arch.ByteOrder.Uint32(s.P[nextPosition:])) & 0xffff0000
		ctxt.Arch.ByteOrder.PutUint32(s.P[nextPosition:], base2+uint32(uint16(lo)))
		return val, true
	case objabi.R_SW64_TPRELHI, objabi.R_SW64_TPRELLO:
		hi, lo := splitSymAddr(r, 16)
		base := uint32(val) & 0xffff0000
		if base != uint32(val) {
			log.Fatalf("The R_SW64_TPRELLO/HI %v has been broken in %v.", r, s)
		}
		if r.Type == objabi.R_SW64_TPRELHI {
			val = int64(base + uint32(uint16(hi)))
		} else {
			val = int64(base + uint32(uint16(lo)))
		}
		return val, true

	case objabi.R_SW64_GPRELLOW, objabi.R_SW64_GPRELHIGH:
		hi, lo := splitGPRelAddr(r)
		base := uint32(val) & 0xffff0000
		if base != uint32(val) {
			log.Fatalf("The R_SW64_GPRELxx %v has been broken in %v.", r, s)
		}
		if r.Type == objabi.R_SW64_GPRELHIGH {
			val = int64(base + uint32(uint16(hi)))
		} else {
			val = int64(base + uint32(uint16(lo)))
		}
		return val, true

	case objabi.R_SW64_BRADDR:
		off := (ld.Symaddr(r.Sym)+r.Add-(ld.Symaddr(s)+int64(r.Off)))/4 - 1
		mask := (int64(1) << 21) - 1
		disp := (int64(1) << 20) - 1
		if off > disp || off < -disp {
			log.Fatalf("BRADDR from %s to %s is too long %v\n",
				s, r.Sym, off)
		}
		off &= mask
		val = off + val
		return val, true
	case objabi.R_SW64_HINT:
		off := (ld.Symaddr(r.Sym)+r.Add-(ld.Symaddr(s)+int64(r.Off)))/4 - 1
		mask := (int64(1) << 16) - 1
		if int64(int16(off)) != off {
			return val, true
		}
		off &= mask
		val = off + val
		return val, true
	}
	return val, false
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

func splitGPRelAddr(r *sym.Reloc) (hi int16, lo int16) {
	addr := ld.Symaddr(r.Sym) + r.Add - gpAddr()
	hi, lo = splitAddr(addr)
	if int64(hi)<<16+int64(lo) != addr {
		log.Fatalf("Symbol %q is out of range when split GP relative address\n",
			r.Sym.Name)
	}
	return
}

// splitSymaddr split address of s to two 16 signed bit
func splitSymAddr(r *sym.Reloc, off int64) (hi int16, lo int16) {
	addr := ld.Symaddr(r.Sym) + r.Add + off
	hi, lo = splitAddr(addr)
	if int64(hi)<<16+int64(lo) != addr {
		log.Fatalf("Symbol %q is out of range when split symbol address\n",
			r.Sym.Name)
	}
	return
}

func archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	log.Fatalf("unexpected relocation variant")
	return -1
}

//zxw new change
func asmb(ctxt *ld.Link) {
	if ctxt.IsELF {
		ld.Asmbelfsetup()
	}

	sect := ld.Segtext.Sections[0]
	ctxt.Out.SeekSet(int64(sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff))
	ld.Codeblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	for _, sect = range ld.Segtext.Sections[1:] {
		ctxt.Out.SeekSet(int64(sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff))
		ld.Datblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	}

	if ld.Segrodata.Filelen > 0 {
		ctxt.Out.SeekSet(int64(ld.Segrodata.Fileoff))
		ld.Datblk(ctxt, int64(ld.Segrodata.Vaddr), int64(ld.Segrodata.Filelen))
	}
	if ld.Segrelrodata.Filelen > 0 {
		ctxt.Out.SeekSet(int64(ld.Segrelrodata.Fileoff))
		ld.Datblk(ctxt, int64(ld.Segrelrodata.Vaddr), int64(ld.Segrelrodata.Filelen))
	}

	ctxt.Out.SeekSet(int64(ld.Segdata.Fileoff))
	ld.Datblk(ctxt, int64(ld.Segdata.Vaddr), int64(ld.Segdata.Filelen))

	ctxt.Out.SeekSet(int64(ld.Segdwarf.Fileoff))
	ld.Dwarfblk(ctxt, int64(ld.Segdwarf.Vaddr), int64(ld.Segdwarf.Filelen))

}

//zxw new add
func asmb2(ctxt *ld.Link) {
	machlink := uint32(0)
	if ctxt.HeadType == objabi.Hdarwin {
		machlink = uint32(ld.Domacholink(ctxt))
	}
	/* output symbol table */
	ld.Symsize = 0

	ld.Lcsize = 0
	symo := uint32(0)
	if !*ld.FlagS {
		// TODO: rationalize
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				symo = uint32(ld.Segdwarf.Fileoff + ld.Segdwarf.Filelen)
				symo = uint32(ld.Rnd(int64(symo), int64(*ld.FlagRound)))
			}

		case objabi.Hplan9:
			symo = uint32(ld.Segdata.Fileoff + ld.Segdata.Filelen)

		case objabi.Hdarwin:
			symo = uint32(ld.Segdwarf.Fileoff + uint64(ld.Rnd(int64(ld.Segdwarf.Filelen), int64(*ld.FlagRound))) + uint64(machlink))
		}

		ctxt.Out.SeekSet(int64(symo))
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				ld.Asmelfsym(ctxt)
				ctxt.Out.Flush()
				ctxt.Out.Write(ld.Elfstrdat)

				if ctxt.LinkMode == ld.LinkExternal {
					ld.Elfemitreloc(ctxt)
				}
			}

		case objabi.Hplan9:
			ld.Asmplan9sym(ctxt)
			ctxt.Out.Flush()

			sym := ctxt.Syms.Lookup("pclntab", 0)
			if sym != nil {
				ld.Lcsize = int32(len(sym.P))
				ctxt.Out.Write(sym.P)
				ctxt.Out.Flush()
			}

		case objabi.Hdarwin:
			if ctxt.LinkMode == ld.LinkExternal {
				ld.Machoemitreloc(ctxt)
			}
		}
	}

	ctxt.Out.SeekSet(0)
	switch ctxt.HeadType {
	default:
	case objabi.Hplan9: /* plan 9 */
		ctxt.Out.Write32(0x647)                      /* magic */
		ctxt.Out.Write32(uint32(ld.Segtext.Filelen)) /* sizes */
		ctxt.Out.Write32(uint32(ld.Segdata.Filelen))
		ctxt.Out.Write32(uint32(ld.Segdata.Length - ld.Segdata.Filelen))
		ctxt.Out.Write32(uint32(ld.Symsize))          /* nsyms */
		ctxt.Out.Write32(uint32(ld.Entryvalue(ctxt))) /* va of entry */
		ctxt.Out.Write32(0)
		ctxt.Out.Write32(uint32(ld.Lcsize))

		//zxw new change
	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd:
		ld.Asmbelf(ctxt, int64(symo))

	case objabi.Hdarwin:
		ld.Asmbmacho(ctxt)
	}

	ctxt.Out.Flush()
	if *ld.FlagC {
		fmt.Printf("textsize=%d\n", ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", ld.Segdata.Length-ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", ld.Symsize)
		fmt.Printf("lcsize=%d\n", ld.Lcsize)
		fmt.Printf("total=%d\n", ld.Segtext.Filelen+ld.Segdata.Length+uint64(ld.Symsize)+uint64(ld.Lcsize))
	}
}
