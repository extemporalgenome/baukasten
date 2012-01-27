package baukasten

import (
	"image/color"
	"testing"
)

func TestConvertColorF(t *testing.T) {
	color := color.White
	r, g, b, a := ConvertColorF(color)
	if r+g+b+a != 4 {
		t.Errorf("%d+%d+%d+%d should be the color.White", r, g, b, a)
	}
}
