package path

import (
	"example/hello/src/constants"
	"example/hello/src/grid"
	"example/hello/src/queue"
	"example/hello/src/vec"
	"fmt"
)

type BFS struct{}

// MenosGrandes there are a looot of bugs in here...
func CalcPath(g *grid.Grid, startCell, endCell grid.Cell) {
	frontier := queue.NewEmptyQueue[grid.Cell]()
	frontier.Enqueue(startCell)
	reached := make(map[constants.ID]bool)
	reached[startCell.ID] = true
	//fmt.Println("START")
	//fmt.Println(frontier)
	for !frontier.IsEmpty() {
		//fmt.Println("frontier not empty")

		current, _ := frontier.Dequeue()

		currentPosition := grid.GetPositionBasedOnCellId(current.ID, g.Size)
		fmt.Println("~!~!~!~!~!")
		fmt.Printf("current.ID = %d \n", current.ID)
		fmt.Printf("currentPosition = %q \n", currentPosition)

		neighbors := g.GetNeighborsIds(current.ID)
		for _, n := range neighbors {
			_, ok := reached[n]
			if ok {
			} else {
				fmt.Printf("ID = %d \n", n)
				nPos := grid.GetPositionBasedOnCellId(n, g.Size)
				fmt.Printf("nPos = %q \n", nPos)

				dx := float64(currentPosition.X - nPos.X)
				dy := float64(currentPosition.Y - nPos.Y)
				fmt.Printf("dx, dy = [%f - %f] \n", dx, dy)

				g.Cells[n].Direction = vec.NewVec2F64(dx, dy)
				frontier.Enqueue(*g.Cells[n])
				reached[n] = true
			}
		}
		fmt.Println("@@@@@")
	}

	// for next in graph.neighbors(current):
	// 	if next not in reached:
	// 		frontier.put(next)
	// 		reached[next] = True
	// 		}
}
