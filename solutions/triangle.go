package solutions

// https://leetcode.com/problems/triangle/

import "math"

func MinimumTotalTriangle(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			prevLeft := math.MaxInt64
			prevRight := math.MaxInt64
			if j != 0 && j < len(triangle[i])-1 {
				prevLeft = triangle[i-1][j-1]
				prevRight = triangle[i-1][j]
			} else if j == 0 {
				prevRight = triangle[i-1][j]
			} else {
				prevLeft = triangle[i-1][j-1]
			}
			if prevLeft < prevRight {
				triangle[i][j] += prevLeft
			} else {
				triangle[i][j] += prevRight
			}
		}
	}
	curMin := math.MaxInt64
	for _, v := range triangle[len(triangle)-1] {
		if v < curMin {
			curMin = v
		}
	}
	return curMin
}
