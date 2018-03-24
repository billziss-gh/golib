/*
 * appdata.go
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

const ErrAppData = "ErrAppData"

type AppData interface {
	ConfigDir() (string, error)
	DataDir() (string, error)
	CacheDir() (string, error)
}

var DefaultAppData AppData

func ConfigDir() (string, error) {
	return DefaultAppData.ConfigDir()
}

func DataDir() (string, error) {
	return DefaultAppData.DataDir()
}

func CacheDir() (string, error) {
	return DefaultAppData.CacheDir()
}
