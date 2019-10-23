package piscine

func IterativeFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	var result int = 1
	for i := 1; i <= nb; i++ {
		result = result * i /// fr=fr*i
	}
	return result
}
