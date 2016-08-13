package ctuniverse

type SpaceObject struct {
	Uuid       string            `json:"uuid" redis:"uuid"`
	Owner      string            `json:"owner" redis:"owner"`
	Type       string            `json:"type" redis:"type"`
	Global_X   int64             `json:"global_x" redis:"global_x"`
	Global_Y   int64             `json:"global_y" redis:"global_y"`
	Velocity_x int64             `json:"velocity_x" redis:"velocity_x"`
	Velocity_y int64             `json:"velocity_y" redis:"velocity_y"`
	Attributes map[string]string `json:"attributes" redis:"attributes"`
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
