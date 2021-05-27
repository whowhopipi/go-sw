// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux
// +build sw64

#include "textflag.h"

TEXT _rt0_sw64_linux(SB), NOFRAME|NOSPLIT, $0
	JMP	_rt0_sw64(SB)

TEXT _rt0_sw64_linux_lib(SB), NOSPLIT, $0
	CALL	_rt0_sw64_lib(SB)
	RET
