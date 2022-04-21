package world

import (
	"errors"

	"github.com/skycoin/cx-space-arena/objects/asteroid"
	"github.com/skycoin/cx-space-arena/objects/bullet"
	"github.com/skycoin/cx-space-arena/objects/player"
)

type World struct {
	Asteroids   []*asteroid.Asteroid
	Bullets     []*bullet.Bullet
	Players     [5]*player.Player
	PlayerCount int
}

// World constructor
func NewWorld() *World {
	world := &World{
		Players:     [5]*player.Player{nil, nil, nil, nil, nil},
		PlayerCount: 0,
	}
	world.AddNewPlayer()
	return world
}

// Add a new player with default settings to the World
func (world *World) AddNewPlayer() (*World, error) {
	if world.PlayerCount == 5 {
		return world, errors.New("Max number of players reached")
	}
	world.PlayerCount++
	world.Players[world.PlayerCount-1] = player.NewPlayer(world.PlayerCount)
	return world, nil
}
