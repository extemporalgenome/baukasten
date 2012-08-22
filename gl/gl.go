package gl

import (
	glcolor "github.com/Agon/baukasten/image/color"
	gl "github.com/chsc/gogl/gl33"
	"image/color"
)

func Clear(c color.Color) {
	r, g, b, a := glcolor.ConvertColorGL(c)
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT) // gl.DEPTH_BUFFER_BIT
}

func Offset(n gl.Pointer, m uintptr) gl.Pointer {
	return gl.Offset(n, m)
}
