package solutions

// https://leetcode.com/problems/minimum-path-sum/

import "math"

func MinPathSum(grid [][]int) int {
	for i := range grid {
		for j := range grid[0] {
			if i == 0 && j == 0 {
				continue
			}
			left := math.MaxInt64
			top := math.MaxInt64
			if i > 0 {
				top = grid[i-1][j]
			}
			if j > 0 {
				left = grid[i][j-1]
			}
			if top < left {
				grid[i][j] += top
			} else {
				grid[i][j] += left
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
