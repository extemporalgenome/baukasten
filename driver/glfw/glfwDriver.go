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
	resizeEvent        chan baukasten.WindowSize
	contextEvent       chan baukasten.ContextEvent
	keyEvent           chan baukasten.Key
	mouseButtonEvent   chan baukasten.MouseButton
	mousePositionEvent chan baukasten.MousePosition
	mouseWheelEvent    chan baukasten.MouseWheel
}

func NewDriver() *Driver {
	return &Driver{
		resizeEvent:        make(chan baukasten.WindowSize, ChanBuffer),
		contextEvent:       make(chan baukasten.ContextEvent, ChanBuffer),
		keyEvent:           make(chan baukasten.Key, ChanBuffer),
		mouseButtonEvent:   make(chan baukasten.MouseButton, ChanBuffer),
		mousePositionEvent: make(chan baukasten.MousePosition, ChanBuffer),
		mouseWheelEvent:    make(chan baukasten.MouseWheel, ChanBuffer),
	}
}

// ### ContextDriver implementations ###

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
	d.graphicSettings = settings
	return nil
}

func (d *Driver) Close() {
	glfw.Terminate()
	glfw.CloseWindow()
}

func (d *Driver) SwapBuffers() {
	glfw.SwapBuffers()
}

func (d *Driver) SetResizeCallback(callback chan<- baukasten.WindowSize) {
	glfw.SetWindowSizeCallback(func(width, height int) {
		callback <- NewWindowSize(width, height)
	})
}

func (d *Driver) SetContextCallback(callback chan<- baukasten.ContextEvent) {
	glfw.SetWindowCloseCallback(func() int {
		callback <- ContextEvent(baukasten.WindowClose)
		return 0
	})
	glfw.SetWindowRefreshCallback(func() {
		callback <- ContextEvent(baukasten.WindowRefresh)
	})
}

// ### InputDriver implementations ###

func (d *Driver) SetKeyCallback(callback chan<- baukasten.Key) {
	glfw.SetKeyCallback(func(key, state int) {
		callback <- NewKeyEvent(key, state)
	})
}

func (d *Driver) SetMouseButtonCallback(callback chan<- baukasten.MouseButton) {
	glfw.SetMouseButtonCallback(func(button, state int) {
		callback <- NewMouseButtonEvent(button, state)
	})
}

func (d *Driver) SetMousePositionCallback(callback chan<- baukasten.MousePosition) {
	glfw.SetMousePosCallback(func(x, y int) {
		callback <- NewMousePositionEvent(x, y)
	})
}

func (d *Driver) SetMouseWheelCallback(callback chan<- baukasten.MouseWheel) {
	glfw.SetMouseWheelCallback(func(delta int) {
		callback <- MouseWheelEvent(delta)
	})
}

func (d *Driver) MousePos() baukasten.MousePosition {
	x, y := glfw.MousePos()
	return NewMousePositionEvent(x, y)
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
