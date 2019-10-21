package piscine

func RecursivePower(nb int, power int) int {
	a := 1
	if nb == 1 {
		return 1
	} else if power < 0 || power > 32 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		return nb * RecursivePower(nb-1)
	}
	return a
}
