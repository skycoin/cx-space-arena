package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (game *Game) UpdateInputs() {
	player1 := game.World.Players[0]
	if player1.Recoil > 0 {
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if player1.Cooldown == 0 {
			game.World.AddBullet(player1)
			player1.Cooldown = 11
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if player1.Velocity < 70 {
			player1.Velocity += 15
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if player1.Velocity > -50 {
			player1.Velocity -= 15
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if player1.Rotation > 0 {
			player1.Rotation -= 4
		} else {
			player1.Rotation = 356
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if player1.Rotation < 360 {
			player1.Rotation += 4
		} else {
			player1.Rotation = 4
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		game.World.AddAst()
	}
}
