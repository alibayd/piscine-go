package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var a rune
	var b rune
	var c rune

	for a = 48; a <= 57; a++ {
		for b = 48; b <= 57; b++ {
			for c = 48; c <= 57; c++ {
				if a < b && b < c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if a <= 54 {
						z01.PrintRune(44)
					}
				}

			}
		}
	}
	z01.PrintRune(10)
}
