package main

import "fmt"

// func main() {

// 	s9 := ""
// 	s10 := "-"
// 	s11 := "--123"
// 	s12 := "1"
// 	s13 := "-3"
// 	s14 := "8292"
// 	s15 := "9223372036854775807"
// 	s16 := "-9223372036854775808"

// 	n9 := Atoi(s9)
// 	n10 := Atoi(s10)
// 	n11 := Atoi(s11)
// 	n12 := Atoi(s12)
// 	n13 := Atoi(s13)
// 	n14 := Atoi(s14)
// 	n15 := Atoi(s15)
// 	n16 := Atoi(s16)

// 	fmt.Println(n9)
// 	fmt.Println(n10)
// 	fmt.Println(n11)
// 	fmt.Println(n12)
// 	fmt.Println(n13)
// 	fmt.Println(n14)
// 	fmt.Println(n15)
// 	fmt.Println(n16)

// }

func Atoi(str string) int {
	runes := []rune(str)
	index := 0
	sign:=1
	var n int
	var result int
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
			n = n*10 + int(runes[i]-'0')
		} else {
			return 0
		}
	}
	result = sign * n
	return result
}
