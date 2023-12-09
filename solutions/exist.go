package solutions

// https://leetcode.com/problems/word-search

func Exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[0] {
			res := helper(i, j, word, board, visited, 0)
			if res {
				return true
			}
		}
	}
	return false
}

func helper(i, j int, word string, board [][]byte, visited [][]bool, counter int) bool {
	var res bool
	if board[i][j] == word[counter] {
		if counter == len(word)-1 {
			return true
		}
		visited[i][j] = true
		if i < len(board)-1 {
			if !visited[i+1][j] {
				res = helper(i+1, j, word, board, visited, counter+1)
				if res {
					return res
				}
			}
		}
		if j < len(board[0])-1 {
			if !visited[i][j+1] {
				res = helper(i, j+1, word, board, visited, counter+1)
				if res {
					return res
				}
			}
		}
		if i > 0 {
			if !visited[i-1][j] {
				res = helper(i-1, j, word, board, visited, counter+1)
				if res {
					return res
				}
			}
		}
		if j > 0 {
			if !visited[i][j-1] {
				res = helper(i, j-1, word, board, visited, counter+1)
				if res {
					return res
				}
			}
		}
		visited[i][j] = false
	}
	return false
}
