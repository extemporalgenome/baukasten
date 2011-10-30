package baukasten

import (
	"image"
	"os"
)

type GraphicDriver interface {
	Init(*GraphicSettings) os.Error
	Close()
	BeginFrame()
	EndFrame()
}

type ContextDriver interface {
	Init(*GraphicSettings) os.Error
	Close()
	SwapBuffers()
	OpenSurface(string) (Surface, os.Error)
	LoadSurface(image.Image) (Surface, os.Error)
}

type InputDriver interface {
	GetKeyEvent()
	GetMouseButtonEvent()
}

type Surface interface {
	Draw(*RectangleF)
	Delete()
}
