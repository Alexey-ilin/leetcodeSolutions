package solutions

// https://leetcode.com/problems/min-cost-climbing-stairs/

func MinCostClimbingStairs(cost []int) int {
	dp := make(map[int]int, len(cost))
	dp[len(cost)] = 0
	for i := len(cost) - 1; i >= 0; i-- {
		dp[i] = cost[i] + min(dp[i+1], dp[i+2])
	}
	return min(dp[0], dp[1])
}
