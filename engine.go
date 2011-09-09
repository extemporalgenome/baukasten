package baukasten

import (
	"os"
)

type Engine interface {
	Init(settings *GraphicSettings) os.Error
	Close()

	ResizeEvent() chan ResizeEvent

	BeginFrame()
	EndFrame()
}
