package main

import (
	"math/rand"
)

type Rect struct {
	p1, p2 Point2d
}

func parametricGenerate(width, height, roomMinW, roomMinH,
	roomMaxW, roomMaxH, sectorlow, sectorhigh int) [][]byte {

	points := generateSectionPoints(width, height, sectorlow, sectorhigh)
	m2d := roomSplash(width, height,
		roomMinW, roomMinH, roomMaxW, roomMaxH, points)
	return byteMapWallProcessing(m2d)
}

func roomSplash(w, h, minW, minH, maxW, maxH int, p []Point2d) [][]byte {
	m2d := makeByteMap(w, h)
	for _, v := range p {
		w, h := rand.Intn(maxW-minW)+minW,
			rand.Intn(maxH-minH)+minH
		m2d = fillRectWithByte(m2d, byte('.'), Rect{
			Point2d{v.x - w/2, v.y - h/2},
			Point2d{v.x + w - w/2 - 1, v.y + h - h/2 - 1},
		})
	}

	return m2d
}

func generateSectionPoints(w, h, low, high int) []Point2d {
	var points []Point2d
	bsp(w/10, h/6, w-w/10, h-h/6, low, high, true, &points)
	if len(points)%2 == 1 {
		points = points[:len(points)-1]
	}

	return points
}

func bsp(x, y, w, h, low, high int, vert bool, p *[]Point2d) {
	if w*h < high && rand.Intn(3) < 2 {
		appendPoint(x, y, w, h, p)
		return
	}

	if w/2*h/2 < low {
		appendPoint(x, y, w, h, p)
		return
	}

	if vert {
		rw := w/2 + rand.Intn(w/4) - w/8
		bsp(x, y, rw, h, low, high, !vert, p)
		bsp(x+rw, y, w-rw, h, low, high, !vert, p)
	} else {
		rh := h/2 + rand.Intn(h/4) - h/8
		bsp(x, y, w, rh, low, high, !vert, p)
		bsp(x, y+rh, w, h-rh, low, high, !vert, p)
	}
}
