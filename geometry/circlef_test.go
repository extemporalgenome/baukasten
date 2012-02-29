package geometry

import (
	"testing"
)

func TestCirclef(t *testing.T) {
	testVec := Vec2(0, 0)
	circle := Circf(Vec2(0, 0), 1.0)

	// IsInside
	if !circle.IsInside(testVec) {
		t.Errorf("%v should be in circle %v", testVec, circle)
	}
	testVec = Vec2(0, 1)
	if circle.IsInside(testVec) == false { // circle contacts point
		t.Errorf("%v should be in circle %v", testVec, circle)
	}
	testVec = Vec2(2, 0)
	if circle.IsInside(testVec) {
		t.Errorf("%v should not be in circle %v", testVec, circle)
	}

	// IsCircleCollision
	circle1 := Circf(Vec2(0, 1), 1.0)
	if circle.IsCircleCollision(circle1) == false {
		t.Errorf("%v should collide with %v", circle, circle1)
	}
	circle1 = Circf(Vec2(1, 1), 1.0) // circle contacts circle1
	if circle.IsCircleCollision(circle1) == false {
		t.Errorf("%v should collide with %v", circle, circle1)
	}
	circle1 = Circf(Vec2(10, 10), 1.0)
	if circle.IsCircleCollision(circle1) {
		t.Errorf("%v should not collide with %v", circle, circle1)
	}

	// IsLineCollision
	line := Lin2f(Vec2(0, 0), Vec2(1, 1))
	circle = Circf(Vec2(0, 0), 1.0)
	if circle.IsLineCollision(line) == false {
		t.Errorf("%v should collide with %v", circle, line)
	}
	line = Lin2f(Vec2(-1, 1), Vec2(1, 1))
	if circle.IsLineCollision(line) == false { // circle contacts with line
		t.Errorf("%v should collide with %v", circle, line)
	}
	line = Lin2f(Vec2(-1, 1.1), Vec2(1, 1.1))
	if circle.IsLineCollision(line) {
		t.Errorf("%v should not collide with %v", circle, line)
	}

	// IsRectangleCollision
	rec := RectF(0, 0, 1, 1)
	circle = Circf(Vec2(0, 0), 1.0)
	if circle.IsRectangleCollision(rec) == false {
		t.Errorf("%v should collide with %v", circle, rec)
	}
	rec = RectF(0, 1, 2, 2)
	circle = Circf(Vec2(0, 0), 1.0)
	if circle.IsRectangleCollision(rec) == false { // circle contacts rec
		t.Errorf("%v should collide with %v", circle, rec)
	}
	rec = RectF(10, 10, 12, 12)
	circle = Circf(Vec2(1, 1), 2.0)
	if circle.IsRectangleCollision(rec) {
		t.Errorf("%v should not collide with %v", circle, rec)
	}
}
