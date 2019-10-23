package piscine

func IsLower(str string) bool {
	runeS := []rune(str)
	for i := range runeS {
		curRune := runeS[i]
		if curRune >= 'a' && curRune <= 'z' {
		} else {
			return false
		}
	}

	return true
}
