package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for i := 0; i < len(str); i++ {
		z01.PrintRune(str[i])
	}
}
