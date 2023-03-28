package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runes := os.Args[1:]

	for i := len(runes) - 1; i >= 0; i-- {
		for _, letter := range runes[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
