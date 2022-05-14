package bullet

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/skycoin/cx-space-arena/constants"
)

func (bullet *Bullet) Update() bool {
	w, h := ebiten.WindowSize()
	if bullet.Lifespan == 0 || bullet.PositionX < 0-constants.BULLET_HITBOX_OFFSET || bullet.PositionY < 0-constants.BULLET_HITBOX_OFFSET || bullet.PositionX > float64(w)+constants.BULLET_HITBOX_OFFSET || bullet.PositionY > float64(h)+constants.BULLET_HITBOX_OFFSET {
		return false
	}
	bullet.PositionX += math.Cos(float64(bullet.Angle)*(math.Pi/180)) * 10
	bullet.PositionY += math.Sin(float64(bullet.Angle)*(math.Pi/180)) * 10
	return true
}
