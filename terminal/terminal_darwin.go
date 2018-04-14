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

func isTerminal(fd uintptr) bool {
	return false
}

func isAnsiTerminal(fd uintptr) bool {
	return false
}
