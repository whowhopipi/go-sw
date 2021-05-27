// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux
// +build sw64

package syscall

const (
	sys_GETEUID = SYS_GETXUID
	sys_GETXUID = SYS_GETXUID
	sys_SETGID  = SYS_SETGID
	sys_SETUID  = SYS_SETUID
)
