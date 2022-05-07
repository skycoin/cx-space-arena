package bullet

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/assets"
	"github.com/skycoin/cx-space-arena/consts"
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
func NewBullet(player *player.Player) *Bullet {
	return &Bullet{
		PositionX: player.PositionX + math.Cos(float64(player.Rotation)*(math.Pi/180))*50, PositionY: player.PositionY + math.Sin(float64(player.Rotation)*(math.Pi/180))*50,
		Angle:    player.Rotation,
		Player:   player,
		Lifespan: consts.BULLET_LIFESPAN,
		Sprite:   assets.SpriteList.Bullet,
	}
}

func IndexOf(bullets []*Bullet, bullet *Bullet) int {
	for i, v := range bullets {
		if v == bullet {
			return i
		}
	}
	return -1
}
