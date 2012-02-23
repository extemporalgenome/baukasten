package geometry

import (
	"strconv"
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

// ### Manipulation functions ###

func (v *Vector4) Set(x, y, z, w float32) {
	v.X = x
	v.Y = y
	v.Z = z
	v.W = w
}

// Adds another vector
func (v *Vector4) Accumulate(vec Vector4) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
	v.W += vec.W
}

// Subtracts another vector
func (v *Vector4) Substract(vec Vector4) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
	v.W -= vec.W
}

func (v *Vector4) Scale(scalar float32) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
	v.W *= scalar
}

func (v *Vector4) Normalize() {
	v.Scale(1 / v.Length())
}

// ### Return functions ###

// Returns the vector added with another vector.
func (v Vector4) Add(vec Vector4) Vector4 {
	return Vector4{v.X + vec.X, v.Y + vec.Y, v.Z + vec.Z, v.W + vec.W}
}

// Returns the vector substracted with another vector.
func (v Vector4) Sub(vec Vector4) Vector4 {
	return Vector4{v.X - vec.X, v.Y - vec.Y, v.Z - vec.Z, v.W - vec.W}
}

// Returns the vector multiplied with another vector.
func (v Vector4) Mul(vec Vector4) Vector4 {
	return Vector4{v.X * vec.X, v.Y * vec.Y, v.Z * vec.Z, v.W * vec.W}
}

// Returns the vector scaled.
func (v Vector4) Scaled(scalar float32) Vector4 {
	return Vector4{v.X * scalar, v.Y * scalar, v.Z * scalar, v.W * scalar}
}

// Returns the vector normalized.
func (v Vector4) Normalized() Vector4 {
	return v.Scaled(1 / v.Length())
}

// Returns the length of the vector.
func (v Vector4) Length() float32 {
	return Sqrt(v.SquaredLength())
}

// Returns the square of the length of the vector.
// Uses X,Y,Z. W is ignored.
func (v Vector4) SquaredLength() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Computes and returns the dot product with another vector.
// Uses X,Y,Z. W is ignored.
func (v Vector4) Dot(vec Vector4) float32 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

// Returns the cross product with another vector.
// Uses X,Y,Z. W is ignored.
func (v Vector4) Cross(vec Vector4) Vector4 {
	return Vector4{
		v.Y*vec.Z - v.Z*vec.Y,
		v.Z*vec.X - v.X*vec.Z,
		v.X*vec.Y - v.Y*vec.X,
		0,
	}
}

// Computes and returns the angle between another vector.
func (v Vector4) AngleBetween(vec Vector4) Angle {
	return Angle(Acos(v.Dot(vec) / (v.Length() * vec.Length())))
}

// Computes and returns the distance to another vector.
func (v Vector4) DistanceBetween(vec Vector4) float32 {
	return (v.Sub(vec)).Length()
}
