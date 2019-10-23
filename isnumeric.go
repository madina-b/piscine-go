package piscine

func IsNumeric(str string) bool {
	runeS := []rune(str)
	for i := range runeS {
		curRune := runeS[i]
		if curRune >= '0' && curRune <= '9' {
		} else {
			return false
		}
	}

	return true
}
