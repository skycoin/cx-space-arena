package game

func (g *Game) UpdateCollisions() {
	for i, b := range g.World.Bullets {
		if b == nil {
			continue
		}
		for _, a := range g.World.Asteroids {
			if a == nil {
				continue
			}
			if overlap(b.PositionX-25, b.PositionY-25, b.PositionX+25, b.PositionY+25,
				a.PositionX-12.5*a.Mass, a.PositionY-12.5*a.Mass, a.PositionX+12.5*a.Mass, a.PositionY+12.5*a.Mass) {
				a.Health -= 30
				g.World.Bullets[i] = nil
			}
		}
	}
}

func overlap(l1x float64, l1y float64, r1x float64, r1y float64,
	l2x float64, l2y float64, r2x float64, r2y float64) bool {

	if l1x >= r2x || r1x <= l2x || l1y >= r2y || r1y <= l2y {
		return false
	}
	return true
}
