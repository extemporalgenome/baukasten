package baukasten

import (
	"image/color"
	"testing"
)

func TestConvertColorF(t *testing.T) {
	color := color.White
	r, g, b, a := ConvertColorF(color)
	if r+g+b+a != 4 {
		t.Errorf("%f+%f+%f+%f should be the color.White", r, g, b, a)
	}
}

func TestConvertFColor(t *testing.T) {
	c := ConvertFColor(1, 1, 1, 1)
	r, g, b, a := c.RGBA()
	if r+g+b+a != 0xFFFF*4 {
		t.Errorf("%d+%d+%d+%d should be the color.White(0xFFFF)", r, g, b, a)
	}
}
