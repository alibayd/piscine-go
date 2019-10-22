package piscine

func Compare(a, b string) int {
	// aa := []rune(a)
	// bb := []rune(b)
	if a == b {
		return 0
	} else if a > b {
		return -1
	} else if a < b {
		return 1
	}
	return 0
}
