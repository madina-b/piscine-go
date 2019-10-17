package piscine

func StrRev(s string) string {
	runes := []rune(s)
	StrLen := len(runes)
	var newRunes []rune
	for i := StrLen - 1; i >= 0; i-- {
		newRunes = append(newRunes, runes[i])
	}
	sNew := string(newRunes)
	return sNew
}
