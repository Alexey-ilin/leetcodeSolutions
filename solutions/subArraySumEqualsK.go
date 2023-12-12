package solutions

func SubarraySum(nums []int, k int) int {
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	sumFreqs := make(map[int]int)
	res := 0
	for i := 0; i < len(nums)+1; i++ {
		res += sumFreqs[prefixSum[i]-k]
		sumFreqs[prefixSum[i]]++
	}
	return res
}
