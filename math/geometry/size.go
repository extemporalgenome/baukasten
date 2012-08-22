package geometry

type Size struct {
	Width, Height int
}

func Sizi(w, h int) Size {
	return Size{Width: w, Height: h}
}

func (s Size) Add(a int) Size {
	s.Width += a
	s.Height += a
	return s
}

func (s Size) Sub(a int) Size {
	s.Width -= a
	s.Height -= a
	return s
}

func (s Size) Zero() bool {
	return s.Width == 0 && s.Height == 0
}

func (s Size) Eq(s1 Size) bool {
	return s.Width == s1.Width && s.Height == s1.Height
}
