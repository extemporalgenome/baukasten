package baukasten

import (
	"math"
)

func Acos(angle float32) float32 {
	return float32(math.Acos(float64(angle)))
}

func Cos(angle float32) float32 {
	return float32(math.Cos(float64(angle)))
}

func Sin(angle float32) float32 {
	return float32(math.Sin(float64(angle)))
}

func Sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}

func Signbit(value float32) bool {
	return math.Signbit(float64(value))
}

func Tan(x float32) float32 {
	return float32(math.Tan(float64(x)))
}
