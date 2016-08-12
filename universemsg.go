// Code by Zachary Giles
// This code is under the MIT License, a copy of which is found in the LICENSE file.

package ctuniverse

/*
import (
	"errors"
)
*/

type UniverseMsg struct {
	Event string `json:"event"`
	Epoch int64  `json:"epoch"`
}

type UniverseMsg_idreq struct {
	UniverseMsg
}

type UniverseMsg_id struct {
	UniverseMsg
	Originepoch int64             `json:"originepoch"`
	Uuid        int64             `json:"uuid"`
	Attributes  map[string]string `json:"attributes"`
}

type UniverseMsg_ping struct {
	UniverseMsg
}

type UniverseMsg_pong struct {
	UniverseMsg
	Originepoch int64 `json:"originepoch"`
}

type UniverseMsg_uos struct {
	UniverseMsg
	O []UniverseObject `json:"o"`
}

type UniverseMsg_uosack struct {
	UniverseMsg
	Originepoch int64 `json:"originepoch"`
}

type UniverseMsg_eorbyspace struct {
	UniverseMsg
	Global_x int64 `json:"global_x"`
	Global_y int64 `json:"global_y"`
	Distance int64 `json:"distance"`
}

type UniverseMsg_uou struct {
	UniverseMsg
	O []UniverseObject `json:"o"`
}

// Not correct syntax for coercing types.. not needed anymore.. yet
/*
func (msg Wsmsg) CoerceToType(map[string]interface{}) (interface{}, error) {
	switch msg.Event {
		case "idreq":
			return Wsmsg_idreq{msg}, nil
		case "id":
			return Wsmsg_id{msg}, nil
		case "ping":
			return Wsmsg_ping{msg}, nil
		case "pong":
			return Wsmsg_pong{msg}, nil
		case "uos":
			return Wsmsg_uos{msg}, nil
		case "uosack":
			return Wsmsg_uosack{msg}, nil
		case "uorbyspace":
			return Wsmsg_uorbyspace{msg}, nil
		case "uou":
			return Wsmsg_uou{msg}, nil
		default:
			return msg, errors.New("Message Type not implemented")
	}
}
*/
