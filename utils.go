package main

import (
	"log"
	"os"
	"strings"
)

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

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
