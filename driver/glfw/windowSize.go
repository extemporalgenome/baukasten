package glfw

type WindowSize struct {
	width, height int
}

func NewWindowSize(w, h int) *WindowSize {
	return &WindowSize{w, h}
}

func (ws *WindowSize) Width() int {
	return ws.width
}

func (ws *WindowSize) Height() int {
	return ws.height
}
