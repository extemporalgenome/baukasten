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
	glimage "github.com/Agon/baukasten/image"
	gl "github.com/chsc/gogl/gl33"
	"image"
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
