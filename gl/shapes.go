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

package gl

import (
	"github.com/Agon/baukasten/math"
	"github.com/Agon/baukasten/math/geometry"
	"image/color"
)

// DrawRectangle draws a Rectanglef as two triangles.
func DrawRectangle(color color.Color, r geometry.Rectanglef) {
	DrawTriangles(color, r.Min, geometry.Vector2{r.Min.X, r.Max.Y}, geometry.Vector2{r.Max.X, r.Min.Y}, geometry.Vector2{r.Max.X, r.Min.Y}, geometry.Vector2{r.Min.X, r.Max.Y}, r.Max)
}

// DrawCircle draws a circle centered at v with a radius of r, with n number of points in color c.
func DrawCircle(c color.Color, r float32, n int, v geometry.Vector2) {
	vectors := make([]geometry.Vector2, n)
	for i := 0; i < n; i++ {
		degInRad := (360 / float32(i)) * math.Pi / 180
		vectors[i] = v.Add(geometry.Vector2{math.Cos(degInRad) * r, math.Sin(degInRad) * r})
	}
	DrawLineLoop(c, vectors...)
}
