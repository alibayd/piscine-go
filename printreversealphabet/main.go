package main

import "github.com/01-edu/z01"

func main() {
	var b rune
	for b = 122; b >= 97; b-- {
		z01.PrintRune(b)
	}
	z01.PrintRune('\n')
}
