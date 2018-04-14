/*
 * replace_test.go
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
	"bytes"
	"math/rand"
	"strings"
	"testing"
)

func testCodes(s string) string {
	return strings.ToUpper(s)
}

var testReplaceDelims = []string{
	`{{{{ }}}}`,
	`{{{{ }}}}`,
	`{{{{ }}}}`,
	`{{{{ }}}}`,
	`{{{{ }}}}`,
	`{{ }}`,
	`{{ }}`,
	`{{ }}`,
}
var testReplaceInputs = []string{
	`hello world`,
	`hello world {{{{red`,
	`hello {{{{red}}}} world, hello {{{{green}}}} world, {{{{}}}}he{{{l}}}lo {{{{blue`,
	`hello {{{{red}}}} world, hello {{{{green}}}} world, {{{{}}}}he{{{l}}}lo {{{{blue}}}} world`,
	`hello {{{{red}}}} world, hello {{{{green}}}} world, {{{{}}}}he{{{l}}}lo {{{{blue}}} }}}} world`,
	`hello {{red}} world, hello {{green}} world, {{}}he{l}lo {{blue`,
	`hello {{red}} world, hello {{green}} world, {{}}he{l}lo {{blue}} world`,
	`hello {{red}} world, hello {{green}} world, {{}}he{l}lo {{blue} }} world`,
}
var testReplaceOutputs = []string{
	`hello world`,
	`hello world `,
	`hello RED world, hello GREEN world, he{{{l}}}lo `,
	`hello RED world, hello GREEN world, he{{{l}}}lo BLUE world`,
	`hello RED world, hello GREEN world, he{{{l}}}lo BLUE}}}  world`,
	`hello RED world, hello GREEN world, he{l}lo `,
	`hello RED world, hello GREEN world, he{l}lo BLUE world`,
	`hello RED world, hello GREEN world, he{l}lo BLUE}  world`,
}

func TestReplace(t *testing.T) {
	for i := range testReplaceDelims {
		r := Replace(testReplaceInputs[i], testReplaceDelims[i], testCodes)
		if r != testReplaceOutputs[i] {
			t.Errorf("%d %#v", i, r)
		}
	}
}

func TestReplaceWriter(t *testing.T) {
	for j := 0; 1000 > j; j++ {
		for i := range testReplaceDelims {
			var buf bytes.Buffer
			w := NewReplaceWriter(&buf, testReplaceDelims[i], testCodes)

			for l, h := 0, 0; len(testReplaceInputs[i]) > l; l = h {
				h = l + rand.Intn(len(testReplaceInputs[i]))
				if len(testReplaceInputs[i]) < h {
					h = len(testReplaceInputs[i])
				}

				b := []byte(testReplaceInputs[i][l:h])
				n, err := w.Write(b)
				if n != len(b) || nil != err {
					t.Error(j, i)
				}
			}

			r := buf.String()
			if r != testReplaceOutputs[i] {
				t.Errorf("%d %d %#v", j, i, r)
			}
		}
	}
}
