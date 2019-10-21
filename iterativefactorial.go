package piscine

func IterativeFactorial(nb int) int {
	var a int
	if nb == 0 || nb == 1 {
		a = 1
	} else if nb > 0 {
		for i := 1; i <= 15; i++ {
			a = a * i
		}

	} else {
		a = 0
	}
	return a
}
