package gl

import (
	"image"
	gl "github.com/chsc/gogl/gl33"
	glimage "github.com/Agon/baukasten/image"
)

type Texture struct {
	id gl.Uint
}

// OpenTexture opens an image, converts it into a OpenGL compatible image format and calls LoadTexture.
// Supported formats: bmp, tiff, gif, jpeg, png
func OpenTexture(name string) (*Texture, error) {
	img, err := glimage.OpenImage(name)
	if err != nil {
		return nil, err
	}
	return LoadTexture(img)
}

// LoadTexture buffers an image.Image into the graphic cards memory.
func LoadTexture(img image.Image) (*Texture, error) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}
	var textureId gl.Uint
	//gl.ActiveTexture(gl.TEXTURE0)
	gl.GenTextures(1, &textureId)
	gl.BindTexture(gl.TEXTURE_2D, textureId)
	//gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, 4, gl.Sizei(rgba.Rect.Dx()), gl.Sizei(rgba.Rect.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&rgba.Pix[0]))

	gl.BindTexture(gl.TEXTURE_2D, 0)
	return &Texture{id: textureId}, nil
}

func (t *Texture) Render(location *UniformLocation) {
	//gl.ActiveTexture(gl.TEXTURE0)
	location.Uniform1i(0) // gl.Texture or 0
	gl.BindTexture(gl.TEXTURE_2D, t.id)
	
}

// Delete deletes the memory buffer on the graphic card.
func (t *Texture) Delete() {
	gl.DeleteTextures(1, &t.id)
}
