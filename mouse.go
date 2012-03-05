package baukasten

type MousePosition interface {
	X() int
	Y() int
}

type MouseButton interface {
	Button() int
	State() int
}

type MouseWheel interface {
	Delta() int
}
