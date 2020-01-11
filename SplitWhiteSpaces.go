package main

import "fmt"

func main() {
	str := "hello \n hjhek"
	fmt.Print(SplitWhiteSpaces(str))
}

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
