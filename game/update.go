package game

// Update function, called 60 times each second
func (game *Game) Update() error {
	game.World.CurrentTick++
	if game.World.CurrentTick%60 == 0 {
		game.World.CurrentTick = 0
	}
	game.UpdateInputs()
	for _, player := range game.World.Players {
		if player != nil {
			player.Update()
		}
	}
	for i, bullet := range game.World.Bullets {
		if bullet != nil {
			if !bullet.Update() {
				game.World.Bullets[i] = nil
			}
		}
	}
	for i, asteroid := range game.World.Asteroids {
		if asteroid != nil {
			if !asteroid.Update() {
				game.World.Asteroids[i] = nil
			}
		}
	}
	game.UpdateCollisions()
	return nil
}
