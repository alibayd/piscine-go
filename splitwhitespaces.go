package piscine

import "fmt"

func SplitWhiteSpaces(str string) []string {
	ar := []rune(str)
	var wrds int
	var len_str int
	for i, j := range ar {
		if j == ' ' || j == ' ' || j == '\n' {
			wrds++
		}
		len_str = i + 1
	}
	wrds++
	var ind_of_prev_sp int
	ans := make([]string, wrds)
	var numb_spaces int = 0
	for i, j := range ar {
		if j == ' ' || j == ' ' || j == '\n' {
			if numb_spaces == 0 { // [word_]
				numb_spaces++
				ru := make([]rune, i)
				for k := 0; k < i; k++ {
					ru[k] = ar[k]
				}
				ans[0] = string(ru)
				ind_of_prev_sp = i
			} else if numb_spaces < (wrds - 1) { //[w1_word2_]
				var l_of_rune int = i - ind_of_prev_sp - 1
				ru := make([]rune, l_of_rune)
				for k := 0; k < l_of_rune; k++ {
					ru[k] = ar[ind_of_prev_sp+1+k]
				}
				ans[numb_spaces] = string(ru)
				numb_spaces++
				ind_of_prev_sp = i
			}
		} else if numb_spaces+1 == wrds { // [.._wordl]
			var l_of_rune int = len_str - ind_of_prev_sp - 1 //5
			ru := make([]rune, l_of_rune)
			for k := 0; k < l_of_rune; k++ {
				ru[k] = ar[ind_of_prev_sp+1+k]
			}
			ans[numb_spaces] = string(ru)
			numb_spaces++
		}

	}
	return ans

}
