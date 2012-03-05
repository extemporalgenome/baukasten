package baukasten

const (
	SystemQuit = iota
	WindowClose
	WindowRefresh
)

type ContextEvent interface {
	Type() int
}
