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

// Package appdata provides access to well known directories for applications.
package appdata

const ErrAppData = "ErrAppData"

type AppData interface {
	ConfigDir() (string, error)
	DataDir() (string, error)
	CacheDir() (string, error)
}

var DefaultAppData AppData

// ConfigDir returns the directory where application configuration files
// should be stored.
func ConfigDir() (string, error) {
	return DefaultAppData.ConfigDir()
}

// DataDir returns the directory where application data files
// should be stored.
func DataDir() (string, error) {
	return DefaultAppData.DataDir()
}

// CacheDir returns the directory where application cache files
// should be stored.
func CacheDir() (string, error) {
	return DefaultAppData.CacheDir()
}
