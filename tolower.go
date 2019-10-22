package piscine

func ToLower(s string) string {
	ss := []rune(s)
	for i := 0; i < len(s); i++ {
		if ss[i] >= 'A' && ss[i] <= 'Z' {
			ss[i] = ss[i] + 32
			//i++
		}
	}
	ret := string(ss)
	return ret
}

//	only = append(only, dg)
