// Baukasten is a game library with focuse on easy game development.
package baukasten

import (
	"errors"
	"image"
	"image/color"
	"time"
)

var NoContextDriverError = errors.New("baukasten.Engine has no loaded ContextDriver.")
var NoGraphicDriverError = errors.New("baukasten.Engine has no loaded GraphicDriver.")
var NoInputDriverError = errors.New("baukasten.Engine has no loaded InputDriver.")
var NoFontDriverError = errors.New("baukasten.Engine has no loaded InputDriver.")

type Engine struct {
	// Drivers
	graphic GraphicDriver
	context ContextDriver
	input   InputDriver
	font    FontDriver

	settings *GraphicSettings

	currentTime time.Time
	lastTime    time.Time
}

func NewEngine(graphic GraphicDriver, context ContextDriver, input InputDriver, font FontDriver) *Engine {
	return &Engine{graphic: graphic, context: context, input: input, font: font}
}

// Initializes the engine and it's drivers.
func (e *Engine) Init(settings *GraphicSettings) (err error) {
	if e.context == nil {
		return NoContextDriverError
	}
	err = e.context.Init(settings)
	if err != nil {
		return err
	}
	if e.input == nil {
		return NoGraphicDriverError
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
	if e.input == nil {
		panic(NoInputDriverError)
	}
	return e.input.KeyEvent()
}

func (e *Engine) MouseButtonEvent() chan MouseButtonEvent {
	if e.input == nil {
		panic(NoInputDriverError)
	}
	return e.input.MouseButtonEvent()
}

func (e *Engine) MousePositionEvent() chan MousePositionEvent {
	if e.input == nil {
		panic(NoInputDriverError)
	}
	return e.input.MousePositionEvent()
}

func (e *Engine) MouseWheelEvent() chan MouseWheelEvent {
	if e.input == nil {
		panic(NoInputDriverError)
	}
	return e.input.MouseWheelEvent()
}

func (e *Engine) DrawPoints(color color.Color, vecs ...Vector2) {
	e.graphic.DrawPoints(color, vecs...)
}

func (e *Engine) DrawLines(color color.Color, vecs ...Vector2) {
	if len(vecs) < 2 {
		panic("Not enough vectors specified.")
	}
	if len(vecs)%2 != 0 {
		panic("Length of vecs is not a power of 2")
	}
	e.graphic.DrawLines(color, vecs...)
}

func (e *Engine) DrawLineStrip(color color.Color, vecs ...Vector2) {
	if len(vecs) < 2 {
		panic("Not enough vectors specified.")
	}
	e.graphic.DrawLineStrip(color, vecs...)
}

func (e *Engine) DrawLineLoop(color color.Color, vecs ...Vector2) {
	if len(vecs) < 2 {
		panic("Not enough vectors specified.")
	}
	e.graphic.DrawLineLoop(color, vecs...)
}

func (e *Engine) DrawTriangle(color color.Color, vec1, vec2, vec3 Vector2) {
	e.graphic.DrawTriangle(color, vec1, vec2, vec3)
}

func (e *Engine) OpenSurface(name string) (Surface, error) {
	return e.graphic.OpenSurface(name)
}

func (e *Engine) LoadSurface(image image.Image) (Surface, error) {
	return e.graphic.LoadSurface(image)
}

func (e *Engine) OpenFont(name string) (Font, error) {
	if e.font == nil {
		return nil, NoFontDriverError
	}
	return e.font.OpenFont(name)
}

func (e *Engine) RenderSurface(text string, width, height int, size float64, color color.Color, font Font) (Surface, error) {
	if e.font == nil {
		return nil, NoFontDriverError
	}
	img := font.Render(text, width, height, size, color)
	return e.LoadSurface(img)
}
