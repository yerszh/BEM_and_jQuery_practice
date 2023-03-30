package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	for i := 1; i < len(arguments); i++ {
		for j := 0; j < i; j++ {
			if arguments[i] < arguments[j] {
				arguments[i], arguments[j] = arguments[j], arguments[i]
			}
		}
	}

	for _, arg := range arguments {
		for _, a := range arg {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
