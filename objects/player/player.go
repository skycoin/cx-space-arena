package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/assets"
	"github.com/skycoin/cx-space-arena/constants"
)

type Player struct {
	PlayerID             int
	PositionX, PositionY float64
	Velocity             int
	Rotation             int
	Health               int
	Recoil               int
	Points               int
	Deaths               int
	Cooldown             int
	Sprite               *ebiten.Image
}

// Player contructor
func NewPlayer(ID int) *Player {
	w, h := ebiten.WindowSize()
	return &Player{
		PlayerID:  ID,
		PositionX: float64(w) / 2, PositionY: float64(h) / 2,
		Velocity: 0,
		Rotation: 0,
		Health:   constants.PLAYER_HEALTH,
		Recoil:   constants.RECOIL_FRAMES,
		Points:   0,
		Deaths:   0,
		Cooldown: 0,
		Sprite:   assets.SpriteList.Player,
	}
}
