package glfw

type WindowSize struct {
	width, height uint
}

func NewWindowSize(w, h uint) *WindowSize {
	return &WindowSize{w, h}
}

func (ws *WindowSize) Width() uint {
	return ws.width
}

func (ws *WindowSize) Height() uint {
	return ws.height
}
