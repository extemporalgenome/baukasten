package baukasten

import (
	"image"
	"image/color"
)

type GraphicDriver interface {
	Init(*GraphicSettings) error
	Close()
	BeginFrame()
	SetClearColor(color.Color)
	EndFrame()
	OpenSurface(string) (Surface, error)
	LoadSurface(image.Image) (Surface, error)
	Resize(int, int)
	DrawTriangle(vec1, vec2, vec3 Vector2, color color.Color)
}

type ContextDriver interface {
	Init(*GraphicSettings) error
	Close()
	SwapBuffers()
	ResizeEvent() chan WindowSizeEvent
	ContextEvent() chan ContextEvent
}

type InputDriver interface {
	KeyEvent() chan KeyEvent
	MouseButtonEvent() chan MouseButtonEvent
	MousePositionEvent() chan MousePositionEvent
	MouseWheelEvent() chan MouseWheelEvent
}

type FontDriver interface {
	OpenFont(fileName string) (Font, error)
	LoadFont(data []byte) (Font, error)
}

type Font interface {
	Render(text string, width, height int, size float64, color color.Color) image.Image
}

type Surface interface {
	Draw(x, y float32)
	Delete()
	Scale(x, y float32)
	Rotate(angle float32)
}