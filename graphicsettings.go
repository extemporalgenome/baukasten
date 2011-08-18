package baukasten

const (
	BitDepthDefault = 16
)

type GraphicSettings struct {
	Width, Height int
	BitDepth      int
	Fullscreen    bool
	Resizeable    bool
	Caption       string
}

func NewGraphicSettings(width, height, bitDepth int, fullscreen, resizeable bool, caption string) *GraphicSettings {
	if bitDepth > 0 && bitDepth <= 32 {
		return &GraphicSettings{width, height, bitDepth, fullscreen, resizeable, caption}
	}
	return &GraphicSettings{width, height, BitDepthDefault, fullscreen, resizeable, caption}
}
