package solutions

func NumSubmatrixSumTarget(matrix [][]int, target int) int {
	count := 0
	rowsPrefix := make([][]int, len(matrix))
	for i := range rowsPrefix {
		rowsPrefix[i] = make([]int, len(matrix[0])+1)
	}
	for i := 0; i < len(rowsPrefix); i++ {
		for j := 1; j < len(rowsPrefix[0]); j++ {
			rowsPrefix[i][j] = rowsPrefix[i][j-1] + matrix[i][j-1]
		}
	}
	for col1 := 0; col1 < len(matrix[0]); col1++ {
		for col2 := col1; col2 < len(matrix[0]); col2++ {
			prefFreqs := make(map[int]int)
			subMatrix := make([]int, len(matrix))
			for i := range subMatrix {
				subMatrix[i] = rowsPrefix[i][col2+1] - rowsPrefix[i][col1]
			}
			subMatrixPrefix := make([]int, len(subMatrix)+1)
			for i := 1; i < len(subMatrixPrefix); i++ {
				subMatrixPrefix[i] = subMatrixPrefix[i-1] + subMatrix[i-1]
			}
			for i := 0; i < len(subMatrixPrefix); i++ {
				count += prefFreqs[subMatrixPrefix[i]-target]
				prefFreqs[subMatrixPrefix[i]]++
			}
		}
	}
	return count
}
