package particles

import (
	"time"

	"github.com/Agon/baukasten"
)

// GravityManipulator manipulates all particles with a gravity.
type GravityManipulator struct {
	Gravity baukasten.Vector2
}

func NewGravityManipulator(gravity baukasten.Vector2) *GravityManipulator {
	return &GravityManipulator{Gravity: gravity}
}

func (m *GravityManipulator) Update(deltaTime time.Duration, particles []Particle) {
	for i := range particles {
		particles[i].Acceleration.Accumulate(m.Gravity.Scaled(float32(deltaTime.Seconds())))
	}
}
