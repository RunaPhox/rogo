package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
}

type Wall struct {
	*tl.Entity
}

func readMap(path string) string {
	file, err := os.Open(path)
	errCheck(err)
	defer file.Close()

	info, err := file.Stat()
	errCheck(err)

	byt := make([]byte, info.Size())
	file.Read(byt)

	return string(byt)
}

func getMap(path string) []string {
	return strings.Split(readMap(path), "\n")
}

func proccessMap(lvl []string) {
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
	})

	player := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
	}
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	level.AddEntity(&player)

	game := tl.NewGame()
	game.Screen().SetLevel(level)
	game.Start()

	arr := getMap("maps/lvl1")
	for _, v := range arr {
		fmt.Println(v)
	}
}
