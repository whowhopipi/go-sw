// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

#include "textflag.h"

TEXT ·Asin(SB), NOFRAME|NOSPLIT, $0
	BR	·asin(SB)

TEXT ·Acos(SB), NOFRAME|NOSPLIT, $0
	BR	·acos(SB)

TEXT ·Asinh(SB), NOFRAME|NOSPLIT, $0
	BR	·asinh(SB)

TEXT ·Acosh(SB), NOFRAME|NOSPLIT, $0
	BR	·acosh(SB)

TEXT ·Atan2(SB), NOFRAME|NOSPLIT, $0
	BR	·atan2(SB)

TEXT ·Atan(SB), NOFRAME|NOSPLIT, $0
	BR	·atan(SB)

TEXT ·Atanh(SB), NOFRAME|NOSPLIT, $0
	BR	·atanh(SB)

TEXT ·Min(SB), NOFRAME|NOSPLIT, $0
	BR	·min(SB)

TEXT ·Max(SB), NOFRAME|NOSPLIT, $0
	BR	·max(SB)

TEXT ·Erf(SB), NOFRAME|NOSPLIT, $0
	BR	·erf(SB)

TEXT ·Erfc(SB), NOFRAME|NOSPLIT, $0
	BR	·erfc(SB)

TEXT ·Exp2(SB), NOFRAME|NOSPLIT, $0
	BR	·exp2(SB)

TEXT ·Expm1(SB), NOFRAME|NOSPLIT, $0
	BR	·expm1(SB)

TEXT ·Exp(SB), NOFRAME|NOSPLIT, $0
	BR	·exp(SB)

TEXT ·Frexp(SB), NOFRAME|NOSPLIT, $0
	BR	·frexp(SB)

TEXT ·Hypot(SB), NOFRAME|NOSPLIT, $0
	BR	·hypot(SB)

TEXT ·Ldexp(SB), NOFRAME|NOSPLIT, $0
	BR	·ldexp(SB)

TEXT ·Log10(SB), NOFRAME|NOSPLIT, $0
	BR	·log10(SB)

TEXT ·Log2(SB), NOFRAME|NOSPLIT, $0
	BR	·log2(SB)

TEXT ·Log1p(SB), NOFRAME|NOSPLIT, $0
	BR	·log1p(SB)

TEXT ·Log(SB), NOFRAME|NOSPLIT, $0
	BR	·log(SB)

TEXT ·Mod(SB), NOFRAME|NOSPLIT, $0
	BR	·mod(SB)

TEXT ·Modf(SB), NOFRAME|NOSPLIT, $0
	BR	·modf(SB)

TEXT ·Remainder(SB), NOFRAME|NOSPLIT, $0
	BR	·remainder(SB)

TEXT ·Sin(SB), NOFRAME|NOSPLIT, $0
	BR	·sin(SB)

TEXT ·Sinh(SB), NOFRAME|NOSPLIT, $0
	BR	·sinh(SB)

TEXT ·Cos(SB), NOFRAME|NOSPLIT, $0
	BR	·cos(SB)

TEXT ·Cosh(SB), NOFRAME|NOSPLIT, $0
	BR	·cosh(SB)

TEXT ·Tan(SB), NOFRAME|NOSPLIT, $0
	BR	·tan(SB)

TEXT ·Tanh(SB), NOFRAME|NOSPLIT, $0
	BR	·tanh(SB)

TEXT ·Cbrt(SB), NOFRAME|NOSPLIT, $0
	BR	·cbrt(SB)

TEXT ·Pow(SB), NOFRAME|NOSPLIT, $0
	BR	·pow(SB)

TEXT ·IsOddInt(SB), NOFRAME|NOSPLIT, $0
	BR	·isOddInt(SB)

TEXT ·Sqrt(SB), NOFRAME|NOSPLIT, $0
	BR	·sqrt(SB)

TEXT ·Floor(SB), NOFRAME|NOSPLIT, $0
	BR	·floor(SB)

TEXT ·Ceil(SB), NOFRAME|NOSPLIT, $0
	BR	·ceil(SB)

TEXT ·Trunc(SB), NOFRAME|NOSPLIT, $0
	BR	·trunc(SB)
