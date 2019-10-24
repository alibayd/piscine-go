package piscine

func MakeRange(min, max int) []int {
	if min > max {
		return nil
	}
	var rg int = max - min
	ans := make ([]int, rg)

	for i := 0; i < rg; i++ {
		ans[i] = min + i
	}
	return ans
}
