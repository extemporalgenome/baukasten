package geometry

type Circlef struct {
	Position Vector2
	Radius   float32
}

func Circf(position Vector2, radius float32) Circlef {
	return Circlef{Position: position, Radius: radius}
}

// Intersect returns true if c intersects c1.
func (c Circlef) Intersect(c1 Circlef) bool {
	return c.Radius+c1.Radius >= c.Position.DistanceBetween(c1.Position)
}

// IntersectLine returns true if c intersects line.
func (c Circlef) IntersectLine(line Line2f) bool {
	closest := closestPointOnSeg(line, c.Position)
	dist := c.Position.Sub(closest)
	return dist.Length() <= c.Radius
}

// IntersectRec returns true if c intersects rec.
func (c Circlef) IntersectRec(rec Rectanglef) bool {
	return c.Position.InRec(rec) || c.IntersectLine(rec.Top()) || c.IntersectLine(rec.Bottom()) || c.IntersectLine(rec.Left()) || c.IntersectLine(rec.Right())
}

// closestPointOnSeg returns the closest point towards vec from line.
func closestPointOnSeg(line Line2f, vec Vector2) Vector2 {
	seg := line.Q.Sub(line.P)
	pt := vec.Sub(line.P)
	if seg.Length() <= 0 {
		panic("Invalid segment length")
	}
	segUnit := seg.Normalized()
	projLength := pt.Dot(segUnit)
	if projLength <= 0 {
		return line.P
	}
	if projLength >= seg.Length() {
		return line.Q
	}
	proj := segUnit.Scaled(projLength)
	return proj.Add(line.P)
}
