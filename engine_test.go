package baukasten

import (
	"fmt"
	"testing"
	"os"
)

func TestEngine(t *testing.T) {
	var err os.Error
	graphicSettings := NewGraphicSettings(640, 480, 32, false, true, "Baukasten testing")
	engine := NewEngine()
	err = engine.Init(graphicSettings)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\n", engine.GetOpenGLVersion())

	err = engine.Resize(1024, 768)
	if err != nil {
		t.Fatal(err)
	}
	/*
		        vertices := []Vertex{ 
				Vertex{ []float32{ 0, 0, -3.0, 1.0 }, []float32{ 1.0, 0.0, 0.0, 1.0 } }, 
				Vertex{ []float32{ 1, 1, -3.0, 1.0 }, []float32{ 0.0, 1.0, 0.0, 1.0 } }, 
				Vertex{ []float32{ 1, 0, -3.0, 1.0 }, []float32{ 0.0, 0.0, 1.0, 1.0 } } }
	*/
	vertices := []float32{
		0, 0, -3, 1,
		1, 1, -3, 1,
		1, 0, -3, 1,
	}
	colors := []float32{
		1, 0, 0, 1,
		0, 1, 0, 1,
		0, 0, 1, 1,
	}

	for i := 0; i < 1000; i++ { // Do 1000 frames
		engine.BeginFrame()
		engine.DrawPolygon(vertices, colors)
		engine.EndFrame()
	}

	err = engine.Quit()
	if err != nil {
		t.Fatal(err)
	}
}
