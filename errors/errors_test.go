/*
 * errors_test.go
 *
 * Copyright 2018 Bill Zissimopoulos
 */
/*
 * This file is part of golib.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

package errors

import (
	goerrors "errors"
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	e0 := New("e0")
	e1 := New("e1", e0)
	e2 := New("e2", e1, 42)

	g0 := goerrors.New("g0")
	g1 := New("g1", g0)
	g2 := New("g2", g1, "42")

	if "e2; e1; e0" != e2.Error() {
		t.Error()
	}
	if "e2; e1; e0" != fmt.Sprintf("%s", e2) {
		t.Error()
	}
	if `"e2; e1; e0"` != fmt.Sprintf("%q", e2) {
		t.Error()
	}
	if "e2; e1; e0" != fmt.Sprintf("%v", e2) {
		t.Error()
	}

	E := `e2 (42)
    	errors.TestErrors:github.com/billziss-gh/golib/errors/errors_test.go:18
e1
    	errors.TestErrors:github.com/billziss-gh/golib/errors/errors_test.go:17
e0
    	errors.TestErrors:github.com/billziss-gh/golib/errors/errors_test.go:16`
	if E != fmt.Sprintf("%+v", e2) {
		t.Error()
	}

	if "g2; g1; g0" != g2.Error() {
		t.Error()
	}
	if "g2; g1; g0" != fmt.Sprintf("%s", g2) {
		t.Error()
	}
	if `"g2; g1; g0"` != fmt.Sprintf("%q", g2) {
		t.Error()
	}

	G := `g2 (42)
    	errors.TestErrors:github.com/billziss-gh/golib/errors/errors_test.go:22
g1
    	errors.TestErrors:github.com/billziss-gh/golib/errors/errors_test.go:21
g0`
	if G != fmt.Sprintf("%+v", g2) {
		t.Error()
	}

	if nil != Cause(e0) {
		t.Error()
	}
	if e0 != Cause(e1) {
		t.Error()
	}
	if e1 != Cause(e2) {
		t.Error()
	}

	if nil != Cause(g0) {
		t.Error()
	}
	if g0 != Cause(g1) {
		t.Error()
	}
	if g1 != Cause(g2) {
		t.Error()
	}

	if nil != Attachment(e0) {
		t.Error()
	}
	if nil != Attachment(e1) {
		t.Error()
	}
	if 42 != Attachment(e2) {
		t.Error()
	}

	if nil != Attachment(g0) {
		t.Error()
	}
	if nil != Attachment(g1) {
		t.Error()
	}
	if "42" != Attachment(g2) {
		t.Error()
	}

	if HasCause(nil, nil) {
		t.Error()
	}

	if !HasCause(e2, e2) {
		t.Error()
	}
	if !HasCause(e2, e1) {
		t.Error()
	}
	if !HasCause(e2, e0) {
		t.Error()
	}
	if HasCause(e2, nil) {
		t.Error()
	}

	if !HasCause(g2, g2) {
		t.Error()
	}
	if !HasCause(g2, g1) {
		t.Error()
	}
	if !HasCause(g2, g0) {
		t.Error()
	}
	if HasCause(g2, nil) {
		t.Error()
	}

	e3 := New("e3", e2)
	g3 := New("g3", g2)

	if HasAttachment(nil, nil) {
		t.Error()
	}

	if !HasAttachment(e3, 42) {
		t.Error()
	}
	if !HasAttachment(e2, 42) {
		t.Error()
	}
	if HasAttachment(e1, 42) {
		t.Error()
	}
	if HasAttachment(e0, 42) {
		t.Error()
	}

	if !HasAttachment(g3, "42") {
		t.Error()
	}
	if !HasAttachment(g2, "42") {
		t.Error()
	}
	if HasAttachment(g1, "42") {
		t.Error()
	}
	if HasAttachment(g0, "42") {
		t.Error()
	}
}

func TestMessage(t *testing.T) {
	e0 := New(": e0")
	e1 := New("", e0)
	e2 := New(": e2", e1)

	if "TestMessage: e2; TestMessage; TestMessage: e0" != e2.Error() {
		t.Error()
	}

	E := `TestMessage: e2
    	errors.TestMessage:github.com/billziss-gh/golib/errors/errors_test.go:173
TestMessage
    	errors.TestMessage:github.com/billziss-gh/golib/errors/errors_test.go:172
TestMessage: e0
    	errors.TestMessage:github.com/billziss-gh/golib/errors/errors_test.go:171`
	if E != fmt.Sprintf("%+v", e2) {
		t.Error()
	}

	func() {
		e0 := New(": e0")
		e1 := New("", e0)
		e2 := New(": e2", e1)

		if "TestMessage: e2; TestMessage; TestMessage: e0" != e2.Error() {
			t.Error()
		}

		E := `TestMessage: e2
    	errors.TestMessage.func1:github.com/billziss-gh/golib/errors/errors_test.go:192
TestMessage
    	errors.TestMessage.func1:github.com/billziss-gh/golib/errors/errors_test.go:191
TestMessage: e0
    	errors.TestMessage.func1:github.com/billziss-gh/golib/errors/errors_test.go:190`
		if E != fmt.Sprintf("%+v", e2) {
			t.Error()
		}
	}()

	myint(42).testMessage(t)
}

type myint int

func (i myint) testMessage(t *testing.T) {
	e0 := New(": e0")
	e1 := New("", e0)
	e2 := New(": e2", e1)

	if "testMessage: e2; testMessage; testMessage: e0" != e2.Error() {
		t.Error()
	}

	E := `testMessage: e2
    	errors.myint.testMessage:github.com/billziss-gh/golib/errors/errors_test.go:217
testMessage
    	errors.myint.testMessage:github.com/billziss-gh/golib/errors/errors_test.go:216
testMessage: e0
    	errors.myint.testMessage:github.com/billziss-gh/golib/errors/errors_test.go:215`
	if E != fmt.Sprintf("%+v", e2) {
		t.Error()
	}

	func() {
		e0 := New(": e0")
		e1 := New("", e0)
		e2 := New(": e2", e1)

		if "testMessage: e2; testMessage; testMessage: e0" != e2.Error() {
			t.Error()
		}

		E := `testMessage: e2
    	errors.myint.testMessage.func1:github.com/billziss-gh/golib/errors/errors_test.go:236
testMessage
    	errors.myint.testMessage.func1:github.com/billziss-gh/golib/errors/errors_test.go:235
testMessage: e0
    	errors.myint.testMessage.func1:github.com/billziss-gh/golib/errors/errors_test.go:234`
		if E != fmt.Sprintf("%+v", e2) {
			t.Error()
		}
	}()
}
