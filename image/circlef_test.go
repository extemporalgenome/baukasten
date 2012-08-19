package image

import (
	"testing"
)

func TestCirclef(t *testing.T) {
	circle := Circf(Vec2(0, 0), 1.0)

	// Intersect
	circle1 := Circf(Vec2(0, 1), 1.0)
	if circle.Intersect(circle1) == false {
		t.Errorf("%v should collide with %v", circle, circle1)
	}
	circle1 = Circf(Vec2(1, 1), 1.0) // circle contacts circle1
	if circle.Intersect(circle1) == false {
		t.Errorf("%v should collide with %v", circle, circle1)
	}
	circle1 = Circf(Vec2(10, 10), 1.0)
	if circle.Intersect(circle1) {
		t.Errorf("%v should not collide with %v", circle, circle1)
	}

	// IntersectLine
	line := Lin2f(Vec2(0, 0), Vec2(1, 1))
	circle = Circf(Vec2(0, 0), 1.0)
	if circle.IntersectLine(line) == false {
		t.Errorf("%v should collide with %v", circle, line)
	}
	line = Lin2f(Vec2(-1, 1), Vec2(1, 1))
	if circle.IntersectLine(line) == false { // circle contacts with line
		t.Errorf("%v should collide with %v", circle, line)
	}
	line = Lin2f(Vec2(-1, 1.1), Vec2(1, 1.1))
	if circle.IntersectLine(line) {
		t.Errorf("%v should not collide with %v", circle, line)
	}

	// IntersectRec
	rec := Rectf(0, 0, 1, 1)
	circle = Circf(Vec2(0, 0), 1.0)
	if circle.IntersectRec(rec) == false {
		t.Errorf("%v should collide with %v", circle, rec)
	}
	rec = Rectf(0, 1, 2, 2)
	circle = Circf(Vec2(0, 0), 1.0)
	if circle.IntersectRec(rec) == false { // circle contacts rec
		t.Errorf("%v should collide with %v", circle, rec)
	}
	rec = Rectf(10, 10, 12, 12)
	circle = Circf(Vec2(1, 1), 2.0)
	if circle.IntersectRec(rec) {
		t.Errorf("%v should not collide with %v", circle, rec)
	}
}
