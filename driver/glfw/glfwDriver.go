// GLFW 2.7 ContextDriver package for baukasten.
package glfw

import (
	"github.com/jteeuwen/glfw"

	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/geometry"
)

const (
	ChanBuffer = 1
)

var DefaultDriver = NewDriver()

type Driver struct {
	graphicSettings    *baukasten.GraphicSettings
	resizeEvent        chan baukasten.WindowSizeEvent
	contextEvent       chan baukasten.ContextEvent
	keyEvent           chan baukasten.KeyEvent
	mouseButtonEvent   chan baukasten.MouseButtonEvent
	mousePositionEvent chan baukasten.MousePositionEvent
	mouseWheelEvent    chan baukasten.MouseWheelEvent
}

func NewDriver() *Driver {
	return &Driver{
		resizeEvent:        make(chan baukasten.WindowSizeEvent, ChanBuffer),
		contextEvent:       make(chan baukasten.ContextEvent, ChanBuffer),
		keyEvent:           make(chan baukasten.KeyEvent, ChanBuffer),
		mouseButtonEvent:   make(chan baukasten.MouseButtonEvent, ChanBuffer),
		mousePositionEvent: make(chan baukasten.MousePositionEvent, ChanBuffer),
		mouseWheelEvent:    make(chan baukasten.MouseWheelEvent, ChanBuffer),
	}
}

func (d *Driver) Init(settings *baukasten.GraphicSettings) (err error) {
	err = glfw.Init()
	if err != nil {
		return
	}
	// Hint OpenGL 3 context
	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 3)
	glfw.OpenWindowHint(glfw.OpenGLProfile, 1)

	windowType := glfw.Windowed
	if settings.Fullscreen {
		windowType = glfw.Fullscreen
	}
	// TODO Stencil
	// TODO RGBA bits
	err = glfw.OpenWindow(settings.Width, settings.Height, 0, 0, 0, 0, settings.BitDepth, 0, windowType)
	if err != nil {
		glfw.Terminate()
		return err
	}
	glfw.SetSwapInterval(1) // VSync
	glfw.SetWindowTitle(settings.Title)
	glfw.SetWindowSizeCallback(func(w, h int) {
		select {
		case d.resizeEvent <- NewWindowSize(uint(w), uint(h)):
		default:
		}
	})
	glfw.SetWindowCloseCallback(func() int {
		select {
		case d.contextEvent <- ContextEvent(baukasten.SystemQuit):
		default:
		}
		return 0
	})
	glfw.SetKeyCallback(func(key, state int) {
		select {
		case d.keyEvent <- NewKeyEvent(key, state):
		default:
		}
	})
	glfw.SetMouseButtonCallback(func(button, state int) {
		select {
		case d.mouseButtonEvent <- NewMouseButtonEvent(button, state):
		default:
		}
	})
	glfw.SetMousePosCallback(func(x, y int) {
		select {
		case d.mousePositionEvent <- NewMousePositionEvent(x, y):
		default:
		}
	})
	glfw.SetMouseWheelCallback(func(delta int) {
		select {
		case d.mouseWheelEvent <- MouseWheelEvent(delta):
		default:
		}
	})

	d.graphicSettings = settings
	return nil
}

func (d *Driver) ResizeEvent() <-chan baukasten.WindowSizeEvent {
	return d.resizeEvent
}

func (d *Driver) ContextEvent() <-chan baukasten.ContextEvent {
	return d.contextEvent
}

func (d *Driver) KeyEvent() <-chan baukasten.KeyEvent {
	return d.keyEvent
}

func (d *Driver) MouseButtonEvent() <-chan baukasten.MouseButtonEvent {
	return d.mouseButtonEvent
}
func (d *Driver) MousePositionEvent() <-chan baukasten.MousePositionEvent {
	return d.mousePositionEvent
}
func (d *Driver) MouseWheelEvent() <-chan baukasten.MouseWheelEvent {
	return d.mouseWheelEvent
}

func (d *Driver) Close() {
	glfw.Terminate()
	glfw.CloseWindow()
}

func (d *Driver) SwapBuffers() {
	glfw.SwapBuffers()
}

func (d *Driver) JoystickParam(joy, param int) int {
	return glfw.JoystickParam(joy, param)
}

// Two axes support
func (d *Driver) JoystickPos(joy int) []geometry.Vector2 {
	axes := []float32{0, 0, 0, 0}
	l := glfw.JoystickPos(joy, axes)
	if l%2 == 0 {
		return []geometry.Vector2{}
	}
	vecAxes := make([]geometry.Vector2, l/2)
	for i := 0; i < l; i += 2 {
		vecAxes[i] = geometry.Vec2(axes[i], axes[i+1])
	}
	return vecAxes
}

// Ten button support
func (d *Driver) JoystickButtons(joy int) []bool {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	l := glfw.JoystickButtons(joy, b)
	if l == 0 {
		return []bool{}
	}
	buttons := make([]bool, len(b))
	for i, x := range b {
		if x == 0 {
			buttons[i] = false // Released
		} else {
			buttons[i] = true // Pressed
		}
	}
	return buttons
}
