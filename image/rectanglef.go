package image

// ZR is the zero Rectanglef.
var ZR Rectanglef

type Rectanglef struct {
	Min, Max Vector2
}

// Rectf is shorthand for Rectanglef{Vector2(x0, y0), Vector2(x1, y1)}.
func Rectf(x0, y0, x1, y1 float32) Rectanglef {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rectanglef{Vector2{x0, y0}, Vector2{x1, y1}}
}

// String returns a string representation of r
func (r Rectanglef) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

// Dx returns r's width.
func (r Rectanglef) Dx() float32 {
	return r.Max.X - r.Min.X
}

// Dy returns r's height.
func (r Rectanglef) Dy() float32 {
	return r.Max.Y - r.Min.Y
}

// Size returns r's width and height
func (r Rectanglef) Size() Vector2 {
	return Vector2{r.Max.X - r.Min.X, r.Max.Y - r.Min.Y}
}

// Center returns r's center.
func (r Rectanglef) Center() Vector2 {
	return Vec2(r.Min.X+r.Dx()/2, r.Min.Y+r.Dy()/2)
}

// Top returns r's top line segment
func (r Rectanglef) Top() Line2f {
	return Lin2f(Vec2(r.Min.X, r.Max.Y), r.Max)
}

// Top returns r's bottom line segment
func (r Rectanglef) Bottom() Line2f {
	return Lin2f(r.Min, Vec2(r.Max.X, r.Min.Y))
}

// Top returns r's left line segment
func (r Rectanglef) Left() Line2f {
	return Lin2f(r.Min, Vec2(r.Min.X, r.Max.Y))
}

// Top returns r's right line segment
func (r Rectanglef) Right() Line2f {
	return Lin2f(Vec2(r.Max.X, r.Min.Y), r.Max)
}

// Add returns the rectangle r translated by p.
func (r Rectanglef) Add(v Vector2) Rectanglef {
	return Rectanglef{
		Vector2{r.Min.X + v.X, r.Min.Y + v.Y},
		Vector2{r.Max.X + v.X, r.Max.Y + v.Y},
	}
}

// Sub returns the rectangle r translated by -p.
func (r Rectanglef) Sub(v Vector2) Rectanglef {
	return Rectanglef{
		Vector2{r.Min.X - v.X, r.Min.Y - v.Y},
		Vector2{r.Max.X - v.X, r.Max.Y - v.Y},
	}
}

// Intersect returns the largest rectangle contained by both r and s. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (r Rectanglef) Intersect(s Rectanglef) Rectanglef {
	if r.Min.X < s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y < s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X > s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y > s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	if r.Min.X > r.Max.X || r.Min.Y > r.Max.Y {
		return ZR
	}
	return r
}

// IntersectCircle returns true if r intersects c.
func (r Rectanglef) IntersectCircle(c Circlef) bool {
	return c.Position.InRec(r) || c.IntersectLine(r.Top()) || c.IntersectLine(r.Bottom()) || c.IntersectLine(r.Left()) || c.IntersectLine(r.Right())
}

// IntersectLine returns true if r intersects l and the intersection points.
func (r Rectanglef) IntersectLine(line Line2f) (bool, []Vector2) {
	return LineRectangleIntersection(line, r)
}

// Union returns the smallest rectangle that contains both r and s.
func (r Rectanglef) Union(s Rectanglef) Rectanglef {
	if r.Min.X > s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y > s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X < s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y < s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	return r
}

// Empty returns whether the rectangle contains no space.
func (r Rectanglef) Empty() bool {
	return r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y
}

// Eq returns whether r and s are equal.
func (r Rectanglef) Eq(s Rectanglef) bool {
	return r.Min.X == s.Min.X && r.Min.Y == s.Min.Y &&
		r.Max.X == s.Max.X && r.Max.Y == s.Max.Y
}

// Overlaps returns whether r and s have a non-empty intersection.
func (r Rectanglef) Overlaps(s Rectanglef) bool {
	return r.Min.X < s.Max.X && s.Min.X < r.Max.X &&
		r.Min.Y < s.Max.Y && s.Min.Y < r.Max.Y
}
