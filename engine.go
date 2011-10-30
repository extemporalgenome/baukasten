package baukasten

import (
	"image"
	"os"
)

type Engine struct {
	graphic GraphicDriver
	context ContextDriver
	input   InputDriver

	settings *GraphicSettings
}

func NewEngine(graphic GraphicDriver, context ContextDriver, input InputDriver) *Engine {
	return &Engine{graphic: graphic, context: context, input: input}
}

func (e *Engine) Init(settings *GraphicSettings) (err os.Error) {
	err = e.context.Init(settings)
	if err != nil {
		return err
	}
	err = e.graphic.Init(settings)
	if err != nil {
		return err
	}
	e.settings = settings
	return nil
}

func (e *Engine) Close() {
	e.graphic.Close()
	e.context.Close()
}

func (e *Engine) BeginFrame() {
	e.graphic.BeginFrame()
}

func (e *Engine) EndFrame() {
	e.graphic.EndFrame()
	e.context.SwapBuffers()
}

func (e *Engine) OpenSurface(name string) (Surface, os.Error) {
	return e.context.OpenSurface(name)
}

func (e *Engine) LoadSurface(image image.Image) (Surface, os.Error) {
	return e.context.LoadSurface(image)
}
