package main

import (
	"log"
	"runtime"
	"time"

	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/driver/glfw"
	"github.com/Agon/baukasten/driver/ogl"
)

const (
	// 16/9
	ScreenWidth  = 853
	ScreenHeight = 480
	BitDepth     = 0
	FullScreen   = false
	Resizeable   = true
	WindowTitle  = "SimpleWindow"
	MaxFPS       = 60
)

var engine *baukasten.Engine

func main() {
	runtime.LockOSThread()
	var err error
	graphicSettings := baukasten.NewGraphicSettings(ScreenWidth, ScreenHeight, BitDepth, FullScreen, Resizeable, WindowTitle)

	engine = baukasten.NewEngine(ogl.DefaultDriver, glfw.DefaultDriver, glfw.DefaultDriver, nil)
	err = engine.Init(graphicSettings)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer engine.Close()

	run := true
	// The timer limits the frames per second
	ticker := time.NewTicker(time.Second / MaxFPS)
	for run {
		select {
		case contextEvent := <-engine.ContextEvent():
			switch contextEvent.Type() {
			case baukasten.SystemQuit:
				run = false
			}
		case windowSize := <-engine.ResizeEvent():
			log.Printf("Window resized to %d,%d\n", windowSize.Width(), windowSize.Height())
			engine.GraphicResize(int(windowSize.Width()), int(windowSize.Height()))
		case keyEvent := <-engine.KeyEvent():
			log.Printf("KeyEvent: Key %d=%s , State %d", keyEvent.Key(), string(keyEvent.Key()), keyEvent.State())
		case <-ticker.C:
			// Graphics:
			engine.BeginFrame()
			// TODO Render code here
			engine.EndFrame()
		}
	}
	runtime.UnlockOSThread()
}
