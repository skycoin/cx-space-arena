package world

import (
	"github.com/skycoin/cx-space-arena/objects/asteroid"
	"github.com/skycoin/cx-space-arena/objects/bullet"
	"github.com/skycoin/cx-space-arena/objects/player"
)

type World struct {
	Asteroids []asteroid.Asteroid
	Bullets   []bullet.Bullet
	Players   []player.Player
}
