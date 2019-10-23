package main

func TrimAtoi(str string) int {
	LetterRunes := []rune(str)
	length := 0
	value := 0
	sign := 1

	for i := range LetterRunes {
		length = i + 1
	}

	for i := 0; i < length; i++ {
		if LetterRunes[i] == 45 {
			sign = -1
			break
		} else if LetterRunes[i] <= 57 && LetterRunes[i] >= 48 {
			sign = 1
			break
		} else if LetterRunes[i] == 43 {
			sign = 1
			break
		}
	}

	for j := 0; j < length; j++ {
		if LetterRunes[j] <= 57 && LetterRunes[j] >= 48 {
			value = value*10 + int(LetterRunes[j]-'0')
		}
	}
	return sign * value
}
