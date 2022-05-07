package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (p *Player) Update() {
	if p.Recoil > 0 {
		p.Recoil--
	}
	if p.Cooldown > 0 {
		p.Cooldown--
	}
	if p.Velocity < 0 {
		p.Velocity += 5
	}
	if p.Velocity > 0 {
		p.Velocity -= 5
	}

	p.PositionX += math.Cos(float64(p.Rotation)*(math.Pi/180)) * float64(p.Velocity) / 10
	p.PositionY += math.Sin(float64(p.Rotation)*(math.Pi/180)) * float64(p.Velocity) / 10

	w, h := ebiten.WindowSize()
	if p.PositionX < 0 {
		p.PositionX = 0
	}
	if p.PositionX > float64(w) {
		p.PositionX = float64(w)
	}
	if p.PositionY < 0 {
		p.PositionY = 0
	}
	if p.PositionY > float64(h) {
		p.PositionY = float64(h)
	}
}
