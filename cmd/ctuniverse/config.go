// Copyright 2016 Zachary Giles
// MIT License (Expat)
//
// Please see the LICENSE file

package main

import (
	"github.com/BurntSushi/toml"
)

type serverconfig struct {
	Maindb       string
	IP           string
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
