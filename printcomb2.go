package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for n1 := '0'; n1 <= '9'; n1++ {
		for n2 := '0'; n2 <= '9'; n2++ {
			for n3 := '0'; n3 <= '9'; n3++ {
				for n4 := '0'; n4 <= '9'; n4++ {
					if n1 == n3 && n2 < n4 || n1 < n3 {
						z01.PrintRune(n1)
						z01.PrintRune(n2)
						z01.PrintRune(' ')
						z01.PrintRune(n3)
						z01.PrintRune(n4)
						if n1 == '9' && n2 == '8' {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
