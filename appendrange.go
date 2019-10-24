package piscine

func AppendRange(min, max int) []int {
	if min > max {
		return nil
	}
	var ans []int
	var rg int = max - min

	for i := 0; i < rg; i++ {
		ans = append(ans, min+i)
	}
	return ans
}
