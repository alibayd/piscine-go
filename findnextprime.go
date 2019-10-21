package piscine

func FindNextPrime(nb int) int {
	for i := 2; i <= nb / 2; i++ {
		if nb%i == 0 {
			nb += 1
			i = 2
		}
	} 
	return nb
}
