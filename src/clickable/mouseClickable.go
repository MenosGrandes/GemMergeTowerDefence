package clickable

import "example/hello/src/vec"

type MouseClickableI interface {
	IsClicked(mouseClickPos vec.Vec2F64) bool
}
