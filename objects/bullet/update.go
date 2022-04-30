package bullet

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (b *Bullet) Update() bool {
	b.PositionX += math.Cos(float64(b.Angle)*(math.Pi/180)) * 10
	b.PositionY += math.Sin(float64(b.Angle)*(math.Pi/180)) * 10
	w, h := ebiten.WindowSize()
	if b.Lifespan == 0 || b.PositionX < 0-25 || b.PositionY < 0-25 || b.PositionX > float64(w)+25 || b.PositionY > float64(h)+25 {
		return false
	}
	return true
}
