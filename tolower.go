package piscine

func ToLower(str string) string {
	word := []rune(str)
	length := 0
	for i := range word {
		length = i + 1
	}
	for i := 0; i < length; i++ {
		if word[i] <= 90 && word[i] >= 65 {
			word[i] = word[i] + 32
		}
	}
	return string(word)
}
