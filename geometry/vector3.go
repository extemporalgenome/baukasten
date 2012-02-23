package geometry

import (
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

// ### Manipulation functions ###

func (v *Vector3) Set(x, y, z float32) {
	v.X = x
	v.Y = y
	v.Z = z
}

// Adds another vector
func (v *Vector3) Accumulate(vec Vector3) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
}

// Subtracts another vector
func (v *Vector3) Substract(vec Vector3) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
}

func (v *Vector3) Scale(scalar float32) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}

func (v *Vector3) Normalize() {
	v.Scale(1 / v.Length())
}

// ### Return functions ###

// Returns the vector added with another vector.
func (v Vector3) Add(vec Vector3) Vector3 {
	return Vector3{v.X + vec.X, v.Y + vec.Y, v.Z + vec.Z}
}

// Returns the vector substracted with another vector.
func (v Vector3) Sub(vec Vector3) Vector3 {
	return Vector3{v.X - vec.X, v.Y - vec.Y, v.Z - vec.Z}
}

// Returns the vector multiplied with another vector.
func (v Vector3) Mul(vec Vector3) Vector3 {
	return Vector3{v.X * vec.X, v.Y * vec.Y, v.Z * vec.Z}
}

// Returns the vector scaled.
func (v Vector3) Scaled(scalar float32) Vector3 {
	return Vector3{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

// Returns the vector normalized.
func (v Vector3) Normalized() Vector3 {
	return v.Scaled(1 / v.Length())
}

// Returns the length of the vector.
func (v Vector3) Length() float32 {
	return Sqrt(v.SquaredLength())
}

// Returns the square of the length of the vector.
func (v Vector3) SquaredLength() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Computes and returns the dot product with another vector.
func (v Vector3) Dot(vec Vector3) float32 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

// Returns the cross product with another vector.
func (v Vector3) Cross(vec Vector3) Vector3 {
	return Vector3{
		v.Y*vec.Z - v.Z*vec.Y,
		v.Z*vec.X - v.X*vec.Z,
		v.X*vec.Y - v.Y*vec.X,
	}
}

// Computes and returns the angle between another vector.
func (v Vector3) AngleBetween(vec Vector3) Angle {
	return Angle(Acos(v.Dot(vec) / (v.Length() * vec.Length())))
}

// Computes and returns the distance to another vector.
func (v Vector3) DistanceBetween(vec Vector3) float32 {
	return (v.Sub(vec)).Length()
}
