package geometry

type Angle float32

// Deg constructs an Angle based on a degrees.
func Deg(a float32) Angle {
	return Angle(a)
}

// Rad constructs an Angle based on a radians.
func Rad(a float32) Angle {
	return Angle(a * 180.0 / Pi)
}

// Degrees returns a as a float32 degree.
func (a Angle) Degrees() float32 {
	return float32(a)
}

// Radians returns a as a float32 radian.
func (a Angle) Radians() float32 {
	return float32(a) * Pi / 180.0
}

// Normalized returns a normalized to a range between -360 and 360.
func (a Angle) Normalized() Angle {
	return Angle(Mod(float32(a), 360.0))
}

// Normalized180 returns a normalized to a range between -180 and 180.
func (a Angle) Normalized180() Angle {
	return Angle(Mod(float32(a), 180.0))
}

// Vec returns a's vector representation aligned to the x-axis.
func (a Angle) Vec() Vector2 {
	return Vec2(Cos(a.Radians()), Sin(a.Radians()))
}
