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

package glfw

import (
	"errors"
	"github.com/jteeuwen/glfw"
	"time"
)

var (
	ErrUnknownOpenGLProfile = errors.New("Unknown OpenGLProfile")
)

type OpenGLProfile int

const (
	OpenGLCoreProfile   = OpenGLProfile(glfw.OpenGLCoreProfile)
	OpenGLCompatProfile = OpenGLProfile(glfw.OpenGLCompatProfile)
)

type Properties struct {
	Width, Height  int
	R, G, B, A     int
	Depth, Stencil int
	Fullscreen     bool
}

type AdvancedProperties struct {
	RefreshRate                                                 int
	AccumRedBits, AccumGreenBits, AccumBlueBits, AccumAlphaBits int
	AuxBuffers                                                  int
	Stereo                                                      bool
	NoWindowResize                                              bool
	FsaaSamples                                                 int
	OpenGLVersionMajor, OpenGLVersionMinor                      int
	OpenGLForwardCompat                                         bool
	OpenGLDebugContext                                          bool
	OpenGLProfile                                               OpenGLProfile
}

// NewWindow initialize glfw and opens a new window with p's properties.
// The AdvancedProperties ap is optional.
func NewWindow(p Properties, ap *AdvancedProperties) (*window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}
	w := &window{t: time.Now()}
	if ap != nil {
		glfw.OpenWindowHint(glfw.RefreshRate, ap.RefreshRate)
		glfw.OpenWindowHint(glfw.AccumRedBits, ap.AccumRedBits)
		glfw.OpenWindowHint(glfw.AccumGreenBits, ap.AccumGreenBits)
		glfw.OpenWindowHint(glfw.AccumBlueBits, ap.AccumBlueBits)
		glfw.OpenWindowHint(glfw.AccumAlphaBits, ap.AccumAlphaBits)
		glfw.OpenWindowHint(glfw.AuxBuffers, ap.AuxBuffers)
		if ap.Stereo {
			glfw.OpenWindowHint(glfw.Stereo, 1)
		} else {
			glfw.OpenWindowHint(glfw.Stereo, 0)
		}
		if ap.NoWindowResize {
			glfw.OpenWindowHint(glfw.WindowNoResize, 1)
		} else {
			glfw.OpenWindowHint(glfw.WindowNoResize, 0)
		}
		glfw.OpenWindowHint(glfw.FsaaSamples, ap.FsaaSamples)
		glfw.OpenWindowHint(glfw.OpenGLVersionMajor, ap.OpenGLVersionMajor)
		glfw.OpenWindowHint(glfw.OpenGLVersionMinor, ap.OpenGLVersionMinor)
		if ap.OpenGLForwardCompat {
			glfw.OpenWindowHint(glfw.OpenGLForwardCompat, 1)
		} else {
			glfw.OpenWindowHint(glfw.OpenGLForwardCompat, 0)
		}
		if ap.OpenGLDebugContext {
			glfw.OpenWindowHint(glfw.OpenGLDebugContext, 1)
		} else {
			glfw.OpenWindowHint(glfw.OpenGLDebugContext, 0)
		}
		switch ap.OpenGLProfile {
		case OpenGLCompatProfile:
			glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
		case OpenGLCoreProfile:
			glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		default:
			return nil, ErrUnknownOpenGLProfile
		}
	}
	mode := glfw.Windowed
	if p.Fullscreen {
		mode = glfw.Fullscreen
	}
	if err := glfw.OpenWindow(p.Width, p.Height, p.R, p.G, p.B, p.A, p.Depth, p.Stencil, mode); err != nil {
		return nil, err
	}
	return w, nil
}

type CloseCallbackHandler func() bool
type SizeCallbackHandler func(width, height int)

type window struct {
	t time.Time
}

func (w *window) SetTitle(title string) {
	glfw.SetWindowTitle(title)
}

func (w *window) Close() {
	glfw.CloseWindow()
	glfw.Terminate()
}

func (w *window) SetVSync(enable bool) {
	if enable {
		glfw.SetSwapInterval(1)
		return
	}
	glfw.SetSwapInterval(0)
}

func (w *window) SetSize(width, height int) {
	glfw.SetWindowSize(width, height)
}

func (w *window) Size() (int, int) {
	return glfw.WindowSize()
}

func (w *window) SetPos(x, y int) {
	glfw.SetWindowPos(x, y)
}

func (w *window) MousePos() (int, int) {
	return glfw.MousePos()
}

func (w *window) SetMousePos(x, y int) {
	glfw.SetMousePos(x, y)
}

func (w *window) MouseWheel() int {
	return glfw.MouseWheel()
}

