package solutions

//https://leetcode.com/problems/minimum-sum-of-mountain-triplets-i/
//https://leetcode.com/problems/minimum-sum-of-mountain-triplets-ii/

func MinimumSumMountainTriplets(nums []int) int {
	lowestToLeft := make([]int, len(nums))
	lowestToRight := make([]int, len(nums))
	lowestToLeft[0] = nums[0]
	lowestToRight[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		if nums[i] < lowestToLeft[i-1] {
			lowestToLeft[i] = nums[i]
		} else {
			lowestToLeft[i] = lowestToLeft[i-1]
		}
	}
	for j := len(nums) - 2; j >= 0; j-- {
		if nums[j] < lowestToRight[j+1] {
			lowestToRight[j] = nums[j]
		} else {
			lowestToRight[j] = lowestToRight[j+1]
		}
	}
	curMin := int(^uint(0) >> 1)
	for i := 1; i < len(nums)-1; i++ {
		tc := lowestToLeft[i-1] + nums[i] + lowestToRight[i+1]
		if lowestToLeft[i-1] >= nums[i] || nums[i] <= lowestToRight[i+1] {
			continue
		}
		if tc < curMin {
			curMin = tc
		}
	}
	if curMin == int(^uint(0)>>1) {
		return -1
	}
	return curMin
}
