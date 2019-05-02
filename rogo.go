package main

import (
	"fmt"
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
	bsp(0, 0, 210, 46, 100, 12, true, &points)

	fmt.Println(len(points))

	byt := make([][]rune, 53)
	for i := range byt {
		byt[i] = make([]rune, 230)
		for j := range byt[i] {
			byt[i][j] = ' '
		}
	}

	for _, v := range points {
		byt[v.y][v.x] = '+'
	}

	for _, v := range byt {
		for _, v2 := range v {
			fmt.Printf("%c", v2)
		}
		fmt.Println()
	}
}
