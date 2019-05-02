package main

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

func countSorroundingBytes(m2d [][]byte, x, y int, by byte) int {
	boundRight := x < len(m2d[0])-1
	boundDown := y < len(m2d)-1
	boundUp := y > 0
	boundLeft := x > 0

	conds := []bool{
		// up
		boundUp && m2d[y-1][x] == by,
		// down
		boundDown && m2d[y+1][x] == by,

		// right
		boundRight && m2d[y][x+1] == by,
		// right-up
		boundRight && boundUp && m2d[y-1][x+1] == by,
		// right-down
		boundRight && boundDown && m2d[y+1][x+1] == by,

		// left
		boundLeft && m2d[y][x-1] == by,
		// left-up
		boundLeft && boundUp && m2d[y-1][x-1] == by,
		// left-down
		boundLeft && boundDown && m2d[y+1][x-1] == by,
	}

	cnt := 0
	for _, v := range conds {
		if v {
			cnt++
		}
	}

	return cnt
}

func surroundedBy(m2d [][]byte, x, y int, by byte) bool {
	return countSorroundingBytes(m2d, x, y, by) > 0
}
