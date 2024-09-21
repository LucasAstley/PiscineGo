package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	}
	for !IsPrime(nb) {
		nb += 1
	}
	return nb
}
