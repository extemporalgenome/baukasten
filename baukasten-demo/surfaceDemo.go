package main

import (
	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/driver/glfw"
	"github.com/Agon/baukasten/driver/ogl"
	"github.com/Agon/baukasten/geometry"
)

type SurfaceDemo struct {
	engine  *baukasten.Engine
	surface baukasten.Surface
}

func NewSurfaceDemo() *SurfaceDemo {
	return &SurfaceDemo{}
}

func (demo *SurfaceDemo) Name() string {
	return "surface"
}

func (demo *SurfaceDemo) Description() string {
	return "Renders surfaces with GLFW and OpenGL 3"
}

func (demo *SurfaceDemo) Load() error {
	width := 853
	height := 480
	graphicSettings := baukasten.NewGraphicSettings(width, height, 0, false, true, "baukasten - Demo - SimpleWindow")

	demo.engine = baukasten.NewEngine(ogl.DefaultDriver, glfw.DefaultDriver, glfw.DefaultDriver, nil)
	err := demo.engine.Init(graphicSettings)
	if err != nil {
		return err
	}
	demo.engine.SetCamera(baukasten.NewTwoDCamera(0, float32(width), float32(height), 0))
	demo.surface, err = demo.engine.OpenSurface("smoke.png")
	if err != nil {
		return err
	}
	return nil
}

func (demo *SurfaceDemo) Unload() error {
	demo.engine.Close()
	return nil
}

func (demo *SurfaceDemo) Update() {
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
	demo.surface.Draw(200, 100)
	demo.surface.DrawRec(geometry.RectF(0, 0, 100, 100))
	demo.surface.DrawRegionRec(geometry.RectF(0, 0, 0.5, 0.5), geometry.RectF(100, 100, 200, 200))
	demo.engine.EndFrame()
}
