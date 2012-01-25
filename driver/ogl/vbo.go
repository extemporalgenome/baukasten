package ogl

import (
	gl "github.com/chsc/gogl/gl33"
)

type VertexBufferObject struct {
	id gl.Uint
}

func NewVBO() *VertexBufferObject {
	vbo := &VertexBufferObject{}
	gl.GenBuffers(1, &vbo.id)
	return vbo
}

func (vbo *VertexBufferObject) BufferData(data []float32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.id)
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(data)*4), gl.Pointer(&data[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (vbo *VertexBufferObject) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.id)
}

func (vbo *VertexBufferObject) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (vbo *VertexBufferObject) DrawArrays(start, end int) {
	gl.DrawArrays(gl.TRIANGLES, gl.Int(start), gl.Sizei(end))
}

func (vbo *VertexBufferObject) Delete() {
	gl.DeleteBuffers(1, &vbo.id)
}
