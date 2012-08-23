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
