package main

import "github.com/01-edu/z01"

func main() {
	var b rune
	for b = 'z'; b >= 'a'; b-- {
		z01.PrintRune(b)
	}
	z01.PrintRune('\n')
}
