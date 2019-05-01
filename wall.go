package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Wall struct {
	*tl.Entity
}

// Physical for Wall
func (w *Wall) Position() (int, int) {
	return w.Entity.Position()
}

func (w *Wall) Size() (int, int) {
	return w.Entity.Size()
}
