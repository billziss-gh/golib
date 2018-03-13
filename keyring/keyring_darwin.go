/*
 * keyring_darwin.go
 *
 * Copyright 2018 Bill Zissimopoulos
 */
/*
 * This file is part of golib.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

package keyring

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/billziss-gh/golib/errors"
)

type DarwinKeyring struct {
}

func (self *DarwinKeyring) Get(service, user string) (string, error) {
	var buf bytes.Buffer
	cmd := exec.Command("security",
		"find-generic-password", "-s", service, "-a", user, "-g")
	cmd.Stderr = &buf
	err := cmd.Run()
	if nil != err {
		return "", errors.New(fmt.Sprintf("cannot get key %s/%s", service, user), err, ErrKeyring)
	}
	out := buf.String()
	var pass string
	_, err = fmt.Sscanf(out, "password: %q", &pass)
	if nil == err {
		return pass, nil
	}
	_, err = fmt.Sscanf(out, "password: 0x%x", &pass)
	if nil == err {
		return pass, nil
	}
	return "", errors.New(fmt.Sprintf("cannot get key %s/%s", service, user), nil, ErrKeyring)
}

func (self *DarwinKeyring) Set(service, user, pass string) error {
	err := exec.Command("security",
		"add-generic-password", "-s", service, "-a", user, "-p", pass, "-U").Run()
	if nil != err {
		return errors.New(fmt.Sprintf("cannot set key %s/%s", service, user), err, ErrKeyring)
	}
	return nil
}

func (self *DarwinKeyring) Delete(service, user string) error {
	err := exec.Command("security",
		"delete-generic-password", "-s", service, "-a", user).Run()
	if nil != err {
		return errors.New(fmt.Sprintf("cannot delete key %s/%s", service, user), err, ErrKeyring)
	}
	return nil
}

func init() {
	DefaultKeyring = &DarwinKeyring{}
}
