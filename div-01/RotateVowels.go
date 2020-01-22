package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
	arguments := os.Args
	activeArgs := arguments[1:]
	lenArgs := 0
	for _ = range arguments {
		lenArgs++
	}
	if lenArgs < 2 {
		z01.PrintRune('\n')
		return
	}
	var totalStr string

	for i := range activeArgs {
		totalStr = totalStr + activeArgs[i]
		if i < lenArgs-2 {
			totalStr = totalStr + " "
		}
	}
	totalRune := []rune(totalStr)
	totLen := 0
	for _ = range totalRune {
		totLen++
	}

	allLetters := []rune(totalStr)
	vowCount := 0
	for i := range allLetters {
		if allLetters[i] == 'a' || allLetters[i] == 'e' || allLetters[i] == 'i' || allLetters[i] == 'o' || allLetters[i] == 'u' || allLetters[i] == 'A' || allLetters[i] == 'E' || allLetters[i] == 'I' || allLetters[i] == 'O' || allLetters[i] == 'U' {
			vowCount++
		}
	}
	vowelRune := make([]rune, vowCount)
	j := 0
	for i := range allLetters {
		if allLetters[i] == 'a' || allLetters[i] == 'e' || allLetters[i] == 'i' || allLetters[i] == 'o' || allLetters[i] == 'u' || allLetters[i] == 'A' || allLetters[i] == 'E' || allLetters[i] == 'I' || allLetters[i] == 'O' || allLetters[i] == 'U' {
			vowelRune[j] = allLetters[i]
			j++
		}
	}
	for i := 0; i < vowCount/2; i++ {
		vow1 := vowelRune[i]
		vow2 := vowelRune[vowCount-1-i]
		vowelRune[i] = vow2
		vowelRune[vowCount-1-i] = vow1
	}
	j = 0
	for i := range allLetters {
		if allLetters[i] == 'a' || allLetters[i] == 'e' || allLetters[i] == 'i' || allLetters[i] == 'o' || allLetters[i] == 'u' || allLetters[i] == 'A' || allLetters[i] == 'E' || allLetters[i] == 'I' || allLetters[i] == 'O' || allLetters[i] == 'U' {
			allLetters[i] = vowelRune[j]
			j++
		}
		z01.PrintRune(allLetters[i])
	}
	z01.PrintRune('\n')

}