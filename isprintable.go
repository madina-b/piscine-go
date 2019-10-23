package piscine

func IsPrintable(str string) bool {
	runeS := []rune(str)
	for i := range runeS {
		curRune := runeS[i]
		if curRune < ' ' {
			return false
		}
	}

	return true
}
