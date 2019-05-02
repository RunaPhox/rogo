package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	/*
		b := parametricGenerate(200, 200, 5, 22, 3, 12)
		b2 := automataDataWallProcessing(b)
		s := byteMapToStringMap(b2)
		writeMap("maps/rand", s)
	*/
	/*
		level := tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorBlack,
		})

		arr := getMap("maps/rand")
		proccessMap(level, arr)

		game := tl.NewGame()
		game.Screen().SetLevel(level)
		game.Start()
	*/

	var points []Point2d
	bsp(0, 0, 210, 46, 150, 92, true, &points)
	printSectorPoints(points)
}
