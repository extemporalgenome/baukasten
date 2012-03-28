package gui

import (
	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/geometry"
)

type Text struct {
	text     string
	fontSize float32
	dirty    bool
	surf     baukasten.Surface
	font     baukasten.Font

	size   geometry.Vector2
	parent Element
}

func NewText(parent Element) *Text {
	return &Text{
		fontSize: 12.0,
		dirty:    true,
		parent:   parent,
	}
}

func (t *Text) Text() string {
	return t.text
}

func (t *Text) SetText(text string) {
	t.text = text
	t.dirty = true
}

func (t *Text) SetFont(font baukasten.Font) {
	t.font = font
	t.dirty = true
}

func (t *Text) Draw(engine *baukasten.Engine, x, y float32) {
	if t.dirty {
		t.size.X = t.fontSize * float32(len(t.text)) * t.fontSize / 20
		t.size.Y = t.fontSize * t.fontSize / 10
		surf, err := engine.RenderText(t.text, int(t.size.X), int(t.size.Y), float64(t.fontSize), geometry.Black, t.font)
		if err != nil {
			return
		}
		t.surf = surf
		t.dirty = false
	}
	t.surf.Draw(x, y)
}

func (t *Text) Size() geometry.Vector2 {
	return t.size
}

func (t *Text) Parent() Element {
	return t.parent
}
