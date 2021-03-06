package main

import (
	tl "github.com/JoelOtter/termloop"
)

func proccessMap(level *tl.BaseLevel, lvl []string) {
	pX, pY := 0, 0

	for y, row := range lvl {
		for x, col := range row {
			switch string(col) {
			case "@":
				pX, pY = x, y

				fallthrough
			case ".":
				floor := tl.NewEntity(x, y, 1, 1)
				floor.SetCell(0, 0, &tl.Cell{Fg: tl.ColorGreen, Ch: '.'})
				level.AddEntity(floor)
			case "#":
				wall := Wall{
					Entity: tl.NewEntity(x, y, 1, 1),
				}
				wall.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlue, Ch: '#'})
				level.AddEntity(&wall)
			}
		}
	}

	// add the player we stored earlier
	player := Player{
		Entity: tl.NewEntity(pX, pY, 1, 1),
		prevX:  pX,
		prevY:  pY,
	}
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	level.AddEntity(&player)
}
