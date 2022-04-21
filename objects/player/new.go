package player

func NewPlayer(ID int) *Player {
	return &Player{
		playerID:  ID,
		positionX: 0, positionY: 0,
		velocityX: 0, velocityY: 0,
		rotationAngle: 0,
		health:        100,
		points:        0,
		deaths:        0,
	}
}
