package baukasten

import (
	"math"
)

func Cos(angle float32) float32 {
	return float32(math.Cos(float64(angle)))
}

func Sin(angle float32) float32 {
	return float32(math.Sin(float64(angle)))
}

func Sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}
