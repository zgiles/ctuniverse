package ctuniverse

type UniverseObject struct {
	Uuid string `json:"uuid" redis:"uuid"`
	Owner string `json:"owner" redis:"owner"`
	Type string `json:"type" redis:"type"`
	Global_X int64 `json:"global_x" redis:"global_x"`
	Global_Y int64 `json:"global_y" redis:"global_y"`
	Fuel int64 `json:"fuel" redis:"fuel"`
	Weight int64 `json:"fuel" redis:"weight"`
	Decay int64 `json:"decay" redis:"decay"`
	Velocity_x int64 `json:"velocity_x" redis:"velocity_x"`
	Velocity_y int64 `json:"velocity_y" redis:"velocity_y"`
}
