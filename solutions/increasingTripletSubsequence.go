package solutions

func IncreasingTriplet(nums []int) bool {
	lowestToLeft := make([]int, len(nums))
	biggestToRight := make([]int, len(nums))
	lowestToLeft[0] = nums[0]
	biggestToRight[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		if lowestToLeft[i-1] < nums[i] {
			lowestToLeft[i] = lowestToLeft[i-1]
		} else {
			lowestToLeft[i] = nums[i]
		}
	}
	for j := len(nums) - 2; j >= 0; j-- {
		if biggestToRight[j+1] > nums[j] {
			biggestToRight[j] = biggestToRight[j+1]
		} else {
			biggestToRight[j] = nums[j]
		}
	}
	for i := 1; i < len(nums)-1; i++ {
		if lowestToLeft[i-1] < nums[i] && nums[i] < biggestToRight[i+1] {
			return true
		}
	}
	return false
}
