package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func intToString(num int) string {
	result := ""

	for ; num > 0; num /= 10 {
		result = string(rune('0'+num%10)) + result
	}

	return result
}

func main() {
	points := &point{}

	setPoint(points)

	result := "x = " + intToString(points.x) + ", y = " + intToString(points.y)

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
