package geometry

type Angle float32

func Deg(angle float32) Angle {
	return Angle(angle)
}

func Rad(angle float32) Angle {
	return Angle(angle * 180 / Pi)
}

func (a Angle) Degrees() float32 {
	return float32(a)
}

func (a Angle) Radians() float32 {
	return float32(a) * Pi / 180
}

func (a Angle) Normalized() Angle {
	for a > 360 {
		a -= 360
	}
	for a < -360 {
		a += 360
	}
	return a
}

func (a Angle) Normalized180() Angle {
	for a > 180 {
		a -= 180
	}
	for a < -180 {
		a += 180
	}
	return a
}
