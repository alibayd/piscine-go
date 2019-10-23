package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	for i := range os.Args {
		ar := []rune(os.Args[i])
		for _, j := range ar {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
