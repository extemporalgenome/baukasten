package baukasten

type Matrix4x4 [16]float32

func TranslationMatrix(x, y, z float32) Matrix4x4 {
	return Matrix4x4{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		x, y, z, 1.0}
}
