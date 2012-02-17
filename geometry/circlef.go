package geometry

type Circlef struct {
	Position Vector2
	Radius   float32
}

func CircF(position Vector2, radius float32) Circlef {
	return Circlef{Position: position, Radius: radius}
}

func (c Circlef) IsInside(vec Vector2) bool {
	return c.Radius >= c.Position.DistanceBetween(vec)
}
