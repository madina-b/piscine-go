package piscine

func AlphaCount(str string) int {
	runes := []rune(str)
	length := len(runes)
	count := 0
	for i := 0; i < length; i++ {
		if runes[i] >= 65 && runes[i] <= 90 {
			count = count + 1
		} else if runes[i] >= 97 && runes[i] <= 122 {
			count = count + 1
		}
	}
	return count
}
