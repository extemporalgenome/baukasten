package baukasten

import (
	"os"
)

type DisplayManager interface {
	Init() os.Error
	Quit() os.Error
}
