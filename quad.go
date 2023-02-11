package piscine

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	PrintQuad(x, y, '-', '|', 'o', 'o', 'o', 'o')
}

func QuadB(x, y int) {
	PrintQuad(x, y, '*', '*', '/', '\\', '\\', '/')
}

func QuadC(x, y int) {
	PrintQuad(x, y, 'B', 'B', 'A', 'A', 'C', 'C')
}

func QuadD(x, y int) {
	PrintQuad(x, y, 'B', 'B', 'A', 'C', 'A', 'C')
}

func QuadE(x, y int) {
	PrintQuad(x, y, 'B', 'B', 'A', 'C', 'C', 'A')
}

func PrintQuad(width, height int, edge1, edge2, topLeft, topRight, botLeft, botRight rune) {
	if width < 0 || height < 0 {
		return
	} else {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				switch {
				case y == 0 && x == 0:
					z01.PrintRune(topLeft)
				case y == 0 && x == width-1:
					z01.PrintRune(topRight)
				case y == height-1 && x == 0:
					z01.PrintRune(botLeft)
				case y == height-1 && x == width-1:
					z01.PrintRune(botRight)
				case y == 0 || y == height-1:
					z01.PrintRune(edge1)
				case x == 0 || x == width-1:
					z01.PrintRune(edge2)
				default:
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
