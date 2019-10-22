package piscine

func ToUpper(s string) string {
	ss := []rune(s)
	for i := 0; i < len(s); i++ {
		if ss[i] >= 'a' && ss[i] <= 'z' {
			ss[i] = ss[i] + 32
			//i++
		}
	}
	ret := string(ss)
	return ret
}

//	only = append(only, dg)
