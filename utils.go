package main

import (
	"log"
	"os"
	"strings"
)

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
