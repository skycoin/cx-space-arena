package player

type Player struct {
	PlayerID             int
	PositionX, PositionY float64
	Velocity             int
	Rotation             int
	Health               int
	Points               int
	Deaths               int
	Cooldown             int
}
