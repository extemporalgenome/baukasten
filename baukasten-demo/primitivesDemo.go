package main

import (
	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/driver/glfw"
	"github.com/Agon/baukasten/driver/ogl"
)

type PrimitivesDemo struct {
	engine *baukasten.Engine
}

func NewPrimitivesDemo() *PrimitivesDemo {
	return &PrimitivesDemo{}
}

func (demo *PrimitivesDemo) Name() string {
	return "primitives"
}

func (demo *PrimitivesDemo) Description() string {
	return "Renders primitives with GLFW and OpenGL 3"
}

func (demo *PrimitivesDemo) Load() error {
	graphicSettings := baukasten.NewGraphicSettings(853, 480, 0, false, true, "baukasten - Demo - SimpleWindow")

	demo.engine = baukasten.NewEngine(ogl.DefaultDriver, glfw.DefaultDriver, glfw.DefaultDriver, nil)
	err := demo.engine.Init(graphicSettings)
	if err != nil {
		return err
	}
	ogl.DefaultDriver.SetClearColor(baukasten.Black)
	return nil
}

func (demo *PrimitivesDemo) Unload() error {
	demo.engine.Close()
	return nil
}

func (demo *PrimitivesDemo) Update() {
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
	demo.engine.DrawPoints(baukasten.White, baukasten.Vector2{0.1, -0.5}, baukasten.Vector2{0.2, -0.5}, baukasten.Vector2{0.3, -0.5}, baukasten.Vector2{0.4, -0.5})
	demo.engine.DrawLines(baukasten.Aqua, baukasten.Vector2{-1, 0}, baukasten.Vector2{0, 0})
	demo.engine.DrawLineStrip(baukasten.Blue, baukasten.Vector2{-1, -0.5}, baukasten.Vector2{0, -0.5}, baukasten.Vector2{0, -1})
	demo.engine.DrawLineLoop(baukasten.Lightblue, baukasten.Vector2{-1, 0.5}, baukasten.Vector2{0, 0.5}, baukasten.Vector2{0.5, 1})
	demo.engine.DrawTriangles(baukasten.Red, baukasten.Vector2{0, 0}, baukasten.Vector2{1, 0}, baukasten.Vector2{1, 1})
	demo.engine.DrawTriangleStrip(baukasten.Lime, baukasten.Vector2{0, -0.2}, baukasten.Vector2{1, -0.2}, baukasten.Vector2{1, 0.8})
	demo.engine.DrawTriangleFan(baukasten.Blue, baukasten.Vector2{0, -0.4}, baukasten.Vector2{1, -0.4}, baukasten.Vector2{1, 0.6})

	demo.engine.EndFrame()
}
