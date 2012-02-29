package geometry

import (
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
	return Sqrt(v.SquaredLength())
}

// Dot returns v's dot product with vec.
func (v Vector2) Dot(vec Vector2) float32 {
	return v.X*vec.X + v.Y*vec.Y
}

// AngleBetween returns the angle between v and vec.
func (v Vector2) AngleBetween(vec Vector2) Angle {
	return Angle(Acos(v.Dot(vec) / (v.Length() * vec.Length())))
}

// DistanceBetween returns the distance between v and vec.
func (v Vector2) DistanceBetween(vec Vector2) float32 {
	return v.Sub(vec).Length()
}
