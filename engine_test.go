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

	fmt.Print(engine.GetOpenGLVersion())

	err = engine.Resize(1024, 768)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 1000; i++ { // Do 1000 frames
		engine.BeginFrame()
		// TODO Drawing goes in here
		engine.EndFrame()
	}

	err = engine.Quit()
	if err != nil {
		t.Fatal(err)
	}
}
