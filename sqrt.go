package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	} else if nb == 1 {
		return 1
	}
	var res int = 1
	var i int = 1
	for res <= nb {
		i++
		res = i * i
		if res == nb {
			return i
		}
	}
	return 0
}
