/*
 * appdata_windows.go
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

// Windows: use environment variables and well-known locations.
//
// A perhaps better solution would be to use SHGetKnownFolderPath.

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

		configDir = os.Getenv("APPDATA")
		dataDir = os.Getenv("APPDATA")
		cacheDir = os.TempDir()

		if "" == configDir {
			configDir = filepath.Join(h, "AppData\\Roaming")
		}

		if "" == dataDir {
			dataDir = filepath.Join(h, "AppData\\Roaming")
		}

		if "" == cacheDir {
			cacheDir = filepath.Join(h, "AppData\\Local\\Temp")
		}
	}

	DefaultAppData = &SystemAppData{configDir, dataDir, cacheDir, e}
}
