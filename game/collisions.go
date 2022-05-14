package game

import "github.com/skycoin/cx-space-arena/constants"

func (game *Game) UpdateCollisions() {
	for _, asteroid := range game.World.Asteroids {
		if asteroid == nil {
			continue
		}
		for i, bullet := range game.World.Bullets {
			if bullet == nil {
				continue
			}
			if overlap(bullet.PositionX-constants.BULLET_HITBOX_OFFSET, bullet.PositionY-constants.BULLET_HITBOX_OFFSET, bullet.PositionX+constants.BULLET_HITBOX_OFFSET, bullet.PositionY+constants.BULLET_HITBOX_OFFSET,
				asteroid.PositionX-constants.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionY-constants.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionX+constants.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionY+constants.ASTEROID_HITBOX_OFFSET*asteroid.Mass) {
				asteroid.Health -= constants.BULLET_DAMAGE
				game.World.Bullets[i] = nil
			}
		}
		for _, player := range game.World.Players {
			if player == nil || player.Recoil > 0 {
				continue
			}
			if overlap(player.PositionX-constants.PLAYER_HITBOX_OFFSET, player.PositionY-constants.PLAYER_HITBOX_OFFSET, player.PositionX+constants.PLAYER_HITBOX_OFFSET, player.PositionY+constants.PLAYER_HITBOX_OFFSET,
				asteroid.PositionX-constants.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionY-constants.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionX+constants.ASTEROID_HITBOX_OFFSET*asteroid.Mass, asteroid.PositionY+constants.ASTEROID_HITBOX_OFFSET*asteroid.Mass) {
				player.Health -= constants.ASTEROID_DAMAGE
				player.Velocity = constants.RECOIL_VELOCITY
				player.Recoil = constants.RECOIL_FRAMES
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
