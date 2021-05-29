package main

import (
	"context"

	"github.com/loupax/roguelike/lib/cursor"
	"github.com/loupax/roguelike/lib/input"
	"github.com/loupax/roguelike/lib/room"
)

const maze = `
############################
#            ##            #
# #### ##### ## ##### #### #
# #  # #   # ## #   # #  # #
# #### ##### ## ##### #### #
#                          #
# #### ## ######## ## #### #
# #### ## ######## ## #### #
#      ##    ##    ##      #
###### ##### ## ##### ######
     # ##### ## ##### #     
     # ##          ## #     
     # ## ###  ### ## #     
###### ## #      # ## ######
          #      #          
###### ## #      # ## ######
     # ## ######## ## #     
     # ##          ## #     
     # ## ######## ## #     
###### ## ######## ## ######
#            ##            #
# #### ##### ## ##### #### #
# #### ##### ## ##### #### #
#   ##                ##   #
### ## ## ######## ## ## ###
### ## ## ######## ## ## ###
#      ##    ##    ##      #
# ########## ## ########## #
# ########## ## ########## #
#                          #
############################`

func Render(m room.Room) error {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			var tile rune
			if m[i][j].Has(room.Walkable) {
				tile = ' '
			}

			if m[i][j].Has(room.Wall) {
				tile = '#'
			}
			cursor.PrintAt(string(tile), i+1, j+1)
		}
	}
	return nil
}

type World struct {
	room   room.Room
	actors room.Entities
}

func (w World) Render() {
	Render(w.room)
	w.actors.Render()
}

func main() {
	cursor.ClearScreen()
	cursor.HideCursor()
	defer cursor.ShowCursor()
	done, err := input.SetupTerminal(0)
	if err != nil {
		panic(err)
	}

	defer done()

	hero := room.Entity{
		Row:       17,
		Col:       14,
		Face:      '@',
		ColorCode: 93,
		Bits:      room.Wall,
	}
	inky := room.Entity{
		Row:       13,
		Col:       13,
		Face:      'I',
		ColorCode: 34,
		Bits:      room.Wall,
	}

	blinky := room.Entity{
		Row:       15,
		Col:       12,
		Face:      'b',
		ColorCode: 91,
		Bits:      room.Wall,
	}
	pinky := room.Entity{
		Row:       15,
		Col:       14,
		Face:      'O',
		ColorCode: 95,
		Bits:      room.Wall,
	}
	clyde := room.Entity{
		Row:       15,
		Col:       15,
		Face:      'C',
		ColorCode: 33,
		Bits:      room.Wall,
	}
	objs := room.Entities{
		&hero,
		&pinky,
		&inky,
		&blinky,
		&clyde,
	}

	r := room.FromString(maze)
	w := World{
		room:   r,
		actors: objs,
	}
	w.Render()

	defer cursor.MoveTo(len(r)+1, 0)
	k := input.NewKeyboard(context.Background())
	var b rune
	for {
		err = k.KeyPress(&b)
		if err != nil {
			panic(err)
		}
		if b == 'q' {
			return
		}
		if b == 'j' {
			rw, cl := hero.Row+1, hero.Col
			if canWalkInto(w, rw, cl) {
				hero.Row++
			} else {
				cursor.Bell()
			}
		}
		if b == 'k' {
			rw, cl := hero.Row-1, hero.Col
			if canWalkInto(w, rw, cl) {
				hero.Row--
			} else {
				cursor.Bell()
			}
		}
		if b == 'h' {
			rw, cl := hero.Row, hero.Col-1
			if cl < 0 {
				cl = len(r[rw]) - 1
			}
			if canWalkInto(w, rw, cl) {
				hero.Col = cl
			} else {
				cursor.Bell()
			}
		}
		if b == 'l' {
			rw, cl := hero.Row, hero.Col+1
			if cl == len(r[rw]) {
				cl = 0
			}
			if canWalkInto(w, rw, cl) {
				hero.Col = cl
			} else {
				cursor.Bell()
			}
		}
		if t := pathfindTowards(w, &inky, &hero); t != nil {
			inky.Row = t.row
			inky.Col = t.col
		}
		if t := pathfindTowards(w, &blinky, &hero); t != nil {
			blinky.Row = t.row
			blinky.Col = t.col
		}
		if t := pathfindTowards(w, &pinky, &hero); t != nil {
			pinky.Row = t.row
			pinky.Col = t.col
		}
		if t := pathfindTowards(w, &clyde, &hero); t != nil {
			clyde.Row = t.row
			clyde.Col = t.col
		}
		w.Render()
	}
}

func canWalkInto(w World, row, col int) bool {
	if w.room[row][col].Has(room.Wall) {
		return false
	}

	return !w.actors.At(row, col).Has(room.Wall)
}

var debugRow int

func DebugAt(msg string, row int) {
	debugRow++
	cursor.PrintAt(msg, row, 50)
}
func Debug(msg string) {
	debugRow++
	cursor.PrintAt(msg, debugRow, 50)
}
