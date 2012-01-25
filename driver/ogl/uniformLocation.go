package ogl

import (
	gl "github.com/chsc/gogl/gl33"
)

type UniformLocation struct {
	id gl.Int
}

func (loc *UniformLocation) Uniform1i(value int) {
	gl.Uniform1i(loc.id, gl.Int(value))
}

func (loc *UniformLocation) UniformMatrix4fv(count int, transpose bool, matrix []float32) {
	gl.UniformMatrix4fv(loc.id, gl.Sizei(count), gl.GLBool(transpose), (*gl.Float)(&matrix[0]))
}
