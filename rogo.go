package main

import (
	"math/rand"
	"time"
)

func main() {
	/*
		level := tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorBlack,
		})

		arr := getMap("maps/lvl1")
		proccessMap(level, arr)

		game := tl.NewGame()
		game.Screen().SetLevel(level)
		game.Start()
	*/

	rand.Seed(time.Now().Unix())
	b := generateAutomataData(70, 30)
	b2 := automataDataWallProcessing(b)
	s := byteMapToStringMap(b2)
	writeMap("maps/rand", s)
}
