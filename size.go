package baukasten

type Size struct {
	Width, Height int
}

func NewSize(w, h int) *Size {
	return &Size{w, h}
}

type SizeF struct {
	Width, Height float32
}

func NewSizeF(w, h float32) *SizeF {
	return &SizeF{w, h}
}
