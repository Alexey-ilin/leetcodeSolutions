package solutions

// https://leetcode.com/problems/minimum-falling-path-sum/

import "math"

func MinFallingPathSum(matrix [][]int) int {
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			prevs := [3]int{math.MaxInt64, matrix[i-1][j], math.MaxInt64}
			if j == 0 {
				prevs[2] = matrix[i-1][j+1]
			} else if j == len(matrix)-1 {
				prevs[0] = matrix[i-1][j-1]
			} else {
				prevs[0] = matrix[i-1][j-1]
				prevs[2] = matrix[i-1][j+1]
			}
			prevMin := math.MaxInt64
			for _, v := range prevs {
				if v < prevMin {
					prevMin = v
				}
			}
			matrix[i][j] += prevMin
		}
	}

	res := math.MaxInt64
	for _, v := range matrix[len(matrix)-1] {
		if v < res {
			res = v
		}
	}
	return res
}
