package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	var len int
	for i := range os.Args {
		len = i + 1
	}
	for i := len; i > 1; i++ {
		ar := []rune(os.Args[i])
		for _, k := range ar {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}
