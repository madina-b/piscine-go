package main

import (
	
	"fmt"
)

// func main() {

// 	fmt.Println(AtoiBase("bcbbbbaabz", "abc"))
// 	fmt.Println(AtoiBase("0001", "01"))
// 	fmt.Println(AtoiBase("00", "01"))
// 	fmt.Println(AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
// 	fmt.Println(AtoiBase("AAho?Ao", "WhoAmI?"))
// 	fmt.Println(AtoiBase("thisinputshouldnotmatter", "abca"))
// 	fmt.Println(AtoiBase("125", "0123456789"))
// 	fmt.Println(AtoiBase("uoi", "choumi"))
// 	fmt.Println(AtoiBase("bbbbbab", "-ab"))

// }

// 12016
// 1
// 0
// 11557277891
// 406772
// 0
// 125
// 125
// 0

func AtoiBase(s string, base string) int {
	sRune := []rune(s)
	baseRune := []rune(base)
	sLen := 0
	baseLen := 0
	n := 0
	index := 0
	sign := 1
	for _ = range sRune {
		sLen++
		if sRune[0] == '-' {
			sign = -1
		}
	}

	for i := range baseRune {
		baseLen++
		if baseRune[i] == '+' || baseRune[i] == '-' {
			return 0
		}
	}
	for i := 0; i < sLen; i++ {
		for j := 0; j < baseLen; j++ {
			if sRune[i] == baseRune[j] {
				index = j
				n = n*baseLen + index
				break
			}
		}
	}
	return n * sign
}
