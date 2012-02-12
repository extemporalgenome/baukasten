package baukasten

import (
	math "github.com/Agon/baukasten/geometry"
)

type Camera interface {
	Get() math.Matrix4
}

type TwoDCamera math.Matrix4

func (c TwoDCamera) Get() math.Matrix4 {
	return math.Matrix4(c)
}

func NewTwoDCamera(left, right, bottom, top float32) Camera {
	return TwoDCamera(math.Ortho2D(left, right, bottom, top))
}
