package piscine

func IsAlpha(str string) bool {
	runeS := []rune(str)
	for i := range runeS {
		curRune := runeS[i]
		if curRune >= 'A' && curRune <= 'Z' {
		} else if curRune >= 'a' && curRune <= 'z' {
		} else if curRune >= '0' && curRune <= '9' {
		} else {
			return false
		}
	}

	return true
}
