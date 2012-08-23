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

type Circlef struct {
	Position Vector2
	Radius   float32
}

func Circf(position Vector2, radius float32) Circlef {
	return Circlef{Position: position, Radius: radius}
}

// Intersect returns true if c intersects c1.
func (c Circlef) Intersect(c1 Circlef) bool {
	return c.Radius+c1.Radius >= c.Position.DistanceBetween(c1.Position)
}

// IntersectLine returns true if c intersects line.
func (c Circlef) IntersectLine(line Line2f) bool {
	closest := closestPointOnSeg(line, c.Position)
	dist := c.Position.Sub(closest)
	return dist.Length() <= c.Radius
}

// IntersectRec returns true if c intersects rec.
func (c Circlef) IntersectRec(rec Rectanglef) bool {
	return c.Position.InRec(rec) || c.IntersectLine(rec.Top()) || c.IntersectLine(rec.Bottom()) || c.IntersectLine(rec.Left()) || c.IntersectLine(rec.Right())
}

// closestPointOnSeg returns the closest point towards vec from line.
func closestPointOnSeg(line Line2f, vec Vector2) Vector2 {
	seg := line.Q.Sub(line.P)
	pt := vec.Sub(line.P)
	if seg.Length() <= 0 {
		panic("Invalid segment length")
	}
	segUnit := seg.Normalized()
	projLength := pt.Dot(segUnit)
	if projLength <= 0 {
		return line.P
	}
	if projLength >= seg.Length() {
		return line.Q
	}
	proj := segUnit.Scaled(projLength)
	return proj.Add(line.P)
}
