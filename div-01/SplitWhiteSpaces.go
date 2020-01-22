package main

import "fmt"

// func main() {
// 	str := "The earliest foundations of what would become   computer science predate the invention of the modern digital computer"
// 	fmt.Println(SplitWhiteSpaces(str))
// 	str = "Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"
// 	fmt.Println(SplitWhiteSpaces(str))
// 	str = "aiding in  computations such as multiplication and division ."
// 	fmt.Println(SplitWhiteSpaces(str))
// 	str = "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
// 	fmt.Println(SplitWhiteSpaces(str))
// 	str = "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
// 	fmt.Println(SplitWhiteSpaces(str))
// 	str = "In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
// 	fmt.Println(SplitWhiteSpaces(str))

}

// [The earliest foundations of what would become computer science predate the invention of the modern digital computer]
// [Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,]
// [aiding in computations such as multiplication and division .]
// [Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment.]
// [Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]]
// [In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,]


func SplitWhiteSpaces(str string) []string {
	srune := []rune(str)
	start := false
	wordcount := 0
	lenthRune := 0

	for i := range srune {
		lenthRune++
		if srune[i] == ' ' || srune[i] == '\n' || srune[i] == '\t' {
			if start == true {
				start = false
			}
		} else {
			if start == false {
				wordcount++
				start = true
			}
		}
	}

	// Initializing variables
	start = false                       //
	output := make([]string, wordcount) //
	j := 0                              //
	startindx := 0                      //
	endindx := 0                        //

	// i = 0
	for i := range srune {

		// Case 1 - separator
		if srune[i] == ' ' || srune[i] == '\n' || srune[i] == '\t' {

			if start {
				endindx = i
				start = false
				output[j] = string(srune[startindx:endindx])
				j++
			}

			// Case 2 - alphanumeric value
		} else {
			if !start {

				start = true
				startindx = i
			}
		}

	}
	if start {
		endindx = lenthRune
		output[wordcount-1] = string(srune[startindx:endindx])
	}
	return output
}