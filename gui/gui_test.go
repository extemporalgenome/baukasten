package gui

import (
	"runtime"
	"testing"
	"time"

	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/driver/freetype"
	"github.com/Agon/baukasten/driver/glfw"
	"github.com/Agon/baukasten/driver/ogl"
	"github.com/Agon/baukasten/geometry"
)

const (
	MaxFPS = 60
)

func TestBasicGui(t *testing.T) {
	width := 853
	height := 480
	graphicSettings := baukasten.NewGraphicSettings(width, height, 0, false, true, "baukasten - GuiTest")

	engine := baukasten.NewEngine(ogl.DefaultDriver, glfw.DefaultDriver, glfw.DefaultDriver, freetype.DefaultDriver)
	err := engine.Init(graphicSettings)
	if err != nil {
		t.Fatalf("%s", err)
	}
	engine.SetCamera(baukasten.NewTwoDCamera(0, float32(width), float32(height), 0))
	engine.SetClearColor(geometry.White)

	font, err := engine.OpenFont("/usr/share/fonts/truetype/FreeMono.ttf")
	if err != nil {
		t.Fatalf("%s", err)
	}

	label := NewLabel(nil)
	label.SetFont(font)
	label.SetText("This is a label.\nWhich contains Text.")
	label.SetBorder(NewBorder(label))

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	ticker := time.NewTicker(time.Second / time.Duration(MaxFPS))
	for {
		select {
		case <-ticker.C:
			engine.BeginFrame()
			label.Draw(engine, 426, 240)
			engine.EndFrame()
		}
	}
}
