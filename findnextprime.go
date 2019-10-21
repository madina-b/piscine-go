package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		for i := 2; i <= nb; i++ {
			if i < nb && nb%i == 0 {
				for j := 0; j < 20; j++ {
					nb = nb + 1
				}
			} else if i == nb {
				return nb
			}

		}

	} else {
		return 1
	}
	return nb
}
