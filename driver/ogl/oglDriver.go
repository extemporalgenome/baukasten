package ogl

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/Agon/baukasten"
	gl "github.com/chsc/gogl/gl33"
)

const (
	VertexShaderType = iota
	FragmentShaderType
)

var DefaultDriver = &OpenGLDriver{}

type OpenGLDriver struct{}

func (driver *OpenGLDriver) Init(graphicSettings *baukasten.GraphicSettings) (err error) {

	err = gl.Init()
	if err != nil {
		err = errors.New(fmt.Sprintf("Init OpenGL extension loading failed with %s.\n", err))
		return
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	return
}

func (driver *OpenGLDriver) Close() {

}

func (driver *OpenGLDriver) BeginFrame() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (driver *OpenGLDriver) SetClearColor(color color.Color) {
	r, g, b, a := baukasten.ConvertColorF(color)
	gl.ClearColor(gl.Clampf(r), gl.Clampf(g), gl.Clampf(b), gl.Clampf(a))
}

func (driver *OpenGLDriver) EndFrame() {

}

func (driver *OpenGLDriver) Resize(w, h int) {
	ScreenWidth = w
	ScreenHeight = h
	gl.Viewport(0, 0, gl.Sizei(w), gl.Sizei(h))
}

func Offset(p gl.Pointer, o uint) gl.Pointer {
	return gl.Pointer(uintptr(p) + uintptr(o))
}
