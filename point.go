package baukasten

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}
