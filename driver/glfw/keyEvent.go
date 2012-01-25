package glfw

type KeyEvent struct {
	key   int
	state int
}

func NewKeyEvent(key, state int) *KeyEvent {
	return &KeyEvent{key, state}
}

func (event *KeyEvent) Key() uint {
	return uint(event.key)
}
func (event *KeyEvent) State() uint {
	return uint(event.state)
}
