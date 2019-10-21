package piscine

func FindNextPrime(nb int) int {
	var primes [2 * nb]int
	for i := range primes {
		primes[i] = 0
	}

	for i := 2; i < nb; i++ {

		if primes[i] != 1 {
			for j := i; j < 2*nb; j = j + i {
				primes[j] = 1
			}
		}
	}

	for i := nb; i < 2*nb; i++ {
		if primes[i] == 0 {
			return i
		}
	}

	return 1
}
