package player

func NewPlayer(ID int) *Player {
	return &Player{
		PlayerID:  ID,
		PositionX: 0, PositionY: 0,
		Velocity: 0,
		Rotation: 0,
		Health:   100,
		Points:   0,
		Deaths:   0,
	}
}
