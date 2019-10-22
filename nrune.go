package piscine

func NRune(s string, n int) rune {
	ru := []rune(s)
	for index, i := range ru {
		if index + 1 == n {
			return i
		}
	}
	return 0
}
