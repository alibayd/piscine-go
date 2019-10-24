package piscine

import "fmt"

func SplitWhiteSpaces(str string) []string {
	ar := []rune(str)
	var len int = 0
	var i int = 0
	var lenar int = 0
	for j, k := range ar {
		if (k == ' ') || (k == '	') || (k == '\n')  {
			len++
		}
		lenar=j+1
	}
	len++

	ind := make([]int, len)
	for h, k := range ar {
		if (k == ' ') || (k == '	') || (k == '\n') {
			ind[i] = h
			i++
		}
	}
	fmt.Println("content iss")
	fmt.Println(ind)
	fmt.Println(len)
	fmt.Println("nnnn")
	ans := make([]string, len)
	var x int = 0
	for x < len { 
		if x == 0 || len == 1{
			ru := make([]rune, lenar)
			var y int = ind[x]
			fmt.Println(y)
			for i:= 0; i < y; i++ {
				ru[i] = ar[i]
			}
			ans[x] = string (ru)
			fmt.Println("ru is")
			fmt.Println(ans[x])
			x++
		}else if len == 2 {
			fmt.Println("indx")
			// var z int = ind[1]
			fmt.Println("x is")
			fmt.Println(x)
			fmt.Println(ind[x])
			var r1 int = ind[0] + 1
			var r2 int = lenar - 1
			ru1 := make([]rune, r2 - r1 + 1)
			for i := 0; i < (r2-r1+1); i++ {
				ru1[i] = rune(ar[r1 + i])
			}
			ans[x] = string(ru1)
			x++
		}else {
			fmt.Println("here x is")
			fmt.Println(x)
			fmt.Println("ind is")
			fmt.Println(ind)
			var r1 int = ind[x-1] + 1
			var r2 int = ind[x] - 1 
			fmt.Println("R is")
			fmt.Println(r1)
			fmt.Println(r2)
			var lg int = r2 - r1 + 1
			ru1 := make([]rune, lg)
			for i := 0; i < (r2-r1+1); i++ {
				ru1[i] = rune(ar[r1 + i])
			}
			ans[x] = string(ru1)
			x++
		}
	}
	return ans
}

