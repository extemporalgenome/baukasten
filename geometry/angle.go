package geometry

type Angle float32

func Deg(angle float32) Angle {
	return Angle(angle)
}

func Rad(angle float32) Angle {
	return Angle(angle * 180.0 / Pi)
}

func (a Angle) Degrees() float32 {
	return float32(a)
}

func (a Angle) Radians() float32 {
	return float32(a) * Pi / 180.0
}

func (a Angle) Normalized() Angle {
	return Angle(Mod(float32(a), 360.0))
}

func (a Angle) Normalized180() Angle {
	return Angle(Mod(float32(a), 180.0))
}

func (a Angle) Vec() Vector2 {
	return Vec2(Cos(a.Radians()), Sin(a.Radians()))
}
