package piscine

func Capitalize(s string) string {
	ar := []rune(s)
	var len int
	for i := range ar {
		len = i + 1
	}
	for i := 0; i < len; i++ {
		if i == 0 || !((ar[i-1] >= 'A' && ar[i-1] <= 'Z') || (ar[i-1] >= 'a' && ar[i-1] <= 'z') || (ar[i-1] >= '0' && ar[i-1] <= '9')) {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] -= 32
			}
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] += 32
		}
	}
	return string(ar)
}
