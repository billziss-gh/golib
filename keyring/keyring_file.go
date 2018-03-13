/*
 * keyring_file.go
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
	"fmt"
	"os"
	"sync"

	"github.com/billziss-gh/golib/config"
	"github.com/billziss-gh/golib/errors"
	"github.com/billziss-gh/golib/util"
)

type FileKeyring struct {
	Path string
	mux  sync.Mutex
}

func (self *FileKeyring) getConf() (config.Config, error) {
	conf, err := util.ReadFunc(self.Path, func(file *os.File) (interface{}, error) {
		return config.Read(file)
	})
	if nil != err {
		if e, ok := err.(*os.PathError); ok && "open" == e.Op {
			return config.Config{}, nil
		}
		return nil, err
	}
	return conf.(config.Config), nil
}

func (self *FileKeyring) setConf(conf config.Config) error {
	return util.WriteFunc(self.Path, 0600, func(file *os.File) error {
		return config.Write(file, conf)
	})
}

func (self *FileKeyring) Get(service, user string) (string, error) {
	self.mux.Lock()
	defer self.mux.Unlock()

	conf, err := self.getConf()
	if nil != err {
		return "", errors.New(fmt.Sprintf("cannot get key %s/%s", service, user), err, ErrKeyring)
	}

	if sect, ok := conf[service]; ok {
		if pass, ok := sect[user]; ok {
			return pass, nil
		}
	}

	return "", errors.New(fmt.Sprintf("cannot get key %s/%s", service, user), nil, ErrKeyring)
}

func (self *FileKeyring) Set(service, user, pass string) error {
	self.mux.Lock()
	defer self.mux.Unlock()

	conf, err := self.getConf()
	if nil != err {
		return errors.New(fmt.Sprintf("cannot set key %s/%s", service, user), err, ErrKeyring)
	}

	if sect, ok := conf[service]; ok {
		sect[user] = pass
	} else {
		sect = config.Section{}
		conf[service] = sect
		sect[user] = pass
	}

	err = self.setConf(conf)
	if nil != err {
		return errors.New(fmt.Sprintf("cannot set key %s/%s", service, user), err, ErrKeyring)
	}

	return nil
}

func (self *FileKeyring) Delete(service, user string) error {
	self.mux.Lock()
	defer self.mux.Unlock()

	conf, err := self.getConf()
	if nil != err {
		return errors.New(fmt.Sprintf("cannot delete key %s/%s", service, user), err, ErrKeyring)
	}

	if sect, ok := conf[service]; ok {
		delete(sect, user)
		if 0 == len(sect) {
			delete(conf, service)
		}
	}

	err = self.setConf(conf)
	if nil != err {
		return errors.New(fmt.Sprintf("cannot delete key %s/%s", service, user), err, ErrKeyring)
	}

	return nil
}
