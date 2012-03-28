package gui

import (
	"github.com/Agon/baukasten"
	"github.com/Agon/baukasten/geometry"
)

type Label struct {
	text   *Text
	border *Border
	parent Element
}

func NewLabel(parent Element) *Label {
	label := &Label{
		parent: parent,
	}
	label.text = NewText(label)
	label.border = NewBorder(label)
	label.border.Width = 0
	return label
}

func (l *Label) Draw(engine *baukasten.Engine, x, y float32) {
	l.text.Draw(engine, x, y)
	l.border.Draw(engine, x, y)
}

func (l *Label) SetBorder(border *Border) {
	l.border = border
}

func (l *Label) SetText(text string) {
	l.text.SetText(text)
}

func (l *Label) SetFont(font baukasten.Font) {
	l.text.SetFont(font)
}

func (l *Label) Parent() Element {
	return l.parent
}

func (l *Label) Size() geometry.Vector2 {
	return l.text.Size()
}
