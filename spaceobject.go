package ctuniverse

type SpaceObject struct {
	Uuid          string     `json:"uuid" redis:"uuid"`
	Owner         string     `json:"owner" redis:"owner"`
	Type          string     `json:"type" redis:"type"`
	Global        []float64  `json:"global"`
	Velocity      []float64  `json:"velocity"`
	Angle         float64    `json:"angle"`
	AngleVelocity float64    `json:"angle_velocity"`
	Boost         int64      `json:"boost"`
	Thrusters     []Thruster `json:"thrusters"`
}

type Thruster struct {
	Type   string `json:"type"`
	Firing int64  `json:"firing"`
}

type SpaceControl struct {
	Uuid   string `json:"uuid" redis:"uuid"`
	Action string `json:"action" redis:"action"`
}

type SpaceID struct {
	Uuid string `json:"uuid" redis:"uuid"`
}

// These will be in attributes for now
// Decay int64 `json:"decay" redis:"decay"`
// Fuel int64 `json:"fuel" redis:"fuel"`
// Weight int64 `json:"weight" redis:"weight"`
// Attributes map[string]string `json:"attributes" redis:"attributes"`
