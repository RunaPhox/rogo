package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
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
			player.SetPosition(x, y+2)
		}
	}
}

func main() {
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		Ch: '.',
	})

	player := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
	}
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	level.AddEntity(&player)

	game := tl.NewGame()
	game.Screen().SetLevel(level)
	game.Start()
}
