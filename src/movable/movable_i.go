package movable

import (
	"example/hello/src/constants"
)

type MovableI interface {
	Move(newPosition constants.Position)
}
