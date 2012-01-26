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

var DefaultDriver = NewDriver()

type Driver struct {
	primitivesProgram        *Program
	primitivesAttributeCoord *AttributeLocation
}

func NewDriver() *Driver {
	return &Driver{}
}

func (d *Driver) Init(graphicSettings *baukasten.GraphicSettings) error {
	var err error
	err = gl.Init()
	if err != nil {
		return errors.New(fmt.Sprintf("Init OpenGL extension loading failed with %s.\n", err))
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// Load primitives shaders
	primitivesVertexShader, err := LoadShader(PrimitiveVertexShaderSource, VertexShaderType)
	if err != nil {
		return err
	}
	primitivesFragmentShader, err := LoadShader(PrimitiveFragmentShaderSource, FragmentShaderType)
	if err != nil {
		return err
	}
	d.primitivesProgram = NewProgram()
	d.primitivesProgram.AttachShaders(primitivesVertexShader, primitivesFragmentShader)
	err = d.primitivesProgram.Link()
	if err != nil {
		return err
	}
	d.primitivesAttributeCoord, err = d.primitivesProgram.GetAttributeLocation(PrimitiveAttributeLocationName)
	if err != nil {
		return err
	}
	return nil
}

func (d *Driver) Close() {}

func (d *Driver) BeginFrame() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (d *Driver) SetClearColor(color color.Color) {
	r, g, b, a := baukasten.ConvertColorF(color)
	gl.ClearColor(gl.Clampf(r), gl.Clampf(g), gl.Clampf(b), gl.Clampf(a))
}

func (d *Driver) EndFrame() {

}

func (d *Driver) Resize(w, h int) {
	ScreenWidth = w
	ScreenHeight = h
	gl.Viewport(0, 0, gl.Sizei(w), gl.Sizei(h))
}