func (w *window) SetMouseWheel(pos int) {
	glfw.SetMouseWheel(pos)
}

// Time returns the time passed since NewWindow was called.
// NOTE: This does not use the equally named glfw function.
func (w *window) Time() time.Duration {
	return w.t.Sub(time.Now())
}

// SetTime sets the current time of the timer to the specified time.
// Subsequent calls to window.Time will be relative to this time.
// NOTE: This does not use the equally named glfw function.
func (w *window) SetTime(d time.Duration) {
	w.t.Add(d)
}

func (w *window) Iconify() {
	glfw.IconifyWindow()
}

func (w *window) Restore() {
	glfw.RestoreWindow()
}

func (w *window) SwapBuffers() {
	glfw.SwapBuffers()
}

func (w *window) Opened() bool {
	return glfw.WindowParam(glfw.Opened) == 1
}

func (w *window) Active() bool {
	return glfw.WindowParam(glfw.Active) == 1
}

func (w *window) Iconified() bool {
	return glfw.WindowParam(glfw.Iconified) == 1
}

func (w *window) Accelerated() bool {
	return glfw.WindowParam(glfw.Accelerated) == 1
}

func (w *window) RedBits() int {
	return glfw.WindowParam(glfw.RedBits)
}

func (w *window) GreenBits() int {
	return glfw.WindowParam(glfw.GreenBits)
}

func (w *window) BlueBits() int {
	return glfw.WindowParam(glfw.BlueBits)
}

func (w *window) AlphaBits() int {
	return glfw.WindowParam(glfw.AlphaBits)
}

func (w *window) DepthBits() int {
	return glfw.WindowParam(glfw.DepthBits)
}

func (w *window) StencilBits() int {
	return glfw.WindowParam(glfw.StencilBits)
}

func (w *window) RefreshRate() int {
	return glfw.WindowParam(glfw.RefreshRate)
}

// TODO Accumulation buffer bits

func (w *window) AuxBuffers() int {
	return glfw.WindowParam(glfw.AuxBuffers)
}

func (w *window) Stereo() bool {
	return glfw.WindowParam(glfw.Stereo) == 1
}

func (w *window) WindowNoResize() bool {
	return glfw.WindowParam(glfw.WindowNoResize) == 1
}

func (w *window) FsaaSamples() int {
	return glfw.WindowParam(glfw.FsaaSamples)
}

func (w *window) OpenGLContextVersion() (major int, minor int) {
	major = glfw.WindowParam(glfw.OpenGLVersionMajor)
	minor = glfw.WindowParam(glfw.OpenGLVersionMinor)
	return
}

func (w *window) OpenGLVersion() (major, minor, rev int) {
	return glfw.GLVersion()
}

func (w *window) OpenGLForwardCompat() bool {
	return glfw.WindowParam(glfw.OpenGLForwardCompat) == 1
}

func (w *window) OpenGLDebugContext() bool {
	return glfw.WindowParam(glfw.OpenGLDebugContext) == 1
}

func (w *window) OpenGLProfile() int {
	return glfw.WindowParam(glfw.OpenGLProfile)
}

func (w *window) SetMouseCursor(enable bool) {
	if enable {
		glfw.Enable(glfw.MouseCursor)
		return
	}
	glfw.Disable(glfw.MouseCursor)
}

func (w *window) SetStickyKeys(enable bool) {
	if enable {
		glfw.Enable(glfw.StickyKeys)
		return
	}
	glfw.Disable(glfw.StickyKeys)
}

func (w *window) SetStickyMouseButtons(enable bool) {
	if enable {
		glfw.Enable(glfw.StickyMouseButtons)
		return
	}
	glfw.Disable(glfw.StickyMouseButtons)
}

func (w *window) SetSystemKeys(enable bool) {
	if enable {
		glfw.Enable(glfw.SystemKeys)
		return
	}
	glfw.Disable(glfw.SystemKeys)
}

func (w *window) SetKeyRepeat(enable bool) {
	if enable {
		glfw.Enable(glfw.KeyRepeat)
		return
	}
	glfw.Disable(glfw.KeyRepeat)
}

func (w *window) SetCloseCallback(f CloseCallbackHandler) {
	glfw.SetWindowCloseCallback(func() int {
		closeWindow := f()
		if closeWindow {
			return 1
		}
		return 0
	})
}

func (w *window) SetSizeCallback(f SizeCallbackHandler) {
	glfw.SetWindowSizeCallback(func(width, height int) {
		f(width, height)
	})
}

// TODO MouseButton, MouseWheel, Key, Char callbacks
