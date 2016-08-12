// Code by Zachary Giles
// This code is under the MIT License, a copy of which is found in the LICENSE file.

package main

import (
	"github.com/BurntSushi/toml"
)

type serverconfig struct {
	Maindb       string
	Ip           string
	Port         int64
	Closetimeout int64
}

type config struct {
	Serverconfig serverconfig
}

func loadConfig(file string) (*config, error) {
	newconfig := new(config)
	if _, err := toml.DecodeFile(file, &newconfig); err != nil {
		return newconfig, err
	}
	return newconfig, nil // config, no error
}
