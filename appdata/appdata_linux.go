/*
 * appdata_linux.go
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
	"os"
	"os/user"
	"path/filepath"

	"github.com/billziss-gh/golib/errors"
)

// Linux: use the XDG Base Directory Specification.
//
// See https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

type LinuxAppData struct {
	configDir string
	dataDir   string
	cacheDir  string
	err       error
}

func (self *LinuxAppData) ConfigDir() (string, error) {
	if nil != self.err {
		return "", errors.New("", self.err, ErrAppData)
	}

	return self.configDir, nil
}

func (self *LinuxAppData) DataDir() (string, error) {
	if nil != self.err {
		return "", errors.New("", self.err, ErrAppData)
	}

	return self.dataDir, nil
}

func (self *LinuxAppData) CacheDir() (string, error) {
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

		configDir = os.Getenv("XDG_CONFIG_HOME")
		dataDir = os.Getenv("XDG_DATA_HOME")
		cacheDir = os.Getenv("XDG_CACHE_HOME")

		if "" == configDir {
			configDir = filepath.Join(h, ".config")
		}

		if "" == dataDir {
			dataDir = filepath.Join(h, ".local/share")
		}

		if "" == cacheDir {
			cacheDir = filepath.Join(h, ".cache")
		}
	}

	DefaultAppData = &LinuxAppData{configDir, dataDir, cacheDir, e}
}
