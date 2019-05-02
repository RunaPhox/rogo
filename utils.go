package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func printSectorPoints(w, h int, p []Point2d) {
	byt := make([][]rune, h)
	for i := range byt {
		byt[i] = make([]rune, w)
		for j := range byt[i] {
			byt[i][j] = ' '
		}
	}

	for _, v := range p {
		byt[v.y][v.x] = '+'
	}

	for i, v := range byt {
		for j, v2 := range v {
			if i == 0 || i == len(byt)-1 ||
				j == 0 || j == len(byt[0])-1 {
				fmt.Printf("#")
			} else {
				fmt.Printf("%c", v2)
			}
		}
		fmt.Println()
	}
	fmt.Println(len(p), "points")
}

func writeMap(path string, lvl []string) {
	file, err := os.Create(path)
	errCheck(err)
	defer file.Close()

	for _, v := range lvl {
		file.WriteString(v)
		file.WriteString("\n")
	}
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

func byteMapToStringMap(bmap [][]byte) []string {
	arr := make([]string, len(bmap))
	for i, v := range bmap {
		arr[i] = string(v)
	}

	return arr
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
