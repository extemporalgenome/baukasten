package baukasten

type Point struct {
	X, Y int
}

func (p *Point) NewPoint(x, y int) *Point {
	return &Point{x, y}
}
