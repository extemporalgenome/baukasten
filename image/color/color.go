// baukasten - Toolkit for OpenGL
// 
// Copyright (c) 2012, Marcel Hauf <marcel.hauf@googlemail.com>
// 
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met: 
// 
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer. 
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution. 
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package color

import (
	gl "github.com/chsc/gogl/gl33"
	"image/color"
)

// Color names taken from http://www.w3.org/TR/SVG/types.html#ColorKeywords
var (
	Aliceblue            = color.RGBA{240, 248, 255, 255}
	Antiquewhite         = color.RGBA{250, 235, 215, 255}
	Aqua                 = color.RGBA{0, 255, 255, 255}
	Aquamarine           = color.RGBA{127, 255, 212, 255}
	Azure                = color.RGBA{240, 255, 255, 255}
	Beige                = color.RGBA{245, 245, 220, 255}
	Bisque               = color.RGBA{255, 228, 196, 255}
	Black                = color.RGBA{0, 0, 0, 255}
	Blanchedalmond       = color.RGBA{255, 235, 205, 255}
	Blue                 = color.RGBA{0, 0, 255, 255}
	Blueviolet           = color.RGBA{138, 43, 226, 255}
	Brown                = color.RGBA{165, 42, 42, 255}
	Burlywood            = color.RGBA{222, 184, 135, 255}
	Cadetblue            = color.RGBA{95, 158, 160, 255}
	Chartreuse           = color.RGBA{127, 255, 0, 255}
	Chocolate            = color.RGBA{210, 105, 30, 255}
	Coral                = color.RGBA{255, 127, 80, 255}
	Cornflowerblue       = color.RGBA{100, 149, 237, 255}
	Cornsilk             = color.RGBA{255, 248, 220, 255}
	Crimson              = color.RGBA{220, 20, 60, 255}
	Cyan                 = color.RGBA{0, 255, 255, 255}
	Darkblue             = color.RGBA{0, 0, 139, 255}
	Darkcyan             = color.RGBA{0, 139, 139, 255}
	Darkgoldenrod        = color.RGBA{184, 134, 11, 255}
	Darkgray             = color.RGBA{169, 169, 169, 255}
	Darkgreen            = color.RGBA{0, 100, 0, 255}
	Darkgrey             = color.RGBA{169, 169, 169, 255}
	Darkkhaki            = color.RGBA{189, 183, 107, 255}
	Darkmagenta          = color.RGBA{139, 0, 139, 255}
	Darkolivegreen       = color.RGBA{85, 107, 47, 255}
	Darkorange           = color.RGBA{255, 140, 0, 255}
	Darkorchid           = color.RGBA{153, 50, 204, 255}
	Darkred              = color.RGBA{139, 0, 0, 255}
	Darksalmon           = color.RGBA{233, 150, 122, 255}
	Darkseagreen         = color.RGBA{143, 188, 143, 255}
	Darkslateblue        = color.RGBA{72, 61, 139, 255}
	Darkslategray        = color.RGBA{47, 79, 79, 255}
	Darkslategrey        = color.RGBA{47, 79, 79, 255}
	Darkturquoise        = color.RGBA{0, 206, 209, 255}
	Darkviolet           = color.RGBA{148, 0, 211, 255}
	Deeppink             = color.RGBA{255, 20, 147, 255}
	Deepskyblue          = color.RGBA{0, 191, 255, 255}
	Dimgray              = color.RGBA{105, 105, 105, 255}
	Dimgrey              = color.RGBA{105, 105, 105, 255}
	Dodgerblue           = color.RGBA{30, 144, 255, 255}
	Firebrick            = color.RGBA{178, 34, 34, 255}
	Floralwhite          = color.RGBA{255, 250, 240, 255}
	Forestgreen          = color.RGBA{34, 139, 34, 255}
	Fuchsia              = color.RGBA{255, 0, 255, 255}
	Gainsboro            = color.RGBA{220, 220, 220, 255}
	Ghostwhite           = color.RGBA{248, 248, 255, 255}
	Gold                 = color.RGBA{255, 215, 0, 255}
	Goldenrod            = color.RGBA{218, 165, 32, 255}
	Gray                 = color.RGBA{128, 128, 128, 255}
	Grey                 = color.RGBA{128, 128, 128, 255}
	Green                = color.RGBA{0, 128, 0, 255}
	Greenyellow          = color.RGBA{173, 255, 47, 255}
	Honeydew             = color.RGBA{240, 255, 240, 255}
	Hotpink              = color.RGBA{255, 105, 180, 255}
	Indianred            = color.RGBA{205, 92, 92, 255}
	Indigo               = color.RGBA{75, 0, 130, 255}
	Ivory                = color.RGBA{255, 255, 240, 255}
	Khaki                = color.RGBA{240, 230, 140, 255}
	Lavender             = color.RGBA{230, 230, 250, 255}
	Lavenderblush        = color.RGBA{255, 240, 245, 255}
	Lawngreen            = color.RGBA{124, 252, 0, 255}
	Lemonchiffon         = color.RGBA{255, 250, 205, 255}
	Lightblue            = color.RGBA{173, 216, 230, 255}
	Lightcoral           = color.RGBA{240, 128, 128, 255}
	Lightcyan            = color.RGBA{224, 255, 255, 255}
	Lightgoldenrodyellow = color.RGBA{250, 250, 210, 255}
	Lightgray            = color.RGBA{211, 211, 211, 255}
	Lightgreen           = color.RGBA{144, 238, 144, 255}
	Lightgrey            = color.RGBA{211, 211, 211, 255}
	Lightpink            = color.RGBA{255, 182, 193, 255}
	Lightsalmon          = color.RGBA{255, 160, 122, 255}
	Lightseagreen        = color.RGBA{32, 178, 170, 255}
	Lightskyblue         = color.RGBA{135, 206, 250, 255}
	Lightslategray       = color.RGBA{119, 136, 153, 255}
	Lightslategrey       = color.RGBA{119, 136, 153, 255}
	Lightsteelblue       = color.RGBA{176, 196, 222, 255}
	Lightyellow          = color.RGBA{255, 255, 224, 255}
	Lime                 = color.RGBA{0, 255, 0, 255}
	Limegreen            = color.RGBA{50, 205, 50, 255}
	Linen                = color.RGBA{250, 240, 230, 255}
	Magenta              = color.RGBA{255, 0, 255, 255}
	Maroon               = color.RGBA{128, 0, 0, 255}
	Mediumaquamarine     = color.RGBA{102, 205, 170, 255}
	Mediumblue           = color.RGBA{0, 0, 205, 255}
	Mediumorchid         = color.RGBA{186, 85, 211, 255}
	Mediumpurple         = color.RGBA{147, 112, 219, 255}
	Mediumseagreen       = color.RGBA{60, 179, 113, 255}
	Mediumslateblue      = color.RGBA{123, 104, 238, 255}
	Mediumspringgreen    = color.RGBA{0, 250, 154, 255}
	Mediumturquoise      = color.RGBA{72, 209, 204, 255}
	Mediumvioletred      = color.RGBA{199, 21, 133, 255}
	Midnightblue         = color.RGBA{25, 25, 112, 255}
	Mintcream            = color.RGBA{245, 255, 250, 255}
	Mistyrose            = color.RGBA{255, 228, 225, 255}
	Moccasin             = color.RGBA{255, 228, 181, 255}
	Navajowhite          = color.RGBA{255, 222, 173, 255}
	Navy                 = color.RGBA{0, 0, 128, 255}
	Oldlace              = color.RGBA{253, 245, 230, 255}
	Olive                = color.RGBA{128, 128, 0, 255}
	Olivedrab            = color.RGBA{107, 142, 35, 255}
	Orange               = color.RGBA{255, 165, 0, 255}
	Orangered            = color.RGBA{255, 69, 0, 255}
	Orchid               = color.RGBA{218, 112, 214, 255}
	Palegoldenrod        = color.RGBA{238, 232, 170, 255}
	Palegreen            = color.RGBA{152, 251, 152, 255}
	Paleturquoise        = color.RGBA{175, 238, 238, 255}
	Palevioletred        = color.RGBA{219, 112, 147, 255}
	Papayawhip           = color.RGBA{255, 239, 213, 255}
	Peachpuff            = color.RGBA{255, 218, 185, 255}
	Peru                 = color.RGBA{205, 133, 63, 255}
	Pink                 = color.RGBA{255, 192, 203, 255}
	Plum                 = color.RGBA{221, 160, 221, 255}
	Powderblue           = color.RGBA{176, 224, 230, 255}
	Purple               = color.RGBA{128, 0, 128, 255}
	Red                  = color.RGBA{255, 0, 0, 255}
	Rosybrown            = color.RGBA{188, 143, 143, 255}
	Royalblue            = color.RGBA{65, 105, 225, 255}
	Saddlebrown          = color.RGBA{139, 69, 19, 255}
	Salmon               = color.RGBA{250, 128, 114, 255}
	Sandybrown           = color.RGBA{244, 164, 96, 255}
	Seagreen             = color.RGBA{46, 139, 87, 255}
	Seashell             = color.RGBA{255, 245, 238, 255}
	Sienna               = color.RGBA{160, 82, 45, 255}
	Silver               = color.RGBA{192, 192, 192, 255}
	Skyblue              = color.RGBA{135, 206, 235, 255}
	Slateblue            = color.RGBA{106, 90, 205, 255}
	Slategray            = color.RGBA{112, 128, 144, 255}
	Slategrey            = color.RGBA{112, 128, 144, 255}
	Snow                 = color.RGBA{255, 250, 250, 255}
	Springgreen          = color.RGBA{0, 255, 127, 255}
	Steelblue            = color.RGBA{70, 130, 180, 255}
	TanColor             = color.RGBA{210, 180, 140, 255}
	Teal                 = color.RGBA{0, 128, 128, 255}
	Thistle              = color.RGBA{216, 191, 216, 255}
	Tomato               = color.RGBA{255, 99, 71, 255}
	Turquoise            = color.RGBA{64, 224, 208, 255}
	Violet               = color.RGBA{238, 130, 238, 255}
	Wheat                = color.RGBA{245, 222, 179, 255}
	White                = color.RGBA{255, 255, 255, 255}
	Whitesmoke           = color.RGBA{245, 245, 245, 255}
	Yellow               = color.RGBA{255, 255, 0, 255}
	Yellowgreen          = color.RGBA{154, 205, 50, 255}
)

