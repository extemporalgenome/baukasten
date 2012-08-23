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
