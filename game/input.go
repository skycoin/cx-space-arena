package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) UpdateInputs() {
	pl1 := g.World.Players[0]
	if pl1.Recoil > 0 {
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if pl1.Cooldown == 0 {
			g.World.AddBullet(pl1)
			pl1.Cooldown = 11
		}
	}

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

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if pl1.Rotation > 0 {
			pl1.Rotation -= 4
		} else {
			pl1.Rotation = 356
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if pl1.Rotation < 360 {
			pl1.Rotation += 4
		} else {
			pl1.Rotation = 4
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.World.AddAst()
	}
}
