package gui

import (
	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/geometry"
)

type Element interface {
	Parent() Element
	Draw(engine *baukasten.Engine, x, y float32)
	Size() geometry.Vector2
}
