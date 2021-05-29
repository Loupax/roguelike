package room

import (
	"math"
	"strings"

	"github.com/loupax/roguelike/lib/cursor"
)

type Room [][]Tile
type Bits uint8
type Tile struct {
	Bits
}

func (b Bits) Clear(f Bits) Bits {
	return b &^ f
}
func (b Bits) Has(f Bits) bool {
	return b&f != 0
}

const Nothing Bits = 0
const (
	Wall Bits = 1 << iota
	Walkable
)

func upsizeRoomTo(r Room, h int, w int) Room {
	out := make(Room, h)
	for i := 0; i < h; i++ {
		out[i] = make([]Tile, w)
		if i >= len(r) {
			continue
		}
		copy(out[i], r[i])
	}
	return out
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func (r Room) Stamp(rw int, cl int, bR Room) Room {
	out := upsizeRoomTo(
		r,
		max(len(r), len(bR)+rw),
		max(len(r[0]), len(bR[0])+cl),
	)

	for i := range bR {
		for j := range bR[i] {
			// Avoid overlapping walls by not
			// overriding floors
			if !out[i+rw][j+cl].Has(Walkable) {
				out[i+rw][j+cl] = bR[i][j]
			}

		}

	}

	return out
}

func Render(r Room) {
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			var tile rune
			if r[i][j].Has(Wall) {
				tile = '#'
			}
			if r[i][j].Has(Walkable) {
				tile = '.'
			}
			if r[i][j].Has(Nothing) {
				tile = ' '
			}
			cursor.PrintAt(string(tile), i+1, j+1)
		}
	}
}
func dist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}
func MakeCircle(r int, fill Bits, stroke Bits) Room {
	c := MakeRectangle((2*r)+1, (2*r)+1, 0, 0)
	for row := range c {
		for col := range c[row] {
			d := dist(col, row, r, r)
			if d < float64(r) {
				c[row][col] = Tile{Bits: fill}
			}
			if int(d) == r {
				c[row][col] = Tile{Bits: stroke}
			}
		}
	}
	return c
}

func MakeRectangle(w int, h int, fill Bits, stroke Bits) Room {
	out := make(Room, h)

	for i := 0; i < h; i++ {
		if i == 0 || i == h-1 {
			out[i] = repeatBit(w, Tile{Bits: stroke})
		} else {
			tmp := repeatBit(w, Tile{Bits: fill})
			tmp[0] = Tile{Bits: stroke}
			tmp[len(tmp)-1] = Tile{Bits: stroke}
			out[i] = tmp
		}

	}

	return out
}

func repeatBit(a int, b Tile) []Tile {
	out := make([]Tile, a)
	for i := 0; i < a; i++ {
		out[i] = b
	}
	return out
}

func FromString(r string) Room {
	w, h, cw := 0, 0, 0
	r = strings.TrimSpace(r)

	for _, char := range r {
		if char == '\n' {
			h++
			cw = 0
			continue
		}

		if cw > w {
			w = cw
		}
		cw++
	}

	rm := make(Room, h+1)

	for ri, rowStr := range strings.Split(r, "\n") {
		rm[ri] = make([]Tile, w+1)
		for ci, char := range rowStr {
			if char == ' ' {
				rm[ri][ci] = Tile{Bits: Walkable}
			} else {
				rm[ri][ci] = Tile{Bits: Wall}
			}
		}
	}
	return rm
}
