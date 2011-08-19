package baukasten

type OGLDriver struct {

}

func (driver *OGLDriver) Init() os.Error {
	if gl.Init() != 0 {
		return os.NewError("OpenGL initialization error.")
	}
	return nil
}

func (driver *OGLDriver) Clear() {
	gl.ClearColor(0, 0, 0, 0)
}

func (driver *OGLDriver) Quit() os.Error {
	return nil
}
