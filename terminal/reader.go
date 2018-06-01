/*
 * reader.go
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
)

// NewReader reads terminal input, including special keys.
func NewReader(r io.Reader) io.Reader {
	return newReader(r)
}
