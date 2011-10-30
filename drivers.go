package baukasten

import (
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
}

type InputDriver interface {
	GetKeyEvent()
	GetMouseButtonEvent()
}

type Surface interface {
	Draw(*RectangleF)
	Delete()
}
