package baukasten

type Gradient struct {
	TopLeft     *Color
	TopRight    *Color
	BottomLeft  *Color
	BottomRight *Color
}

func NewGradient(color *Color) *Gradient {
	return &Gradient{color, color, color, color}
}

func NewComplexGradient(topLeft, topRight, bottomLeft, bottomRight *Color) *Gradient {
	return &Gradient{topLeft, topRight, bottomLeft, bottomRight}
}
