package main

import tl "github.com/JoelOtter/termloop"

type Player struct {
	*tl.Entity
	prevX int
	prevY int
}

// Drawable
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		x, y := player.Position()

		switch event.Key {
		case tl.KeyArrowRight:
			player.SetPosition(x+1, y)
		case tl.KeyArrowLeft:
			player.SetPosition(x-1, y)
		case tl.KeyArrowUp:
			player.SetPosition(x, y-1)
		case tl.KeyArrowDown:
			player.SetPosition(x, y+1)
		}
	}
}
