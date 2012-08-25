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
	"github.com/Agon/baukasten/math/matrix"
	gl "github.com/chsc/gogl/gl33"
	"image/color"
)

const (
	FloatType = gl.FLOAT
)

var (
	DefaultCamera matrix.Matrix4
)

func Init() error {
	if err := gl.Init(); err != nil {
		return err
	}
	// Load default vertex- and fragmentshader for surfaces
	surfaceVS, err := LoadShader(surfaceVertexShaderData, VertexShader)
	if err != nil {
		return err
	}
	surfaceFS, err := LoadShader(surfaceFragmentShaderData, FragmentShader)
	if err != nil {
		return err
	}
	DefaultSurfaceShaderProgram = NewProgram()
	DefaultSurfaceShaderProgram.AttachShaders(surfaceVS, surfaceFS)
	if err = DefaultSurfaceShaderProgram.Link(); err != nil {
		return err
	}
	// Surface AttributeLocations
	surfaceCoordAttrib, err = DefaultSurfaceShaderProgram.GetAttributeLocation("coord2d")
	if err != nil {
		return err
	}
	surfaceTexAttrib, err = DefaultSurfaceShaderProgram.GetAttributeLocation("texcoord")
	if err != nil {
		return err
	}
	// UniformLocations
	surfaceColorUni, err = DefaultSurfaceShaderProgram.GetUniformLocation("v_color")
	if err != nil {
		return err
	}
	surfaceMvpUni, err = DefaultSurfaceShaderProgram.GetUniformLocation("mvp")
	if err != nil {
		return err
	}
	surfaceTexUni, err = DefaultSurfaceShaderProgram.GetUniformLocation("texture")
	if err != nil {
		return err
	}
	// TODO Load primitives shaders, program, attribute- and uniformlocations
	return nil
}

func Close() {
	// TODO Both surface shaders aren't getting deleted in this func (?)
	DefaultSurfaceShaderProgram.Delete()
}

// TODO Allow different blend functions
func EnableBlend() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
}

func Clear(c color.Color) {
	r, g, b, a := glcolor.ConvertColorGL(c)
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT) // gl.DEPTH_BUFFER_BIT
}

func Offset(n gl.Pointer, m uintptr) gl.Pointer {
	return gl.Offset(n, m)
}
