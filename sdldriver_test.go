package baukasten

import (
	"testing"
	"os"
)

func TestSDLDriver(t *testing.T) {
	fullScreen := true
	resizeAble := true
	var wm WindowManager
	var err os.Error

	graphicSettings := NewGraphicSettings(640, 480, 32, fullScreen, resizeAble, "Baukasten testing...")
	wm = NewSDLDriver()

	err = wm.Init(graphicSettings)
	if err != nil {
		t.Fatal(err)
	}

	err = wm.Resize(1024, 768)
	if err != nil {
		t.Fatal(err)
	}

	err = wm.Quit()
	if err != nil {
		t.Fatal(err)
	}
}
