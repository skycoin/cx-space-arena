package world

import "github.com/skycoin/cx-space-arena/objects"

type World struct {
	Players   []objects.Player
	Asteroids []objects.Asteroid
	Bullets   []objects.Bullet
}
