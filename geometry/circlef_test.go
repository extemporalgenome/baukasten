package geometry

import (
	"testing"
)

func TestCirclef(t *testing.T) {
	testVec := Vec2(0, 0)
	circle := Circf(Vec2(0, 0), 1.0)
	if !circle.IsInside(testVec) {
		t.Errorf("%v should be in circle %v", testVec, circle)
	}
	testVec = Vec2(0, 1)
	if !circle.IsInside(testVec) {
		t.Errorf("%v should be in circle %v", testVec, circle)
	}
	testVec = Vec2(2, 0)
	if circle.IsInside(testVec) {
		t.Errorf("%v should not be in circle %v", testVec, circle)
	}
}
