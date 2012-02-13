package geometry

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

func TestOverflow(t *testing.T) {
	c := ConvertFColor(1, 1, 1, 10)
	_, _, _, a := c.RGBA()
	if a != 0xFFFF {
		t.Errorf("Input should reach a maximum of %d not %d", 0xFFFF, a)
	}
}

func TestUnderflow(t *testing.T) {
	c := ConvertFColor(1, 1, 1, -10)
	_, _, _, a := c.RGBA()
	if a != 0 {
		t.Errorf("Input should reach a minimum of %d not %d", 0, a)
	}
}
