package main

import 
	
	"fmt"


// func main() {

// 	result := ConvertBase("4506C", "0123456789ABCDEF", "choumi")
// 	fmt.Println(result)
// 	result = ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
// 	fmt.Println(result)
// 	result = ConvertBase("256850", "0123456789", "01")
// 	fmt.Println(result)
// 	result = ConvertBase("uuhoumo", "choumi", "Zone01")
// 	fmt.Println(result)
// 	result = ConvertBase("683241", "0123456789", "0123456789")
// 	fmt.Println(result)

// }

func ConvertBase(s, baseFrom, baseTo string) string {

	sRune := []rune(s)
	sLen := 0
	nbr := 0
	index := 0
	// sign:=1
	for _ = range sRune {
		sLen++
		// if sRune[0] == '-' {
		// 	sign = -1
		// }
	}

	baseFrLen := 0
	baseFrRune := []rune(baseFrom)
	for i := range baseFrRune {
		baseFrLen++
		if baseFrRune[i] == '+' || baseFrRune[i] == '-' {
			return "0"
		}
	}
	for i := 0; i < sLen; i++ {
		for j := 0; j < baseFrLen; j++ {
			if sRune[i] == baseFrRune[j] {
				index = j
				nbr = nbr*baseFrLen + index
				break
			}
		}
	}

	baseTorune := []rune(baseTo)
	lenbaseTo := 0

	for _ = range baseTorune {
		lenbaseTo++
	}
	div := lenbaseTo
	if lenbaseTo <= 1 {

		return "NV"
	}
	for i := 0; i <= lenbaseTo-1; i++ {
		if baseTorune[i] == '+' || baseTorune[i] == '-' {

			return "NV"
		}
	}
	for i := 0; i <= lenbaseTo-2; i++ {

		for j := i + 1; j <= lenbaseTo-1; j++ {
			if baseTorune[i] == baseTorune[j] {

				return "NV"
			}
		}
	}

	var result string
	for i := nbr; i > 0; i = i / div {
		digit := baseTorune[i%div]
		result = string(digit) + result
	}
	if nbr < 0 {
		nbr = nbr * -1
		result = "-" + result
	}
	return result
}
