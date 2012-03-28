// Freetype FontDriver package for baukasten.
package freetype

import (
	"io/ioutil"

	"github.com/Agon/baukasten"

	"code.google.com/p/freetype-go/freetype"
)

var DefaultDriver = NewDriver()

type Driver struct{}

func NewDriver() *Driver {
	return &Driver{}
}

func (d *Driver) OpenFont(fileName string) (baukasten.Font, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return d.LoadFont(b)
}

func (d *Driver) LoadFont(data []byte) (baukasten.Font, error) {
	font, err := freetype.ParseFont(data)
	if err != nil {
		return nil, err
	}
	// Context
	c := freetype.NewContext()
	c.SetFont(font)
	return &FontContext{context: c, units: font.UnitsPerEm()}, nil
}
