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
	"github.com/Agon/baukasten/math"
	"strconv"
)

// Vector2 is a two dimensional vector.
type Vector2 struct{ X, Y float32 }

func Vec2(x, y float32) Vector2 {
	return Vector2{x, y}
}

// String returns a string representation of v like "(3.0,4.1223e12)".
func (v *Vector2) String() string {
	return "(" + strconv.FormatFloat(float64(v.X), 'e', 3, 32) + "," + strconv.FormatFloat(float64(v.Y), 'e', 3, 32) + ")"
}

// Add returns v added on vec.
func (v Vector2) Add(vec Vector2) Vector2 {
	return Vector2{v.X + vec.X, v.Y + vec.Y}
}

// Sub returns v subtracted from vec.
func (v Vector2) Sub(vec Vector2) Vector2 {
	return Vector2{v.X - vec.X, v.Y - vec.Y}
}

// Mul returns v multiplied by vec.
func (v Vector2) Mul(vec Vector2) Vector2 {
	return Vector2{v.X * vec.X, v.Y * vec.Y}
}

// Scaled returns v scaled by scalar.
func (v Vector2) Scaled(scalar float32) Vector2 {
	return Vector2{v.X * scalar, v.Y * scalar}
}

// Normalized returns v normalized.
func (v Vector2) Normalized() Vector2 {
	return v.Scaled(1 / v.Length())
}

// SquaredLength returns v's length squared.
func (v Vector2) SquaredLength() float32 {
	return v.X*v.X + v.Y*v.Y
}

// Length returns v's length.
func (v Vector2) Length() float32 {
	return math.Sqrt(v.SquaredLength())
}

// Dot returns v's dot product with vec.
func (v Vector2) Dot(vec Vector2) float32 {
	return v.X*vec.X + v.Y*vec.Y
}

// InCirc returns whether v is in c. 
func (v Vector2) InCirc(c Circlef) bool {
	return c.Radius >= c.Position.DistanceBetween(v)
}

// InRec returns whether v is in r.
func (v Vector2) InRec(r Rectanglef) bool {
	return r.Min.X <= v.X && v.X < r.Max.X &&
		r.Min.Y <= v.Y && v.Y < r.Max.Y
}

// AngleBetween returns the angle between v and vec.
func (v Vector2) AngleBetween(vec Vector2) Angle {
	return Angle(math.Acos(v.Dot(vec) / (v.Length() * vec.Length())))
}

// DistanceBetween returns the distance between v and vec.
func (v Vector2) DistanceBetween(vec Vector2) float32 {
	return v.Sub(vec).Length()
}
