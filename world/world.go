package world

import (
	"errors"

	"github.com/skycoin/cx-space-arena/objects/asteroid"
	"github.com/skycoin/cx-space-arena/objects/bullet"
	"github.com/skycoin/cx-space-arena/objects/player"
)

type World struct {
	Asteroids    []*asteroid.Asteroid
	Bullets      []*bullet.Bullet
	Players      [5]*player.Player
	PlayerCount  int
	BulletsOnScr int
}

// World constructor
func NewWorld() *World {
	world := &World{
		Players:     [5]*player.Player{nil, nil, nil, nil, nil},
		PlayerCount: 0,
	}
	world.AddPlayer()
	world.AddAst()
	world.AddAst()
	world.AddAst()
	world.AddAst()
	world.AddAst()
	return world
}

// Add a new player to the World
func (world *World) AddPlayer() (*World, error) {
	if world.PlayerCount >= len(world.Players) {
		return world, errors.New("Max player count reached")
	}
	world.Players[world.PlayerCount] = player.NewPlayer(world.PlayerCount + 1)
	world.PlayerCount++
	return world, nil
}

// Add a new bullet belonging to a player to the World
func (world *World) AddBullet(p *player.Player) (*World, error) {
	world.BulletsOnScr++
	world.Bullets = append(world.Bullets, bullet.NewBullet(p))
	return world, nil
}

// Remove bullet at the specified index.
// index will come from the for loop from which the function is called
func (world *World) RemoveBullet(i int) (*World, error) {
	if i >= len(world.Bullets) {
		return world, errors.New("Bullet does not exist")
	}
	world.Bullets = append(world.Bullets[:i], world.Bullets[i+1:]...)
	world.BulletsOnScr--
	return world, nil
}

func (world *World) AddAst() (*World, error) {
	if asteroid.IndexOf(world.Asteroids, nil) == -1 {
		world.Asteroids = append(world.Asteroids, asteroid.NewAsteroid())
	} else {
		world.Asteroids[asteroid.IndexOf(world.Asteroids, nil)] = asteroid.NewAsteroid()
	}
	return world, nil
}
