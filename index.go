package piscine

func Index(s string, toFind string) int {
	runeS := []rune(s)
	runeF := []rune(toFind)
	lenS := len(runeS)
	lenF := len(runeF)

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
