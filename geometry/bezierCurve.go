package geometry

type BezierCurve []Vector2

func BCurve(points ...Vector2) BezierCurve {
	return BezierCurve(points)
}

func (b BezierCurve) RelativePoint(pos float32) Vector2 {
	if pos < 0 || pos > 1 {
		return Vec2(0, 0)
	}
	vecs := relativePoint([]Vector2(b), pos)
	if len(vecs) == 0 {
		return Vec2(0, 0)
	}
	if len(vecs) > 1 {
		panic("relativePoint should reduce all points to one.")
	}
	return vecs[0]
}

func relativePoint(points []Vector2, pos float32) []Vector2 {
	if len(points) < 2 {
		if len(points) == 1 {
			return points[:1]
		}
		return []Vector2{}
	}
	subPoints := make([]Vector2, len(points)-1)
	for i := 0; i < len(points)-1; i++ {
		subPoints[i] = points[i].Add(points[i+1].Sub(points[i]).Scaled(pos))
	}
	if len(subPoints) > 1 {
		return relativePoint(subPoints, pos)
	}
	return subPoints
}
