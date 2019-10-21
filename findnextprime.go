package piscine

func IsPrime(nb int) bool {
	if nb > 0 {
		for i := 2; i <= nb; i++ {
			if i < nb && nb%i == 0 {
				return false
			} else if nb%i == 0 {
				return true
			}
		}
	}

	return false
}

func FindNextPrime(nb int) int {
	for i := nb; i <= 2*nb; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 1
}
