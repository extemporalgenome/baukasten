package baukasten

import (
	"fmt"
	"math"
	"os"

	"github.com/banthar/gl"
	"sdl"
)

type Engine struct {
	graphicSettings *GraphicSettings
	screen          *sdl.Surface
}

func NewEngine() *Engine {
	return &Engine{}
}

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

	// change to the projection matrix and set viewing volume
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	// aspect ratio
	aspect := float64(width) / float64(height)

	// Set perspective
	var fov, near, far float64
	fov = 45.0
	near = 0.1
	far = 100.0
	top := math.Tan(float64(fov*math.Pi/360.0)) * near
	bottom := -top
	left := aspect * bottom
	right := aspect * top
	gl.Frustum(left, right, bottom, top, near, far)

	// Make sure we're changing the model view and not the projection
	gl.MatrixMode(gl.MODELVIEW)

	// Reset the view
	gl.LoadIdentity()
	return nil
}

func (e *Engine) BeginFrame() {
	e.Clear()
	gl.LoadIdentity()
}

func (e *Engine) EndFrame() {
	sdl.GL_SwapBuffers()
	// TODO Frames
}

func (e *Engine) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
}
