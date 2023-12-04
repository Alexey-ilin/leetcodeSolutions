package solutions

// https://leetcode.com/problems/unique-paths-ii/

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i == 0 && j == 0 {
				obstacleGrid[0][0] = 1
				continue
			} else if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			left := 0
			top := 0
			if i > 0 {
				top = obstacleGrid[i-1][j]
			}
			if j > 0 {
				left = obstacleGrid[i][j-1]
			}
			obstacleGrid[i][j] = top + left
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
