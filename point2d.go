package main

import (
	"math"
	"math/rand"
)

type Point2d struct {
	x, y int
}

func appendPoint(x, y, w, h int, p *[]Point2d) {
	x1, y1 := x+w/2+rand.Intn(w/2)-w/4, y+h/2+rand.Intn(h/2)-h/4
	for sameXOrY(*p, Point2d{x1, y1}) {
		x1, y1 = x+w/2+rand.Intn(w/2)-w/2, y+h/2+rand.Intn(h/2)-h/2
	}

	*p = append(*p, Point2d{x1, y1})
}

func orderedPointsFromMin(p1, p2 Point2d) (minX, minY, maxX, maxY int) {
	minX = int(math.Min(float64(p1.x), float64(p2.x)))
	minY = int(math.Min(float64(p1.y), float64(p2.y)))
	maxX = int(math.Max(float64(p1.x), float64(p2.x)))
	maxY = int(math.Max(float64(p1.y), float64(p2.y)))

	return
}

func sameXOrY(points []Point2d, p Point2d) bool {
	for _, v := range points {
		if p.x == v.x || p.y == v.y {
			return true
		}
	}
	return false
}
