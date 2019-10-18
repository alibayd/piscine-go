package piscine

func StrLen(str string) int {
	ar := []rune(str)
	var len int
	for i  := range ar {
		len = i + 1
	}
	return len
}
