package solutions

// https://leetcode.com/problems/find-pivot-index

func PivotIndex(nums []int) int {
	sumLeft := make([]int, len(nums)+1)
	sumRight := make([]int, len(nums)+1)
	for i := 1; i < len(sumLeft); i++ {
		sumLeft[i] = sumLeft[i-1] + nums[i-1]
	}
	for j := len(nums) - 1; j >= 0; j-- {
		sumRight[j] = sumRight[j+1] + nums[j]
	}
	for i := 0; i < len(sumLeft)-1; i++ {
		if sumLeft[i] == sumRight[i+1] {
			return i
		}
	}
	return -1

}
