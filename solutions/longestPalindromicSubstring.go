package solutions

//https://leetcode.com/problems/longest-palindromic-substring/

func LongestPalindromicSubstring(s string) string {
	maxPalindrome := string(s[0])
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for r := len(s) - 1; r >= 0; r-- {
		for c := r + 1; c < len(s); c++ {
			if s[r] == s[c] {
				if c-r == 1 || dp[r+1][c-1] {
					dp[r][c] = true
					newFound := s[r : c+1]
					if len(newFound) > len(maxPalindrome) {
						maxPalindrome = newFound
					}
				}
			}
		}
	}
	return maxPalindrome
}
