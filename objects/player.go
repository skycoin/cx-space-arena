package objects

type Player struct {
	positionX, positionY float64
	velocityX, velocityY float64
	rotationAngle        float64
	health               int
	points               int
	deaths               int
}
