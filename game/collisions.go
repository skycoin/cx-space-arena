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
			if overlap(b.PositionX-25, b.PositionY-25, a.PositionX+25, a.PositionY+25,
				a.PositionX-12.5*a.Mass, a.PositionY-12.5*a.Mass, a.PositionX+12.5*a.Mass, a.PositionY+12.5*a.Mass) {
				a.Health -= 30
				g.World.Bullets[i] = nil
			}
		}
	}
}

func overlap(l1x float64, l1y float64, r1x float64, r1y float64,
	l2x float64, l2y float64, r2x float64, r2y float64) bool {
	if l1x >= r2x || l2x >= r1x {
		return false
	}

	if r1y >= l2y || r2y >= l1y {
		return false
	}

	return true
}