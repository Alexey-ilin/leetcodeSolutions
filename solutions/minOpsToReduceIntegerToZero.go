package solutions

import "strconv"

func MinOpsToReduceToZero(n int) int {
	binstr := strconv.FormatInt(int64(n), 2)
	var ones int
	var res int
	for i := len(binstr) - 1; i >= 0; i-- {
		if binstr[i] == '1' {
			ones += 1
		} else {
			if ones == 1 {
				ones = 0
				res += 1
			} else if ones > 1 {
				ones = 1
				res += 1
			}
		}
	}
	if ones == 1 {
		res += 1
	} else if ones > 1 {
		res += 2
	}
	return res
}
