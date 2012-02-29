package geometry

import (
	"testing"
)

func TestLine2f(t *testing.T) {
	a := Vec2(0, 0)
	b := Vec2(1, 1)
	l := Lin2f(a, b)
	c := Vec2(0, 1)
	d := Vec2(1, 0)
	l2 := Lin2f(c, d)

	pos, intersect := l.Intersection(l2)
	if !intersect {
		t.Errorf("%v should intersect with %v", l, l2)
	}
	if pos.X != 0.5 && pos.Y != 0.5 {
		t.Errorf("%v should intersect with %v at %v", l, l2, Vec2(0.5, 0.5))
	}

	// IsRectangleCollision
	line := Lin2f(Vec2(0, 0), Vec2(1, 1))
	rec := RectF(0, 0, 1, 1)
	if line.IsRectangleCollision(rec) == false {
		t.Errorf("%v should collide with %v", line, rec)
	}
	rec = RectF(0.5, 0.5, 2, 2)
	if line.IsRectangleCollision(rec) == false {
		t.Errorf("%v should collide with %v", line, rec)
	}
	rec = RectF(1.1, 1.1, 2, 2)
	if line.IsRectangleCollision(rec) {
		t.Errorf("%v should not collide with %v", line, rec)
	}
}
