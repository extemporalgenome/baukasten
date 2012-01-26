// Baukasten is a game library with focuse on easy game development.
package baukasten

import (
	"image"
	"image/color"
	"time"
)

type Engine struct {
	// Drivers
	graphic GraphicDriver
	context ContextDriver
	input   InputDriver

	settings *GraphicSettings

	currentTime time.Time
	lastTime    time.Time
}

func NewEngine(graphic GraphicDriver, context ContextDriver, input InputDriver) *Engine {
	return &Engine{graphic: graphic, context: context, input: input}
}

// Initializes the engine and it's drivers.
func (e *Engine) Init(settings *GraphicSettings) (err error) {
	err = e.context.Init(settings)
	if err != nil {
		return err
	}
	err = e.graphic.Init(settings)
	if err != nil {
		return err
	}
	e.settings = settings
	e.currentTime = time.Now()
	return nil
}

// Shuts down the engine and it's drivers.
func (e *Engine) Close() {
	e.graphic.Close()
	e.context.Close()
}

// Returns the duration (delta time) between the last call to DeltaTime and now in seconds.
// If DeltaTime is called for the first time the duration between the Init call and now is returned.
func (e *Engine) DeltaTime() float32 {
	e.currentTime = time.Now()
	duration := e.currentTime.Sub(e.lastTime)
	e.lastTime = e.currentTime
	return float32(duration.Seconds())
}

// This should be called before each frame is rendered.
func (e *Engine) BeginFrame() {
	e.graphic.BeginFrame()
}

// This ends the current rendering to a frame.
func (e *Engine) EndFrame() {
	e.graphic.EndFrame()
	e.context.SwapBuffers()
}

// Resizes the graphic screen.
func (e *Engine) GraphicResize(w, h int) {
	e.graphic.Resize(w, h)
}

func (e *Engine) ResizeEvent() chan WindowSizeEvent {
	return e.context.ResizeEvent()
}

func (e *Engine) ContextEvent() chan ContextEvent {
	return e.context.ContextEvent()
}

func (e *Engine) KeyEvent() chan KeyEvent {
	return e.input.KeyEvent()
}

func (e *Engine) MouseButtonEvent() chan MouseButtonEvent {
	return e.input.MouseButtonEvent()
}

func (e *Engine) MousePositionEvent() chan MousePositionEvent {
	return e.input.MousePositionEvent()
}

func (e *Engine) MouseWheelEvent() chan MouseWheelEvent {
	return e.input.MouseWheelEvent()
}

func (e *Engine) OpenSurface(name string) (Surface, error) {
	return e.graphic.OpenSurface(name)
}

func (e *Engine) LoadSurface(image image.Image) (Surface, error) {
	return e.graphic.LoadSurface(image)
}

func (e *Engine) OpenFont(name string) (Font, error) {
	return OpenFont(name)
}

func (e *Engine) RenderSurface(text string, width, height int, size float64, color color.Color, font Font) (Surface, error) {
	img := font.Render(text, width, height, size, color)
	return e.LoadSurface(img)
}
