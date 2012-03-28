package freetype

import (
	"image"
	"image/color"
	"image/draw"

	"code.google.com/p/freetype-go/freetype"
)

type FontContext struct {
	context *freetype.Context
	units   int
}

func (f *FontContext) Render(text string, width, height int, size float64, c color.Color) image.Image {
	fg := image.NewUniform(c)
	bg := image.Transparent
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	f.context.SetFontSize(size)
	f.context.SetClip(rgba.Bounds())
	f.context.SetDst(rgba)
	f.context.SetSrc(fg)

	// Draw the text.
	pt := freetype.Pt(0, f.context.FUnitToPixelRU(f.units))
	f.context.DrawString(text, pt)

	return rgba
}
