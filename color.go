package baukasten

type Color3b struct {
	R, G, B byte
}

func (c *Color3b) To3f() *Color3f {
	return &Color3f{255 / c.R, 255 / c.G, 255 / c.B}
}

type Color3f struct {
	R, G, B float32
}

type Color4b struct {
	R, G, B, A byte
}

func (c *Color4b) To4f() *Color4f {
	return &Color4f{255 / c.R, 255 / c.G, 255 / c.B, 255 / c.A}
}

type Color4f struct {
	R, G, B, A float32
}
