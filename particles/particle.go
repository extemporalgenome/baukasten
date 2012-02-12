// A simple, two dimensional particle simulator package for baukasten.
package particles

import (
	"time"

	math "github.com/Agon/baukasten/geometry"
)

type Particle struct {
	Acceleration math.Vector2
	Position     math.Vector2
	Velocity     math.Vector2
	Condition    Condition
	Life         time.Duration
}

func Par(a, p, v math.Vector2, life time.Duration, c Condition) Particle {
	return Particle{Acceleration: a, Position: p, Velocity: v, Life: life, Condition: c}
}
