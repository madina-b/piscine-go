package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	lenArg := 0
	for i := range arguments {
		lenArg = i + 1
	}
	for i := lenArg - 1; i > 0; i-- {
		for _, letter := range arguments[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
