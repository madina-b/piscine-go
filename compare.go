package piscine

func Compare(a, b string) int {
	runeA := []rune(a)
	runeB := []rune(b)
	lenA := 0
	lenB := 0
	for i := range runeA {
		lenA = i + 1
	}
	for j := range runeB {
		lenB = j + 1
	}
	if lenA > lenB {
		return -Compare(b, a)
	}

	for i := 0; i < lenA; i++ {
		if runeA[i] < runeB[i] {
			return -1
		} else if runeA[i] > runeB[i] {
			return 1
		}
	}

	if lenA == lenB {
		return 0
	} else {
		return -1
	}
}
