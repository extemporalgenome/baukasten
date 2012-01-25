package baukasten

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"

	"code.google.com/p/freetype-go/freetype"
)

type FontContext struct {
	context *freetype.Context
	units   int
}

func OpenFont(name string) (*FontContext, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(b)
	if err != nil {
		return nil, err
	}
	// Context
	c := freetype.NewContext()
	c.SetFont(font)
	return &FontContext{c, font.UnitsPerEm()}, nil
}

func (f *FontContext) Render(text string, width, height int, size float64, color color.Color) image.Image {
	fg := image.NewUniform(color)
	bg := image.Transparent
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	f.context.SetFontSize(size)
	f.context.SetClip(rgba.Bounds())
	f.context.SetDst(rgba)
	f.context.SetSrc(fg)

	// Draw the text.
	pt := freetype.Pt(10, 10+f.context.FUnitToPixelRU(f.units))
	f.context.DrawString(text, pt)

	return rgba
}
