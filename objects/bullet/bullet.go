package bullet

import "github.com/skycoin/cx-space-arena/objects/player"

type Bullet struct {
	PositionX, PositionY float64
	FireAngle            int
	FromPlayer           *player.Player
	Lifespan             int
}

// Bullet constructor
func NewBullet(p *player.Player) *Bullet {
	return &Bullet{
		PositionX: p.PositionX, PositionY: p.PositionY,
		FireAngle:  p.Rotation,
		FromPlayer: p,
		Lifespan:   10,
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
