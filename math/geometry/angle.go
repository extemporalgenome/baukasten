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
)

type Angle float32

// Deg constructs an Angle based on a degrees.
func Deg(a float32) Angle {
	return Angle(a)
}

// Rad constructs an Angle based on a radians.
func Rad(a float32) Angle {
	return Angle(a * 180.0 / math.Pi)
}

// Degrees returns a as a float32 degree.
func (a Angle) Degrees() float32 {
	return float32(a)
}

// Radians returns a as a float32 radian.
func (a Angle) Radians() float32 {
	return float32(a) * math.Pi / 180.0
}

// Normalized returns a normalized to a range between -360 and 360.
func (a Angle) Normalized() Angle {
	return Angle(math.Mod(float32(a), 360.0))
}

// Normalized180 returns a normalized to a range between -180 and 180.
func (a Angle) Normalized180() Angle {
	return Angle(math.Mod(float32(a), 180.0))
}

// Vec returns a's vector representation aligned to the x-axis.
func (a Angle) Vec() Vector2 {
	return Vec2(math.Cos(a.Radians()), math.Sin(a.Radians()))
}
