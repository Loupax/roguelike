package main

import (
	"github.com/beefsack/go-astar"
	"github.com/loupax/roguelike/lib/room"
)

type pfTile struct {
	room.Bits
	row   int
	col   int
	world *pfWorld
}
type pfWorld struct {
	tiles [][]*pfTile
}

func (t *pfTile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}

	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		rw, cl := t.row+offset[1], t.col+offset[0]
		if cl < 0 || cl >= len(t.world.tiles[0]) {
			continue
		}
		if rw < 0 || rw >= len(t.world.tiles) {
			continue
		}

		neighbors = append(neighbors, t.world.tiles[rw][cl])
	}
	return neighbors
}

func (t *pfTile) PathNeighborCost(to astar.Pather) float64 {
	if t.Bits.Has(room.Wall) {
		return 50000000
	}
	return 1
}

func (t *pfTile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*pfTile)
	absX := toT.row - t.row
	if absX < 0 {
		absX = -absX
	}
	absY := toT.col - t.col
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}
func pathfindTowards(w World, from, to *room.Entity) *pfTile {
	pfw := &pfWorld{}

	pfw.tiles = make([][]*pfTile, len(w.room))
	for i, rw := range w.room {
		pfw.tiles[i] = make([]*pfTile, len(rw))
		for j, cl := range rw {
			pfw.tiles[i][j] = &pfTile{
				Bits:  cl.Bits,
				row:   i,
				col:   j,
				world: pfw,
			}
		}
	}
	for _, a := range w.actors {
		pfw.tiles[a.Row][a.Col] = &pfTile{
			Bits:  a.Bits,
			row:   a.Row,
			col:   a.Col,
			world: pfw,
		}
	}
	p, _, found := astar.Path(pfw.tiles[from.Row][from.Col], pfw.tiles[to.Row][to.Col])

	if found {
		nsi := len(p) - 2 //Index of the next step in path
		if nsi > 0 {
			return p[nsi].(*pfTile)
		}
	}
	return nil

}
