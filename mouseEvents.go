package baukasten

type MousePositionEvent interface {
	X() int
	Y() int
}

type MouseButtonEvent interface {
	Button() int
	State() int
}

type MouseWheelEvent interface {
	Delta() int
}
