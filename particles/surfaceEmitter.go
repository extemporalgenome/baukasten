package particles

import (
	"time"

	"github.com/Agon/baukasten"
)

type SurfaceEmitter struct {
	Particle
	particles []Particle
	surface   baukasten.Surface

	emitFrequency    time.Duration
	emitLife         time.Duration
	emitTime         time.Duration
	emitAcceleration baukasten.Vector2
	emitVelocity     baukasten.Vector2

	manipulators []Manipulator
}

func NewSurfaceEmitter(s baukasten.Surface) *SurfaceEmitter {
	return &SurfaceEmitter{particles: make([]Particle, 0), surface: s, manipulators: make([]Manipulator, 0)}
}

func (e *SurfaceEmitter) Particles() []Particle {
	return e.particles
}

func (e *SurfaceEmitter) AddManipulator(m Manipulator) {
	for i := range e.manipulators {
		if e.manipulators[i] == nil {
			e.manipulators[i] = m
		}
	}
	e.manipulators = append(e.manipulators, m)
}

func (e *SurfaceEmitter) Emit(frequency, life time.Duration, acceleration baukasten.Vector2, velocity baukasten.Vector2) {
	e.emitFrequency = frequency
	e.emitLife = life
	e.emitAcceleration = acceleration
	e.emitVelocity = velocity
}

func (e *SurfaceEmitter) Update(deltaTime time.Duration) {
	sec := float32(deltaTime.Seconds())
	if e.emitFrequency > 0 {
		e.emitTime += deltaTime
		for e.emitTime > 0 {
			e.EmitParticle(e.emitAcceleration, e.emitVelocity, e.emitLife, Alive)
			e.emitTime -= e.emitFrequency
		}
	}
	for _, m := range e.manipulators {
		m.Update(deltaTime, e.particles)
	}
	for i := range e.particles {
		if e.particles[i].Life <= 0 {
			e.particles[i].Condition = Dead
		}
		switch e.particles[i].Condition {
		case Dead, Paused:
			continue
		case Frozen:
			e.particles[i].Life -= deltaTime
			continue
		case Alive:
			// Euler method
			// v = velocity
			// a = acceleration
			// dt = delta time in seconds
			// x = position
			// v = v + a * dt
			// x = x + v * dt
			e.particles[i].Velocity.Accumulate(e.particles[i].Acceleration.Scaled(sec))
			e.particles[i].Position.Accumulate(e.particles[i].Velocity.Scaled(sec))
			e.particles[i].Life -= deltaTime
		default:
			// Unknown condition, set to dead
			e.particles[i].Condition = Dead
		}
	}
}

func (e *SurfaceEmitter) EmitParticle(acceleration, velocity baukasten.Vector2, life time.Duration, condition Condition) {
	for i := range e.particles {
		if e.particles[i].Condition == Dead {
			e.particles[i] = Par(acceleration, e.Position, velocity, life, condition)
			return
		}
	}
	e.particles = append(e.particles, Par(acceleration, e.Position, velocity, life, condition))
}

func (e *SurfaceEmitter) Draw() {
	for i := range e.particles {
		switch e.particles[i].Condition {
		case Frozen, Alive:
			e.surface.Draw(e.particles[i].Position.X, e.particles[i].Position.Y)
		}
	}
}
