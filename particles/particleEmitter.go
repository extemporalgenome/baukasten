package particles

import (
	"time"

	math "github.com/Agon/baukasten/geometry"
)

type ParticleEmitter interface {
	Particles() []Particle
	Emit(duration, life time.Duration, acceleration math.Vector2, velocity math.Vector2)
	Update(deltaTime time.Duration)
	AddManipulator(*Manipulator)
	Draw()
}
