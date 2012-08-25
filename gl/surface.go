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
	glimage "github.com/Agon/baukasten/image"
	"github.com/Agon/baukasten/math/geometry"
	"github.com/Agon/baukasten/math/matrix"
	"image"
)

const (
	surfaceVertexShaderData   = "#version 120\nattribute vec2 coord2d;\nattribute vec2 texcoord;\nvarying vec2 f_texcoord;\nvarying vec4 f_color;\nuniform mat4 mvp;\nuniform vec4 v_color;\nvoid main(void) {\ngl_Position = mvp * vec4(coord2d, 0.0, 1.0);\nf_texcoord = texcoord;\nf_color = v_color;\n}"
	surfaceFragmentShaderData = "#version 120\nvarying vec2 f_texcoord;\nvarying vec4 f_color;\nuniform sampler2D texture;\nvoid main(void) {\n" +
		"gl_FragColor = f_color * texture2D(texture, f_texcoord);\n}"
)

var (
	DefaultSurfaceShaderProgram *Program
	DefaultSurfaceCamera        matrix.Matrix4
	surfaceCoordAttrib          *AttributeLocation
	surfaceTexAttrib            *AttributeLocation
	surfaceColorUni             *UniformLocation
	surfaceMvpUni               *UniformLocation
	surfaceTexUni               *UniformLocation
)

type Surface struct {
	// Scale
	scaleX float32
	scaleY float32
	// Rotate
	angle float32
	// Size
	width, height float32
	// Color
	r, g, b, a float32
	// OpenGL
	texture *Texture
	vbo     *VertexBufferObject
}

func OpenSurface(name string) (*Surface, error) {
	img, err := glimage.OpenImage(name)
	if err != nil {
		return nil, err
	}
	return LoadSurface(img)
}

func LoadSurface(img image.Image) (*Surface, error) {
	texture, err := LoadTexture(img)
	if err != nil {
		return nil, err
	}
	// Generate triangles
	width := float32(img.Bounds().Dx())
	height := float32(img.Bounds().Dy())

	x := width / 2
	y := height / 2
	triangles := []float32{
		-x, -y, 0, 0,
		x, -y, 1, 0,
		x, y, 1, 1,
		x, y, 1, 1,
		-x, y, 0, 1,
		-x, -y, 0, 0,
	}

	vbo := NewVBO()
	vbo.BufferData(triangles)

	surface := &Surface{
		texture: texture,
		vbo:     vbo,
		width:   width,
		height:  height,
	}
	return surface, nil
}

func (s *Surface) Delete() {
	s.texture.Delete()
	s.vbo.Delete()
}

func (s *Surface) Draw(x, y float32) {
	DefaultSurfaceShaderProgram.Use()
	model := matrix.TranslationMatrix(x, y, 0)
	model = model.Mul(matrix.ScaleMatrix(s.scaleX, s.scaleY, 1))
	if s.angle != 0 {
		model = model.Mul(matrix.RotationMatrix(s.angle, geometry.Vector3{0, 0, 1}))
	}
	matrix := DefaultSurfaceCamera.Mul(model)
	surfaceMvpUni.UniformMatrix4fv(1, false, matrix.Transposed())

	surfaceTexUni.Uniform1i(0)
	surfaceColorUni.Uniform4f(s.r, s.g, s.b, s.a)

	surfaceCoordAttrib.Enable()
	defer surfaceCoordAttrib.Disable()

	s.vbo.Bind()
	defer s.vbo.Unbind()
	surfaceCoordAttrib.AttribPointer(2, FloatType, false, 16, Offset(nil, 0))

	s.texture.Bind()
	defer s.texture.Unbind()

	surfaceTexAttrib.Enable()
	defer surfaceTexAttrib.Disable()
	surfaceTexAttrib.AttribPointer(2, FloatType, false, 16, Offset(nil, 8))

	s.vbo.DrawArrays(0, 6)
}
