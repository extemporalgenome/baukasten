package main

import (
	"runtime"

	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/driver/glfw"
	"github.com/Agon/baukasten/driver/ogl"
)

type SimpleWindowDemo struct {
	engine *baukasten.Engine
}

func NewSimpleWindowDemo() *SimpleWindowDemo {
	return &SimpleWindowDemo{}
}

func (demo *SimpleWindowDemo) Name() string {
	return "simplewindow"
}

func (demo *SimpleWindowDemo) Description() string {
	return "Shows a simple window using GLFW and OpenGL 3"
}

func (demo *SimpleWindowDemo) Load() error {
	runtime.LockOSThread()
	graphicSettings := baukasten.NewGraphicSettings(853, 480, 0, false, true, "baukasten - Demo - SimpleWindow")

	demo.engine = baukasten.NewEngine(ogl.DefaultDriver, glfw.DefaultDriver, glfw.DefaultDriver, nil)
	err := demo.engine.Init(graphicSettings)
	if err != nil {
		return err
	}
	runtime.UnlockOSThread()
	return nil
}

func (demo *SimpleWindowDemo) Unload() error {
	demo.engine.Close()
	return nil
}

func (demo *SimpleWindowDemo) Update() {
	runtime.LockOSThread()
	select {
	case contextEvent := <-demo.engine.ContextEvent():
		switch contextEvent.Type() {
		case baukasten.SystemQuit:
			demo.Unload() // Hackish
		}
	case <-demo.engine.KeyEvent():
	case windowSize := <-demo.engine.ResizeEvent():
		demo.engine.GraphicResize(int(windowSize.Width()), int(windowSize.Height()))
	default:
	}
	demo.engine.BeginFrame()
	// TODO Render code here
	demo.engine.EndFrame()
	runtime.UnlockOSThread()
}
