package image

import (
	"strconv"
	"github.com/Agon/baukasten/math"
)

// Vector4 is a four dimensional vector.
type Vector4 struct{ X, Y, Z, W float32 }

func Vec4(x, y, z, w float32) Vector4 {
	return Vector4{x, y, z, w}
}

// String returns a string representation of v like "(3.0,4.1223e12,2.0,22.0)".
func (v Vector4) String() string {
	return "(" + strconv.FormatFloat(float64(v.X), 'e', 3, 32) + "," + strconv.FormatFloat(float64(v.Y), 'e', 3, 32) + "," + strconv.FormatFloat(float64(v.Z), 'e', 3, 32) + "," + strconv.FormatFloat(float64(v.W), 'e', 3, 32) + ")"
}

// ### Return functions ###

// Add returns v added on vec.
func (v Vector4) Add(vec Vector4) Vector4 {
	return Vector4{v.X + vec.X, v.Y + vec.Y, v.Z + vec.Z, v.W + vec.W}
}

// Sub returns v subtracted from vec.
func (v Vector4) Sub(vec Vector4) Vector4 {
	return Vector4{v.X - vec.X, v.Y - vec.Y, v.Z - vec.Z, v.W - vec.W}
}

// Mul returns v multiplied by vec.
func (v Vector4) Mul(vec Vector4) Vector4 {
	return Vector4{v.X * vec.X, v.Y * vec.Y, v.Z * vec.Z, v.W * vec.W}
}

// Scaled returns v scaled by scalar.
func (v Vector4) Scaled(scalar float32) Vector4 {
	return Vector4{v.X * scalar, v.Y * scalar, v.Z * scalar, v.W * scalar}
}

// Normalized returns v normalized.
func (v Vector4) Normalized() Vector4 {
	return v.Scaled(1 / v.Length())
}

// SquaredLength returns v's length squared.
// Uses X,Y,Z. W is ignored.
func (v Vector4) SquaredLength() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Length returns v's length.
func (v Vector4) Length() float32 {
	return math.Sqrt(v.SquaredLength())
}

// Dot returns v's dot product with vec.
// Uses X,Y,Z. W is ignored.
func (v Vector4) Dot(vec Vector4) float32 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

// Cross returns v's cross product with vec.
// Uses X,Y,Z. W is ignored.
func (v Vector4) Cross(vec Vector4) Vector4 {
	return Vector4{
		v.Y*vec.Z - v.Z*vec.Y,
		v.Z*vec.X - v.X*vec.Z,
		v.X*vec.Y - v.Y*vec.X,
		0,
	}
}

// AngleBetween returns the angle between v and vec.
// Uses X,Y,Z. W is ignored.
func (v Vector4) AngleBetween(vec Vector4) Angle {
	return Angle(math.Acos(v.Dot(vec) / (v.Length() * vec.Length())))
}

// DistanceBetween returns the distance between v and vec.
// Uses X,Y,Z. W is ignored.
func (v Vector4) DistanceBetween(vec Vector4) float32 {
	return (v.Sub(vec)).Length()
}
