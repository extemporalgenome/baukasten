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
	"errors"
	"io/ioutil"

	gl "github.com/chsc/gogl/gl33"
)

type ShaderType byte

const (
	VertexShader ShaderType = iota
	FragmentShader
)

var (
	ErrUnknownShader = errors.New("Unknown shader type.")
)

type Shader struct {
	id gl.Uint
}

func OpenShader(name string, shaderType ShaderType) (*Shader, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return LoadShader(string(b), shaderType)
}

func LoadShader(data string, shaderType ShaderType) (*Shader, error) {
	var id gl.Uint
	switch shaderType {
	case VertexShader:
		id = gl.CreateShader(gl.VERTEX_SHADER)
	case FragmentShader:
		id = gl.CreateShader(gl.FRAGMENT_SHADER)
	default:
		return nil, ErrUnknownShader
	}
	shader := &Shader{id: id}
	src := gl.GLStringArray(data)
	length := gl.Int(-1)
	gl.ShaderSource(shader.id, gl.Sizei(1), &src[0], &length)
	gl.GLStringArrayFree(src)
	return shader, nil
}

func (s *Shader) Delete() {
	gl.DeleteShader(s.id)
}

func (s *Shader) Id() gl.Uint {
	return s.id
}

func (s *Shader) Compile() error {
	compileOk := gl.Int(gl.FALSE)
	gl.CompileShader(s.id)
	gl.GetShaderiv(s.id, gl.COMPILE_STATUS, &compileOk)
	if compileOk == gl.FALSE {
		return s.getError()
	}
	return nil
}

func (s *Shader) getError() error {
	var logLength gl.Int
	gl.GetShaderiv(s.id, gl.INFO_LOG_LENGTH, &logLength)
	log := gl.GLStringAlloc(gl.Sizei(logLength))
	defer gl.GLStringFree(log)
	gl.GetShaderInfoLog(s.id, gl.Sizei(logLength), nil, log)
	err := gl.GoString(log)
	return errors.New(err)
}
