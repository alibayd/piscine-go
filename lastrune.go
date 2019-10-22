package piscine

func LastRune(s string) rune {
	ru := []rune(s)
	var len int = 0
	for i := range ru {
		len = i + 1
	}
	for x, y := range ru {
		if x+1 == len {
			return y
		}
	}
	return 0
}
