package piscine

func StrLen(str string) int {
	word := []rune(str)
	counter := 0
	for i := range word {
		counter = i
	}
	return counter + 1

}
