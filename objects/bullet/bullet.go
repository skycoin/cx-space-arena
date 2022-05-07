package bullet

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/assets"
	"github.com/skycoin/cx-space-arena/objects/player"
)

type Bullet struct {
	PositionX, PositionY float64
	Angle                int
	Player               *player.Player
	Lifespan             int
	Sprite               *ebiten.Image
}

// Bullet constructor
func NewBullet(p *player.Player) *Bullet {
	return &Bullet{
		PositionX: p.PositionX + math.Cos(float64(p.Rotation)*(math.Pi/180))*50, PositionY: p.PositionY + math.Sin(float64(p.Rotation)*(math.Pi/180))*50,
		Angle:    p.Rotation,
		Player:   p,
		Lifespan: 10,
		Sprite:   assets.SpriteList.Bullet,
	}
}

func IndexOf(bullets []*Bullet, b *Bullet) int {
	for i, v := range bullets {
		if v == b {
			return i
		}
	}
	return -1
}
