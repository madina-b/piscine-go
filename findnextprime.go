package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		if nb != 2 && nb%2 == 0 {
			nb = nb + 1
			FindNextPrime(nb)
		} else if nb != 3 && nb%3 == 0 {
			nb = nb + 1
			FindNextPrime(nb)
		} else if nb != 5 && nb%5 == 0 {
			nb = nb + 1
			FindNextPrime(nb)
		} else if nb != 7 && nb%7 == 0 {
			nb = nb + 1
			FindNextPrime(nb)
		}
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
