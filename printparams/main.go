package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runes := os.Args[1:]
	for _, word := range runes {
		for _, letter := range word {
			z01.PrintRune(letter)
		}
		z01.PrintRune(rune('\n'))
	}
}
