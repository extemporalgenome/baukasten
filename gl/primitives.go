// baukasten - Toolkit for OpenGL
// 
// Copyright (c) 2012, Marcel Hauf <marcel.hauf@googlemail.com>
// 
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met: 
// 
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer. 
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution. 
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package gl

import (
	glcolor "github.com/Agon/baukasten/image/color"
	"github.com/Agon/baukasten/math/geometry"
	gl "github.com/chsc/gogl/gl33"
	"image/color"
)

const (
	primitivesVertexShaderData = "#version 120\n" +
		"attribute vec2 coord;\n" +
		"attribute vec4 v_color;\n" +
		"varying vec4 f_color;\n" +
		"uniform mat4 mvp;\n" +
		"void main(void) {\n" +
		"gl_Position = mvp * vec4(coord, 0.0, 1.0);\n" +
		"f_color = v_color;\n" +
		"}"
	primitivesFragmentShaderData = "#version 120\n" +
		"varying vec4 f_color;\n" +
		"void main(void) {\n" +
		" gl_FragColor = f_color;\n" +
		"}"
)

var (
	DefaultPrimitivesShaderProgram *Program
	primitivesAttributeCoord       *AttributeLocation
	primitivesAttributeColor       *AttributeLocation
	primitivesUniformMatrix        *UniformLocation
)

func DrawPoints(color color.Color, vecs ...geometry.Vector2) {
	drawPrimitives(color, gl.POINTS, vecs...)
}

func DrawLines(color color.Color, vecs ...geometry.Vector2) {
	drawPrimitives(color, gl.LINES, vecs...)
}

func DrawLineStrip(color color.Color, vecs ...geometry.Vector2) {
	drawPrimitives(color, gl.LINE_STRIP, vecs...)
}

func DrawLineLoop(color color.Color, vecs ...geometry.Vector2) {
	drawPrimitives(color, gl.LINE_LOOP, vecs...)
}

func DrawTriangles(color color.Color, vecs ...geometry.Vector2) {
	drawPrimitives(color, gl.TRIANGLES, vecs...)
}

func DrawTriangleStrip(color color.Color, vecs ...geometry.Vector2) {
	drawPrimitives(color, gl.TRIANGLE_STRIP, vecs...)
}

func DrawTriangleFan(color color.Color, vecs ...geometry.Vector2) {
	drawPrimitives(color, gl.TRIANGLE_FAN, vecs...)
}

func drawPrimitives(color color.Color, mode gl.Enum, vecs ...geometry.Vector2) {
	vertices := make([]float32, len(vecs)*2)
	r, g, b, a := glcolor.ConvertColorF(color)
	var colors []float32
	for i := range vecs {
		vertices[i*2] = vecs[i].X
		vertices[i*2+1] = vecs[i].Y
		colors = append(colors, r, g, b, a)
	}
	DefaultPrimitivesShaderProgram.Use()
	primitivesUniformMatrix.UniformMatrix4fv(1, false, DefaultCamera.Transposed())
	primitivesAttributeCoord.Enable()
	primitivesAttributeCoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices[0]))

	primitivesAttributeColor.Enable()
	primitivesAttributeColor.AttribPointer(4, gl.FLOAT, false, 0, gl.Pointer(&colors[0]))

	gl.DrawArrays(mode, 0, gl.Sizei(len(vecs)))

	primitivesAttributeColor.Disable()
	primitivesAttributeCoord.Disable()
}
