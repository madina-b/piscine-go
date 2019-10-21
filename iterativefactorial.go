package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 || nb == 1 {
		arg = 1
	} else if nb > 0 {
		for i := 1; i <= nb; i++ {
			arg = arg * i
		}

	} else {
		arg = 0
	}
	return arg
}
