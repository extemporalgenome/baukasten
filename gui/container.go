package gui

type Container interface {
	Element
	Childs() []Element
}
