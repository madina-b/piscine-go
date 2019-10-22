package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(nb int) {
	if nb > 0 {
		digit := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i := nb; i > 0; i = i / 10 {
			remainder := i % 10
			digit[remainder] = digit[remainder] + 1
		}
		for i := '0'; i <= '9'; i++ {
			for j := 0; j < digit[i]; j++ {
				z01.PrintRune(i)
			}
		}
	}
}
