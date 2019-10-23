package piscine

func IsUpper(str string) bool {
	runeS := []rune(str)
	for i := range runeS {
		curRune := runeS[i]
		if curRune >= 'A' && curRune <= 'Z' {
		} else {
			return false
		}
	}

	return true
}
