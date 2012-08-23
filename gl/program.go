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
	"fmt"

	gl "github.com/chsc/gogl/gl33"
)

var (
	ErrProgramError = errors.New("Error in program.")
)

type Program struct {
	id gl.Uint
}

func NewProgram() *Program {
	id := gl.CreateProgram()
	return &Program{id: id}
}

func (p *Program) AttachShaders(shaders ...*Shader) {
	for _, shader := range shaders {
		gl.AttachShader(p.id, shader.Id())
	}
}

func (p *Program) Link() error {
	var compileOk gl.Int
	gl.LinkProgram(p.id)
	gl.GetProgramiv(p.id, gl.LINK_STATUS, &compileOk)
	if compileOk == 0 {
		// TODO Get errror
		return ErrProgramError
	}
	return nil
}

func (p *Program) Use() {
	gl.UseProgram(p.id)
}

func (p *Program) GetAttributeLocation(name string) (*AttributeLocation, error) {
	attributeName := gl.GLString(name)
	defer gl.GLStringFree(attributeName)
	attributeTemp := gl.GetAttribLocation(p.id, attributeName)
	if attributeTemp == -1 {
		return nil, errors.New(fmt.Sprintf("Could not bind attribute %s\n", gl.GoString(attributeName)))
	}

	return &AttributeLocation{id: gl.Uint(attributeTemp)}, nil
}

func (p *Program) GetUniformLocation(name string) (*UniformLocation, error) {
	attributeName := gl.GLString(name)
	defer gl.GLStringFree(attributeName)
	id := gl.GetUniformLocation(p.id, attributeName)
	if id == -1 {
		return nil, errors.New(fmt.Sprintf("Could not bind uniform %s\n", gl.GoString(attributeName)))
	}
	return &UniformLocation{id: id}, nil
}

func (p *Program) Delete() {
	gl.DeleteProgram(p.id)
}
