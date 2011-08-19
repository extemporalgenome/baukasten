package baukasten

type DisplayManager interface {
	Init() os.Error
	Quit() os.Error
}