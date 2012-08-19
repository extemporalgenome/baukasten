package image

type BezierCurve []Vector2

func BCurve(points ...Vector2) BezierCurve {
	return BezierCurve(points)
}

// RelativePoint returns the point on pos position on b's curve.
func (b BezierCurve) RelativePoint(pos float32) Vector2 {
	controlPoints := []Vector2(b)
	if len(controlPoints) == 0 {
		return Vec2(0, 0)
	}
	p := 1.0 - pos
	for j := len(controlPoints) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			controlPoints[i].X = p*controlPoints[i].X + pos*controlPoints[i+1].X
			controlPoints[i].Y = p*controlPoints[i].Y + pos*controlPoints[i+1].Y
		}
	}
	return controlPoints[0]
}
