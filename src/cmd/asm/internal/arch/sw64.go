// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the
// sw64 instruction set, to minimize its interaction
// with the core of the assembler.

package arch

import (
	"cmd/internal/obj"
	"cmd/internal/obj/sw64"
	"fmt"
)

func jumpSW64(word string) bool {
	switch word {
	case "CALL", "JMP",
		"BR", "BSR",
		"BEQ", "BNE", "BLT", "BLE", "BGT", "BGE",
		"BLBC", "BLBS", "FBEQ", "FBNE", "FBLT",
		"FBGT", "FBGE":
		return true
	}
	return false
}

func sw64RegisterNumber(name string, n int16) (int16, bool) {
	//snyh_TODO: update by cmd/internal/obj/sw64/a.out.go
	switch name {
	case "R":
		if 0 <= n && n <= 32 {
			return sw64.REG_R0 + n, true
		}
	case "F":
		if 0 <= n && n <= 32 {
			return sw64.REG_F0 + n, true
		}
	}
	return 0, false
}

func archSW64() *Arch {
	register := make(map[string]int16)
	for i := sw64.REG_R0; i <= sw64.REG_R31; i++ {
		register[obj.Rconv(i)] = int16(i)
	}
	for i := sw64.REG_R0; i <= sw64.REG_R31; i++ {
		register[fmt.Sprintf("R%d", i-sw64.REG_R0)] = int16(i)
	}
	for i := sw64.REG_F0; i <= sw64.REG_F31; i++ {
		register[obj.Rconv(i)] = int16(i)
	}

	// Pseudo-registers.
	register["SB"] = RSB
	register["FP"] = RFP
	register["PC"] = RPC

	registerPrefix := map[string]bool{
		"T": true,
		"S": true,
		"A": true,
		"F": true,
		"R": true,
	}

	instructions := make(map[string]obj.As)
	for i, s := range obj.Anames {
		instructions[s] = obj.As(i)
	}
	for i, s := range sw64.Anames {
		if obj.As(i) >= obj.A_ARCHSPECIFIC {
			instructions[s] = obj.As(i) + obj.ABaseSW64
		}
	}

	return &Arch{
		LinkArch:       &sw64.LinkSW64,
		Instructions:   instructions,
		Register:       register,
		RegisterPrefix: registerPrefix,
		RegisterNumber: sw64RegisterNumber,
		IsJump:         jumpSW64,
	}
}
