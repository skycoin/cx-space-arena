package game

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	log.Print(g.World.Players[0].Rotation)
	if g.World.Players[0].Velocity < 0 {
		g.World.Players[0].Velocity += 1
	}
	if g.World.Players[0].Velocity > 0 {
		g.World.Players[0].Velocity -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if g.World.Players[0].Velocity < 7 {
			g.World.Players[0].Velocity += 2
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if g.World.Players[0].Velocity > -5 {
			g.World.Players[0].Velocity -= 2
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if g.World.Players[0].Rotation > 0 {
			g.World.Players[0].Rotation -= 0.01
		} else {
			g.World.Players[0].Rotation = 0.99
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.World.Players[0].Rotation < 1.0 {
			g.World.Players[0].Rotation += 0.01
		} else {
			g.World.Players[0].Rotation = 0.01
		}
	}
	g.World.Players[0].Rotation = math.Round(g.World.Players[0].Rotation*100) / 100
	g.World.Players[0].PositionX += math.Cos(g.World.Players[0].Rotation*360*(math.Pi/180)) * float64(g.World.Players[0].Velocity)
	g.World.Players[0].PositionY += math.Sin(g.World.Players[0].Rotation*360*(math.Pi/180)) * float64(g.World.Players[0].Velocity)
	return nil
}
