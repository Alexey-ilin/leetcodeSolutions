package solutions

func MaximalSquare(matrix [][]byte) int {
	// Get the height and width of the matrix
	H := len(matrix)
	W := len(matrix[0])

	// Create a 2D slice to store the dynamic programming results
	dp := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
	}

	// Initialize the maximum size of the square to 0
	maxSize := 0

	// Iterate over each cell in the matrix
	for r := 0; r < H; r++ {
		for c := 0; c < W; c++ {
			// If the cell contains a '1'
			if matrix[r][c] == '1' {
				// Set the corresponding cell in the dp slice to 1
				dp[r][c] = 1

				// If the cell is not on the top row or left column
				if r > 0 && c > 0 {
					// Update the dp value based on the minimum of its top, left and top-left neighbors
					dp[r][c] += min(dp[r][c-1], dp[r-1][c-1], dp[r-1][c])
				}

				// Update the maximum size of the square if necessary
				if dp[r][c] > maxSize {
					maxSize = dp[r][c]
				}
			}
		}
	}

	// Return the area of the maximum square
	return maxSize * maxSize
}

// Helper function to find the minimum value in a slice of integers
func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
