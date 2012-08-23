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

package matrix

import (
	"github.com/Agon/baukasten/math"
	"github.com/Agon/baukasten/math/geometry"
)

// A 4x4 matrix.
type Matrix4 []float32

// IdentityMatrix returns the identity 4x4 matrix.
func IdentityMatrix() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

// TranslationMatrix returns the translation matrix of x, y, z.
func TranslationMatrix(x, y, z float32) Matrix4 {
	return Matrix4{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1}
}

// ScaleMatrix returns the scalation matrix of x, y, z.
func ScaleMatrix(x, y, z float32) Matrix4 {
	return Matrix4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1}
}

// XRotationMatrix returns the x-axis rotation matrix based on angle x.
func XRotationMatrix(x float32) Matrix4 {
	cos := math.Cos(x)
	sin := math.Sin(x)
	return Matrix4{
		1, 0, 0, 0,
		0, cos, -sin, 0,
		0, sin, cos, 0,
		0, 0, 0, 1}
}

// YRotationMatrix returns the y-axis rotation matrix based on angle y.
func YRotationMatrix(y float32) Matrix4 {
	cos := math.Cos(y)
	sin := math.Sin(y)
	return Matrix4{
		cos, 0, sin, 0,
		0, 1, 0, 0,
		-sin, 0, cos, 0,
		0, 0, 0, 1}
}

