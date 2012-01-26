package ogl

import (
	"image/color"

	"github.com/Agon/baukasten"

	gl "github.com/chsc/gogl/gl33"
)

const (
	PrimitiveCoordAttribLocationName = "coord"
	PrimitiveColorAttribLocationName = "v_color"
	PrimitiveVertexShaderSource      = "#version 120\n" +
		"attribute vec2 " + PrimitiveCoordAttribLocationName + ";\n" +
		"attribute vec4 " + PrimitiveColorAttribLocationName + ";\n" +
		"attribute vec4 f_color;\n" +
		"void main(void) {\n  gl_Position = vec4(" + PrimitiveCoordAttribLocationName + ", 0.0, 1.0);" +
		"f_color = " + PrimitiveColorAttribLocationName + ";\n}"
	PrimitiveFragmentShaderSource = "#version 120\n" +
		"varying vec4 f_color;\n" +
		"void main(void) {\n" +
		"  gl_FragColor = f_color;\n" +
		"}"
)

func (d *Driver) DrawTriangle(vec1, vec2, vec3 baukasten.Vector2, color color.Color) {
	vertices := []float32{vec1.X, vec1.Y, vec2.X, vec2.Y, vec3.X, vec3.Y}
	r, g, b, a := baukasten.ConvertColorF(color)
	colors := []float32{r, g, b, a}
	colors = append(colors, append(colors, colors...)...) // colors += colors + colors

	d.primitivesProgram.Use()
	d.primitivesAttributeCoord.Enable()
	d.primitivesAttributeCoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices))

	d.primitivesAttributeColor.Enable()
	d.primitivesAttributeColor.AttribPointer(4, gl.FLOAT, false, 0, gl.Pointer(&colors))

	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	d.primitivesAttributeCoord.Disable()
	d.primitivesAttributeColor.Disable()

}
