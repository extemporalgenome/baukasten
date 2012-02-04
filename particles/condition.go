package particles

const (
	// Particle is dead and is ready to get recycled.
	Dead = Condition(iota)
	// Particle is paused and will not be simulated
	Paused
	// Particle is frozen and will not be simulated but drawn
	Frozen
	// Particle is alive and will be simulated, drawn.
	Alive
)

type Condition int
