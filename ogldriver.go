package baukasten

import (
	"image"
	// For image loading
	_ "image/bmp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	_ "image/tiff"
	_ "image/ycbcr"
	"os"

	"gl"
)

type OglDriver struct {
	settings *GraphicSettings
}

func NewOglDriver() *OglDriver {
	return &OglDriver{}
}

func (driver *OglDriver) Init(graphicSettings *GraphicSettings) os.Error {
	driver.settings = graphicSettings

	gl.ShadeModel(gl.SMOOTH)
	gl.ClearColor(0, 0, 0, 0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE_2D)
	return nil
}

func (driver *OglDriver) Close() {
	// TODO Release all loaded memory assets (textures)
}

func (driver *OglDriver) BeginFrame() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()
}

func (driver *OglDriver) EndFrame() {}

func (driver *OglDriver) OpenSurface(name string) (surface Surface, err os.Error) {
	file, err := os.Open(name)
	if err != nil {
		return
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}
	return driver.LoadSurface(img)
}

func (driver *OglDriver) LoadSurface(img image.Image) (surface Surface, err os.Error) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	rgba := image.NewRGBA(w, h)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}
	texture := gl.GenTexture()
	texture.Bind(gl.TEXTURE_2D)

	gl.TexImage2D(gl.TEXTURE_2D, 0, 4, rgba.Rect.Dx(), rgba.Rect.Dy(), 0, gl.RGBA, gl.UNSIGNED_BYTE, rgba.Pix)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	texture.Unbind(gl.TEXTURE_2D)

	return &OglSurface{texture}, nil
}
