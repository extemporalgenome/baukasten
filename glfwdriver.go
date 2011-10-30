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
	"github.com/jteeuwen/glfw"
)

type GlfwDriver struct {
	graphicSettings *GraphicSettings
}

func NewGlfwDriver() *GlfwDriver {
	return &GlfwDriver{}
}

func (driver *GlfwDriver) Init(settings *GraphicSettings) os.Error {
	err := glfw.Init()
	if err != nil {
		return err
	}
	windowType := glfw.Windowed
	if settings.Fullscreen {
		windowType = glfw.Fullscreen
	}
	// TODO BitDepth
	err = glfw.OpenWindow(settings.Width, settings.Height, 8, 8, 8, 8, 0, 8, windowType)
	if err != nil {
		glfw.Terminate()
		return err
	}
	glfw.SetSwapInterval(1) // VSync
	glfw.SetWindowTitle(settings.Title)

	glfw.SetWindowSizeCallback(onResize)

	driver.graphicSettings = settings
	return nil
}

func onResize(w, h int) {
	gl.Viewport(0, 0, w, h)
}

func (driver *GlfwDriver) Close() {
	glfw.Terminate()
	glfw.CloseWindow()
}

func (driver *GlfwDriver) SwapBuffers() {
	glfw.SwapBuffers()
}

func (driver *GlfwDriver) OpenSurface(name string) (surface Surface, err os.Error) {
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

func (driver *GlfwDriver) LoadSurface(img image.Image) (surface Surface, err os.Error) {
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

// InputDriver implementations
func (driver *GlfwDriver) GetKeyEvent()         {}
func (driver *GlfwDriver) GetMouseButtonEvent() {}
