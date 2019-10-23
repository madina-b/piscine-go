package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	for i, word := range arguments {
		if i > 0 {
			for _, letter := range word {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}
	}
}
