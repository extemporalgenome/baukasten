package baukasten

const (
	SystemEvent_Quit = iota
)

type FrameEvent struct {
	DeltaTime uint
	Ticks     uint32
}

type SystemEvent int

type ResizeEvent struct {
	Width, Height int
}

type MouseMotionEvent interface {
	Position() (int, int)
	Difference() (int, int)
}

type KeyboardEvent interface {
	// TODO Keyboard
}

type MouseButtonEvent interface {
	// TODO MouseButton
}
