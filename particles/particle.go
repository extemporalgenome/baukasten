// A simple, two dimensional particle simulator package for baukasten.
package particles

import (
	"time"

	"github.com/Agon/baukasten"
)

type Particle struct {
	Acceleration baukasten.Vector2
	Position     baukasten.Vector2
	Velocity     baukasten.Vector2
	Condition    Condition
	Life         time.Duration
}

func Par(a, p, v baukasten.Vector2, life time.Duration, c Condition) Particle {
	return Particle{Acceleration: a, Position: p, Velocity: v, Life: life, Condition: c}
}
