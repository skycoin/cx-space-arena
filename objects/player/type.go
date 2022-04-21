package player

type Player struct {
	playerID             int
	positionX, positionY float64
	velocityX, velocityY float64
	rotationAngle        float64
	health               int
	points               int
	deaths               int
}

func (p *Player) GetID() int {
	return p.playerID
}

func (p *Player) IsAlive() bool {
	return (p.health > 0)
}
