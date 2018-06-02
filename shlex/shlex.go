/*
 * shlex.go
 *
 * Copyright 2018 Bill Zissimopoulos
 */
/*
 * This file is part of golib.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

// Package shlex is used for simple command line splitting.
//
// Both POSIX and Windows dialects are provided.
package shlex

import (
	"unicode/utf8"
)

const (
	Space       = rune(' ')
	Word        = rune('A')
	DoubleQuote = rune('"')
	SingleQuote = rune('\'')
	EmptyRune   = rune(-2)
	NoEscape    = rune(-1)
)

// Dialect represents a dialect of command line splitting.
type Dialect struct {
	IsSpace func(r rune) bool
	IsQuote func(r rune) bool
	Escape  func(s rune, r, r0 rune) rune
}

// Posix is the POSIX dialect of command line splitting.
// See https://tinyurl.com/26man79 for guidelines.
var Posix = Dialect{
	IsSpace: func(r rune) bool {
		return ' ' == r || '\t' == r || '\n' == r
	},
	IsQuote: func(r rune) bool {
		return '"' == r || '\'' == r
	},
	Escape: func(s rune, r, r0 rune) rune {
		if '\\' != r {
			return NoEscape
		}
		switch s {
		case Space, Word:
			if '\n' == r0 || EmptyRune == r0 {
				return EmptyRune
			}
			return r0
		case DoubleQuote:
			if '\n' == r0 || EmptyRune == r0 {
				return EmptyRune
			}
			if '$' == r0 || '`' == r0 || '"' == r0 || '\\' == r0 {
				return r0
			}
			return NoEscape
		default:
			return NoEscape
		}
	},
}

// Windows is the Windows dialect of command line splitting.
// See https://tinyurl.com/ycdj5ghh for guidelines.
var Windows = Dialect{
	IsSpace: func(r rune) bool {
		return ' ' == r || '\t' == r || '\r' == r || '\n' == r
	},
	IsQuote: func(r rune) bool {
		return '"' == r
	},
	Escape: func(s rune, r, r0 rune) rune {
		switch s {
		case Space, Word:
			if '\\' == r && '"' == r0 {
				return r0
			}
			return NoEscape
		case DoubleQuote:
			if ('\\' == r || '"' == r) && '"' == r0 {
				return r0
			}
			return NoEscape
		default:
			return NoEscape
		}
	},
}

// Split splits a command line into tokens according to the chosen dialect.
func (dialect *Dialect) Split(line string) (tokens []string) {
	tokens = make([]string, 0, 8)

	var token []rune
	var state = []rune{' '}

	for r, w := utf8.DecodeRuneInString(line); 0 != w; {
		line = line[w:]
		r0, w0 := utf8.DecodeRuneInString(line)

		s := state[len(state)-1]
		e := NoEscape
		if 0 != w0 {
			e = dialect.Escape(s, r, r0)
			if NoEscape != e {
				line = line[w0:]
				r0, w0 = utf8.DecodeRuneInString(line)
				if EmptyRune == e {
					r, w = r0, w0
					continue
				}
			}
		} else {
			e = dialect.Escape(s, r, EmptyRune)
			if EmptyRune == e {
				r, w = r0, w0
				continue
			}
		}

		switch s {
		case Space:
			switch {
			case NoEscape != e:
				state = append(state, Word)
				token = append(token, e)
			case dialect.IsQuote(r):
				state = append(state, Word, r)
				token = make([]rune, 0)
			case !dialect.IsSpace(r):
				state = append(state, Word)
				token = append(token, r)
			}
		case Word:
			switch {
			case NoEscape != e:
				token = append(token, e)
			case dialect.IsQuote(r):
				state = append(state, r)
			case dialect.IsSpace(r):
				state = state[:len(state)-1]
				tokens = append(tokens, string(token))
				token = nil
			default:
				token = append(token, r)
			}
		default: // quote
			switch {
			case NoEscape != e:
				token = append(token, e)
			case s == r:
				state = state[:len(state)-1]
			default:
				token = append(token, r)
			}
		}

		r, w = r0, w0
	}

	if nil != token {
		tokens = append(tokens, string(token))
	}

	return
}
