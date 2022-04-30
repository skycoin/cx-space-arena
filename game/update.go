package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update function, called 60 times each second
func (g *Game) Update() error {
	pl1 := g.World.Players[0]

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if pl1.Cooldown == 0 {
			g.World.AddBullet(pl1)
			pl1.Cooldown = 11
		}
	}

	if pl1.Cooldown > 0 {
		pl1.Cooldown--
	}

	for _, b := range g.World.Bullets {
		if b != nil {
			b.PositionX += math.Cos(float64(b.Angle)*(math.Pi/180)) * 10
			b.PositionY += math.Sin(float64(b.Angle)*(math.Pi/180)) * 10
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

	// Update asteroid
	for _, v := range g.World.Asteroids {
		if v != nil {
			v.PositionX += float64(v.VelocityX) * 3 / v.Mass // Velocity scales opposite of mass
			v.PositionY += float64(v.VelocityY) * 3 / v.Mass
			v.Rotation += v.VelocityR * 3 / int(v.Mass)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.World.AddAst()
	}

	return nil
}
