package main

import (
	"context"

	"github.com/loupax/roguelike/lib/cursor"
	"github.com/loupax/roguelike/lib/input"
	"github.com/loupax/roguelike/lib/room"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cursor.ClearScreen()
	defer cursor.MoveTo(10000, 0)

	r := room.MakeRectangle(10, 10, room.Walkable, room.Wall)
	r = r.Stamp(5, 5, room.MakeRectangle(10, 10, room.Walkable, room.Wall))

	done, err := input.SetupTerminal(0)
	if err != nil {
		return err
	}
	defer done()

	cursor.HideCursor()
	defer cursor.ShowCursor()

	k := input.NewKeyboard(context.TODO())

	hero := room.Entity{
		Row:  1,
		Col:  1,
		Face: '@',
		Bits: room.Wall,
	}

	var b rune
	room.Render(r)
	objs := room.Entities{
		&room.Entity{
			Row:  5,
			Col:  5,
			Face: '(',
			Bits: room.Wall,
		},
		&room.Entity{
			Row:  8,
			Col:  8,
			Face: 'f',
			Bits: room.Walkable,
		},
		&hero,
	}
	objs.Render()
	for {
		err := k.KeyPress(&b)
		if err != nil {
			return err
		}
		if b == 'q' {
			return nil
		}
		if b == 'j' {
			rw, cl := hero.Row+1, hero.Col
			if !r[rw][cl].Has(room.Wall) && !objs.At(rw, cl).Has(room.Wall) {
				hero.Row++
			} else {
				cursor.Bell()
			}
		}
		if b == 'k' {
			rw, cl := hero.Row-1, hero.Col
			if !r[rw][cl].Has(room.Wall) && !objs.At(rw, cl).Has(room.Wall) {
				hero.Row--
			} else {
				cursor.Bell()
			}
		}
		if b == 'h' {
			rw, cl := hero.Row, hero.Col-1
			if !r[rw][cl].Has(room.Wall) && !objs.At(rw, cl).Has(room.Wall) {
				hero.Col--
			} else {
				cursor.Bell()
			}
		}
		if b == 'l' {
			rw, cl := hero.Row, hero.Col+1
			if !r[rw][cl].Has(room.Wall) && !objs.At(rw, cl).Has(room.Wall) {
				hero.Col++
			} else {
				cursor.Bell()
			}
		}
		room.Render(r)
		objs.Render()
	}
}
