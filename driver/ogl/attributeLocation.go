package ogl

import (
	gl "github.com/chsc/gogl/gl33"
)

type AttributeLocation struct {
	id gl.Uint
}

func (loc *AttributeLocation) Enable() {
	gl.EnableVertexAttribArray(loc.id)
}

func (loc *AttributeLocation) Disable() {
	gl.DisableVertexAttribArray(loc.id)
}

func (loc *AttributeLocation) AttribPointer(size int, t gl.Enum, normalized bool, stride int, pointer gl.Pointer) {
	gl.VertexAttribPointer(loc.id, gl.Int(size), t, gl.GLBool(normalized), gl.Sizei(stride), pointer)
}
