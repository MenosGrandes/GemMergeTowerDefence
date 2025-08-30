package grid

import (
	"example/hello/src/constants"
	"example/hello/src/drawable"
	"example/hello/src/vec"

	"github.com/hajimehoshi/ebiten/v2"
)

type Grid struct {
	Size  vec.Vec2UI16
	Cells []*Cell
}

func (g *Grid) DrawFlowField(screen *ebiten.Image) {
	for _, c := range g.Cells {
		// Skip if no direction
		if c.Direction.X == 0 && c.Direction.Y == 0 {
			continue
		}

		// Center of the cell

		x := c.ID % constants.ID(g.Size.Y)
		y := c.ID / constants.ID(g.Size.Y)
		cx := float64(x)*constants.SpriteSize + constants.SpriteSize/2
		cy := float64(y)*constants.SpriteSize + constants.SpriteSize/2
		// Draw arrow in that direction
		drawable.DrawSolidArrow(screen, cx, cy, c.Direction.X, c.Direction.Y, constants.SpriteSize*0.4, 0xFF00FF00)
	}
}
func NewGridEmpty(size vec.Vec2UI16) Grid {
	return Grid{Size: size, Cells: []*Cell{}}
}
func NewGrid(size vec.Vec2UI16) Grid {
	c := []*Cell{}
	_size := size.X * size.Y
	for id := range _size {
		c = append(c, NewCellPtr(constants.ID(id)))
	}
	return Grid{Size: size, Cells: c}
}

// MenosGrandes Width is Y right?
func GetCellIdBasedOnPosition(position, size vec.Vec2UI16) constants.ID {
	return constants.ID(position.X*size.Y + position.Y)
}
func GetPositionBasedOnCellId(id constants.ID, size vec.Vec2UI16) vec.Vec2UI16 {
	y := uint16(id) / size.Y
	x := uint16(id) % size.Y
	return vec.NewVec2UI16(x, y)
}

// MenosGrandes There is an X and Y miscalculated
func (g *Grid) GetNeighborsIds(c constants.ID) []constants.ID {

	cPosition := GetPositionBasedOnCellId(c, g.Size)
	result := make([]constants.ID, 0)
	if c <= constants.ID(g.Size.X*g.Size.Y) {
		/*
			Those are ID's of grid
			x →
			y ↓
				1	2	3	4
				5	6	7	8
				9	10	11	12
				13	14	15	16
		*/
		// top
		if cPosition.Y > 0 {
			result = append(result, GetCellIdBasedOnPosition(vec.NewVec2UI16(cPosition.X, cPosition.Y-1), g.Size))
		}
		// bottom
		if cPosition.Y < g.Size.X-1 {
			result = append(result, GetCellIdBasedOnPosition(vec.NewVec2UI16(cPosition.X, cPosition.Y+1), g.Size))
		}
		// left
		if cPosition.X > 0 {
			result = append(result, GetCellIdBasedOnPosition(vec.NewVec2UI16(cPosition.X-1, cPosition.Y), g.Size))
		}
		// right
		if cPosition.X < g.Size.Y-1 {
			result = append(result, GetCellIdBasedOnPosition(vec.NewVec2UI16(cPosition.X+1, cPosition.Y), g.Size))
		}
	}
	return result

}
