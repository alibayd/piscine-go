package piscine

func Index(s string, toFind string) int {
	gv := []rune(s)
	fd := []rune(toFind)
	var count int
	var len int
	for i := range fd {
		len = i + 1
	}
	for _, i := range gv {
		for _, j := range toFind {
			if i == j {
				count++
			}
		}
	}
	if count < len {
		return -1 // case when the given string is absent
	}
	if len == 1 {
		for index, letter := range gv {
			for _, letter1 := range fd {
				if letter == letter1 { // case when given string is one letter
					return index
				}
			}
		}
	}
	return 1
}
