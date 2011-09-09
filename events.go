package baukasten

// TODO Add keys

type QuitEvent struct {
	Type int
}

type ErrorEvent struct {
	msg string
}

func (e *ErrorEvent) String() string {
	return e.msg
}

type ResizeEvent struct {
	Width, Height int
}

type MouseMotionEvent struct {
	Position   Point
	Difference Point
}

type KeyboardEvent struct {
	State uint
	Key   uint
	Type  uint
}

type MouseButtonEvent struct {
	// TODO MouseButton
}
