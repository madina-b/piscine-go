package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i < '9'; i++ {
		for j := '0'; j < '9'; j++ {
			z01.PrintRune(i)
			z01.PrintRune(j)
			z01.PrintRune(44)
			z01.PrintRune(32)
		}
	}
}
