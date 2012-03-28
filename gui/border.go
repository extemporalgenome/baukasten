package gui

import (
	"image/color"

	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/geometry"
)

type BorderStyle byte

const (
	BorderStyleSolid = iota
)

type Border struct {
	Style BorderStyle
	Width float32
	Color color.Color

	size geometry.Vector2

	parent Element
}

func NewBorder(parent Element) *Border {
	return &Border{
		Style:  BorderStyleSolid,
		Width:  4.0,
		Color:  geometry.Black,
		parent: parent,
	}
}

func (b *Border) Parent() Element {
	return b.parent
}

func (b *Border) Draw(engine *baukasten.Engine, x, y float32) {
	if b.parent != nil {
		b.size = b.parent.Size()
	}
	if b.Width == 0 {
		return
	}
	topLeft := geometry.Vec2(x-b.size.X/2, y+b.size.Y/2)
	topRight := geometry.Vec2(x+b.size.X/2, y+b.size.Y/2)
	botRight := geometry.Vec2(x+b.size.X/2, y-b.size.Y/2)
	botLeft := geometry.Vec2(x-b.size.X/2, y-b.size.Y/2)
	engine.DrawLineLoop(b.Color, topLeft, topRight, botRight, botLeft)
}
