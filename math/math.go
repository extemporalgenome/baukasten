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

package math

import (
	"math"
)

// Pi is math.Pi but as a float32 type
const Pi = float32(math.Pi)

// Similar to math.Acos but takes and returns a float32 instead of a float64.
func Acos(angle float32) float32 {
	return float32(math.Acos(float64(angle)))
}

// Similar to math.Cos but takes and returns a float32 instead of a float64.
func Cos(angle float32) float32 {
	return float32(math.Cos(float64(angle)))
}

// Similar to math.Sin but takes and returns a float32 instead of a float64.
func Sin(angle float32) float32 {
	return float32(math.Sin(float64(angle)))
}

// Similar to math.Sqrt but takes and returns a float32 instead of a float64.
func Sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}

// Similar to math.Signbit but takes and returns a float32 instead of a float64.
func Signbit(value float32) bool {
	return math.Signbit(float64(value))
}

// Similar to math.Tan but takes and returns a float32 instead of a float64.
func Tan(x float32) float32 {
	return float32(math.Tan(float64(x)))
}

// Similar to math.Mod but takes and returns a float32 instead of a float64.
func Mod(x, y float32) float32 {
	return float32(math.Mod(float64(x), float64(y)))
}
