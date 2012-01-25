package glfw

type ContextEvent int

func (event ContextEvent) Type() int { return int(event) }
