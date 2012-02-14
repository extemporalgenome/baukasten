package main

import (
	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/driver/glfw"
	"github.com/Agon/baukasten/driver/ogl"
	"github.com/Agon/baukasten/geometry"
	"github.com/Agon/baukasten/particles"
	"time"
)

type ParticlesDemo struct {
	engine  *baukasten.Engine
	emitter *particles.SurfaceEmitter
	surface baukasten.Surface
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
	width := 853
	height := 480
	graphicSettings := baukasten.NewGraphicSettings(width, height, 0, false, true, "baukasten - Demo - Particles")

	demo.engine = baukasten.NewEngine(ogl.DefaultDriver, glfw.DefaultDriver, glfw.DefaultDriver, nil)
	err := demo.engine.Init(graphicSettings)
	if err != nil {
		return err
	}
	demo.surface, err = demo.engine.OpenSurface("smoke.png")
	if err != nil {
		return err
	}
	c := geometry.Blue
	c.A = 175
	demo.surface.SetColor(c)
	demo.engine.SetCamera(baukasten.NewTwoDCamera(0, float32(width), float32(height), 0))
	gravity := particles.NewGravityManipulator(geometry.Vector2{10, -10})
	pointGravity := particles.NewGravityPointManipulator(geometry.Vector2{400, 200}, -10, 200)
	demo.emitter = particles.NewSurfaceEmitter(demo.surface)
	demo.emitter.Position = geometry.Vector2{400, 480}
	demo.emitter.Emit(time.Second/25, time.Second*10, geometry.Vector2{0, 0}, geometry.Vector2{0, 0})
	demo.emitter.AddManipulator(gravity)
	demo.emitter.AddManipulator(pointGravity)
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
	case mouse := <-demo.engine.MousePositionEvent():
		demo.emitter.Position = geometry.Vector2{float32(mouse.X()), float32(mouse.Y())}
	case windowSize := <-demo.engine.ResizeEvent():
		demo.engine.GraphicResize(int(windowSize.Width()), int(windowSize.Height()))
		demo.engine.SetCamera(baukasten.NewTwoDCamera(0, float32(windowSize.Width()), float32(windowSize.Height()), 0))
	default:
	}
	demo.emitter.Update(demo.engine.DeltaTime())
	demo.engine.BeginFrame()
	demo.emitter.Draw()
	demo.engine.EndFrame()
}