package baukasten

type Camera interface {
	Get() Matrix4
}

type TwoDCamera Matrix4

func (c TwoDCamera) Get() Matrix4 {
	return Matrix4(c)
}

func NewTwoDCamera(left, right, bottom, top float32) Camera {
	return TwoDCamera(Ortho2D(left, right, bottom, top))
}
