package baukasten

import (
	"strconv"
)

type Vector2 struct{ X, Y float32 }

func Vec2(x, y float32) Vector2 {
	return Vector2{x, y}
}

// String returns a string representation of v like "(3.0,4.1223e12)".
func (v Vector2) String() string {
	return "(" + strconv.FormatFloat(float64(v.X), 'e', 3, 32) + "," + strconv.FormatFloat(float64(v.Y), 'e', 3, 32) + ")"
}

// ### Manipulation functions ###

func (v Vector2) Set(x, y float32) {
	v.X = x
	v.Y = y
}

// Adds another vector
func (v Vector2) Accumulate(vec Vector2) {
	v.X += vec.X
	v.Y += vec.Y
}

// Subtracts another vector
func (v Vector2) Substract(vec Vector2) {
	v.X -= vec.X
	v.Y -= vec.Y
}

func (v Vector2) Scale(scalar float32) {
	v.X *= scalar
	v.Y *= scalar
}

func (v Vector2) Normalize() {
	v.Scale(1 / v.Magnitude())
}

// ### Return functions ###

// Returns the vector added with another vector.
func (v Vector2) Add(vec Vector2) Vector2 {
	return Vector2{v.X + vec.X, v.Y + vec.Y}
}

// Returns the vector substracted with another vector.
func (v Vector2) Sub(vec Vector2) Vector2 {
	return Vector2{v.X - vec.X, v.Y - vec.Y}
}

// Returns the vector multiplied with another vector.
func (v Vector2) Mul(vec Vector2) Vector2 {
	return Vector2{v.X * vec.X, v.Y * vec.Y}
}

// Returns the vector scaled.
func (v Vector2) Scaled(scalar float32) Vector2 {
	return Vector2{v.X * scalar, v.Y * scalar}
}

// Returns the vector normalized.
func (v Vector2) Normalized() Vector2 {
	return v.Scaled(1 / v.Magnitude())
}

// Returns the square of the length of the vector.
func (v Vector2) MagnitudeSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

// Returns the length of the vector.
func (v Vector2) Magnitude() float32 {
	return Sqrt(v.MagnitudeSquared())
}

// Computes and returns the dot product with another vector.
func (v Vector2) DotProduct(vec Vector2) float32 {
	return v.X*vec.X + v.Y*vec.Y
}

// Computes and returns the angle between another vector.
func (v Vector2) AngleBetween(vec Vector2) float32 {
	return Acos(v.DotProduct(vec) / (v.Magnitude() * vec.Magnitude()))
}

// Computes and returns the distance between another vector.
func (v Vector2) DistanceBetween(vec Vector2) float32 {
	return (v.Sub(vec)).Magnitude()
}
