package piscine

func IterativePower(nb int, power int) int {
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
		for i := 1; i <= power; i++ {
			a = a * nb
		}
	}
	return a
}
