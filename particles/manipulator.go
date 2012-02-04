package particles

import (
	"time"
)

type Manipulator interface {
	Update(time.Duration, []Particle)
}
