package baukasten

type Matrix4x4 struct {
	M11, M12, M13, M14 float32
	M21, M22, M23, M24 float32
	M31, M32, M33, M34 float32
	M41, M42, M43, M44 float32
}

func TranslationMatrix(x, y, z float32) Matrix4x4 {
	return Matrix4x4{
		1.0, 0.0, 0.0, x,
		0.0, 1.0, 0.0, y,
		0.0, 0.0, 1.0, z,
		0.0, 0.0, 0.0, 1.0}
}

func RotationMatrix(vec Vector3, angle float32) Matrix4x4 {
	c := Cos(angle)
	s := Sin(angle)
	norm := vec.Normalized()
	x := norm.X
	y := norm.Y
	z := norm.Z
	return Matrix4x4{
		x*x*(1-c) + c, x*y*(1-c) - z*s, x*z*(1-c) + y*s, 0,
		y*x*(1-c) + z*s, y*y*(1-c) + c, y*z*(1-c) - x*s, 0,
		x*z*(1-c) - y*s, y*z*(1-c) + x*s, z*z*(1-c) + c, 0,
		0, 0, 0, 1}
}
