package main

import (
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
	width := 853
	height := 480
	graphicSettings := baukasten.NewGraphicSettings(width, height, 0, false, true, "baukasten - Demo - SimpleWindow")

	demo.engine = baukasten.NewEngine(ogl.DefaultDriver, glfw.DefaultDriver, glfw.DefaultDriver, nil)
	err := demo.engine.Init(graphicSettings)
	if err != nil {
		return err
	}
	demo.engine.SetCamera(baukasten.NewTwoDCamera(0, float32(width), float32(height), 0))
	return nil
}

func (demo *SimpleWindowDemo) Unload() error {
	demo.engine.Close()
	return nil
}

func (demo *SimpleWindowDemo) Update() {
	select {
	case contextEvent := <-demo.engine.ContextEvent():
		switch contextEvent.Type() {
		case baukasten.SystemQuit:
			demo.Unload() // Hackish
		}
	case <-demo.engine.KeyEvent():
	case windowSize := <-demo.engine.ResizeEvent():
		demo.engine.GraphicResize(int(windowSize.Width()), int(windowSize.Height()))
		demo.engine.SetCamera(baukasten.NewTwoDCamera(0, float32(windowSize.Width()), float32(windowSize.Height()), 0))
	default:
	}
	demo.engine.BeginFrame()
	// TODO Render code here
	demo.engine.EndFrame()
}
