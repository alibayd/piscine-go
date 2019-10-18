package piscine

func StrLen(str string) int {
	var len int
	for index  := range str {
		len = index + 1
	}
	return len
}
