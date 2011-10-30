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
	OpenSurface(string) (Surface, os.Error)
	LoadSurface(image.Image) (Surface, os.Error)
}

type ContextDriver interface {
	Init(*GraphicSettings) os.Error
	Close()
	SwapBuffers()
}

type InputDriver interface {
	GetKeyEvent()
	GetMouseButtonEvent()
}

type Surface interface {
	Draw(*RectangleF)
	Delete()
}
