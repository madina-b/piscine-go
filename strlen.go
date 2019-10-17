package piscine

func StrLen(str string) int {
	counter := 0
	for i, letter := range str {
		counter = i
	}
	return counter + 1

}
