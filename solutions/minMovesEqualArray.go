package solutions

import "math"

func MinMovesToEqualArray(nums []int) int {
	sum := 0
	min := math.MaxInt
	for _, v := range nums {
		sum += v
		if v < min {
			min = v
		}
	}
	return sum - min*len(nums)
}
