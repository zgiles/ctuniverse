// Copyright 2016 Zachary Giles
// MIT License (Expat)
//
// Please see the LICENSE file

// Package ctuniverse contains all the ctuniverse types
package ctuniverse

// SpaceMessage is a wrapper for data coming from clients
type SpaceMessage struct {
	Messagetype string      `json:"messagetype"`
	O           interface{} `json:"o"`
}
