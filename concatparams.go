package piscine

func ConcatParams(args []string) string {
	var len int
	for i := range args {
		len = i
	}

	var res string

	for i := range args {
		res += args[i]
		if i != len {
			res += "\n"
		}
	}
	return res
}
