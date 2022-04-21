package world

import (
	"errors"

	"github.com/skycoin/cx-space-arena/objects/asteroid"
	"github.com/skycoin/cx-space-arena/objects/bullet"
	"github.com/skycoin/cx-space-arena/objects/player"
)

type World struct {
	Asteroids []*asteroid.Asteroid
	Bullets   []*bullet.Bullet
	Players   [5]*player.Player
	PlayerNum int
}

// Add a new player with default settings to the World
func (world *World) AddNewPlayer() (*World, error) {
	if world.PlayerNum == 5 {
		return world, errors.New("Max number of players reached")
	}
	world.PlayerNum++
	world.Players[world.PlayerNum-1] = player.NewPlayer(world.PlayerNum)
	return world, nil
}

// Remove a player from the World
func (world *World) DespawnPlayer(p *player.Player) (*World, error) {
	if world.Players[p.GetID()] == nil {
		return world, errors.New("Cannot despawn empty slot")
	}
	world.Players[p.GetID()] = nil
	return world, nil
}

// Add an existing player with its current settings back to the World
func (world *World) RespawnPlayer(p *player.Player) (*World, error) {
	if world.Players[p.GetID()] != nil {
		return world, errors.New("Player slot already occupied")
	}
	world.Players[p.GetID()] = p
	return world, nil
}
