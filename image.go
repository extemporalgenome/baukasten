package baukasten

import (
	"image"
	// For image loading
	_ "code.google.com/p/go.image/bmp"
	_ "code.google.com/p/go.image/tiff"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func OpenImage(name string) (image.Image, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(file)
	return img, err
}
