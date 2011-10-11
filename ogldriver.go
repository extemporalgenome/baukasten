package baukasten

import (
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
