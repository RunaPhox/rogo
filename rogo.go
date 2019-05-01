package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	rand.Seed(time.Now().Unix())
	b := generateAutomataData(70, 30)
	b2 := automataDataWallProcessing(b)
	s := byteMapToStringMap(b2)
	writeMap("maps/rand", s)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	arr := getMap("maps/rand")
	proccessMap(level, arr)

	game := tl.NewGame()
	game.Screen().SetLevel(level)
	game.Start()
}
