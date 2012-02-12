package baukasten

import (
	"math"
)

// Pi returns the math.Pi number as a float32.
func Pi() float32 {
	return math.Pi
}

// Similar to math.Acos but takes and returns a float32 instead of a float64.
func Acos(angle float32) float32 {
	return float32(math.Acos(float64(angle)))
}

// Similar to math.Cos but takes and returns a float32 instead of a float64.
func Cos(angle float32) float32 {
	return float32(math.Cos(float64(angle)))
}

// Similar to math.Sin but takes and returns a float32 instead of a float64.
func Sin(angle float32) float32 {
	return float32(math.Sin(float64(angle)))
}

// Similar to math.Sqrt but takes and returns a float32 instead of a float64.
func Sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}

// Similar to math.Signbit but takes and returns a float32 instead of a float64.
func Signbit(value float32) bool {
	return math.Signbit(float64(value))
}

// Similar to math.Tan but takes and returns a float32 instead of a float64.
func Tan(x float32) float32 {
	return float32(math.Tan(float64(x)))
}
