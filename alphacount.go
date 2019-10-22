package piscine

func AlphaCount(str string) int {
	length := 0
	runes := []rune(str)
	for j := range runes {
		length = j
	}
	count := 0
	for i := 0; i < length+1; i++ {
		if runes[i] >= 65 && runes[i] <= 90 {
			count = count + 1
		} else if runes[i] >= 97 && runes[i] <= 122 {
			count = count + 1
		}
	}
	return count
}
