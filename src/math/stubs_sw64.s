// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "textflag.h"

TEXT ·Asin(SB), NOFRAME|NOSPLIT, $0
	JMP	·asin(SB)

TEXT ·Acos(SB), NOFRAME|NOSPLIT, $0
	JMP	·acos(SB)

TEXT ·Asinh(SB), NOFRAME|NOSPLIT, $0
	JMP	·asinh(SB)

TEXT ·Acosh(SB), NOFRAME|NOSPLIT, $0
	JMP	·acosh(SB)

TEXT ·Atan2(SB), NOFRAME|NOSPLIT, $0
	JMP	·atan2(SB)

TEXT ·Atan(SB), NOFRAME|NOSPLIT, $0
	JMP	·atan(SB)

TEXT ·Atanh(SB), NOFRAME|NOSPLIT, $0
	JMP	·atanh(SB)

TEXT ·Min(SB), NOFRAME|NOSPLIT, $0
	JMP	·min(SB)

TEXT ·Max(SB), NOFRAME|NOSPLIT, $0
	JMP	·max(SB)

TEXT ·Erf(SB), NOFRAME|NOSPLIT, $0
	JMP	·erf(SB)

TEXT ·Erfc(SB), NOFRAME|NOSPLIT, $0
	JMP	·erfc(SB)

TEXT ·Exp2(SB), NOFRAME|NOSPLIT, $0
	JMP	·exp2(SB)

TEXT ·Expm1(SB), NOFRAME|NOSPLIT, $0
	JMP	·expm1(SB)

TEXT ·Exp(SB), NOFRAME|NOSPLIT, $0
	JMP	·exp(SB)

TEXT ·Frexp(SB), NOFRAME|NOSPLIT, $0
	JMP	·frexp(SB)

TEXT ·Hypot(SB), NOFRAME|NOSPLIT, $0
	JMP	·hypot(SB)

TEXT ·Ldexp(SB), NOFRAME|NOSPLIT, $0
	JMP	·ldexp(SB)

TEXT ·Log10(SB), NOFRAME|NOSPLIT, $0
	JMP	·log10(SB)

TEXT ·Log2(SB), NOFRAME|NOSPLIT, $0
	JMP	·log2(SB)

TEXT ·Log1p(SB), NOFRAME|NOSPLIT, $0
	JMP	·log1p(SB)

TEXT ·Log(SB), NOFRAME|NOSPLIT, $0
	JMP	·log(SB)

TEXT ·Mod(SB), NOFRAME|NOSPLIT, $0
	JMP	·mod(SB)

TEXT ·Modf(SB), NOFRAME|NOSPLIT, $0
	JMP	·modf(SB)

TEXT ·Remainder(SB), NOFRAME|NOSPLIT, $0
	JMP	·remainder(SB)

TEXT ·Sin(SB), NOFRAME|NOSPLIT, $0
	JMP	·sin(SB)

TEXT ·Sinh(SB), NOFRAME|NOSPLIT, $0
	JMP	·sinh(SB)

TEXT ·Cos(SB), NOFRAME|NOSPLIT, $0
	JMP	·cos(SB)

TEXT ·Cosh(SB), NOFRAME|NOSPLIT, $0
	JMP	·cosh(SB)

TEXT ·Tan(SB), NOFRAME|NOSPLIT, $0
	JMP	·tan(SB)

TEXT ·Tanh(SB), NOFRAME|NOSPLIT, $0
	JMP	·tanh(SB)

TEXT ·Cbrt(SB), NOFRAME|NOSPLIT, $0
	JMP	·cbrt(SB)

TEXT ·Pow(SB), NOFRAME|NOSPLIT, $0
	JMP	·pow(SB)

TEXT ·IsOddInt(SB), NOFRAME|NOSPLIT, $0
	JMP	·isOddInt(SB)

TEXT ·Sqrt(SB), NOFRAME|NOSPLIT, $0
	JMP	·sqrt(SB)

TEXT ·Floor(SB), NOFRAME|NOSPLIT, $0
	JMP	·floor(SB)

TEXT ·Ceil(SB), NOFRAME|NOSPLIT, $0
	JMP	·ceil(SB)

TEXT ·Trunc(SB), NOFRAME|NOSPLIT, $0
	JMP	·trunc(SB)
