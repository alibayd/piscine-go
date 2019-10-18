package piscine

func StrRev(s string) string {
	ru := []rune(s)
	revru := []rune(s)
	var i int = 11
	for index := 0; index < 12; index++ {
		revru[index] = ru[i]
		i--
	}
	ss := string(revru)
	return ss
}
