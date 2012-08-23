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
