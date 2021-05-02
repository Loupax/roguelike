package room

import (
	"strings"
)

type Room [][]Bits
type Bits uint8

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


func upsizeRoomTo(r Room, h int, w int)Room{
	out := make(Room, h)
	for i:= 0 ; i<h;i++{
		out[i]=make([]Bits, w)
		if i >= len(r){
			continue
		}
		copy(out[i], r[i])
	}  
	return out
}
func max(x,y int)int{
	if x > y{
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
	
	for i := range bR{
		for j := range bR[i]{
			out[i+rw][j+cl]= bR[i][j]
		}

	}

	return out
}

func (r Room) Render() string {
	var b strings.Builder
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			if r[i][j].Has(Wall) {
				b.WriteRune('#')
			}
			if r[i][j].Has(Walkable) {
				b.WriteRune('.')
			}
			if r[i][j].Has(Nothing) {
				b.WriteRune(' ')
			}

		}
		b.WriteRune('\n')
	}

	return b.String()
}

func GenerateRectangle(w int, h int, fill Bits, stroke Bits) Room {
	out := make(Room, h)

	for i := 0; i < h; i++ {
		if i == 0 || i == h-1 {
			out[i] = repeatBit(w, stroke)
		} else {
			tmp := repeatBit(w, fill)
			tmp[0] = stroke
			tmp[len(tmp)-1]=stroke
			out[i]=tmp
		}

	}

	return out
}

func repeatBit(a int, b Bits) []Bits {
	out := make([]Bits, a)
	for i := 0; i < a; i++ {
		out[i] = b
	}
	return out
}
