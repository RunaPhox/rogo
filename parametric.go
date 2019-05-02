package main

import "math/rand"

type Point2d struct {
	x, y int
}

/*
func parametricGenerate(width, height, roomhigh, sectorhight, sectorlow int) [][]byte {
	m2d := makeByteMap(width, height)

	var points []Point2d
	bsp(0, 0, width, height, sectorhight, sectorlow, true, &points)
}
*/

func bsp(x, y, w, h, high, low int, vert bool, p *[]Point2d) {
	if w*h < high {
		if true {
			x1, y1 := rand.Intn(w)+x, rand.Intn(h)+y
			*p = append(*p, Point2d{x1, y1})
			return
		}
	}

	if w*h <= low {
		x1, y1 := rand.Intn(w)+x, rand.Intn(h)+y
		*p = append(*p, Point2d{x1, y1})
		return
	} else {
		if vert {
			bsp(x, y, w/2, h, high, low, !vert, p)
			bsp(x+w/2, y, w/2, h, high, low, !vert, p)
		} else {
			bsp(x, y, w, h/2, high, low, !vert, p)
			bsp(x, y+h/2, w, h/2, high, low, !vert, p)
		}
	}

}
