// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build sw64

package syscall

var RawSyscallNoError = rawSyscallNoError

const Sys_GETEUID = sys_GETXUID
