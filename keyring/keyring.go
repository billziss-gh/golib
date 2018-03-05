/*
 * keyring.go
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

const ErrKeyring = "ErrKeyring"

type Keyring interface {
	Get(service, user string) (string, error)
	Set(service, user, pass string) error
	Delete(service, user string) error
}

var DefaultKeyring Keyring

func Get(service, user string) (string, error) {
	return DefaultKeyring.Get(service, user)
}

func Set(service, user, pass string) error {
	return DefaultKeyring.Set(service, user, pass)
}

func Delete(service, user string) error {
	return DefaultKeyring.Delete(service, user)
}
