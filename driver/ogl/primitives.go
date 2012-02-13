package ogl

import (
	"image/color"

	gl "github.com/chsc/gogl/gl33"

	"github.com/Agon/baukasten/geometry"
)

const (
	PrimitiveCoordAttribLocationName = "coord"
	PrimitiveColorAttribLocationName = "v_color"
	PrimitiveMatrixLocationName      = "mvp"
	PrimitiveVertexShaderSource      = "#version 120\n" +
		"attribute vec2 " + PrimitiveCoordAttribLocationName + ";\n" +
		"attribute vec4 " + PrimitiveColorAttribLocationName + ";\n" +
		"varying vec4 f_color;\n" +
		"uniform mat4 " + PrimitiveMatrixLocationName + ";\n" +
		"void main(void) {\n  gl_Position = " + PrimitiveMatrixLocationName + " * vec4(" + PrimitiveCoordAttribLocationName + ", 0.0, 1.0);\n" +
		"f_color = " + PrimitiveColorAttribLocationName + ";\n}"
	PrimitiveFragmentShaderSource = "#version 120\n" +
		"varying vec4 f_color;\n" +
		"void main(void) {\n" +
		"  gl_FragColor = f_color;\n" +
		"}"
)

func (d *Driver) DrawPoints(color color.Color, vecs ...geometry.Vector2) {
	d.drawPrimitives(color, gl.POINTS, vecs...)
}

func (d *Driver) DrawLines(color color.Color, vecs ...geometry.Vector2) {
	d.drawPrimitives(color, gl.LINES, vecs...)
}

func (d *Driver) DrawLineStrip(color color.Color, vecs ...geometry.Vector2) {
	d.drawPrimitives(color, gl.LINE_STRIP, vecs...)
}

func (d *Driver) DrawLineLoop(color color.Color, vecs ...geometry.Vector2) {
	d.drawPrimitives(color, gl.LINE_LOOP, vecs...)
}

func (d *Driver) DrawTriangles(color color.Color, vecs ...geometry.Vector2) {
	d.drawPrimitives(color, gl.TRIANGLES, vecs...)
}

func (d *Driver) DrawTriangleStrip(color color.Color, vecs ...geometry.Vector2) {
	d.drawPrimitives(color, gl.TRIANGLE_STRIP, vecs...)
}

func (d *Driver) DrawTriangleFan(color color.Color, vecs ...geometry.Vector2) {
	d.drawPrimitives(color, gl.TRIANGLE_FAN, vecs...)
}

func (d *Driver) drawPrimitives(color color.Color, mode gl.Enum, vecs ...geometry.Vector2) {
	vertices := make([]float32, len(vecs)*2)
	r, g, b, a := geometry.ConvertColorF(color)
	var colors []float32
	for i := range vecs {
		vertices[i*2] = vecs[i].X
		vertices[i*2+1] = vecs[i].Y
		colors = append(colors, r, g, b, a)
	}
	d.primitivesProgram.Use()
	d.primitivesUniformMatrix.UniformMatrix4fv(1, false, d.Camera().Get().Transposed())
	d.primitivesAttributeCoord.Enable()
	d.primitivesAttributeCoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices[0]))

	d.primitivesAttributeColor.Enable()
	d.primitivesAttributeColor.AttribPointer(4, gl.FLOAT, false, 0, gl.Pointer(&colors[0]))

	gl.DrawArrays(mode, 0, gl.Sizei(len(vecs)))

	d.primitivesAttributeColor.Disable()
	d.primitivesAttributeCoord.Disable()
}
