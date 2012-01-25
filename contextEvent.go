package baukasten

const (
	SystemQuit = iota
)

type ContextEvent interface {
	Type() int
}
