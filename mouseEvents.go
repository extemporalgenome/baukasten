package baukasten

import (
	"image"
)

type MousePositionEvent interface {
	Position() image.Point
}

type MouseButtonEvent interface {
	Button() uint
	State() uint
}

type MouseWheelEvent interface {
	Delta() int
}