// ConvertFColor returns r, g, b, a float32 types converted to color.Color.
// Values below zero are capped at zero. Values greater than one are capped at one.
func ConvertFColor(r, g, b, a float32) color.Color {
	if r > 1 {
		r = 1
	}
	if g > 1 {
		g = 1
	}
	if b > 1 {
		b = 1
	}
	if a > 1 {
		a = 1
	}
	if r < 0 {
		r = 0
	}
	if g < 0 {
		g = 0
	}
	if b < 0 {
		b = 0
	}
	if a < 0 {
		a = 0
	}
	return color.RGBA{uint8(r * 0xFFFF), uint8(g * 0xFFFF), uint8(b * 0xFFFF), uint8(a * 0xFFFF)}
}

// ConvertColorF returns the converted colors from c a Go type color.Color to red, green, blue and alpha float32 types within a range of 0 to 1.
func ConvertColorF(c color.Color) (r float32, g float32, b float32, a float32) {
	red, green, blue, alpha := c.RGBA()
	if red == 0 {
		r = 0
	} else {
		r = float32(red) / float32(0xFFFF)
	}
	if green == 0 {
		g = 0
	} else {
		g = float32(green) / float32(0xFFFF)
	}
	if blue == 0 {
		b = 0
	} else {
		b = float32(blue) / float32(0xFFFF)
	}
	if alpha == 0 {
		a = 0
	} else {
		a = float32(alpha) / float32(0xFFFF)
	}
	return
}

func ConvertColorGL(c color.Color) (gl.Float, gl.Float, gl.Float, gl.Float) {
	r, g, b, a := ConvertColorF(c)
	return gl.Float(r), gl.Float(g), gl.Float(b), gl.Float(a)
}

func ConvertColorGLClampf(c color.Color) (gl.Clampf, gl.Clampf, gl.Clampf, gl.Clampf) {
	r, g, b, a := ConvertColorF(c)
	return gl.Clampf(r), gl.Clampf(g), gl.Clampf(b), gl.Clampf(a)
}
