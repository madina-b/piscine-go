package piscine

func NRune(s string, n int) rune {
	Runes := []rune(s)
	for i := range Runes {
		length = i + 1
	}
	if n <= length {
		return Runes[n-1]
	} else {
		return 32
	}
}
