package piscine

func AlphaCount(str string) int {
	var len int = 0
	for _, letter := range str {
		if (letter >= 'a' && letter < 'z') || (letter <= 'Z' && letter > 'A') {
			len += 1
		}
	}
	return len
}
