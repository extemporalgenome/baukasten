package main

import (
	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/driver/glfw"
	"github.com/Agon/baukasten/driver/ogl"
	"github.com/Agon/baukasten/particles"
	"time"
)

type ParticlesDemo struct {
	engine  *baukasten.Engine
	emitter *particles.SurfaceEmitter
}

func NewParticlesDemo() *ParticlesDemo {
	return &ParticlesDemo{}
}

func (demo *ParticlesDemo) Name() string {
	return "particles"
}

func (demo *ParticlesDemo) Description() string {
	return "Shows particle flying around using GLFW and OpenGL 3"
}

func (demo *ParticlesDemo) Load() error {
	graphicSettings := baukasten.NewGraphicSettings(853, 480, 0, false, true, "baukasten - Demo - Particles")

	demo.engine = baukasten.NewEngine(ogl.DefaultDriver, glfw.DefaultDriver, glfw.DefaultDriver, nil)
	err := demo.engine.Init(graphicSettings)
	if err != nil {
		return err
	}
	surface, err := demo.engine.OpenSurface("smoke.png")
	if err != nil {
		return err
	}
	gravity := particles.NewGravityManipulator(baukasten.Vector2{10, -10})
	demo.emitter = particles.NewSurfaceEmitter(surface)
	demo.emitter.Position = baukasten.Vector2{400, 480}
	demo.emitter.Emit(time.Second/2, time.Second*5, baukasten.Vector2{0, -10}, baukasten.Vector2{0, -10})
	demo.emitter.AddManipulator(gravity)
	return nil
}

func (demo *ParticlesDemo) Unload() error {
	demo.engine.Close()
	return nil
}

func (demo *ParticlesDemo) Update() {
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
	demo.emitter.Update(demo.engine.DeltaTime())
	demo.engine.BeginFrame()
	demo.emitter.Draw()
	demo.engine.EndFrame()
}
