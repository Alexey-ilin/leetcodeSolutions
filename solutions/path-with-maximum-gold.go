package solutions

// https://leetcode.com/problems/path-with-maximum-gold/

func GetMaximumGold(grid [][]int) int {
	maxGold := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			maxGold = max(maxGold, backtrack(&grid, i, j))
		}
	}
	return maxGold
}

func backtrack(grid *[][]int, i, j int) int {
	if i < 0 || i > len((*grid))-1 || j < 0 || j > len((*grid)[0])-1 || (*grid)[i][j] == 0 {
		return 0
	}
	cell_value := (*grid)[i][j]
	(*grid)[i][j] = 0
	max_gold := max(backtrack(grid, i-1, j), backtrack(grid, i, j+1), backtrack(grid, i+1, j), backtrack(grid, i, j-1))
	(*grid)[i][j] = cell_value
	return cell_value + max_gold

}
