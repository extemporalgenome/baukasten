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
		"varying vec4 f_color;\n" +
		"void main(void) {\n  gl_Position = vec4(" + PrimitiveCoordAttribLocationName + ", 0.0, 1.0);\n" +
		"f_color = " + PrimitiveColorAttribLocationName + ";\n}"
	PrimitiveFragmentShaderSource = "#version 120\n" +
		"varying vec4 f_color;\n" +
		"void main(void) {\n" +
		"  gl_FragColor = f_color;\n" +
		"}"
)

func (d *Driver) DrawPoints(color color.Color, vecs ...baukasten.Vector2) {
	vertices := make([]float32, len(vecs)*2)
	r, g, b, a := baukasten.ConvertColorF(color)
	var colors []float32
	for i := range vecs {
		vertices[i*2] = vecs[i].X
		vertices[i*2+1] = vecs[i].Y
		colors = append(colors, r, g, b, a)
	}
	d.primitivesProgram.Use()
	d.primitivesAttributeCoord.Enable()
	d.primitivesAttributeCoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices[0]))

	d.primitivesAttributeColor.Enable()
	d.primitivesAttributeColor.AttribPointer(4, gl.FLOAT, false, 0, gl.Pointer(&colors[0]))

	gl.DrawArrays(gl.POINTS, 0, gl.Sizei(len(vecs)))

	d.primitivesAttributeColor.Disable()
	d.primitivesAttributeCoord.Disable()
}

func (d *Driver) DrawLines(color color.Color, vecs ...baukasten.Vector2) {
	vertices := make([]float32, len(vecs)*2)
	r, g, b, a := baukasten.ConvertColorF(color)
	var colors []float32
	for i := range vecs {
		vertices[i*2] = vecs[i].X
		vertices[i*2+1] = vecs[i].Y
		colors = append(colors, r, g, b, a)
	}
	d.primitivesProgram.Use()
	d.primitivesAttributeCoord.Enable()
	d.primitivesAttributeCoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices[0]))

	d.primitivesAttributeColor.Enable()
	d.primitivesAttributeColor.AttribPointer(4, gl.FLOAT, false, 0, gl.Pointer(&colors[0]))

	gl.DrawArrays(gl.LINES, 0, gl.Sizei(len(vecs)))

	d.primitivesAttributeColor.Disable()
	d.primitivesAttributeCoord.Disable()
}

func (d *Driver) DrawLineStrip(color color.Color, vecs ...baukasten.Vector2) {
	vertices := make([]float32, len(vecs)*2)
	r, g, b, a := baukasten.ConvertColorF(color)
	var colors []float32
	for i := range vecs {
		vertices[i*2] = vecs[i].X
		vertices[i*2+1] = vecs[i].Y
		colors = append(colors, r, g, b, a)
	}
	d.primitivesProgram.Use()
	d.primitivesAttributeCoord.Enable()
	d.primitivesAttributeCoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices[0]))

	d.primitivesAttributeColor.Enable()
	d.primitivesAttributeColor.AttribPointer(4, gl.FLOAT, false, 0, gl.Pointer(&colors[0]))

	gl.DrawArrays(gl.LINE_STRIP, 0, gl.Sizei(len(vecs)))

	d.primitivesAttributeColor.Disable()
	d.primitivesAttributeCoord.Disable()
}

func (d *Driver) DrawLineLoop(color color.Color, vecs ...baukasten.Vector2) {
	vertices := make([]float32, len(vecs)*2)
	r, g, b, a := baukasten.ConvertColorF(color)
	var colors []float32
	for i := range vecs {
		vertices[i*2] = vecs[i].X
		vertices[i*2+1] = vecs[i].Y
		colors = append(colors, r, g, b, a)
	}
	d.primitivesProgram.Use()
	d.primitivesAttributeCoord.Enable()
	d.primitivesAttributeCoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices[0]))

	d.primitivesAttributeColor.Enable()
	d.primitivesAttributeColor.AttribPointer(4, gl.FLOAT, false, 0, gl.Pointer(&colors[0]))

	gl.DrawArrays(gl.LINE_LOOP, 0, gl.Sizei(len(vecs)))

	d.primitivesAttributeColor.Disable()
	d.primitivesAttributeCoord.Disable()
}

func (d *Driver) DrawTriangle(color color.Color, vec1, vec2, vec3 baukasten.Vector2) {
	vertices := []float32{vec1.X, vec1.Y, vec2.X, vec2.Y, vec3.X, vec3.Y}
	r, g, b, a := baukasten.ConvertColorF(color)
	colors := []float32{r, g, b, a, r, g, b, a, r, g, b, a, r, g, b, a}

	d.primitivesProgram.Use()
	d.primitivesAttributeCoord.Enable()
	d.primitivesAttributeCoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices[0]))

	d.primitivesAttributeColor.Enable()
	d.primitivesAttributeColor.AttribPointer(4, gl.FLOAT, false, 0, gl.Pointer(&colors[0]))

	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	d.primitivesAttributeColor.Disable()
	d.primitivesAttributeCoord.Disable()
}
