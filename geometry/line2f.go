package geometry

type Line2f struct {
	P Vector2
	Q Vector2
}

func Lin2f(p, q Vector2) Line2f {
	return Line2f{P: p, Q: q}
}

// Intersection returns the intersection point with another line and true if it intersects.
// If it does not intersect a zero value Vector2 and false is returned.
func (l Line2f) Intersection(line Line2f) (Vector2, bool) {
	return DoLinesIntersect(l, line)
}

func (l Line2f) IsRectangleCollision(rec Rectanglef) bool {
	return LineRectangleIntersection(l, rec)
}

// DirectionVector returns l's direction vector.
func (l Line2f) DirectionVector() Vector2 {
	return (l.Q.Sub(l.P)).Normalized()
}

// DoLinesIntersect returns the intersection point between two lines and true if it intersects.
// If it does not intersect a zero value Vector2 and false is returned.
func DoLinesIntersect(L1, L2 Line2f) (Vector2, bool) {
	d := (L2.Q.Y-L2.P.Y)*(L1.Q.X-L1.P.X) - (L2.Q.X-L2.P.X)*(L1.Q.Y-L1.P.Y)
	if d == 0 {
		return Vec2(0, 0), false
	}
	n_a := (L2.Q.X-L2.P.X)*(L1.P.Y-L2.P.Y) - (L2.Q.Y-L2.P.Y)*(L1.P.X-L2.P.X)
	n_b := (L1.Q.X-L1.P.X)*(L1.P.Y-L2.P.Y) - (L1.Q.Y-L1.P.Y)*(L1.P.X-L2.P.X)

	ua := n_a / d
	ub := n_b / d

	var ptIntersection Vector2
	if ua >= 0 && ua <= 1 && ub >= 0 && ub <= 1 {
		ptIntersection.X = L1.P.X + (ua * (L1.Q.X - L1.P.X))
		ptIntersection.Y = L1.P.Y + (ua * (L1.Q.Y - L1.P.Y))
		return ptIntersection, true
	}
	return Vec2(0, 0), false
}

func LineRectangleIntersection(line Line2f, rec Rectanglef) bool {
	if _, intersects := DoLinesIntersect(line, rec.Top()); intersects {
		return true
	}
	if _, intersects := DoLinesIntersect(line, rec.Bottom()); intersects {
		return true
	}
	if _, intersects := DoLinesIntersect(line, rec.Left()); intersects {
		return true
	}
	_, intersects := DoLinesIntersect(line, rec.Right())
	return intersects
}
