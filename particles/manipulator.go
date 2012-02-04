package particles

import (
	"time"
)

type Manipulator interface {
	Manipulate(*Particle)
	Update(time.Duration)
}
