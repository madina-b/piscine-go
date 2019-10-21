package piscine

func IterativeFactorial(nb int) int {
	a := 1
	if nb == 0 || nb == 1 {
		a = 1
	} else if nb > 0 && nb < 20 {
		for i := 1; i <= nb; i++ {
			a = a * i
		}

	} else {
		a = 0
	}
	return a
}
