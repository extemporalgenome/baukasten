package baukasten

import (
	"fmt"
	"os"

	"sdl"
)

type SDLDriver struct {
	graphicSettings *GraphicSettings
	screen          *sdl.Surface
}

func NewSDLDriver() *SDLDriver {
	return &SDLDriver{}
}

func (driver *SDLDriver) Init(graphicSettings *GraphicSettings) os.Error {
	driver.graphicSettings = graphicSettings
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		return fmt.Errorf("SDL initialize error: %s", sdl.GetError())
	}
	settings := uint32(sdl.OPENGL)
	if driver.graphicSettings.Resizeable {
		settings |= sdl.RESIZABLE
	}
	driver.screen = sdl.SetVideoMode(driver.graphicSettings.Width, driver.graphicSettings.Height, driver.graphicSettings.BitDepth, settings)
	if driver.screen == nil {
		return fmt.Errorf("SDL video mode set error: %s", sdl.GetError())
	}
	sdl.WM_SetCaption(driver.graphicSettings.Caption, driver.graphicSettings.Caption)
	return nil
}

func (driver *SDLDriver) Quit() os.Error {
	sdl.Quit()
	return nil
}

func (driver *SDLDriver) Resize(width, height int) os.Error {
	driver.screen = sdl.SetVideoMode(width, height, driver.graphicSettings.BitDepth, sdl.OPENGL|sdl.RESIZABLE)
	if driver.screen == nil {
		return fmt.Errorf("SDL video mode set error on resize: %s", sdl.GetError())
	}
	return nil
}
