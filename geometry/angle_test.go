package geometry

import (
	"testing"
)

const (
	TestDegree     = float32(90)
	TestRadians    = 1.5707963267948966
	TestOverDegree = float32(370)
)

var (
	TestOverDegreeResult = Deg(10)
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
	// Normalize
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
}
