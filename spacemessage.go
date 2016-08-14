// Copyright 2016 Zachary Giles
// MIT License (Expat)
//
// Please see the LICENSE file
package ctuniverse

type SpaceMessage struct {
	Messagetype string      `json:"messagetype"`
	O           interface{} `json:"o"`
}
