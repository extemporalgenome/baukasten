package glfw

import (
	"github.com/Agon/baukasten"
	"github.com/jteeuwen/glfw"
)

const (
	ChanBuffer = 1
)

var DefaultDriver = NewGlfwDriver()

type GlfwDriver struct {
	graphicSettings    *baukasten.GraphicSettings
	resizeEvent        chan baukasten.WindowSizeEvent
	contextEvent       chan baukasten.ContextEvent
	keyEvent           chan baukasten.KeyEvent
	mouseButtonEvent   chan baukasten.MouseButtonEvent
	mousePositionEvent chan baukasten.MousePositionEvent
	mouseWheelEvent    chan baukasten.MouseWheelEvent
}

func NewGlfwDriver() *GlfwDriver {
	return &GlfwDriver{
		resizeEvent:        make(chan baukasten.WindowSizeEvent, ChanBuffer),
		contextEvent:       make(chan baukasten.ContextEvent, ChanBuffer),
		keyEvent:           make(chan baukasten.KeyEvent, ChanBuffer),
		mouseButtonEvent:   make(chan baukasten.MouseButtonEvent, ChanBuffer),
		mousePositionEvent: make(chan baukasten.MousePositionEvent, ChanBuffer),
		mouseWheelEvent:    make(chan baukasten.MouseWheelEvent, ChanBuffer),
	}
}

func (driver *GlfwDriver) Init(settings *baukasten.GraphicSettings) (err error) {
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
		case driver.resizeEvent <- NewWindowSize(uint(w), uint(h)):
		default:
		}
	})
	glfw.SetWindowCloseCallback(func() int {
		select {
		case driver.contextEvent <- ContextEvent(baukasten.SystemQuit):
		default:
		}
		return 0
	})
	glfw.SetKeyCallback(func(key, state int) {
		select {
		case driver.keyEvent <- NewKeyEvent(key, state):
		default:
		}
	})

	driver.graphicSettings = settings
	return nil
}

func (driver *GlfwDriver) ResizeEvent() chan baukasten.WindowSizeEvent {
	return driver.resizeEvent
}

func (driver *GlfwDriver) ContextEvent() chan baukasten.ContextEvent {
	return driver.contextEvent
}

func (driver *GlfwDriver) KeyEvent() chan baukasten.KeyEvent {
	return driver.keyEvent
}

func (driver *GlfwDriver) MouseButtonEvent() chan baukasten.MouseButtonEvent {
	return driver.mouseButtonEvent
}
func (driver *GlfwDriver) MousePositionEvent() chan baukasten.MousePositionEvent {
	return driver.mousePositionEvent
}
func (driver *GlfwDriver) MouseWheelEvent() chan baukasten.MouseWheelEvent {
	return driver.mouseWheelEvent
}

func (driver *GlfwDriver) Close() {
	glfw.Terminate()
	glfw.CloseWindow()
}

func (driver *GlfwDriver) SwapBuffers() {
	glfw.SwapBuffers()
}
