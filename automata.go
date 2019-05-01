package main

import "math/rand"

func makeByteMap(width, height int) [][]byte {
	m2d := make([][]byte, height)

	for i := range m2d {
		m2d[i] = make([]byte, width)
		for j := range m2d[i] {
			m2d[i][j] = byte(' ')
		}
	}

	return m2d
}

func generateAutomataData() [][]byte {
	const width = 40
	const height = 30

	m2d := makeByteMap(width, height)
	x, y := width/2, height/2

	const steps = 400
	for i := 0; i < steps; i++ {
		if boundViolation(x, y, 1, width-1, height-1) {
			x, y = width/2, height/2
		}

		dir := rand.Intn(4)

		if m2d[y][x] == byte(' ') {
			m2d[y][x] = byte('.')
		} else {
			m2d[y][x] = byte(' ')
		}

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

	return m2d
}

func automataDataWallProcessing(m2d [][]byte) [][]byte {
	for i, row := range m2d {
		for j, col := range row {
			if col == byte(' ') && surroundedBy(m2d, j, i, byte('.')) {
				m2d[i][j] = byte('#')
			}
		}
	}

	return m2d
}

func surroundedBy(m2d [][]byte, x, y int, by byte) bool {
	boundRight := x < len(m2d[0])-1
	boundDown := y < len(m2d)-1
	boundUp := y > 0
	boundLeft := x > 0

	up := boundUp && m2d[y-1][x] == by
	down := boundDown && m2d[y+1][x] == by

	right := boundRight && m2d[y][x+1] == by
	rightUp := boundRight && boundUp && m2d[y-1][x+1] == by
	rightDown := boundRight && boundDown && m2d[y+1][x+1] == by

	left := boundLeft && m2d[y][x-1] == by
	leftUp := boundUp && boundLeft && m2d[y-1][x-1] == by
	leftDown := boundDown && boundLeft && m2d[y+1][x-1] == by

	return up || down || right || rightUp || rightDown || left || leftUp || leftDown
}

func boundViolation(x, y, size, width, height int) bool {
	return x < 0 || x+size >= width || y < 0 || y+size >= height
}
