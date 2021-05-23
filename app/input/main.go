package main

import (
	"context"
	"fmt"
	"time"

	"github.com/loupax/roguelike/lib/cursor"
	"github.com/loupax/roguelike/lib/input"
)

func main() {
	done, err := input.SetupTerminal(0)
	if err != nil {
		panic(err)
	}
	defer done()

	ctx, d := context.WithTimeout(context.Background(), 5*time.Second)
	defer d()
	k := input.NewKeyboard(ctx)

	var b rune
	for {
		err := k.KeyPress(&b)
		cursor.MoveLeft(100)
		fmt.Printf("%s", string(b))
		if err != nil {
			panic(err)
		}
		if b == 'q' {
			return
		}
	}
}
