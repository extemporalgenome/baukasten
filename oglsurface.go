package baukasten

import (
	"gl"
)

type OglSurface struct {
	texture gl.Texture
}

func (surface *OglSurface) Draw(rec *RectangleF) {
	surface.texture.Bind(gl.TEXTURE_2D)

	gl.Begin(gl.QUADS)

	gl.TexCoord2f(0, 0)
	gl.Vertex3f(rec.Position.X-rec.Size.Width/2, rec.Position.Y+rec.Size.Height/2, 0) // Bottom left

	gl.TexCoord2f(1, 0)
	gl.Vertex3f(rec.Position.X+rec.Size.Width/2, rec.Position.Y+rec.Size.Height/2, 0) // Bottom right

	gl.TexCoord2f(1, 1)
	gl.Vertex3f(rec.Position.X+rec.Size.Width/2, rec.Position.Y-rec.Size.Height/2, 0) // Top right

	gl.TexCoord2f(0, 1)
	gl.Vertex3f(rec.Position.X-rec.Size.Width/2, rec.Position.Y-rec.Size.Height/2, 0) // Top left

	gl.End()

	surface.texture.Unbind(gl.TEXTURE_2D)
}

func (surface *OglSurface) Delete() {
	surface.texture.Delete()
}
