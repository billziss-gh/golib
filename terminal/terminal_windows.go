/*
 * terminal_windows.go
 *
 * Copyright 2018 Bill Zissimopoulos
 */
/*
 * This file is part of golib.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

package terminal

import (
	"syscall"
	"unsafe"
)

const (
	_ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
)

var (
	dll            = syscall.NewLazyDLL("kernel32.dll")
	getConsoleMode = dll.NewProc("GetConsoleMode")
	setConsoleMode = dll.NewProc("SetConsoleMode")
)

func isTerminal(fd uintptr) bool {
	var mode uint32
	res, _, _ := getConsoleMode.Call(fd, uintptr(unsafe.Pointer(&mode)))
	return 0 != res
}

func isAnsiTerminal(fd uintptr) bool {
	// This is a bit hacky. On Win10 (later versions) ANSI support exists,
	// but is disabled by default. So we enable it, when asked for it!

	var mode uint32
	res, _, _ := getConsoleMode.Call(fd, uintptr(unsafe.Pointer(&mode)))
	if 0 != res {
		if 0 != mode&_ENABLE_VIRTUAL_TERMINAL_PROCESSING {
			return true
		}

		mode |= _ENABLE_VIRTUAL_TERMINAL_PROCESSING
		res, _, _ := setConsoleMode.Call(fd, uintptr(mode))
		if 0 != res {
			res, _, _ := getConsoleMode.Call(fd, uintptr(unsafe.Pointer(&mode)))
			if 0 != res {
				return 0 != mode&_ENABLE_VIRTUAL_TERMINAL_PROCESSING
			}
		}
	}

	return false
}
