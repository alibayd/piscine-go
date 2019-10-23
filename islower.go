package piscine

func IsLower(str string) bool {
	ar := []rune(str)
	var len int
	var count int = 0
	for i := range ar {
		len = i + 1
	}
	for i := 0; i < len; i++ {
		if ar[i] >= 'a' && ar[i] <= 'z' {
			count++
		}
	}
	if count == len {
		return true
	}
	return false
}
