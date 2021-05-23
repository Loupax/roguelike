package main

import (
	"fmt"
	"github.com/loupax/roguelike/lib/room"
	"os"
	"os/exec"
)

func main() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	seedRoom := room.GenerateRectangle(10, 10, room.Walkable, room.Wall)

	ch := make(chan string)
	go func(ch chan string) {
		var s = make([]byte, 1)
		for {
			os.Stdin.Read(s)
			ch <- string(s)
		}
	}(ch)
	ClearEntireScreen()
	fmt.Println(seedRoom.Render())
	for {
		s := <-ch
		ClearEntireScreen()
		fmt.Printf("Got input %s\n", s)
		fmt.Println(seedRoom.Render())
	}
}

func ClearEntireScreen() {
	fmt.Print("\033[H\033[2J")
}
