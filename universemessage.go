// Code by Zachary Giles
// This code is under the MIT License, a copy of which is found in the LICENSE file.

package ctuniverse

type UniverseMessage struct {
	Messagetype string      `json:"messagetype"`
	O           interface{} `json:o`
}
