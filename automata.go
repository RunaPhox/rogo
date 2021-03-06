package main

import "math/rand"

func generateAutomataData(width, height int) [][]byte {
	const lower = 1
	const high = 9
	size := 5

	m2d := makeByteMap(width, height)
	x, y := width/2, height/2
	prevX, prevY := x, y
	prevSize := size

	const steps = 1000
	for i := 0; i < steps; i++ {
		dir := rand.Intn(4)

		b := byte('.')

		prevSize = size
		if !boundViolation(x, y, size, width, height) && m2d[y][x] == b {
			size--
			if size < lower {
				size = lower
			}
		} else {
			size++
		}

		if boundViolation(x, y, size, width, height) {
			x, y, size = prevX, prevY, prevSize
		}

		for n := y; n < y+size-1; n++ {
			for m := x; m < x+size-1; m++ {
				m2d[n][m] = b
			}
		}

		prevX, prevY = x, y
		switch dir {
		case 0:
			x++
		case 1:
			x--
		case 2:
			y++
		case 3:
			y--
		}
	}

	m2d[height/2][width/2] = byte('@')

	return m2d
}

func boundViolation(x, y, size, width, height int) bool {
	return x < 0 || x+size >= width || y < 0 || y+size >= height
}
