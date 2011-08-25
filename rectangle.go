package baukasten

type Rectangle struct {
	Position *Point
	Size     *Size
}

func NewRectangle(pos *Point, size *Size) *Rectangle {
	return &Rectangle{pos, size}
}

type RectangleF struct {
	Position *Vector2
	Size     *SizeF
}

func NewRectangleF(pos *Vector2, size *SizeF) *RectangleF {
	return &RectangleF{pos, size}
}
