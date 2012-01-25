package baukasten

import (
	"image/color"
)

// Converts a Go type color.Color to red, green, blue and alpha float32 values with a range of 0 to 1.
func ConvertColorF(color color.Color) (r float32, g float32, b float32, a float32) {
	red, green, blue, alpha := color.RGBA()
	if red == 0 {
		r = 0
	} else {
		r = float32(0xFFFF / red)
	}
	if green == 0 {
		g = 0
	} else {
		g = float32(0xFFFF / green)
	}
	if blue == 0 {
		b = 0
	} else {
		b = float32(0xFFFF / blue)
	}
	if alpha == 0 {
		a = 0
	} else {
		a = float32(0xFFFF / alpha)
	}
	return
}
