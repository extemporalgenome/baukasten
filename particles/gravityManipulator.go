package particles

import (
	"time"

	math "github.com/Agon/baukasten/geometry"
)

// GravityManipulator manipulates all particles with a gravity.
type GravityManipulator struct {
	Gravity math.Vector2
}

func NewGravityManipulator(gravity math.Vector2) *GravityManipulator {
	return &GravityManipulator{Gravity: gravity}
}

func (m *GravityManipulator) Update(deltaTime time.Duration, particles []Particle) {
	for i := range particles {
		particles[i].Acceleration = particles[i].Acceleration.Add(m.Gravity.Scaled(float32(deltaTime.Seconds())))
	}
}
