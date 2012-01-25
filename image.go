package baukasten

import (
	"image"
	// For image loading
	_ "image/bmp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	_ "image/tiff"
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
