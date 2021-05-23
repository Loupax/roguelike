package main

import (
	"github.com/loupax/roguelike/lib/cursor"
	"github.com/loupax/roguelike/lib/input"
	"github.com/loupax/roguelike/lib/room"
)

func main() {
	cursor.ClearScreen()
	done, err := input.SetupTerminal(0)
	if err != nil {
		panic(err)
	}

	defer done()

	hero := room.Entity{
		Row:  7,
		Col:  76,
		Face: '@',
		Bits: room.Wall,
	}

	c := room.MakeCircle(7, room.Walkable, room.Wall)
	c = c.Stamp(4, 2, room.MakeRectangle(70, 7, room.Walkable, room.Wall))
	c = c.Stamp(3, 69, room.MakeCircle(4, room.Walkable, room.Wall))
	c.Render()
	hero.Render()
	cursor.MoveTo(10000, 0)
}
