package asteroid

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	w, h := ebiten.WindowSize()
	asteroid, _, err := ebitenutil.NewImageFromFile("assets/asteroid.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Asteroid{
		PositionX: float64(rand.Intn(w)),
		PositionY: float64(rand.Intn(h)),
		VelocityX: rand.Intn(10) - 5,
		VelocityY: rand.Intn(10) - 5,
		Rotation:  rand.Intn(360),
		VelocityR: rand.Intn(10) - 5,
		Mass:      float64(rand.Intn(100)+1) * 0.1,
		Health:    100,
		Sprite:    asteroid,
	}
}
