package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		var arrInt []int
		for n != 0 {
			arrInt = append(arrInt, n%10)
			n = n / 10
		}
		for i := 1; i < len(arrInt); i++ {
			for j := 0; j < i; j++ {
				if arrInt[j] > arrInt[i] {
					arrInt[j], arrInt[i] = arrInt[i], arrInt[j]
				}
			}
		}
		for _, i := range arrInt {
			z01.PrintRune(rune(i + '0'))
		}
	}
}
