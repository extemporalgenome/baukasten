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

	poly1 := []Vector2{
		Vector2{0, 0},
		Vector2{1, 1},
		Vector2{1, 0},
	}
	poly2 := []Vector2{
		Vector2{0, 0},
		Vector2{-1, -1},
		Vector2{-1, 0},
	}

	list := engine.StartList()
	engine.FillPolygon2(poly2)
	engine.EndList()

	for i := 0; i < 1000; i++ { // Do 1000 frames
		engine.BeginFrame()
		engine.DrawPolygon2(poly1)
		engine.DrawList(list)
		engine.EndFrame()
	}

	err = engine.Quit()
	if err != nil {
		t.Fatal(err)
	}
}
