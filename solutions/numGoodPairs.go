package solutions

// https://leetcode.com/problems/number-of-good-pairs/

func NumIdenticalPairs(nums []int) int {
	var nPairs int
	nCount := make(map[int]int)
	for _, n := range nums {
		nCount[n] += 1
	}
	for _, v := range nCount {
		nPairs += v * (v - 1) / 2
	}
	return nPairs
}
