package main

import (
	"os"

	"github.com/01-edu/z01"
)

func parseInteger(str string) int32 {
	var num int32

	for _, digit := range str {
		num = num*10 + int32(digit-'0')
	}

	return num
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 2 {
		return
	}
	flag := false

	if arguments[0] == "--upper" {
		flag = true
	}
	for _, arg := range arguments {
		if parseInteger(arg) <= 26 {
			if flag {
				z01.PrintRune(parseInteger(arg) + 64)
			} else {
				z01.PrintRune(parseInteger(arg) + 96)
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
