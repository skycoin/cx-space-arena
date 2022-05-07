package asteroid

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/assets"
)

type Asteroid struct {
	PositionX, PositionY float64
	VelocityX, VelocityY int
	Rotation             int
	VelocityR            int
	Mass                 float64
	Health               int
	Sprite               *ebiten.Image
}

// Asteroid constructor
func NewAsteroid() *Asteroid {
	direction := [2]int{-1, 1}
	w, h := ebiten.WindowSize()
	return &Asteroid{
		PositionX: float64(rand.Intn(w)),
		PositionY: float64(rand.Intn(h)),
		VelocityX: direction[rand.Intn(2)], // Velocity is determined by mass later
		VelocityY: direction[rand.Intn(2)],
		Rotation:  rand.Intn(361),
		VelocityR: direction[rand.Intn(2)],         // Same goes for rotation speed
		Mass:      float64(rand.Intn(21)+30) * 0.1, // Asteroid mass is between 3 and 5
		Health:    90,
		Sprite:    assets.SpriteList.Asteroid[0],
	}
}

func IndexOf(asteroids []*Asteroid, a *Asteroid) int {
	for i, v := range asteroids {
		if v == a {
			return i
		}
	}
	return -1
}
