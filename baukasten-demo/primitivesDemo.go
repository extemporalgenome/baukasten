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
	width := 853
	height := 480
	graphicSettings := baukasten.NewGraphicSettings(width, height, 0, false, true, "baukasten - Demo - Particles")

	demo.engine = baukasten.NewEngine(ogl.DefaultDriver, glfw.DefaultDriver, glfw.DefaultDriver, nil)
	err := demo.engine.Init(graphicSettings)
	if err != nil {
		return err
	}
	demo.engine.SetCamera(baukasten.NewTwoDCamera(0, float32(width), float32(height), 0))
	demo.engine.SetClearColor(baukasten.Black)
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

	demo.engine.DrawPoints(baukasten.White, baukasten.Vector2{0, 10}, baukasten.Vector2{10, 10}, baukasten.Vector2{20, 10}, baukasten.Vector2{30, 10})
	demo.engine.DrawLines(baukasten.Aqua, baukasten.Vector2{0, 50}, baukasten.Vector2{100, 50})
	demo.engine.DrawLineStrip(baukasten.Blue, baukasten.Vector2{0, 100}, baukasten.Vector2{100, 100}, baukasten.Vector2{100, 150})
	demo.engine.DrawLineLoop(baukasten.Lightblue, baukasten.Vector2{0, 200}, baukasten.Vector2{100, 200}, baukasten.Vector2{100, 250})

	demo.engine.DrawTriangles(baukasten.Red, baukasten.Vector2{200, 100}, baukasten.Vector2{300, 100}, baukasten.Vector2{300, 200})
	demo.engine.DrawTriangleStrip(baukasten.Lime, baukasten.Vector2{200, 200}, baukasten.Vector2{300, 200}, baukasten.Vector2{300, 300})
	demo.engine.DrawTriangleFan(baukasten.Blue, baukasten.Vector2{200, 300}, baukasten.Vector2{300, 300}, baukasten.Vector2{300, 400})

	demo.engine.EndFrame()
}
