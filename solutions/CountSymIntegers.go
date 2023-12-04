package solutions

import (
	"strconv"
)

func CountSymmetricIntegers(low int, high int) int {
	var counter int
	for n := low; n <= high; n++ {
		nstr := strconv.Itoa(n)
		if len(nstr)%2 == 1 {
			continue
		}
		mid := len(nstr) / 2
		var sum int
		for i := 0; i < mid; i++ {
			cn, _ := strconv.Atoi(string(nstr[i]))
			sum += cn
		}
		for i := mid; i < len(nstr); i++ {
			cn, _ := strconv.Atoi(string(nstr[i]))
			sum -= cn
		}
		if sum == 0 {
			counter++
		}
	}
	return counter

}
