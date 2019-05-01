package main

import (
	tl "github.com/JoelOtter/termloop"
)

func main() {
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	arr := getMap("maps/lvl1")
	proccessMap(level, arr)

	game := tl.NewGame()
	game.Screen().SetLevel(level)
	game.Start()

}
