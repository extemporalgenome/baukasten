package glfw

type MouseButtonEvent struct {
	button int
	state  int
}

func NewMouseButtonEvent(button, state int) *MouseButtonEvent {
	return &MouseButtonEvent{button: button, state: state}
}

func (e *MouseButtonEvent) Button() int {
	return e.button
}

func (e *MouseButtonEvent) State() int {
	return e.state
}

type MousePositionEvent struct {
	x, y int
}

func NewMousePositionEvent(x, y int) *MousePositionEvent {
	return &MousePositionEvent{x:x, y:y}
}

func (e *MousePositionEvent) X() int {
	return e.x
}

func (e *MousePositionEvent) Y() int {
	return e.y
}

type MouseWheelEvent int

func (e MouseWheelEvent) Delta() int { return int(e) }
