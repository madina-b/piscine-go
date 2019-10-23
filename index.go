package piscine

func Index(s string, toFind string) int {
	runeS := []rune(s)
	runeF := []rune(toFind)
	lenS := 0
	lenF := 0
	for i := range runeS {
		lenS = i + 1
	}
	for j := range runeF {
		lenF = j + 1
	}
	if lenF == 0 {
		return 0
	}

	for i := 0; i < lenS-lenF+1; i++ {
		ind := i
		for j := 0; j < lenF; j++ {
			if runeS[ind] != runeF[j] {
				break
			} else {
				ind++
			}
			if j == lenF-1 {
				return i
			}
		}
	}
	return -1
}
