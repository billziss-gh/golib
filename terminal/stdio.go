/*
 * stdio.go
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
	"io"
	"os"
)

var Stdout io.Writer
var Stderr io.Writer

func init() {
	escape := NullEscapeCode
	if IsAnsiTerminal(os.Stdout.Fd()) {
		escape = AnsiEscapeCode
	}

	Stdout = NewEscapeWriter(os.Stdout, "{{ }}", escape)
	Stderr = NewEscapeWriter(os.Stderr, "{{ }}", escape)
}
