// Copyright 2016 Zachary Giles
// MIT License (Expat)
//
// Please see the LICENSE file

package ctuniverse

// SpaceObject is an object in the universe.
type SpaceObject struct {
	UUID          string     `json:"uuid" redis:"uuid"`
	Owner         string     `json:"owner" redis:"owner"`
	Type          string     `json:"type" redis:"type"`
	Global        []float64  `json:"global"`
	Velocity      []float64  `json:"velocity"`
	Angle         float64    `json:"angle"`
	AngleVelocity float64    `json:"angle_velocity"`
	Boost         int64      `json:"boost"`
	Thrusters     []Thruster `json:"thrusters"`
}

// Thruster describes thrusters on a SpaceObject, if it has any.
type Thruster struct {
	Type   string `json:"type"`
	Firing int64  `json:"firing"`
}

// SpaceControl is given when an event occurs in the universe such as a ship blowing up.
type SpaceControl struct {
	UUID   string `json:"uuid" redis:"uuid"`
	Action string `json:"action" redis:"action"`
}

// SpaceID will hopefully be used to identify new clients as they come in.
type SpaceID struct {
	UUID string `json:"uuid" redis:"uuid"`
}

// These will be in attributes for now
// Decay int64 `json:"decay" redis:"decay"`
// Fuel int64 `json:"fuel" redis:"fuel"`
// Weight int64 `json:"weight" redis:"weight"`
// Attributes map[string]string `json:"attributes" redis:"attributes"`
