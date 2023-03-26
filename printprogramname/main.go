package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runes := os.Args[0]

	for _, letter := range runes {
		if letter == '.' || letter == '/' {
			continue
		}
		z01.PrintRune(letter)
	}
	z01.PrintRune(rune('\n'))
}
