// OpenGL 3.3 GraphicDriver package for baukasten.
package ogl

import (
	"fmt"
	"image/color"

	gl "github.com/chsc/gogl/gl33"

	"github.com/Agon/baukasten"
	math "github.com/Agon/baukasten/geometry"
)

const (
	VertexShaderType = iota
	FragmentShaderType
)

var DefaultDriver = NewDriver()

type Driver struct {
	primitivesProgram        *Program
	primitivesAttributeCoord *AttributeLocation
	primitivesAttributeColor *AttributeLocation
	primitivesUniformMatrix  *UniformLocation
	camera                   baukasten.Camera
}

func NewDriver() *Driver {
	return &Driver{}
}

func (d *Driver) Init(graphicSettings *baukasten.GraphicSettings) error {
	var err error
	err = gl.Init()
	if err != nil {
		return fmt.Errorf("Init OpenGL extension loading failed with %s.\n", err)
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// Load primitives shaders
	primitivesVertexShader, err := LoadShader(PrimitiveVertexShaderSource, VertexShaderType)
	if err != nil {
		return fmt.Errorf("Init OpenGL extension loading failed at loading PrimitivesVertexShader with %s.\n", err)
	}
	err = primitivesVertexShader.Compile()
	if err != nil {
		return fmt.Errorf("Init OpenGL extension loading failed at compiling PrimitivesVertexShader with %s.\n", err)
	}
	primitivesFragmentShader, err := LoadShader(PrimitiveFragmentShaderSource, FragmentShaderType)
	if err != nil {
		return fmt.Errorf("Init OpenGL extension loading failed at loading PrimitivesFragmentShader with %s.\n", err)
	}
	err = primitivesFragmentShader.Compile()
	if err != nil {
		return fmt.Errorf("Init OpenGL extension loading failed at compiling FragmentShaderType with %s.\n", err)
	}
	d.primitivesProgram = NewProgram()
	d.primitivesProgram.AttachShaders(primitivesVertexShader, primitivesFragmentShader)
	err = d.primitivesProgram.Link()
	if err != nil {
		return fmt.Errorf("Init OpenGL extension loading failed at linking PrimitivesProgram with %s.\n", err)
	}
	d.primitivesAttributeCoord, err = d.primitivesProgram.GetAttributeLocation(PrimitiveCoordAttribLocationName)
	if err != nil {
		return fmt.Errorf("Init OpenGL extension loading failed at getting PrimitiveCoordAttribLocation with %s.\n", err)
	}
	d.primitivesAttributeColor, err = d.primitivesProgram.GetAttributeLocation(PrimitiveColorAttribLocationName)
	if err != nil {
		return fmt.Errorf("Init OpenGL extension loading failed at getting PrimitiveColorAttribLocation with %s.\n", err)
	}
	d.primitivesUniformMatrix, err = d.primitivesProgram.GetUniformLocation(PrimitiveMatrixLocationName)
	if err != nil {
		return fmt.Errorf("Init OpenGL extension loading failed at getting PrimitiveMatrixLocationName with %s.\n", err)
	}
	return nil
}

func (d *Driver) Close() {}

func (d *Driver) BeginFrame() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (d *Driver) SetClearColor(color color.Color) {
	r, g, b, a := math.ConvertColorF(color)
	gl.ClearColor(gl.Clampf(r), gl.Clampf(g), gl.Clampf(b), gl.Clampf(a))
}

func (d *Driver) Camera() baukasten.Camera {
	return d.camera
}

func (d *Driver) SetCamera(c baukasten.Camera) {
	d.camera = c
}

func (d *Driver) EndFrame() {

}

func (d *Driver) Resize(w, h int) {
	ScreenWidth = w
	ScreenHeight = h
	gl.Viewport(0, 0, gl.Sizei(w), gl.Sizei(h))
}
