package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		for i := 2; i <= nb; i++ {
			if i < nb && nb%i == 0 {
				nb = nb + 1
				FindNextPrime(nb)
			} else if i == nb {
				return nb
			}

		}

	} else {
		return 2
	}
	return nb
}
