package bullet

import "github.com/skycoin/cx-space-arena/objects/player"

// Bullet constructor
func NewBullet(p *player.Player) *Bullet {
	return &Bullet{
		PositionX: p.PositionX, PositionY: p.PositionY,
		FireAngle:  p.Rotation,
		FromPlayer: p,
		Lifespan:   10,
	}
}
