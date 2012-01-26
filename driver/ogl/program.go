package ogl

import (
	"errors"
	"fmt"

	gl "github.com/chsc/gogl/gl33"
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
		return errors.New("Error in program.\n")
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
