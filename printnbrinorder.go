package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var len int = 0
	var x int = n
	var Rune []rune
	for {
		x /= 10
		len++
		if x <= 0 {
			break
		}
	}

	for {
		y := n % 10
		n /= 10
		Rune = append(Rune, rune(48+y))
		if n <= 0 {
			break
		}
	}
	for i := range Rune {
		for j := range Rune {
			if Rune[j] > Rune[i] {
				Rune[i], Rune[j] = Rune[j], Rune[i]
			}
		}
	}
	for i := 0; i < len; i++ {
		z01.PrintRune(rune(Rune[i]))
	}
}
