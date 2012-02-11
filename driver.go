package baukasten

import (
	"image"
	"image/color"
)

type GraphicDriver interface {
	Init(*GraphicSettings) error
	Close()
	BeginFrame()
	EndFrame()
	SetClearColor(color.Color)
	Camera() Camera
	SetCamera(Camera)
	Resize(int, int)
	OpenSurface(string) (Surface, error)
	LoadSurface(image.Image) (Surface, error)
	DrawPoints(color color.Color, vecs ...Vector2)
	DrawLines(color color.Color, vecs ...Vector2)
	DrawLineStrip(color color.Color, vecs ...Vector2)
	DrawLineLoop(color color.Color, vecs ...Vector2)
	DrawTriangles(color color.Color, vecs ...Vector2)
	DrawTriangleStrip(color color.Color, vecs ...Vector2)
	DrawTriangleFan(color color.Color, vecs ...Vector2)
}

type ContextDriver interface {
	Init(*GraphicSettings) error
	Close()
	SwapBuffers()
	ResizeEvent() <-chan WindowSizeEvent
	ContextEvent() <-chan ContextEvent
}

type InputDriver interface {
	KeyEvent() <-chan KeyEvent
	MouseButtonEvent() <-chan MouseButtonEvent
	MousePositionEvent() <-chan MousePositionEvent
	MouseWheelEvent() <-chan MouseWheelEvent
	JoystickPos(joy int) []Vector2
	JoystickButtons(joy int) []bool
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
	GetScale() (float32, float32)
	Rotate(angle float32)
	Rotation() float32
}
