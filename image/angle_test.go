package image

import (
	"testing"
)

const (
	TestDegree     = float32(90)
	TestRadians    = 1.5707963267948966
	TestOverDegree = float32(370.10)
	TestZero       = float32(0)
)

var (
	TestOverDegreeResult = Deg(TestOverDegree - 360.0)
)

func TestAngle(t *testing.T) {
	d := Deg(TestDegree)
	r := Rad(TestRadians)
	if d != r {
		t.Errorf("Deg(%f) should equal Rad(%f)", d, r)
	}
	if d.Degrees() != TestDegree {
		t.Errorf("d.Degrees() %f should equal %f", d.Degrees(), TestDegree)
	}
	if d.Radians() != TestRadians {
		t.Errorf("d.Radians() %f should equal %f", d.Radians(), TestRadians)
	}

	// Vec
	a := Rad(0)
	vec := Vec2(1, 0)
	if a.Vec() != vec {
		t.Errorf("%v.Vec() should result in %v not in %v", a, vec, a.Vec())
	}
	a = Deg(180)
	vec = Vec2(-1, 0)
	if a.Vec().X != vec.X {
		t.Errorf("%v.Vec().X should result in %f not in %f", a, vec.X, a.Vec().X)
	}

	// Zero test
	zeroDegree := Deg(TestZero)
	if zeroDegree.Degrees() != TestZero {
		t.Errorf("zeroDegree.Degrees() %f should equal %f", zeroDegree.Degrees(), TestZero)
	}
	if zeroDegree.Radians() != TestZero {
		t.Errorf("zeroDegree.Radians() %f should equal %f", zeroDegree.Radians(), TestZero)
	}

	// Normalize
	d = Deg(TestDegree)
	if d.Normalized() != d {
		t.Errorf("d.Normalized() %f should equal %f", d.Normalized(), d)
	}
	if d.Normalized180() != d {
		t.Errorf("d.Normalized180() %f should equal %f", d.Normalized180(), d)
	}
	overDegree := Deg(TestOverDegree)
	if overDegree.Normalized() != TestOverDegreeResult {
		t.Errorf("overDegree.Normalized() %f should equal %f", overDegree.Normalized(), TestOverDegreeResult)
	}
	if overDegree.Normalized180() != TestOverDegreeResult {
		t.Errorf("overDegree.Normalized180() %f should equal %f", overDegree.Normalized180(), TestOverDegreeResult)
	}
	if zeroDegree.Normalized() != zeroDegree {
		t.Errorf("zeroDegree.Normalized() %f should equal %f", zeroDegree.Normalized(), zeroDegree)
	}
	if zeroDegree.Normalized180() != zeroDegree {
		t.Errorf("zeroDegree.Normalized180() %f should equal %f", zeroDegree.Normalized180(), zeroDegree)
	}
}
