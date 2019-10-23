package piscine

func NRune(s string, n int) rune {
	length := 0
	Runes := []rune(s)
	for i := range Runes {
		length = i + 1
	}
	if n <= length && n > 0 {
		return Runes[n-1]
	} else {
		return 0
	}
}
