package piscine

func NRune(s string) rune {
	length := 0
	Runes := []rune(s)
	for i := range Runes {
		length = i + 1
	}
	return Runes[length-1]
}
