package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
	prevX int
	prevY int
}

func (p *Player) Position() (int, int) {
	return p.Entity.Position()
}

func (p *Player) Size() (int, int) {
	return p.Entity.Size()
}

func (p *Player) Collide(phys tl.Physical) {
	if _, ok := phys.(*Wall); ok {
		p.SetPosition(p.prevX, p.prevY)
	}
}

// Drawable
func (p *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		p.prevX, p.prevY = p.Position()

		switch event.Key {
		case tl.KeyArrowRight:
			p.SetPosition(p.prevX+1, p.prevY)
		case tl.KeyArrowLeft:
			p.SetPosition(p.prevX-1, p.prevY)
		case tl.KeyArrowUp:
			p.SetPosition(p.prevX, p.prevY-1)
		case tl.KeyArrowDown:
			p.SetPosition(p.prevX, p.prevY+1)
		}
	}
}
