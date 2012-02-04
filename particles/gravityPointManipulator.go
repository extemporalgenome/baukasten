package particles

import (
	"time"

	"github.com/Agon/baukasten"
)

// GravityPointManipulator manipulates particles, which are in range of the manipulator,
// towards or away from the position of the manipulator.
type GravityPointManipulator struct {
	Position baukasten.Vector2
	Force    float32
	Range    float32
}

func NewGravityPointManipulator(position baukasten.Vector2, force, gravityRange float32) *GravityPointManipulator {
	return &GravityPointManipulator{Position: position, Force: force, Range: gravityRange}
}

func (m *GravityPointManipulator) Update(deltaTime time.Duration, particles []Particle) {
	if m.Range < 0 { // Global effect
		for i := range particles {
			way := particles[i].Position.Sub(m.Position)
			way.Normalize()
			way.Scale(m.Force)
			particles[i].Velocity.Accumulate(way)
		}
		return
	}
	for i := range particles {
		if particles[i].Position.DistanceBetween(m.Position) <= m.Range {
			way := particles[i].Position.Sub(m.Position)
			way.Normalize()
			way.Scale(m.Force)
			particles[i].Velocity.Accumulate(way)
		}
	}
}
