package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprites struct {
	Player   *ebiten.Image
	Asteroid []*ebiten.Image
	Bullet   *ebiten.Image
}

var SpriteList *Sprites

func init() {
	player, _, err := ebitenutil.NewImageFromFile("assets/ship1.png")
	if err != nil {
		log.Fatal(err)
	}
	bullet, _, err := ebitenutil.NewImageFromFile("assets/bullet.png")
	if err != nil {
		log.Fatal(err)
	}
	asteroid, _, err := ebitenutil.NewImageFromFile("assets/asteroid.png")
	if err != nil {
		log.Fatal(err)
	}
	asteroidDamaged, _, err := ebitenutil.NewImageFromFile("assets/asteroidDamaged.png")
	if err != nil {
		log.Fatal(err)
	}
	asteroidBroken, _, err := ebitenutil.NewImageFromFile("assets/asteroidBroken.png")
	if err != nil {
		log.Fatal(err)
	}
	SpriteList = &Sprites{

		Player:   player,
		Asteroid: []*ebiten.Image{asteroid, asteroidDamaged, asteroidBroken},
		Bullet:   bullet,
	}
}
