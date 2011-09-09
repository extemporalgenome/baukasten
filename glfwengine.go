packae baukasten

import (
	"os"
	
	"github.com/jteeuwen/glfw"
)

type GlfwEngine struct {
	graphicSettings *GraphicSettings
	
	resizeEvent chan ResizeEvent
}

func NewGlfwEngine() *GlfwEngine {
	return &GlfwEngine{resizeEvent:make(chan ResizeEvent, 0)}
}

func (engine *GlfwEngine) Init(settings *GraphicSettings) (err os.Error) {
	err = glfw.Init()
	if err != nil {
		return
	}
	windowType := glfw.Windowed
	if settings.Fullscreen {
		windowType = glfw.Fullscreen
	}
	err = glfw.OpenWindow(settings.Width, settings.Height, 0, 0, 0, 0, 0, 0, windowType)
	if err != nil {
		glfw.Terminate()
		return
	}
	glfw.SetSwapInterval(1) // VSync
	glfw.SetWindowTitel(settings.Caption)
	
	engine.graphicSettings = settings
}

func (engine *GlfwEngine) Close() {
	glfw.Terminate()
	glfw.CloseWindow()
}

func (engine *GlfwEngine) ResizeEvent() chan ResizeEvent {
	return engine.resizeEvent
}

func (engine *GlfwEngine) BeginFrame() {
	
}

func (engine *GlfwEngine) EndFrame() {
	glfw.SwapBuffers()
}