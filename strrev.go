package piscine

func StrRev(s string) string {
	ru := []rune(s)
	revru := []rune(s)

	ar := []rune(s)
	var len int
	for k := range ar {
		len = k 
	}
	var d int = len

	for index := 0; index <= d; index++ {
		revru[index] = ru[len]
		len--
	}
	ss := string(revru)
	return ss
}
