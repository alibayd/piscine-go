package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for i = 0; i < 12; i++ {
		j = str[i]
		z01.PrintRune(j)
	}
}
