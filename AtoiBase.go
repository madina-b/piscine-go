package main

import "fmt"

// func main() {
// 	fmt.Println(AtoiBase("125", "0123456789"))
// 	fmt.Println(AtoiBase("1111101", "01"))
// 	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
// 	fmt.Println(AtoiBase("uoi", "choumi"))
// 	fmt.Println(AtoiBase("bbbbbab", "ab"))
// }

func AtoiBase(s string, base string) int {
	sRune := []rune(s)
	baseRune := []rune(base)
	sLen := 0
	baseLen := 0
	n := 0
	index := 0
	for _ = range sRune {
		sLen++
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
			// else {
			// 	return 0
			// }
		}
	}
	return n
}
