// cmd/9l/list.c from Vita Nuova.
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2008 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2008 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sw64

import (
	"cmd/internal/obj"
	"fmt"
)

func init() {
	obj.RegisterRegister(obj.RBaseSW64, REG_R0+1024, rconv)
	//	obj.RegisterRegister(obj.RBaseSW64, REG_LAST+1, rconv)
	obj.RegisterOpcode(obj.ABaseSW64, Anames)
}

func rconv(r int) string {
	if r == 0 {
		return "NONE"
	}
	//if r == REGG {
	if r == REG_R15 {
		// Special case.
		return "g"
	}
	//if r == REGSB {
	if r == REG_R29 {
		// Special case.
		return "RSB"
	}
	//if r == REGSP {
	if r == REG_R30 {
		// Special case.
		return "SP"
	}
	if r == REGCTXT {
		return "REGCTXT"
	}
	if r == REGZERO {
		return "ZERO"
	}
	if IsRReg(r) {
		return fmt.Sprintf("R%d", r-REG_R0)
	}
	if IsFReg(r) {
		return fmt.Sprintf("F%d", r-REG_F0)
	}
	return fmt.Sprintf("badreg(%d)", r)

}

func DRconv(a int) string {
	s := "C_??"
	if a >= C_NONE && a <= C_NCLASS {
		s = cnames77[a]
	}
	var fp string
	fp += s
	return fp
}
