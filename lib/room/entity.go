package room

import (
	"fmt"

	"github.com/loupax/roguelike/lib/cursor"
)

type Entity struct {
	Row  int
	Col  int
	Bits Bits
	Face rune
}

type Entities []*Entity

func (os Entities) Has(b Bits) bool {
	for _, t := range os {
		if t.Bits.Has(b) {
			return true
		}
	}
	return false
}
func (os Entities) At(line, col int) Entities {
	out := make(Entities, 0)
	for _, t := range os {
		if t.Row == line && t.Col == col {
			out = append(out, t)
		}
	}
	return out
}

func (os Entities) Render() {
	for _, o := range os {
		o.Render()
	}
}
func (o Entity) Render() {
	cursor.MoveTo(o.Row+1, o.Col+1)
	fmt.Printf("%s", string(o.Face))
}
