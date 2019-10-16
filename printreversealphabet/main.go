package main

import "github.com/01-edu/z01"

func main() {
	var i rune
	for i = 'z'; i <= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
