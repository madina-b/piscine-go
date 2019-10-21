package piscine

func IterativeFactorial(c int) int {
	var a int
	if c == 0 || c == 1 {
		a = 1
	} else if nb > 0 {
		for i := 1; i <= c; i++ {
			a = a * i
		}

	} else {
		a = 0
	}
	return a
}
