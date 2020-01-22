package main

import "fmt"

func main() {
	str := "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished"
	fmt.Println(Split(str, "|=choumi=|"))
	str = "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator"
	fmt.Println(Split(str, "!==!"))
	str = "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator"
	fmt.Println(Split(str, "AFJ"))
	str = "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer"
	fmt.Println(Split(str, "<<==123==>>"))
}

func Split(str, charset string) []string {
	strAsRune := []rune(str)
	charsetAsRune := []rune(charset)
	lenStr := 0
	lenCharset := 0
	size := 0

	for i := range strAsRune {
		lenStr = i + 1
	}
	for i := range charsetAsRune {
		lenCharset = i + 1
	}
	if lenStr > 0 && lenCharset > 0 {
		size++
		for i := 0; i < lenStr; i++ {
			if str[i] == charset[0] {
				if i+lenCharset+1 < lenStr {
					k := 0
					for j := i; j < i+lenCharset; j++ {
						if strAsRune[j] == charsetAsRune[k] {
							k++
						}
					}
					if k == lenCharset {
						size++
					}
				}
			}
		}
	}
	answer := make([]string, size)
	word := ""
	index := 0
	if lenStr > 0 && lenCharset > 0 {
		for i := 0; i < lenStr; i++ {
			if str[i] == charset[0] {
				if i+lenCharset+1 < lenStr {
					k := 0
					for j := i; j < i+lenCharset; j++ {
						if strAsRune[j] == charsetAsRune[k] {
							k++
						}
					}
					if k == lenCharset {
						answer[index] = word
						i += lenCharset - 1
						index++
						word = ""
					} else {
						word += string(strAsRune[i])
					}
				} else {
					word += string(strAsRune[i])
				}
			} else {
				word += string(strAsRune[i])
			}
		}
		answer[size-1] = word
	}
	return answer
}
