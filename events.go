package baukasten

import (
	"sdl"
)

const (
	// TODO: Add missing SDL key codes
	KeyW = uint(sdl.K_w)
	KeyA = uint(sdl.K_a)
	KeyS = uint(sdl.K_s)
	KeyD = uint(sdl.K_d)
	KeyL = uint(sdl.K_l)

	KeyF1 = uint(sdl.K_F1)
	KeyF2 = uint(sdl.K_F2)

	KeyUp       = uint(sdl.K_UP)
	KeyDown     = uint(sdl.K_DOWN)
	KeyRight    = uint(sdl.K_RIGHT)
	KeyLeft     = uint(sdl.K_LEFT)
	KeyPageUp   = uint(sdl.K_PAGEUP)
	KeyPageDown = uint(sdl.K_PAGEDOWN)

	// KeyState
	KeyPressed  = uint(sdl.KEYDOWN)
	KeyReleased = uint(sdl.KEYUP)
)

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
