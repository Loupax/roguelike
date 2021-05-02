package main

import (
	"fmt"
	"github.com/loupax/roguelike/lib/room"
)

func main() {
	// fmt.Println(room.GenerateRectangle(20,10).Render())
	bg := room.GenerateRectangle(25, 70, room.Walkable, room.Wall)
	merge := bg.Stamp(10, 5, room.GenerateRectangle(70, 5, room.Walkable, room.Wall))
	fmt.Println(merge.Render())
	// fmt.Println(merge.Stamp(5, 5, room.GenerateRectangle(20, 10, room.Walkable, room.Wall)).Render())
	// fmt.Println(room.GenerateRectangle(20,10, room.Walkable, room.Wall).Stamp(5,5, merge).Render())
}
