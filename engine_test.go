package baukasten

import (
	"fmt"
	"testing"
	"os"
)

func TestEngine(t *testing.T) {
	var err os.Error
	graphicSettings := NewGraphicSettings(640, 480, 32, false, true, "Baukasten testing")

	// Drivers
	graphicDriver := NewOglDriver()
	glfwDriver := NewGlfwDriver()

	engine := NewEngine(graphicDriver, glfwDriver, glfwDriver)
	err = engine.Init(graphicSettings)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer engine.Close()
	fmt.Println("Engine loaded.")

	var logo Surface
	logo, err = engine.OpenSurface("./baukastenlogo.tga")
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println("Surface loaded.")

	rec := NewRectangleF(&Vector2{0, 0}, NewSizeF(1, 1))
	for i := 0; i < 1000; i++ {
		engine.BeginFrame()
		logo.Draw(rec)
		engine.EndFrame()
	}
}
