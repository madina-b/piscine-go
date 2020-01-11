package main

import (
	// "fmt"
	"os"
)

// func main() {
// 	fmt.Println(Atoi())
// }

func Atoi() int64 {
	arguments := os.Args
	if len(arguments) != 2 {
		return 0
	}
	str := arguments[1]
	runes := []rune(str)
	index := 0
	var sign int64
	var n int64
	var result int64
	for i := range runes {
		if runes[i] <= '9' && runes[i] >= '0' {
			index = i
			break
		}
	}
	if index > 1 {
		return 0
	} else if index == 1 {
		if runes[0] == '-' {
			sign = -1
		} else if runes[0] == '+' {
			sign = 1
		} else {
			return 0
		}
	}

	for i := index; i < len(str); i++ {
		if runes[i] <= '9' && runes[i] >= '0' {
			n = n*10 + int64(runes[i]-'0')
		} else {
			return 0
		}
	}
	result = sign * n
	return result
}
