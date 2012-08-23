// baukasten - Toolkit for OpenGL
// 
// Copyright (c) 2012, Marcel Hauf <marcel.hauf@googlemail.com>
// 
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met: 
// 
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer. 
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution. 
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package geometry

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
