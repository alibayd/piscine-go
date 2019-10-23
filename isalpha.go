package piscine

func IsAlpha(str string) bool {
	ar := []rune(str)
	var len int
	var count int = 0
	for i := range ar {
		len = i + 1
	}
	for i := 0; i < len; i++ {
		if (ar[i] >= 'A' && ar[i] <= 'Z') || (ar[i] >= 'a' && ar[i] <= 'z') || (ar[i] >= '0' && ar[i] <= '9') {
			count++
		}
	}
	if count == len {
		return true
	}
	return false
}
