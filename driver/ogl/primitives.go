package ogl

import (
	"image/color"

	"github.com/Agon/baukasten"

	gl "github.com/chsc/gogl/gl33"
)

const (
	PrimitiveAttributeLocationName = "coord"
	PrimitiveVertexShaderSource    = "#version 120\nattribute vec2 " + PrimitiveAttributeLocationName + ";\nvoid main(void) {\n  gl_Position = vec4(" + PrimitiveAttributeLocationName + ", 0.0, 1.0);\n}"
	PrimitiveFragmentShaderSource  = "#version 120\nvoid main(void) {\n  gl_FragColor[0] = 0.0;\n  gl_FragColor[1] = 0.0;\n  gl_FragColor[2] = 1.0;\n}"
)

func (d *Driver) DrawTriangle(vec1, vec2, vec3 baukasten.Vector2, color color.Color) {
	vertices := []float32{vec1.X, vec1.Y, vec2.X, vec2.Y, vec3.X, vec3.Y}

	d.primitivesProgram.Use()
	d.primitivesAttributeCoord.Enable()
	d.primitivesAttributeCoord.AttribPointer(2, gl.FLOAT, false, 0, gl.Pointer(&vertices))
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	d.primitivesAttributeCoord.Disable()

}
