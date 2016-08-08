package main

import (
	"errors"
	"github.com/zgiles/ctuniverse"
)

type Wsmsg struct {
	Event string `json:"event"`
	Epoch int64 `json:"epoch"`
}

type Wsmsg_idreq struct {
	Wsmsg
}

type Wsmsg_id struct {
	Wsmsg
	Originepoch int64 `json:"originepoch"`
	Uuid int64 `json:"uuid"`
	Attributes map[string]string `json:"attributes"`
}

type Wsmsg_ping struct {
	Wsmsg
}

type Wsmsg_pong struct {
	Wsmsg
	Originepoch int64 `json:"originepoch"`
}

type Wsmsg_uos struct {
	Wsmsg
	O []ctuniverse.UniverseObject `json:"o"`
}

type Wsmsg_uosack struct {
	Wsmsg
	Originepoch int64 `json:"originepoch"`
}

type Wsmsg_eorbyspace struct {
	Wsmsg
	Global_x int64 `json:"global_x"`
	Global_y int64 `json:"global_y"`
	Distance int64 `json:"distance"`
}

type Wsmsg_uou struct {
	Wsmsg
	O []ctuniverse.UniverseObject `json:"o"`
}

func (msg Wsmsg) CoerceToType() (interface{}, error) {
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
