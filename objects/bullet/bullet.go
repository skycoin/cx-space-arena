package bullet

import (
	"github.com/skycoin/cx-space-arena/objects/player"
)

type Bullet struct {
	PositionX, PositionY float64
	FireAngle            int
	FromPlayer           *player.Player
	Lifespan             int
}
