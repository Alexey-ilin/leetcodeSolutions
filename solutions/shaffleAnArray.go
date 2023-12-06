package solutions

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (sol *Solution) Reset() []int {
	return sol.nums
}

func (sol *Solution) Shuffle() []int {
	copyNums := make([]int, len(sol.nums))
	copy(copyNums, sol.nums)
	for i := len(copyNums) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		copyNums[i], copyNums[j] = copyNums[j], copyNums[i]
	}
	return copyNums
}
