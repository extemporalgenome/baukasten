package baukasten

type Color3b struct {
	R, G, B byte
}

func (c *Color3b) To3f() *Color3f {
	return &Color3f{255 / float32(c.R), 255 / float32(c.G), 255 / float32(c.B)}
}

type Color3f struct {
	R, G, B float32
}

type Color4b struct {
	R, G, B, A byte
}

func (c *Color4b) To4f() *Color4f {
	return &Color4f{255 / float32(c.R), 255 / float32(c.G), 255 / float32(c.B), 255 / float32(c.A)}
}

type Color4f struct {
	R, G, B, A float32
}

type Color struct {
	R, G, B, A byte
}

func NewColor(r, g, b, a byte) *Color {
	return &Color{r, g, b, a}
}

func NewColorF(r, g, b, a float32) *Color {
	return &Color{byte(255 * r), byte(255 * g), byte(255 * b), byte(255 * a)}
}
