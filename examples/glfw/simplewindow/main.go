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

package main

import (
	"fmt"
	"github.com/Agon/baukasten/glfw"
	"github.com/Agon/baukasten/time/timeutil"
	"time"
)

var (
	Width, Height                           = 256, 256
	RedBits, GreenBits, BlueBits, AlphaBits = 8, 8, 8, 8
	DepthBits, StencilBits                  = 0, 0
	Fullscreen                              = false
	WindowTitle                             = "baukasten - gl, glfw - SimpleWindow"
	FPSLimit                                = time.Duration(60)
	CloseChan                               = make(chan bool, 1)
)

func main() {
	properties := glfw.Properties{
		Width, Height,
		RedBits, GreenBits, BlueBits, AlphaBits,
		DepthBits, StencilBits,
		Fullscreen,
	}
	window, err := glfw.NewWindow(properties, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Close()
	window.SetTitle(WindowTitle)
	window.SetVSync(true)
	window.SetCloseCallback(onClose)
	window.SetSizeCallback(onResize)
	window.SetMouseButtonCallback(onMouseButton)
	window.SetMouseWheelCallback(onMouseWheel)
	window.SetKeyCallback(onKey)
	window.SetCharCallback(onChar)

	timer := timeutil.NewDeltaTimer()
	ticker := time.NewTicker(time.Second / FPSLimit)
	run := true
	for run {
		select {
		case <-ticker.C:
			timer.DeltaTime()
			//fmt.Println("Frame time: " + deltaTime.String())
			window.SwapBuffers()
		case <-CloseChan:
			run = false
			fmt.Println("Closed")
		}
	}
}

func onResize(w, h int) {
	fmt.Printf("New size: %d,%d\n", w, h)
}

func onClose() bool {
	CloseChan <- true
	return true // Close window
}

func onMouseButton(btn glfw.MouseButton, pressed bool) {
	fmt.Printf("Button %d is pressed %v\n", btn, pressed)
}

func onMouseWheel(delta int) {
	fmt.Printf("Mouse wheel's delta: %d\n", delta)
}

func onKey(k glfw.Key, pressed bool) {
	fmt.Printf("Key %s , pressed=%v\n", k, pressed)
}

func onChar(k glfw.Key, pressed bool) {
	fmt.Printf("Char %s , pressed=%v\n", k, pressed)
}
