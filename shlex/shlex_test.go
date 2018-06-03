/*
 * shlex_test.go
 *
 * Copyright 2018 Bill Zissimopoulos
 */
/*
 * This file is part of golib.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

package shlex

import (
	"reflect"
	"testing"
)

func TestPosix(t *testing.T) {
	tests := []struct {
		i string
		o []string
	}{
		{"", []string{}},
		{"foo", []string{"foo"}},
		{" foo", []string{"foo"}},
		{"   foo", []string{"foo"}},
		{"foo ", []string{"foo"}},
		{"foo   ", []string{"foo"}},
		{"foo bar", []string{"foo", "bar"}},
		{"foo   bar", []string{"foo", "bar"}},
		{`foo"   "bar`, []string{"foo   bar"}},
		{`foo"   `, []string{"foo   "}},
		{`"   "foo`, []string{"   foo"}},
		{`foo "bar   baz"`, []string{"foo", "bar   baz"}},
		{`foo "bar   baz`, []string{"foo", "bar   baz"}},
		{`foo "bar   baz `, []string{"foo", "bar   baz "}},
		{`foo "bar   baz" `, []string{"foo", "bar   baz"}},
		{`foo "bar   baz" bag`, []string{"foo", "bar   baz", "bag"}},
		{`foo "bar   baz"bag`, []string{"foo", "bar   bazbag"}},
		{`foo ""`, []string{"foo", ""}},
		{`foo "" bar`, []string{"foo", "", "bar"}},
		{`foo'   'bar`, []string{"foo   bar"}},
		{`foo'   `, []string{"foo   "}},
		{`'   foo`, []string{"   foo"}},
		{`foo 'bar   baz'`, []string{"foo", "bar   baz"}},
		{`foo 'bar   baz`, []string{"foo", "bar   baz"}},
		{`foo 'bar   baz `, []string{"foo", "bar   baz "}},
		{`foo 'bar   baz' `, []string{"foo", "bar   baz"}},
		{`foo 'bar   baz' bag`, []string{"foo", "bar   baz", "bag"}},
		{`foo 'bar   baz'bag`, []string{"foo", "bar   bazbag"}},
		{`foo ''`, []string{"foo", ""}},
		{`foo '' bar`, []string{"foo", "", "bar"}},
		{`\foo`, []string{`foo`}},
		{`f\oo`, []string{`foo`}},
		{`foo\`, []string{`foo`}},
		{`\"foo`, []string{`"foo`}},
		{`f\"oo`, []string{`f"oo`}},
		{`foo\"`, []string{`foo"`}},
		{`\'foo`, []string{`'foo`}},
		{`f\'oo`, []string{`f'oo`}},
		{`foo\'`, []string{`foo'`}},
		{"foo\\\nbar", []string{"foobar"}},
		{`foo "b\ar"`, []string{`foo`, `b\ar`}},
		{`foo "b\"ar"`, []string{`foo`, `b"ar`}},
		{`foo "b'ar"`, []string{`foo`, `b'ar`}},
		{`foo "b'a'r"`, []string{`foo`, `b'a'r`}},
		{`foo"\`, []string{`foo`}},
		{"foo\"\\\nbar\"", []string{"foobar"}},
		{"foo\"\\\nbar", []string{"foobar"}},
		{`foo 'b\ar'`, []string{`foo`, `b\ar`}},
		{`foo 'b\"ar'`, []string{`foo`, `b\"ar`}},
		{`foo 'b"ar'`, []string{`foo`, `b"ar`}},
		{`foo 'b"a"r'`, []string{`foo`, `b"a"r`}},
		{`foo'\`, []string{`foo\`}},
		{"foo'\\\nbar'", []string{"foo\\\nbar"}},
		{"foo'\\\nbar", []string{"foo\\\nbar"}},
	}

	for _, test := range tests {
		o := Posix.Split(test.i)
		if !reflect.DeepEqual(test.o, o) {
			t.Errorf("expect %#v, got %#v", test.o, o)
		}
	}
}

func TestWindows(t *testing.T) {
	tests := []struct {
		i string
		o []string
	}{
		{"", []string{}},
		{"foo", []string{"foo"}},
		{" foo", []string{"foo"}},
		{"   foo", []string{"foo"}},
		{"foo ", []string{"foo"}},
		{"foo   ", []string{"foo"}},
		{"foo bar", []string{"foo", "bar"}},
		{"foo   bar", []string{"foo", "bar"}},
		{`foo"   "bar`, []string{"foo   bar"}},
		{`foo"   `, []string{"foo   "}},
		{`"   "foo`, []string{"   foo"}},
		{`foo "bar   baz"`, []string{"foo", "bar   baz"}},
		{`foo "bar   baz`, []string{"foo", "bar   baz"}},
		{`foo "bar   baz `, []string{"foo", "bar   baz "}},
		{`foo "bar   baz" `, []string{"foo", "bar   baz"}},
		{`foo "bar   baz" bag`, []string{"foo", "bar   baz", "bag"}},
		{`foo "bar   baz"bag`, []string{"foo", "bar   bazbag"}},
		{`foo ""`, []string{"foo", ""}},
		{`foo "" bar`, []string{"foo", "", "bar"}},
		{`\foo`, []string{`\foo`}},
		{`f\oo`, []string{`f\oo`}},
		{`foo\`, []string{`foo\`}},
		{`\"foo`, []string{`"foo`}},
		{`f\"oo`, []string{`f"oo`}},
		{`foo\"`, []string{`foo"`}},
		{`foo "b\ar"`, []string{`foo`, `b\ar`}},
		{`foo "b\"ar"`, []string{`foo`, `b"ar`}},
		{`foo "b""ar"`, []string{`foo`, `b"ar`}},
		{`foo\\`, []string{`foo\\`}},
		{`foo\\\`, []string{`foo\\\`}},
		{`foo\\\\`, []string{`foo\\\\`}},
		{`foo\\\\\`, []string{`foo\\\\\`}},
		{`foo\\"`, []string{`foo\`}},
		{`foo\\\"`, []string{`foo\"`}},
		{`foo\\\\"`, []string{`foo\\`}},
		{`foo\\\\\"`, []string{`foo\\"`}},
		{`"foo\\"   bar`, []string{`foo\`, `bar`}},
		{`"foo\\\"   bar"`, []string{`foo\"   bar`}},
		{`"foo\\\\"   bar"`, []string{`foo\\`, `bar`}},
		{`"foo\\\\\"   bar"`, []string{`foo\\"   bar`}},
		{`"foo\\   bar"`, []string{`foo\\   bar`}},
		{`"foo\\\   bar"`, []string{`foo\\\   bar`}},
		{`"foo\\\\   bar"`, []string{`foo\\\\   bar`}},
		{`"foo\\\\\   bar"`, []string{`foo\\\\\   bar`}},

		{`CallMeIshmael`, []string{`CallMeIshmael`}},
		{`"Call Me Ishmael"`, []string{`Call Me Ishmael`}},
		{`Cal"l Me I"shmael`, []string{`Call Me Ishmael`}},
		{`CallMe\"Ishmael`, []string{`CallMe"Ishmael`}},
		{`"CallMe\"Ishmael"`, []string{`CallMe"Ishmael`}},
		{`"Call Me Ishmael\\"`, []string{`Call Me Ishmael\`}},
		{`"CallMe\\\"Ishmael"`, []string{`CallMe\"Ishmael`}},
		{`a\\\b`, []string{`a\\\b`}},
		{`"a\\\b"`, []string{`a\\\b`}},
		{`"\"Call Me Ishmael\""`, []string{`"Call Me Ishmael"`}},
		{`"C:\TEST A\\"`, []string{`C:\TEST A\`}},
		{`"\"C:\TEST A\\\""`, []string{`"C:\TEST A\"`}},
		{`"a b c"  d  e`, []string{`a b c`, `d`, `e`}},
		{`"ab\"c"  "\\"  d`, []string{`ab"c`, `\`, `d`}},
		{`a\\\b d"e f"g h`, []string{`a\\\b`, `de fg`, `h`}},
		{`a\\\"b c d`, []string{`a\"b`, `c`, `d`}},
		{`a\\\\"b c" d e`, []string{`a\\b c`, `d`, `e`}},
		{`"a b c""`, []string{`a b c"`}},
		{`"""CallMeIshmael"""  b  c`, []string{`"CallMeIshmael"`, `b`, `c`}},
		{`"""Call Me Ishmael"""`, []string{`"Call Me Ishmael"`}},
		{`""""Call Me Ishmael"" b c`, []string{`"Call`, `Me`, `Ishmael`, `b`, `c`}},
	}

	for _, test := range tests {
		o := Windows.Split(test.i)
		if !reflect.DeepEqual(test.o, o) {
			t.Errorf("expect %#v, got %#v", test.o, o)
		}
	}
}
