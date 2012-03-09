// Baukasten is a game library with focuse on easy game development.
package baukasten

import (
	"container/list"
	"errors"
	"image"
	"image/color"
	"time"

	math "github.com/Agon/baukasten/geometry"
)

var NoContextDriverError = errors.New("baukasten.Engine has no loaded ContextDriver.")
var NoGraphicDriverError = errors.New("baukasten.Engine has no loaded GraphicDriver.")
var NoInputDriverError = errors.New("baukasten.Engine has no loaded InputDriver.")
var NoFontDriverError = errors.New("baukasten.Engine has no loaded InputDriver.")

// Engine handles loading, unloading of drivers and is able to call general functions to the drivers.
type Engine struct {
	// Drivers
	graphic GraphicDriver
	context ContextDriver
	input   InputDriver
	font    FontDriver

	settings *GraphicSettings

	currentTime time.Time
	lastTime    time.Time

	contextEvent chan ContextEvent
	resizeEvent  chan WindowSize

	keyEvent           chan Key
	mouseButtonEvent   chan MouseButton
	mousePositionEvent chan MousePosition
	mouseWheelEvent    chan MouseWheel
}

func NewEngine(graphic GraphicDriver, context ContextDriver, input InputDriver, font FontDriver) *Engine {
	return &Engine{graphic: graphic, context: context, input: input, font: font, lastTime: time.Now()}
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
	if e.graphic == nil {
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

// Returns the duration (delta time) between the last call to DeltaTime and now.
// If DeltaTime is called for the first time the duration between the Init call and now is returned.
func (e *Engine) DeltaTime() time.Duration {
	e.currentTime = time.Now()
	duration := e.currentTime.Sub(e.lastTime)
	e.lastTime = e.currentTime
	return duration
}

func (e *Engine) SetClearColor(c color.Color) {
	e.graphic.SetClearColor(c)
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

func (e *Engine) Camera() Camera {
	return e.graphic.Camera()
}

func (e *Engine) SetCamera(c Camera) {
	e.graphic.SetCamera(c)
}

// Resizes the graphic screen.
func (e *Engine) GraphicResize(w, h int) {
	e.graphic.Resize(w, h)
}

func (e *Engine) MousePos() MousePosition {
	return e.input.MousePos()
}

func (e *Engine) JoystickButtons(joy int) []bool {
	return e.input.JoystickButtons(joy)
}

func (e *Engine) JoystickPos(joy int) []math.Vector2 {
	return e.input.JoystickPos(joy)
}

func (e *Engine) ResizeEvent() <-chan WindowSize {
	if e.resizeEvent == nil {
		event := make(chan WindowSize, 0)
		in := make(chan WindowSize, 0)
		e.resizeEvent = event
		go func() {
			l := list.New()
			for {
				if l.Len() > 0 { // Can receive->saved/saved->send
					element := l.Front()
					value := element.Value.(WindowSize)
					select {
					case e.resizeEvent <- value:
						l.Remove(element)
					case v := <-in:
						l.PushBack(v)
					}
				} else { // Can send through / receive->saved
					v := <-in
					select {
					case e.resizeEvent <- v:
					default:
						l.PushBack(v)
					}
				}
			}
		}()
		e.context.SetResizeCallback(in)
	}
	return e.resizeEvent
}

func (e *Engine) ContextEvent() <-chan ContextEvent {
	if e.contextEvent == nil {
		event := make(chan ContextEvent, 0)
		in := make(chan ContextEvent, 0)
		e.contextEvent = event
		go func() {
			l := list.New()
			for {
				if l.Len() > 0 { // Can receive->saved/saved->send
					element := l.Front()
					value := element.Value.(ContextEvent)
					select {
					case e.contextEvent <- value:
						l.Remove(element)
					case v := <-in:
						l.PushBack(v)
					}
				} else { // Can send through / receive->saved
					v := <-in
					select {
					case e.contextEvent <- v:
					default:
						l.PushBack(v)
					}
				}
			}
		}()
		e.context.SetContextCallback(in)
	}
	return e.contextEvent
}

func (e *Engine) KeyEvent() <-chan Key {
	if e.keyEvent == nil {
		event := make(chan Key, 0)
		in := make(chan Key, 0)
		e.keyEvent = event
		go func() {
			l := list.New()
			for {
				if l.Len() > 0 { // Can receive->saved/saved->send
					element := l.Front()
					value := element.Value.(Key)
					select {
					case e.keyEvent <- value:
						l.Remove(element)
					case v := <-in:
						l.PushBack(v)
					}
				} else { // Can send through / receive->saved
					v := <-in
					select {
					case e.keyEvent <- v:
					default:
						l.PushBack(v)
					}
				}
			}
		}()
		e.input.SetKeyCallback(in)
	}
	return e.keyEvent
}

func (e *Engine) MouseButtonEvent() <-chan MouseButton {
	if e.mouseButtonEvent == nil {
		event := make(chan MouseButton, 0)
		in := make(chan MouseButton, 0)
		e.mouseButtonEvent = event
		go func() {
			l := list.New()
			for {
				if l.Len() > 0 { // Can receive->saved/saved->send
					element := l.Front()
					value := element.Value.(MouseButton)
					select {
					case e.mouseButtonEvent <- value:
						l.Remove(element)
					case v := <-in:
						l.PushBack(v)
					}
				} else { // Can send through / receive->saved
					v := <-in
					select {
					case e.mouseButtonEvent <- v:
					default:
						l.PushBack(v)
					}
				}
			}
		}()
		e.input.SetMouseButtonCallback(in)
	}
	return e.mouseButtonEvent
}

func (e *Engine) MousePositionEvent() <-chan MousePosition {
	if e.mousePositionEvent == nil {
		event := make(chan MousePosition, 0)
		in := make(chan MousePosition, 0)
		e.mousePositionEvent = event
		go func() {
			l := list.New()
			for {
				if l.Len() > 0 { // Can receive->saved/saved->send
					element := l.Front()
					value := element.Value.(MousePosition)
					select {
					case e.mousePositionEvent <- value:
						l.Remove(element)
					case v := <-in:
						l.PushBack(v)
					}
				} else { // Can send through / receive->saved
					v := <-in
					select {
					case e.mousePositionEvent <- v:
					default:
						l.PushBack(v)
					}
				}
			}
		}()
		e.input.SetMousePositionCallback(in)
	}
	return e.mousePositionEvent
}

func (e *Engine) MouseWheelEvent() <-chan MouseWheel {
	if e.mouseWheelEvent == nil {
		event := make(chan MouseWheel, 0)
		in := make(chan MouseWheel, 0)
		e.mouseWheelEvent = event
		go func() {
			l := list.New()
			for {
				if l.Len() > 0 { // Can receive->saved/saved->send
					element := l.Front()
					value := element.Value.(MouseWheel)
					select {
					case e.mouseWheelEvent <- value:
						l.Remove(element)
					case v := <-in:
						l.PushBack(v)
					}
				} else { // Can send through / receive->saved
					v := <-in
					select {
					case e.mouseWheelEvent <- v:
					default:
						l.PushBack(v)
					}
				}
			}
		}()
		e.input.SetMouseWheelCallback(in)
	}
	return e.mouseWheelEvent
}

// DrawPoints draws each vector as a single point.
func (e *Engine) DrawPoints(color color.Color, vecs ...math.Vector2) {
	e.graphic.DrawPoints(color, vecs...)
}

// DrawLines draws each pair of vectors as an independent line segment.
// The length of vecs needs to be a power of 2.
func (e *Engine) DrawLines(color color.Color, vecs ...math.Vector2) {
	if len(vecs) < 2 {
		panic("Not enough vectors specified.")
	}
	if len(vecs)%2 != 0 {
		panic("Length of vecs is not a power of 2")
	}
	e.graphic.DrawLines(color, vecs...)
}

// DrawLineStrip draws a connected group of line segments from the first vector to the last.
// The length of vecs needs to be greater than one.
func (e *Engine) DrawLineStrip(color color.Color, vecs ...math.Vector2) {
	if len(vecs) < 2 {
		panic("Not enough vectors specified.")
	}
	e.graphic.DrawLineStrip(color, vecs...)
}

// DrawLineLoop draws a connected group of line segments from the first vector to the last, then back to the frist.
// The length of vecs needs to be greater than one.
func (e *Engine) DrawLineLoop(color color.Color, vecs ...math.Vector2) {
	if len(vecs) < 2 {
		panic("Not enough vectors specified.")
	}
	e.graphic.DrawLineLoop(color, vecs...)
}

// DrawTriangles draws three vectors as an independent triangle.
// The length of vecs needs to be a power of 3.
func (e *Engine) DrawTriangles(color color.Color, vecs ...math.Vector2) {
	if len(vecs)%3 != 0 {
		panic("Length of vecs is not a power of 3.")
	}
	e.graphic.DrawTriangles(color, vecs...)
}

// DrawTriangleStrip draws a connected group of triangles.
// The length of vecs needs to be greater than two.
func (e *Engine) DrawTriangleStrip(color color.Color, vecs ...math.Vector2) {
	if len(vecs) < 3 {
		panic("Not enough vectors specified.")
	}
	e.graphic.DrawTriangleStrip(color, vecs...)
}

// DrawTriangleFan draws a connected group of triangles, centering around the second vector.
// The length of vecs needs to be greater than two.
func (e *Engine) DrawTriangleFan(color color.Color, vecs ...math.Vector2) {
	if len(vecs) < 3 {
		panic("Not enough vectors specified.")
	}
	e.graphic.DrawTriangleFan(color, vecs...)
}

// DrawRectangle draws a Rectanglef as two triangles.
func (e *Engine) DrawRectangle(color color.Color, r math.Rectanglef) {
	e.graphic.DrawTriangles(color, r.Min, math.Vector2{r.Min.X, r.Max.Y}, math.Vector2{r.Max.X, r.Min.Y}, math.Vector2{r.Max.X, r.Min.Y}, math.Vector2{r.Min.X, r.Max.Y}, r.Max)
}

// DrawCircle draws a circle centered at v with a radius of r, with n number of points in color c.
func (e *Engine) DrawCircle(c color.Color, r float32, n int, v math.Vector2) {
	vectors := make([]math.Vector2, n)
	for i := 0; i < n; i++ {
		degInRad := (360 / float32(i)) * math.Pi / 180
		vectors[i] = v.Add(math.Vector2{math.Cos(degInRad) * r, math.Sin(degInRad) * r})
	}
	e.DrawLineLoop(c, vectors...)
}

// OpenSurface loads and decodes an image, then creates a Surface of it.
// Following image formats are supported: bmp, gif, jpeg, png, tiff
func (e *Engine) OpenSurface(name string) (Surface, error) {
	return e.graphic.OpenSurface(name)
}

// LoadSurface loads a Surface from a type which implements Go image.Image.
func (e *Engine) LoadSurface(image image.Image) (Surface, error) {
	return e.graphic.LoadSurface(image)
}

// OpenFont loads a font file and creates a Font.
func (e *Engine) OpenFont(name string) (Font, error) {
	if e.font == nil {
		return nil, NoFontDriverError
	}
	return e.font.OpenFont(name)
}

// RenderText renders a text to an image and creates a Surface of it.
func (e *Engine) RenderText(text string, width, height int, size float64, color color.Color, font Font) (Surface, error) {
	if e.font == nil {
		return nil, NoFontDriverError
	}
	img := font.Render(text, width, height, size, color)
	return e.LoadSurface(img)
}
