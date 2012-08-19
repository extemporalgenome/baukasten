package image

type Sizef struct {
	Width, Height float32
}

func Sizf(w, h float32) Sizef {
	return Sizef{Width: w, Height: h}
}

func (s Sizef) Add(a float32) Sizef {
	s.Width += a
	s.Height += a
	return s
}

func (s Sizef) Sub(a float32) Sizef {
	s.Width -= a
	s.Height -= a
	return s
}

func (s Sizef) Zero() bool {
	return s.Width == 0 && s.Height == 0
}

func (s Sizef) Eq(s1 Sizef) bool {
	return s.Width == s1.Width && s.Height == s1.Height
}
