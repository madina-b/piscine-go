package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		for i := 2; i <= nb; i++ {
			if i < nb && nb%i == 0 {
				nb = nb + 1
				FindNextPrime(nb)
			} else if i == nb {
				return nb
			}

		}

	} else {
		return 1
	}
	return nb
}
