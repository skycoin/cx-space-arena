package asteroid

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Asteroid struct {
	PositionX, PositionY float64
	VelocityX, VelocityY int
	Rotation             int
	VelocityR            int
	Mass                 float64
	Health               int
}

// Asteroid constructor
func NewAsteroid() *Asteroid {
	w, h := ebiten.WindowSize()
	return &Asteroid{
		PositionX: float64(rand.Intn(w)),
		PositionY: float64(rand.Intn(h)),
		VelocityX: rand.Intn(10) - 5,
		VelocityY: rand.Intn(10) - 5,
		Rotation:  rand.Intn(360),
		VelocityR: rand.Intn(10) - 5,
		Mass:      rand.Float64() * 10,
		Health:    100,
	}
}
