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

func TestLine2f(t *testing.T) {
	a := Vec2(0, 0)
	b := Vec2(1, 1)
	l := Lin2f(a, b)
	c := Vec2(0, 1)
	d := Vec2(1, 0)
	l2 := Lin2f(c, d)

	// Intersect
	pos, intersect := l.Intersect(l2)
	if !intersect {
		t.Errorf("%v should intersect with %v", l, l2)
	}
	if pos.X != 0.5 && pos.Y != 0.5 {
		t.Errorf("%v should intersect with %v at %v", l, l2, Vec2(0.5, 0.5))
	}

	// IntersectRec
	line := Lin2f(Vec2(0, 0), Vec2(1, 1))
	rec := Rectf(0, 0, 1, 1)
	if intersects, _ := line.IntersectRec(rec); intersects == false {
		t.Errorf("%v should collide with %v", line, rec)
	}
	rec = Rectf(0.5, 0.5, 2, 2)
	if intersects, _ := line.IntersectRec(rec); intersects == false {
		t.Errorf("%v should collide with %v", line, rec)
	}
	rec = Rectf(1.1, 1.1, 2, 2)
	if intersects, _ := line.IntersectRec(rec); intersects {
		t.Errorf("%v should not collide with %v", line, rec)
	}
}
