package grid

import (
	"example/hello/src/constants"
	"example/hello/src/vec"
	"slices"
	"testing"
)

/*
0	1	2	3
4	5	6	7
8	9	10	11
12	13	14	15
*/
func TestCreationOfGrid(t *testing.T) {
	g := NewGrid(vec.NewVec2UI16(4, 4))
	idToCheck := constants.ID(0)
	for _, c := range g.Cells {
		if c.ID != idToCheck {
			t.Errorf("Wrong ID - got %d, wanted %d", c.ID, idToCheck)
		}
		idToCheck++
	}
}
func TestGetNeighborsIds(t *testing.T) {
	g := NewGrid(vec.NewVec2UI16(4, 4))
	{
		pos := GetPositionBasedOnCellId(10, g.Size)
		wanted := vec.NewVec2UI16(2, 2)
		if pos != wanted {
			t.Errorf("Wrong position - got %d, wanted %d", pos, wanted)
		}
	}
	{
		pos := GetPositionBasedOnCellId(12, g.Size)
		wanted := vec.NewVec2UI16(0, 3)
		if pos != wanted {
			t.Errorf("Wrong position - got %d, wanted %d", pos, wanted)
		}
	}
	{
		pos := GetPositionBasedOnCellId(1, g.Size)
		wanted := vec.NewVec2UI16(1, 0)
		if pos != wanted {
			t.Errorf("Wrong position - got %d, wanted %d", pos, wanted)
		}
	}
	{
		pos := GetPositionBasedOnCellId(13, g.Size)
		wanted := vec.NewVec2UI16(1, 3)
		if pos != wanted {
			t.Errorf("Wrong position - got %d, wanted %d", pos, wanted)
		}
	}
	{
		pos := GetPositionBasedOnCellId(0, g.Size)
		wanted := vec.NewVec2UI16(0, 0)
		if pos != wanted {
			t.Errorf("Wrong position - got %d, wanted %d", pos, wanted)
		}
	}
	{
		g := NewGrid(vec.NewVec2UI16(17, 20))
		{
			pos := GetPositionBasedOnCellId(0, g.Size)
			wanted := vec.NewVec2UI16(0, 0)
			if pos != wanted {
				t.Errorf("Wrong position - got %d, wanted %d", pos, wanted)
			}
		}
		{
			pos := GetPositionBasedOnCellId(19, g.Size)
			wanted := vec.NewVec2UI16(19, 0)
			if pos != wanted {
				t.Errorf("Wrong position - got %d, wanted %d", pos, wanted)
			}
		}
		{
			pos := GetPositionBasedOnCellId(20, g.Size)
			wanted := vec.NewVec2UI16(0, 1)
			if pos != wanted {
				t.Errorf("Wrong position - got %d, wanted %d", pos, wanted)
			}
		}
	}
}
func TestGetCellIdBasedOnPosition(t *testing.T) {
	g := NewGrid(vec.NewVec2UI16(17, 20))
	{
		got := GetCellIdBasedOnPosition(vec.NewVec2UI16(0, 0), g.Size)
		wanted := constants.ID(0)
		if got != wanted {
			t.Errorf("got %d, wanted %d", got, wanted)
		}
	}
	{
		got := GetCellIdBasedOnPosition(vec.NewVec2UI16(0, 19), g.Size)
		wanted := constants.ID(19)
		if got != wanted {
			t.Errorf("got %d, wanted %d", got, wanted)
		}
	}
	{
		got := GetCellIdBasedOnPosition(vec.NewVec2UI16(1, 0), g.Size)
		wanted := constants.ID(20)
		if got != wanted {
			t.Errorf("got %d, wanted %d", got, wanted)
		}
	}
}
func TestGetNeigborsId(t *testing.T) {
	g := NewGrid(vec.NewVec2UI16(4, 4))
	{
		g := NewGrid(vec.NewVec2UI16(17, 20))

		neib := g.GetNeighborsIds(0)
		if len(neib) != 2 {
			t.Errorf("Should have 3 neigbours, but contains %d", neib)
		}
		if !slices.Contains(neib, 1) {
			t.Errorf("Should contain neigh with ID 1 , but contains %d", neib)
		}
		if !slices.Contains(neib, 20) {
			t.Errorf("Should contain neigh with ID 20 , but contains %d", neib)
		}
	}
	{
		neib := g.GetNeighborsIds(0)
		if len(neib) != 2 {
			t.Errorf("Should have two neigbours, but contains %d", neib)
		}
		if !slices.Contains(neib, 1) {
			t.Errorf("Should contain neigh with ID 1 , but contains %d", neib)
		}
		if !slices.Contains(neib, 4) {
			t.Errorf("Should contain neigh with ID 4 , but contains %d", neib)
		}
	}
	{
		neib := g.GetNeighborsIds(10)
		if len(neib) != 4 {
			t.Errorf("Should have 4 neigbours, but contains %d", neib)
		}
		if !slices.Contains(neib, 6) {
			t.Errorf("Should contain neigh with ID 6 , but contains %d", neib)
		}
		if !slices.Contains(neib, 9) {
			t.Errorf("Should contain neigh with ID 9 , but contains %d", neib)
		}
		if !slices.Contains(neib, 14) {
			t.Errorf("Should contain neigh with ID 14 , but contains %d", neib)
		}
		if !slices.Contains(neib, 11) {
			t.Errorf("Should contain neigh with ID 11 , but contains %d", neib)
		}
	}
	{
		neib := g.GetNeighborsIds(15)
		if len(neib) != 2 {
			t.Errorf("Should have 2 neigbours, but contains %d", neib)
		}
		if !slices.Contains(neib, 11) {
			t.Errorf("Should contain neigh with ID 11 , but contains %d", neib)
		}
		if !slices.Contains(neib, 14) {
			t.Errorf("Should contain neigh with ID 14 , but contains %d", neib)
		}

	}
	{
		neib := g.GetNeighborsIds(12)
		if len(neib) != 2 {
			t.Errorf("Should have 2 neigbours, but contains %d", neib)
		}
		if !slices.Contains(neib, 8) {
			t.Errorf("Should contain neigh with ID 8 , but contains %d", neib)
		}
		if !slices.Contains(neib, 13) {
			t.Errorf("Should contain neigh with ID 13 , but contains %d", neib)
		}

	}
	{
		neib := g.GetNeighborsIds(3)
		if len(neib) != 2 {
			t.Errorf("Should have 2 neigbours, but contains %d", neib)
		}
		if !slices.Contains(neib, 2) {
			t.Errorf("Should contain neigh with ID 2 , but contains %d", neib)
		}
		if !slices.Contains(neib, 7) {
			t.Errorf("Should contain neigh with ID 7 , but contains %d", neib)
		}

	}
	{
		neib := g.GetNeighborsIds(8)
		if len(neib) != 3 {
			t.Errorf("Should have 3 neigbours, but contains %d", neib)
		}
		if !slices.Contains(neib, 4) {
			t.Errorf("Should contain neigh with ID 4 , but contains %d", neib)
		}
		if !slices.Contains(neib, 9) {
			t.Errorf("Should contain neigh with ID 9 , but contains %d", neib)
		}
		if !slices.Contains(neib, 12) {
			t.Errorf("Should contain neigh with ID 12 , but contains %d", neib)
		}
	}
}
