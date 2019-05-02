package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	rand.Seed(time.Now().Unix())
	w, h := 180, 60
	roomMinW, roomMinH, roomMaxW, roomMaxH := 7, 4, 18, 10
	sectorlow, sectorhigh := 100, 600

	b := parametricGenerate(
		w, h,
		roomMinW, roomMinH, roomMaxW, roomMaxH,
		sectorlow, sectorhigh,
	)
	s := byteMapToStringMap(b)
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
