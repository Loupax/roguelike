package cursor

import "fmt"

func ClearScreen() {
	fmt.Print("\033[2J")
}

func MoveTo(line, col int) {
	fmt.Printf("\033[%d;%dH", line, col)
}

func MoveLeft(i int) {
	fmt.Printf("\033[%dN", i)
}

func HideCursor() {
	fmt.Printf("\033[?25l")
}

func ShowCursor() {
	fmt.Printf("\033[?25h")
}

func Bell() {
	fmt.Printf("\033\007")
}

func PrintAt(c rune, line, col int) {
	fmt.Printf("\0337")
	fmt.Printf("\033[%d;%dH%s", line, col, string(c))
	fmt.Printf("\0338")
}
