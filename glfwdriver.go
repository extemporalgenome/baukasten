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
// TODO Move image loading to the Go package image
func (driver *GlfwDriver) LoadSurface(name string) (surface Surface, err os.Error) {
	texture := gl.GenTexture()
	texture.Bind(gl.TEXTURE_2D)

	if glfw.LoadTexture2D(name, 0) {
		return nil, os.NewError("Failed to load texture: " + name)
	}
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	texture.Unbind(gl.TEXTURE_2D)
	return &OglSurface{texture}, nil
}

// InputDriver implementations
func (driver *GlfwDriver) GetKeyEvent()         {}
func (driver *GlfwDriver) GetMouseButtonEvent() {}