// ZRotationMatrix returns the z-axis rotation matrix based on angle z.
func ZRotationMatrix(z float32) Matrix4 {
	cos := math.Cos(z)
	sin := math.Cos(z)
	return Matrix4{
		cos, -sin, 0, 0,
		sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

// RotationMatrix returns the rotation matrix based on vec with a amount.
func RotationMatrix(a float32, vec geometry.Vector3) Matrix4 {
	c := math.Cos(a)
	s := math.Sin(a)
	return Matrix4{
		vec.X*vec.X*(1-c) + c, vec.X*vec.Y*(1-c) - vec.Z*s, vec.X*vec.Z*(1-c) + vec.Y*s, 0,
		vec.X*vec.Y*(1-c) + vec.Z*s, vec.Y*vec.Y*(1-c) + c, vec.Y*vec.Z*(1-c) - vec.X*s, 0,
		vec.X*vec.Z*(1-c) - vec.Y*s, vec.Y*vec.Z*(1-c) + vec.X*s, vec.Z*vec.Z*(1-c) + c, 0,
		0, 0, 0, 1,
	}
}

// ReflectXMatrix returns the reflection matrix for the x-axis.
func ReflectXMatrix() Matrix4 {
	return Matrix4{
		-1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// ReflectYMatrix returns the reflection matrix for the y-axis.
func ReflectYMatrix() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, -1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// ReflectZMatrix returns the reflection matrix for the z-axis.
func ReflectZMatrix() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, 1,
	}
}

// LookAtMatrix returns the view matrix derived from an eye point, a reference point indicating the center of the scene and an up vector.
// Similar to gluLookAt.
func LookAtMatrix(eye, center, up geometry.Vector3) Matrix4 {
	z := (center.Normalized()).Scaled(-1.0)
	dk := center.Cross(up)
	x := dk.Normalized()
	y := z.Cross(x)
	return Matrix4{
		x.X, y.X, z.X, -eye.X,
		x.Y, y.Y, z.Y, -eye.X,
		x.Z, y.Z, z.Z, -eye.X,
		0, 0, 0, 1,
	}
}

// PerspectiveMatrix returns the projection matrix for perspective corrected views.
// Similar to gluPerspective.
func PerspectiveMatrix(fovy, aspect, zNear, zFar float32) Matrix4 {
	if zNear == zFar {
		panic("zFar and zNear must not be the same.")
	}
	f := math.Tan(math.Pi/2 - fovy)
	return Matrix4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (zFar + zNear) / (zNear - zFar), (2 * zFar * zNear) / (zNear - zFar),
		0, 0, -1, 0,
	}
}

// Ortho returns the projection matrix for an orthogonal perspective.
// Similar to gluOrtho.
func Ortho(left, right, bottom, top, near, far float32) Matrix4 {
	tx := -((right + left) / (right - left))
	ty := -((top + bottom) / (top - bottom))
	tz := -((far + near) / (far - near))
	return Matrix4{
		2 / (right - left), 0, 0, tx,
		0, 2 / (top - bottom), 0, ty,
		0, 0, -2 / (far - near), tz,
		0, 0, 0, 1,
	}
}

// Ortho2D returns the projection matrix for an orthogonal perspective with 2D drawing in mind.
// Equal to Ortho(left, right, bottom, top, -1, 1).
// Similar to gluOrtho2D.
func Ortho2D(left, right, bottom, top float32) Matrix4 {
	return Ortho(left, right, bottom, top, -1, 1)
}

// Mul returns m1 multiplied by m2.
func (m1 Matrix4) Mul(m2 Matrix4) Matrix4 {
	return Matrix4{
		m1[0]*m2[0] + m1[1]*m2[4] + m1[2]*m2[8] + m1[3]*m2[12],
		m1[0]*m2[1] + m1[1]*m2[5] + m1[2]*m2[9] + m1[3]*m2[13],
		m1[0]*m2[2] + m1[1]*m2[6] + m1[2]*m2[10] + m1[3]*m2[14],
		m1[0]*m2[3] + m1[1]*m2[7] + m1[2]*m2[11] + m1[3]*m2[15],

		m1[4]*m2[0] + m1[5]*m2[4] + m1[6]*m2[8] + m1[7]*m2[12],
		m1[4]*m2[1] + m1[5]*m2[5] + m1[6]*m2[9] + m1[7]*m2[13],
		m1[4]*m2[2] + m1[5]*m2[6] + m1[6]*m2[10] + m1[7]*m2[14],
		m1[4]*m2[3] + m1[5]*m2[7] + m1[6]*m2[11] + m1[7]*m2[15],

		m1[8]*m2[0] + m1[9]*m2[4] + m1[10]*m2[8] + m1[11]*m2[12],
		m1[8]*m2[1] + m1[9]*m2[5] + m1[10]*m2[9] + m1[11]*m2[13],
		m1[8]*m2[2] + m1[9]*m2[6] + m1[10]*m2[10] + m1[11]*m2[14],
		m1[8]*m2[3] + m1[9]*m2[7] + m1[10]*m2[11] + m1[11]*m2[15],

		m1[12]*m2[0] + m1[13]*m2[4] + m1[14]*m2[8] + m1[15]*m2[12],
		m1[12]*m2[1] + m1[13]*m2[5] + m1[14]*m2[9] + m1[15]*m2[13],
		m1[12]*m2[2] + m1[13]*m2[6] + m1[14]*m2[10] + m1[15]*m2[14],
		m1[12]*m2[3] + m1[13]*m2[7] + m1[14]*m2[11] + m1[15]*m2[15],
	}
}

// Transposed returns m transposed.
// The transposed matrix is used in OpenGL.
func (m Matrix4) Transposed() Matrix4 {
	return Matrix4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}

// Determinant returns m's determinant.
func (m Matrix4) Determinant() float32 {
	a0 := m[0]*m[5] - m[1]*m[4]
	a1 := m[0]*m[6] - m[2]*m[4]
	a2 := m[0]*m[7] - m[3]*m[4]
	a3 := m[1]*m[6] - m[2]*m[5]
	a4 := m[1]*m[7] - m[3]*m[5]
	a5 := m[2]*m[7] - m[3]*m[6]
	b0 := m[8]*m[13] - m[9]*m[12]
	b1 := m[8]*m[14] - m[10]*m[12]
	b2 := m[8]*m[15] - m[11]*m[12]
	b3 := m[9]*m[14] - m[10]*m[13]
	b4 := m[9]*m[15] - m[11]*m[13]
	b5 := m[10]*m[15] - m[11]*m[14]
	return a0*b5 - a1*b4 + a2*b3 + a3*b2 - a4*b1 + a5*b0
}
