package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update function, called 60 times each second
func (g *Game) Update() error {
	pl1 := g.World.Players[0]

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if pl1.Cooldown == 0 {
			g.World.AddNewBullet(pl1)
			pl1.Cooldown = 11
		}
	}

	if pl1.Cooldown > 0 {
		pl1.Cooldown--
	}

	for i, v := range g.World.Bullets {
		if v != nil {
			if v.Lifespan == 0 {
				g.World.RemoveBullet(i)
			}
			v.Lifespan--
		}
	}

	// Constant force against the ship
	if pl1.Velocity < 0 {
		pl1.Velocity += 5
	}
	if pl1.Velocity > 0 {
		pl1.Velocity -= 5
	}

	// Velocity increases and decreases the longer the button is held down
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if pl1.Velocity < 70 {
			pl1.Velocity += 15
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if pl1.Velocity > -50 {
			pl1.Velocity -= 15
		}
	}

	// Rotation changes at a constant rate
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if pl1.Rotation > 0 {
			pl1.Rotation -= 5
		} else {
			pl1.Rotation = 355
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if pl1.Rotation < 360 {
			pl1.Rotation += 5
		} else {
			pl1.Rotation = 5
		}
	}

	// Use sine and cosine functions to get X and Y positions based on rotation and velocity
	pl1.PositionX += math.Cos(float64(pl1.Rotation)*(math.Pi/180)) * float64(pl1.Velocity) / 10
	pl1.PositionY += math.Sin(float64(pl1.Rotation)*(math.Pi/180)) * float64(pl1.Velocity) / 10

	w, h := ebiten.WindowSize()
	if pl1.PositionX < 0 {
		pl1.PositionX = 0
	}
	if pl1.PositionX > float64(w) {
		pl1.PositionX = float64(w)
	}
	if pl1.PositionY < 0 {
		pl1.PositionY = 0
	}
	if pl1.PositionY > float64(h) {
		pl1.PositionY = float64(h)
	}

	return nil
}
