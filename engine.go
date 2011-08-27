package baukasten

import (
	"fmt"
	"os"

	"github.com/banthar/gl"
	"sdl"
)

type Engine struct {
	graphicSettings *GraphicSettings
	screen          *sdl.Surface
	ticks           uint32

	onKeyboardEvent chan KeyboardEvent
	onQuitEvent     chan QuitEvent
	onResizeEvent   chan ResizeEvent
}

func NewEngine() *Engine {
	return &Engine{
		onKeyboardEvent: make(chan KeyboardEvent),
		onResizeEvent:   make(chan ResizeEvent),
		onQuitEvent:     make(chan QuitEvent)}
}

// WARNING contains deprecated code as of OpenGL 3
func (e *Engine) Init(graphicSettings *GraphicSettings) os.Error {
	e.graphicSettings = graphicSettings
	// SDL
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		return fmt.Errorf("SDL initialize error: %s", sdl.GetError())
	}
	err := e.Resize(e.graphicSettings.Width, e.graphicSettings.Height)
	if err != nil {
		return err
	}
	sdl.WM_SetCaption(e.graphicSettings.Caption, e.graphicSettings.Caption)
	// OpenGL
	if gl.Init() != 0 {
		return os.NewError("OpenGL initialization error.")
	}
	gl.Enable(gl.DEPTH_TEST)
	return nil
}

func (e *Engine) Quit() os.Error {
	sdl.Quit()
	return nil
}

func (e *Engine) PollEvent() {
	ev := sdl.PollEvent()
	switch event := ev.(type) {
	case *sdl.ResizeEvent:
		e.onResizeEvent <- ResizeEvent{int(event.W), int(event.H)}
	case *sdl.QuitEvent:
		e.onQuitEvent <- QuitEvent{int(event.Type)}
	case *sdl.KeyboardEvent:
		e.onKeyboardEvent <- KeyboardEvent{
			Key:   uint(event.Keysym.Sym),
			State: uint(event.State),
			Type:  uint(event.Type)}
	}
}

func (e *Engine) OnKeyboardEvent() chan KeyboardEvent {
	return e.onKeyboardEvent
}

func (e *Engine) OnQuitEvent() chan QuitEvent {
	return e.onQuitEvent
}

func (e *Engine) OnResizeEvent() chan ResizeEvent {
	return e.onResizeEvent
}

// WARNING contains deprecated code as of OpenGL 3
func (e *Engine) GetOpenGLVersion() string {
	return gl.GetString(gl.VERSION)
}

// WARNING contains deprecated code as of OpenGL 3
func (e *Engine) Resize(width, height int) os.Error {
	if height == 0 {
		height = 1
	}
	// SDL
	settings := uint32(sdl.OPENGL)
	if e.graphicSettings.Resizeable {
		settings |= sdl.RESIZABLE
	}
	e.screen = sdl.SetVideoMode(width, height, e.graphicSettings.BitDepth, settings)
	if e.screen == nil {
		return fmt.Errorf("SDL video mode set error on resize: %s", sdl.GetError())
	}
	// OpenGL	
	// Setup our viewport
	gl.Viewport(0, 0, width, height)

	// TODO proper Camera and view handling

	// Reset the view
	gl.LoadIdentity()
	return nil
}

// Delta time in seconds.
func (e *Engine) DeltaTime() float32 {
	t := sdl.GetTicks()
	delta := t - e.ticks
	e.ticks = t
	return float32(delta) / 1000.0
}

// WARNING contains deprecated code as of OpenGL 3
func (e *Engine) BeginFrame() {
	e.Clear()
	gl.LoadIdentity()
}

func (e *Engine) EndFrame() {
	sdl.GL_SwapBuffers()
}

func (e *Engine) Clear() {
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// WARNING contains deprecated code as of OpenGL 3
func (e *Engine) StartList() uint {
	list := gl.GenLists(1)
	gl.NewList(list, gl.COMPILE)
	return list
}

// WARNING contains deprecated code as of OpenGL 3
func (e *Engine) EndList() {
	gl.EndList()
}

// WARNING contains deprecated code as of OpenGL 3
func (e *Engine) DrawList(list uint) {
	gl.CallList(list)
}

func (e *Engine) PushMatrix() {
	gl.PushMatrix()
}

func (e *Engine) PopMatrix() {
	gl.PopMatrix()
}

func (e *Engine) Translate(vec Vector3) {
	gl.Translatef(vec.X, vec.Y, vec.Z)
}

func (e *Engine) Rotate(vec Vector3, amount float32) {
	gl.Rotatef(vec.X, vec.Y, vec.Z, amount)
}

// WARNING contains deprecated code as of OpenGL 3
func (e *Engine) DrawPolygon2(vertices []Vector2) {
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	gl.Begin(gl.TRIANGLES)
	for _, vector := range vertices {
		gl.Vertex2f(vector.X, vector.Y)
	}
	gl.End()
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
}

// WARNING contains deprecated code as of OpenGL 3
func (e *Engine) FillPolygon2(vertices []Vector2) {
	gl.Begin(gl.TRIANGLES)
	for _, vector := range vertices {
		gl.Vertex2f(vector.X, vector.Y)
	}
	gl.End()
}
