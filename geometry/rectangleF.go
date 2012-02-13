package geometry

// ZR is the zero RectangleF.
var ZR RectangleF

type RectangleF struct {
	Min, Max Vector2
}

// RectF is shorthand for RectangleF{Vector2(x0, y0), Vector2(x1, y1)}.
func RectF(x0, y0, x1, y1 float32) RectangleF {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return RectangleF{Vector2{x0, y0}, Vector2{x1, y1}}
}

// String returns a string representation of r
func (r RectangleF) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

// Dx returns r's width.
func (r RectangleF) Dx() float32 {
	return r.Max.X - r.Min.X
}

// Dy returns r's height.
func (r RectangleF) Dy() float32 {
	return r.Max.Y - r.Min.Y
}

// Size returns r's width and height
func (r RectangleF) Size() Vector2 {
	return Vector2{r.Max.X - r.Min.X, r.Max.Y - r.Min.Y}
}

// Add returns the rectangle r translated by p.
func (r RectangleF) Add(v Vector2) RectangleF {
	return RectangleF{
		Vector2{r.Min.X + v.X, r.Min.Y + v.Y},
		Vector2{r.Max.X + v.X, r.Max.Y + v.Y},
	}
}

// Sub returns the rectangle r translated by -p.
func (r RectangleF) Sub(v Vector2) RectangleF {
	return RectangleF{
		Vector2{r.Min.X - v.X, r.Min.Y - v.Y},
		Vector2{r.Max.X - v.X, r.Max.Y - v.Y},
	}
}

// Intersect returns the largest rectangle contained by both r and s. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (r RectangleF) Intersect(s RectangleF) RectangleF {
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

// Union returns the smallest rectangle that contains both r and s.
func (r RectangleF) Union(s RectangleF) RectangleF {
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
func (r RectangleF) Empty() bool {
	return r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y
}

// Eq returns whether r and s are equal.
func (r RectangleF) Eq(s RectangleF) bool {
	return r.Min.X == s.Min.X && r.Min.Y == s.Min.Y &&
		r.Max.X == s.Max.X && r.Max.Y == s.Max.Y
}

// Overlaps returns whether r and s have a non-empty intersection.
func (r RectangleF) Overlaps(s RectangleF) bool {
	return r.Min.X < s.Max.X && s.Min.X < r.Max.X &&
		r.Min.Y < s.Max.Y && s.Min.Y < r.Max.Y
}
