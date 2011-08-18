package baukasten

import (
	"os"
)

type WindowManager interface {
	Init(graphicSettings *GraphicSettings) os.Error
	Resize(width, height int) os.Error
	Quit() os.Error
}
