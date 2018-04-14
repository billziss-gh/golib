/*
 * terminal_darwin.go
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
	"os"
	"syscall"
	"unsafe"
)

func isTerminal(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		fd, syscall.TIOCGETA, uintptr(unsafe.Pointer(&termios)))
	return 0 == err
}

func isAnsiTerminal(fd uintptr) bool {
	return isTerminal(fd) && "dumb" != os.Getenv("TERM")
}
