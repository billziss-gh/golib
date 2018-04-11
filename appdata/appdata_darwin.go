/*
 * appdata_darwin.go
 *
 * Copyright 2018 Bill Zissimopoulos
 */
/*
 * This file is part of golib.
 *
 * It is licensed under the MIT license. The full license text can be found
 * in the License.txt file at the root of this project.
 */

package appdata

import (
	"os/user"
	"path/filepath"

	"github.com/billziss-gh/golib/errors"
)

// Darwin: use well-known locations.
//
// We should probably be using API's like NSSearchPathForDirectoriesInDomains or FSFindFolder,
// but cannot without cgo.

type SystemAppData struct {
	configDir string
	dataDir   string
	cacheDir  string
	err       error
}

func (self *SystemAppData) ConfigDir() (string, error) {
	if nil != self.err {
		return "", errors.New("", self.err, ErrAppData)
	}

	return self.configDir, nil
}

func (self *SystemAppData) DataDir() (string, error) {
	if nil != self.err {
		return "", errors.New("", self.err, ErrAppData)
	}

	return self.dataDir, nil
}

func (self *SystemAppData) CacheDir() (string, error) {
	if nil != self.err {
		return "", errors.New("", self.err, ErrAppData)
	}

	return self.cacheDir, nil
}

func init() {
	configDir := ""
	dataDir := ""
	cacheDir := ""

	h := ""
	u, e := user.Current()
	if nil == e {
		h = u.HomeDir

		configDir = filepath.Join(h, "Library/Preferences")
		dataDir = filepath.Join(h, "Library/Application Support")
		cacheDir = filepath.Join(h, "Library/Caches")
	}

	DefaultAppData = &SystemAppData{configDir, dataDir, cacheDir, e}
}
