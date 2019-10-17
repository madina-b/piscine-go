package piscine

func StrRev(s string) string {
	runes := []rune(s)
	counter := 0
	for i := range runes {
		counter = i
	}
	StrLen = counter + 1

	for i := StrLen - 1; i >= 0; i-- {
		var1 = runes[i]
		var2 = runes[StrLen-1-i]
		runes[i] = var2
		runes[StrLen-1-i] = var1
	}
	sNew := string(runes)
	return sNew
}
