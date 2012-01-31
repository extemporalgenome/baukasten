package main

import (
	"image/color"

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
	ogl.DefaultDriver.SetClearColor(color.Black)
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
	demo.engine.DrawPoints(color.White, baukasten.Vector2{0.1, -0.5}, baukasten.Vector2{0.2, -0.5}, baukasten.Vector2{0.3, -0.5}, baukasten.Vector2{0.4, -0.5})
	demo.engine.DrawLines(color.White, baukasten.Vector2{-1, 0}, baukasten.Vector2{0, 0})
	demo.engine.DrawLineStrip(color.White, baukasten.Vector2{-1, -0.5}, baukasten.Vector2{0, -0.5}, baukasten.Vector2{0, -1})
	demo.engine.DrawLineLoop(color.White, baukasten.Vector2{-1, 0.5}, baukasten.Vector2{0, 0.5}, baukasten.Vector2{0.5, 1})
	demo.engine.DrawTriangles(color.White, baukasten.Vector2{0, 0}, baukasten.Vector2{1, 0}, baukasten.Vector2{1, 1})
	demo.engine.EndFrame()
}
