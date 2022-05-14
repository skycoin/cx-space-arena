package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/constants"
)

func (player *Player) Update() {
	if player.Recoil > 0 {
		player.Recoil--
	}
	if player.Cooldown > 0 {
		player.Cooldown--
	}
	if player.Velocity < 0 {
		player.Velocity += constants.VELOCITY_PER_FRAME
	}
	if player.Velocity > 0 {
		player.Velocity -= constants.VELOCITY_PER_FRAME
	}

	player.PositionX += math.Cos(float64(player.Rotation)*(math.Pi/180)) * float64(player.Velocity) / 10
	player.PositionY += math.Sin(float64(player.Rotation)*(math.Pi/180)) * float64(player.Velocity) / 10

	w, h := ebiten.WindowSize()
	if player.PositionX < 0 {
		player.PositionX = 0
	}
	if player.PositionX > float64(w) {
		player.PositionX = float64(w)
	}
	if player.PositionY < 0 {
		player.PositionY = 0
	}
	if player.PositionY > float64(h) {
		player.PositionY = float64(h)
	}
}
