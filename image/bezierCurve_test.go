package image

import (
	"testing"
)

func TestBezierCurve(t *testing.T) {
	a := Vec2(0, 0)
	b := Vec2(1, 1)
	c := Vec2(2, 0)
	d := Vec2(3, 2)
	curve := BCurve(a, b, c, d)

	// input 0 (min)
	vec := curve.RelativePoint(0)
	if vec != a {
		t.Errorf("RelativePoint(0) should return %v not %v, since the curve should start at the first point", a, vec)
	}

	// input 0.5
	vec = curve.RelativePoint(0.5)
	midle := Vec2(1.5, 0.625)
	if vec != midle {
		t.Errorf("RelativePoint(0) should return %v not %v, since the curve should start at the first point", midle, vec)
	}

	// input 1 (max)
	vec = curve.RelativePoint(1)
	if vec != d {
		t.Errorf("RelativePoint(1) should return %v not %v, since the curve should start at the first point", d, vec)
	}
}
