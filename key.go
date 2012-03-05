package baukasten

const (
	KeyRelease = iota
	KeyPress
)

const (
	KeyUnknown = -1
	KeySpace   = 32
	KeySpecial = 256
	_          = (KeySpecial + iota)
	KeyEsc
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyLshift
	KeyRshift
	KeyLctrl
	KeyRctrl
	KeyLalt
	KeyRalt
	KeyTab
	KeyEnter
	KeyBackspace
	KeyInsert
	KeyDel
	KeyPageup
	KeyPagedown
	KeyHome
	KeyEnd
	KeyKP0
	KeyKP1
	KeyKP2
	KeyKP3
	KeyKP4
	KeyKP5
	KeyKP6
	KeyKP7
	KeyKP8
	KeyKP9
	KeyKPDidivde
	KeyKPMultiply
	KeyKPSubtract
	KeyKPAdd
	KeyKPDecimal
	KeyKPEqual
	KeyKPEnter
	KeyKPNumlock
	KeyCapslock
	KeyScrolllock
	KeyPause
	KeyLsuper
	KeyRsuper
	KeyMenu
)

type Key interface {
	Key() uint
	State() uint
}
