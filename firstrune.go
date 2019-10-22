package piscine

func FirstRune(s string) rune {
	ru := []rune(s)
	for _, i := range ru {
		return i
	}
	return 0
}
