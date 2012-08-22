package geometry

import (
	"github.com/Agon/baukasten/math"
	"strconv"
)

// Vector3 is a three dimensional vector.
type Vector3 struct{ X, Y, Z float32 }

func Vec3(x, y, z float32) Vector3 {
	return Vector3{x, y, z}
}

// String returns a string representation of v like "(3.0,4.1223e12,2.0)".
func (v Vector3) String() string {
	return "(" + strconv.FormatFloat(float64(v.X), 'e', 3, 32) + "," + strconv.FormatFloat(float64(v.Y), 'e', 3, 32) + "," + strconv.FormatFloat(float64(v.Z), 'e', 3, 32) + ")"
}

// Add returns v added on vec.
func (v Vector3) Add(vec Vector3) Vector3 {
	return Vector3{v.X + vec.X, v.Y + vec.Y, v.Z + vec.Z}
}

// Sub returns v subtracted from vec.
func (v Vector3) Sub(vec Vector3) Vector3 {
	return Vector3{v.X - vec.X, v.Y - vec.Y, v.Z - vec.Z}
}

// Mul returns v multiplied by vec.
func (v Vector3) Mul(vec Vector3) Vector3 {
	return Vector3{v.X * vec.X, v.Y * vec.Y, v.Z * vec.Z}
}

// Scaled returns v scaled by scalar.
func (v Vector3) Scaled(scalar float32) Vector3 {
	return Vector3{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

// Normalized returns v normalized.
func (v Vector3) Normalized() Vector3 {
	return v.Scaled(1 / v.Length())
}

// SquaredLength returns v's length squared.
func (v Vector3) SquaredLength() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Length returns v's length.
func (v Vector3) Length() float32 {
	return math.Sqrt(v.SquaredLength())
}

// Dot returns v's dot product with vec.
func (v Vector3) Dot(vec Vector3) float32 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

// Cross returns v's cross product with vec.
func (v Vector3) Cross(vec Vector3) Vector3 {
	return Vector3{
		v.Y*vec.Z - v.Z*vec.Y,
		v.Z*vec.X - v.X*vec.Z,
		v.X*vec.Y - v.Y*vec.X,
	}
}

// AngleBetween returns the angle between v and vec.
func (v Vector3) AngleBetween(vec Vector3) Angle {
	return Angle(math.Acos(v.Dot(vec) / (v.Length() * vec.Length())))
}

// DistanceBetween returns the distance between v and vec.
func (v Vector3) DistanceBetween(vec Vector3) float32 {
	return (v.Sub(vec)).Length()
}
