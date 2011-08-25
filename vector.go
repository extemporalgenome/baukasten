package baukasten

import (
	"math"
)

type Vector2 struct {
	X, Y float32
}

func NewVector2(x, y float32) *Vector2 {
	return &Vector2{x, y}
}

func (v *Vector2) Copy() Vector2 {
	return Vector2{v.X, v.Y}
}

// Return functions

func (v *Vector2) Add(vec *Vector2) *Vector2 {
	return &Vector2{v.X + vec.X, v.Y + vec.Y}
}

func (v *Vector2) Sub(vec *Vector2) *Vector2 {
	return &Vector2{v.X - vec.X, v.Y - vec.Y}
}

func (v *Vector2) Mul(z float32) *Vector2 {
	return &Vector2{v.X * z, v.Y * z}
}

func (v *Vector2) Div(z float32) *Vector2 {
	return &Vector2{v.X / z, v.Y / z}
}

func (v *Vector2) Normalized() *Vector2 {
	l := 1.0 / v.Magnitude()
	return &Vector2{v.X * l, v.Y * l}
}

// Signbit returns true and true if x and y are negative or negative zero. 
func (v *Vector2) Signbit() (bool, bool) {
	return math.Signbit(float64(v.X)), math.Signbit(float64(v.Y))
}

// Distance between two points.
func (v *Vector2) Distance(vec Vector2) float32 {
	x := v.X - vec.X
	y := v.Y - vec.Y
	return float32(math.Sqrt(float64(x*x + y*y)))
}

// Modify functions
func (v *Vector2) Set(x, y float32) {
	v.X = x
	v.Y = y
}

func (v *Vector2) Accumulate(vec *Vector2) {
	v.X += vec.X
	v.Y += vec.Y
}

func (v *Vector2) Subtract(vec *Vector2) {
	v.X -= vec.X
	v.Y -= vec.Y
}

func (v *Vector2) Scale(z float32) {
	v.X *= z
	v.Y *= z
}

func (v *Vector2) Divide(z float32) {
	v.X /= z
	v.Y /= z
}

func (v *Vector2) LengthSqrt() float32 {
	return v.X*v.X + v.Y*v.Y
}

// Returns the length of the vector.
func (v *Vector2) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.LengthSqrt())))
}

func (v *Vector2) Normalize() {
	l := 1.0 / v.Magnitude()
	v.X *= l
	v.Y *= l
}

// Utility functions

// Computes and returns the dot product.
func DotProduct(a, b Vector2) float32 {
	return a.X*b.X + a.Y*b.Y
}

func AngleBetween(a, b Vector2) float32 {
	return float32(math.Acos(float64(DotProduct(a, b) / (a.Magnitude() * b.Magnitude()))))
}

// ### Vector3 ###

type Vector3 struct {
	X, Y, Z float32
}

func NewVector3(x, y, z float32) *Vector3 {
	return &Vector3{x, y, z}
}

// ### Vector4 ###

type Vector4 struct {
	X, Y, Z, W float32
}

func NewVector4(x, y, z, w float32) *Vector4 {
	return &Vector4{x, y, z, w}
}
