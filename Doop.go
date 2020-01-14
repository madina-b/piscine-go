//Doop

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args
	lenArgs := 0
	var EndResult int64
	error1 := "No division by 0"
	errorRune1 := []rune(error1)
	error2 := "No modulo by 0"
	errorRune2 := []rune(error2)
	for _ = range arguments {
		lenArgs++
	}

	if lenArgs != 4 {
		return
	}

	var value1 int64
	var value2 int64
	if IsValid(arguments[1]) && IsValid(arguments[3]) {
		value1 = Atoi(arguments[1])
		value2 = Atoi(arguments[3])
	} else {
		z01.PrintRune('0')
		return
	}

	var upperBound int64
	upperBound = 9223372036854775807
	var lowerBound int64
	lowerBound = -9223372036854775807

	//condition if it is not int
	if value1 > upperBound || value1 < lowerBound {
		z01.PrintRune('0')
		return

	} else if value2 > upperBound || value2 < lowerBound {
		z01.PrintRune('0')
		return

	}

	//case1 - "+" addition
	if arguments[2] == "+" {
		EndResult = value1 + value2

		if value1 > 0 && value2 > 0 && EndResult < 0 {
			z01.PrintRune('0')
			return
		}
		if value1 < 0 && value2 < 0 && EndResult > 0 {
			z01.PrintRune('0')
			return
		}
		if value1 != EndResult-value2 {
			z01.PrintRune('0')
			return
		}

		//case2 - "-" subtraction
	} else if arguments[2] == "-" {
		EndResult = value1 - value2
		if value1 < 0 && value2 > 0 && EndResult > 0 {
			z01.PrintRune('0')
			return
		} else if value1 > 0 && value2 < 0 && EndResult < 0 {
			z01.PrintRune('0')
			return
		}
		if value2 != EndResult-value1 {
			z01.PrintRune('0')
			return
		}

		//case3 - "/" division
	} else if arguments[2] == "/" {
		if value2 == 0 {
			for i := range errorRune1 {
				z01.PrintRune(errorRune1[i])
			}
			return
		}
		EndResult = value1 / value2

		//case4 - "%" modulo
	} else if arguments[2] == "%" {
		if value2 == 0 {
			for i := range errorRune2 {
				z01.PrintRune(errorRune2[i])
			}
			return
		}
		EndResult = value1 % value2

		//case5 - "*" multiplication
	} else if arguments[2] == "*" {
		if value1 == 0 || value2 == 0 {
			z01.PrintRune('0')
			return
		}
		EndResult = value1 * value2
		if value1 > 0 && value2 > 0 && EndResult < 0 {
			z01.PrintRune('0')
			return
		}
		if value1 < 0 && value2 < 0 && EndResult < 0 {
			z01.PrintRune('0')
			return
		}
		if value1 > 0 && value2 < 0 && EndResult > 0 {
			z01.PrintRune('0')
			return
		}
		if value1 < 0 && value2 > 0 && EndResult > 0 {
			z01.PrintRune('0')
			return
		}
		if value1 != EndResult/value2 {
			z01.PrintRune('0')
			return
		}

	} else {
		z01.PrintRune('0')
		return
	}

	answer := []rune(Itoa(EndResult))
	for i := range answer {
		z01.PrintRune(answer[i])
	}

}

func IsValid(input string) bool {
	inputRune := []rune(input)
	for i := range inputRune {
		if inputRune[i] < '0' || inputRune[i] > '9' {
			return false
		}
	}
	return true
}

func Atoi(str string) int64 {
	runes := []rune(str)
	index := 0
	var sign int64
	sign = 1
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

func Itoa(num int64) string {
	var sign string
	var str string
	var str1 string
	if num == 0 {
		str = "0"
	} else if num < 0 {
		num = num * -1
		sign = "-"
	}
	for i := num; i > 0; i = i / 10 {
		str1 = string(rune(i%10) + '0')
		str = str1 + str
	}

	return sign + str
}
