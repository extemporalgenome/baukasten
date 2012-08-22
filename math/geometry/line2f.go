package geometry

type Line2f struct {
	P Vector2
	Q Vector2
}

func Lin2f(p, q Vector2) Line2f {
	return Line2f{P: p, Q: q}
}

// Intersect returns the intersection point with another line and true if it intersects.
// If it does not intersect a zero value Vector2 and false is returned.
func (l Line2f) Intersect(line Line2f) (Vector2, bool) {
	return DoLinesIntersect(l, line)
}

// IntersectCircle returns true if l intersects c.
func (l Line2f) IntersectCircle(c Circlef) bool {
	closest := closestPointOnSeg(l, c.Position)
	dist := c.Position.Sub(closest)
	return dist.Length() <= c.Radius
}

// IntersectRec returns true if l intersects rec and the intersection points
func (l Line2f) IntersectRec(rec Rectanglef) (bool, []Vector2) {
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

// LineRectangleIntersection returns true if line intersects rec and the intersection points.
func LineRectangleIntersection(line Line2f, rec Rectanglef) (bool, []Vector2) {
	points := make([]Vector2, 0)
	intersection := false
	if p, intersects := DoLinesIntersect(line, rec.Top()); intersects {
		intersection = true
		points = append(points, p)
	}
	if p, intersects := DoLinesIntersect(line, rec.Bottom()); intersects {
		intersection = true
		points = append(points, p)
	}
	if p, intersects := DoLinesIntersect(line, rec.Left()); intersects {
		intersection = true
		points = append(points, p)
	}
	if p, intersects := DoLinesIntersect(line, rec.Right()); intersects {
		intersection = true
		points = append(points, p)
	}
	return intersection, points
}
