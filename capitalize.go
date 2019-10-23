package piscine

func Capitalize(s string) string {
	runeS := []rune(s)
	start := false
	for i := range runeS {
		curRune := runeS[i]
		if curRune >= 'A' && curRune <= 'Z' {
			if start {
				runeS[i] = curRune + 32
			} else {
				start = true
			}
		} else if curRune >= 'a' && curRune <= 'z' {
			if !start {
				runeS[i] = curRune - 32
				start = true
			}
		} else if curRune >= '0' && curRune <= '9' {
			if !start {
				start = true
			}
		} else {
			start = false
		}
	}

	return string(runeS)
}
