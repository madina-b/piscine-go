package piscine

func Sqrt(nb int) int {
	c := 1
	if nb <= 0 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			c = i * i
			if c == nb {
				return i
			}

		}
	}
	return c
}
