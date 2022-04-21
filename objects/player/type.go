package player

type Player struct {
	PlayerID             int
	PositionX, PositionY float64
	Velocity             int
	Rotation             float64
	Health               int
	Points               int
	Deaths               int
}
