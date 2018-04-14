/*
 * codes.go
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

import "strings"

func NullEscapeCode(code string) string {
	return ""
}

func AnsiEscapeCode(code string) string {
	parts := strings.Split(code, " ")
	for i, p := range parts {
		parts[i] = ansiColors[p]
	}
	return "\033[" + strings.Join(parts, ";") + "m"
}

var ansiColors = map[string]string{
	"reset":     "0",
	"bold":      "1",
	"black":     "30",
	"red":       "31",
	"green":     "32",
	"yellow":    "33",
	"blue":      "34",
	"magenta":   "35",
	"cyan":      "36",
	"white":     "37",
	"bgblack":   "40",
	"bgred":     "41",
	"bggreen":   "42",
	"bgyellow":  "43",
	"bgblue":    "44",
	"bgmagenta": "45",
	"bgcyan":    "46",
	"bgwhite":   "47",
}
