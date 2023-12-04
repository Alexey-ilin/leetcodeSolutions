package solutions

import "math"

type ordered interface {
	int | float64 | ~string
}

func FindKOr(nums []int, k int) int {
	var ans int
	for i := 0; ; i++ {
		cnt := 0
		bit := int(math.Pow(2, float64(i)))
		if bit > maxOf(nums) {
			break
		}
		for _, n := range nums {
			if (bit & n) == bit {
				cnt++
			}
		}
		if cnt >= k {
			ans += bit
		}
	}
	return ans
}

func maxOf[T ordered](slice []T) T {
	var curMax T
	if len(slice) == 0 {
		return curMax
	}
	curMax = slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > curMax {
			curMax = slice[i]
		}
	}
	return curMax
}
