package baukasten

import (
	"fmt"
	"image"
	_ "image/png"
	"io"
	"testing"
	"os"
)

const (
	LogoFileNameTGA = "./baukastenlogo.tga"
	LogoFileNamePNG = "./baukastenlogo.png"
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
	logo, err = engine.OpenSurface(LogoFileNameTGA)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println("Surface1 opened.")

	var logo2 Surface
	var logoStream io.Reader
	var logo2Img image.Image
	logoStream, err = os.Open(LogoFileNamePNG)
	if err != nil {
		t.Fatal(err)
		return
	}
	logo2Img, _, err = image.Decode(logoStream)
	if err != nil {
		t.Fatal(err)
		return
	}
	logo2, err = engine.LoadSurface(logo2Img)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println("Surface2 loaded.")

	rec1 := NewRectangleF(&Vector2{-0.5, 0}, NewSizeF(1, 1))
	rec2 := NewRectangleF(&Vector2{0.5, 0}, NewSizeF(1, 1))
	for i := 0; i < 1000; i++ {
		engine.BeginFrame()
		logo.Draw(rec1)
		logo2.Draw(rec2)
		engine.EndFrame()
	}
}
