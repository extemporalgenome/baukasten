package particles

import (
	"time"

	"github.com/Agon/baukasten"
)

type ParticleEmitter interface {
	Particles() []Particle
	Emit(duration, life time.Duration, acceleration baukasten.Vector2, velocity baukasten.Vector2)
	Update(deltaTime time.Duration)
	AddManipulator(*Manipulator)
	Draw()
}
