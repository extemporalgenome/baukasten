package main

import (
	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/driver/glfw"
	"github.com/Agon/baukasten/driver/ogl"
	"github.com/Agon/baukasten/geometry"
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
	demo.engine.SetClearColor(geometry.Black)
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
		case baukasten.WindowClose:
			demo.Unload() // Hackish
		}
	case <-demo.engine.KeyEvent():
	case windowSize := <-demo.engine.ResizeEvent():
		demo.engine.GraphicResize(int(windowSize.Width()), int(windowSize.Height()))
		demo.engine.SetCamera(baukasten.NewTwoDCamera(0, float32(windowSize.Width()), float32(windowSize.Height()), 0))
	default:
	}
	demo.engine.BeginFrame()

	demo.engine.DrawPoints(geometry.White, geometry.Vector2{0, 10}, geometry.Vector2{10, 10}, geometry.Vector2{20, 10}, geometry.Vector2{30, 10})
	demo.engine.DrawLines(geometry.Aqua, geometry.Vector2{0, 50}, geometry.Vector2{100, 50})
	demo.engine.DrawLineStrip(geometry.Blue, geometry.Vector2{0, 100}, geometry.Vector2{100, 100}, geometry.Vector2{100, 150})
	demo.engine.DrawLineLoop(geometry.Lightblue, geometry.Vector2{0, 200}, geometry.Vector2{100, 200}, geometry.Vector2{100, 250})

	demo.engine.DrawTriangles(geometry.Red, geometry.Vector2{200, 100}, geometry.Vector2{300, 100}, geometry.Vector2{300, 200})
	demo.engine.DrawTriangleStrip(geometry.Lime, geometry.Vector2{200, 200}, geometry.Vector2{300, 200}, geometry.Vector2{300, 300})
	demo.engine.DrawTriangleFan(geometry.Blue, geometry.Vector2{200, 300}, geometry.Vector2{300, 300}, geometry.Vector2{300, 400})
	a, b, c, d := geometry.Vec2(400, 100), geometry.Vec2(500, 200), geometry.Vec2(600, 100), geometry.Vec2(700, 300)
	curve := geometry.BCurve(a, b, c, d)
	curveDetail := 100
	curveVecs := make([]geometry.Vector2, curveDetail)
	for i := 0; i < curveDetail; i++ {
		curveVecs[i] = curve.RelativePoint(float32(i) / float32(curveDetail))
	}
	demo.engine.DrawLineStrip(geometry.Orange, curveVecs...)
	demo.engine.DrawPoints(geometry.Green, a, b, c, d)
	demo.engine.EndFrame()
}
