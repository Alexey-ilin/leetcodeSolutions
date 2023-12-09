package solutions

import "strconv"

// https://leetcode.com/problems/score-after-flipping-matrix/

func MatrixScore(grid [][]int) int {
	var res int
	for i := range grid {
		// var strRepr string
		if grid[i][0] == 0 {
			for j := range grid[i] {
				grid[i][j] = 1 - grid[i][j]
			}
		}

	}
	for j := range grid[0] {
		countZeros := 0
		for i := range grid {
			if grid[i][j] == 0 {
				countZeros++
			}
		}
		if countZeros > len(grid)/2 {
			for i := range grid {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}
	for i := range grid {
		var strRepr string
		for j := range grid[i] {
			strRepr += strconv.Itoa(grid[i][j])
		}
		n, _ := strconv.ParseInt(strRepr, 2, 64)
		res += int(n)
	}
	return res
}
