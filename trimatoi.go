package piscine

// import (
// 	"fmt"
// )

func TrimAtoi(s string) int {
	ru := []rune(s)
	var myint int = 0
	var only []int
	var sn int = 1
	for _, letter := range ru {
		if letter == '-' && only == nil {
			sn = -sn
		}
		dg := 0
		for j := '0'; j <= '9'; j++ {
			// if only == nil && dg == 0 {
			// 	break
			// }else
			if letter == j {
				only = append(only, dg)
				break
			}
			dg++
		}
	}

	var len int = 0
	for i, _ := range only {
		len = i + 1
	}

}
