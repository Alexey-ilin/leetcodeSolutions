package solutions_test

func generateMatrix(rows, cols int, symbol byte) [][]byte {
	matrix := make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
		for j := range matrix[i] {
			matrix[i][j] = symbol
		}
	}
	return matrix
}

func generateAlternatingMatrix(rows, cols int) [][]byte {
	matrix := make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
		for j := range matrix[i] {
			if (i+j)%2 == 0 {
				matrix[i][j] = '1'
			} else {
				matrix[i][j] = '0'
			}
		}
	}
	return matrix
}

func generateBorderMatrix(rows, cols int) [][]byte {
	matrix := make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
		for j := range matrix[i] {
			if i == 0 || j == 0 || i == rows-1 || j == cols-1 {
				matrix[i][j] = '1'
			} else {
				matrix[i][j] = '0'
			}
		}
	}
	return matrix
}
