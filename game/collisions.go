package game

import "github.com/skycoin/cx-space-arena/consts"

func (g *Game) UpdateCollisions() {
	for _, asteroid := range g.World.Asteroids {
		if asteroid == nil {
			continue
		}
		for i, bullet := range g.World.Bullets {
			if bullet == nil {
				continue
			}
			if overlap(bullet.PositionX-consts.BULLET_HITBOX_OFFSET, bullet.PositionY-consts.BULLET_HITBOX_OFFSET, bullet.PositionX+consts.BULLET_HITBOX_OFFSET, bullet.PositionY+consts.BULLET_HITBOX_OFFSET,
				asteroid.PositionX-consts.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionY-consts.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionX+consts.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionY+consts.ASTEROID_HITBOX_OFFSET*asteroid.Mass) {
				asteroid.Health -= consts.BULLET_DAMAGE
				g.World.Bullets[i] = nil
			}
		}
		for _, player := range g.World.Players {
			if player == nil || player.Recoil > 0 {
				continue
			}
			if overlap(player.PositionX-consts.PLAYER_HITBOX_OFFSET, player.PositionY-consts.PLAYER_HITBOX_OFFSET, player.PositionX+consts.PLAYER_HITBOX_OFFSET, player.PositionY+consts.PLAYER_HITBOX_OFFSET,
				asteroid.PositionX-consts.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionY-consts.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionX+consts.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionY+consts.ASTEROID_HITBOX_OFFSET*asteroid.Mass) {
				player.Health -= consts.ASTEROID_DAMAGE
				player.Velocity = consts.RECOIL_VELOCITY
				player.Recoil = consts.RECOIL_FRAMES
			}
		}
	}
}

func overlap(point1x float64, point1y float64, point2x float64, point2y float64,
	point3x float64, point3y float64, point4x float64, point4y float64) bool {

	if point1x >= point4x || point2x <= point3x || point1y >= point4y || point2y <= point3y {
		return false
	}
	return true
}
