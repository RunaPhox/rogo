package main

import (
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

// Physical for Wall
func (w *Wall) Position() (int, int) {
	return w.Position()
}

func (w *Wall) Size() (int, int) {
	return w.Size()
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

	player := Player{
		Entity: tl.NewEntity(pX, pY, 1, 1),
	}
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	level.AddEntity(&player)

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
			player.SetPosition(x, y+1)
		}
	}
}

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
