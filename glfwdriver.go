package baukasten

import (
	"os"

	"gl"
	"github.com/jteeuwen/glfw"
)

type GlfwDriver struct {
	graphicSettings *GraphicSettings
}

func NewGlfwDriver() *GlfwDriver {
	return &GlfwDriver{}
}

func (driver *GlfwDriver) Init(settings *GraphicSettings) os.Error {
	err := glfw.Init()
	if err != nil {
		return err
	}
	windowType := glfw.Windowed
	if settings.Fullscreen {
		windowType = glfw.Fullscreen
	}
	// TODO BitDepth
	err = glfw.OpenWindow(settings.Width, settings.Height, 8, 8, 8, 8, 0, 8, windowType)
	if err != nil {
		glfw.Terminate()
		return err
	}
	glfw.SetSwapInterval(1) // VSync
	glfw.SetWindowTitle(settings.Title)

	glfw.SetWindowSizeCallback(onResize)

	driver.graphicSettings = settings
	return nil
}

func onResize(w, h int) {
	gl.Viewport(0, 0, w, h)
}

func (driver *GlfwDriver) Close() {
	glfw.Terminate()
	glfw.CloseWindow()
}

func (driver *GlfwDriver) SwapBuffers() {
	glfw.SwapBuffers()
}

// InputDriver implementations
func (driver *GlfwDriver) GetKeyEvent()         {}
func (driver *GlfwDriver) GetMouseButtonEvent() {}
