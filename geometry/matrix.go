package baukasten

import (
	"math"
)

// A 4x4 matrix.
type Matrix4 []float32

// The identity 4x4 matrix.
func IdentityMatrix() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

// Creates a 4x4 matrix which represents a translation.
func TranslationMatrix(x, y, z float32) Matrix4 {
	return Matrix4{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1}
}

// Create a 4x4 matrix which represents a scaling operation.
func ScaleMatrix(x, y, z float32) Matrix4 {
	return Matrix4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1}
}

// Create a 4x4 matrix which rotates about the x-axis.
func XRotationMatrix(angle float32) Matrix4 {
	cos := Cos(angle)
	sin := Sin(angle)
	return Matrix4{
		1, 0, 0, 0,
		0, cos, -sin, 0,
		0, sin, cos, 0,
		0, 0, 0, 1}
}

// Creates a 4x4 matrix which rotates about the y-axis.
func YRotationMatrix(angle float32) Matrix4 {
	cos := Cos(angle)
	sin := Sin(angle)
	return Matrix4{
		cos, 0, sin, 0,
		0, 1, 0, 0,
		-sin, 0, cos, 0,
		0, 0, 0, 1}
}

// Creates a 4x4 matrix which rotates about the z-axis.
func ZRotationMatrix(angle float32) Matrix4 {
	cos := Cos(angle)
	sin := Cos(angle)
	return Matrix4{
		cos, -sin, 0, 0,
		sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

// Creates a 4x4 matrix which rotates in the direction of vec with the amount of angle.
func RotationMatrix(angle float32, vec Vector3) Matrix4 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	return Matrix4{
		vec.X*vec.X*(1-c) + c, vec.X*vec.Y*(1-c) - vec.Z*s, vec.X*vec.Z*(1-c) + vec.Y*s, 0,
		vec.X*vec.Y*(1-c) + vec.Z*s, vec.Y*vec.Y*(1-c) + c, vec.Y*vec.Z*(1-c) - vec.X*s, 0,
		vec.X*vec.Z*(1-c) - vec.Y*s, vec.Y*vec.Z*(1-c) + vec.X*s, vec.Z*vec.Z*(1-c) + c, 0,
		0, 0, 0, 1,
	}
}

// Creates a 4x4 matrix which reflects x coordinates.
func ReflectXMatrix() Matrix4 {
	return Matrix4{
		-1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Creates a 4x4 matrix which reflects y coordinates.
func ReflectYMatrix() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, -1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Creates a 4x4 matrix which reflects z coordinates.
func ReflectZMatrix() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, 1,
	}
}

// Creates a 4x4 view matrix derived from an eye point, a reference point indicating the center of the scene and an up vector.
// Similar to gluLookAt.
func LookAtMatrix(eye, center, up Vector3) Matrix4 {
	z := (center.Normalized()).Scaled(-1.0)
	dk := center.CrossProduct(up)
	x := dk.Normalized()
	y := z.CrossProduct(x)
	return Matrix4{
		x.X, y.X, z.X, -eye.X,
		x.Y, y.Y, z.Y, -eye.X,
		x.Z, y.Z, z.Z, -eye.X,
		0, 0, 0, 1,
	}
}

// Creates a 4x4 projection matrix for perspective corrected views.
// Similar to gluPerspective.
func PerspectiveMatrix(fovy, aspect, zNear, zFar float32) Matrix4 {
	if zNear == zFar {
		panic("zFar and zNear must not be the same.")
	}
	f := float32(math.Tan(math.Pi/2 - float64(fovy)))
	return Matrix4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (zFar + zNear) / (zNear - zFar), (2 * zFar * zNear) / (zNear - zFar),
		0, 0, -1, 0,
	}
}

// Creates a 4x4 projection matrix for an orthogonal perspective.
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

// Creates a 4x4 projection matrix for an orthogonal perspective with 2D drawing in mind.
// Equal to Ortho(left, right, bottom, top, -1, 1).
// Similar to gluOrtho2D.
func Ortho2D(left, right, bottom, top float32) Matrix4 {
	return Ortho(left, right, bottom, top, -1, 1)
}

func (m1 Matrix4) Multiplied(m2 Matrix4) Matrix4 {
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

func (m Matrix4) Transposed() Matrix4 {
	return Matrix4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}

/*
func MakePerspectiveMatrix(fovy, aspect, zNear, zFar float32) Matrix4 {
	f := 1 / float32(math.Tan(float64(fovy/2)))
	a := 1 / (zNear - zFar)
	return Matrix4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (zFar + zNear) * a, 2 * zFar * zNear * a,
		0, 0, -1, 0,
	}
}
*/
