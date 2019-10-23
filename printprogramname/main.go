package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	nm := os.Args[0]
	for _, i := range nm {
		z01.PrintRune(i)
	}
}
