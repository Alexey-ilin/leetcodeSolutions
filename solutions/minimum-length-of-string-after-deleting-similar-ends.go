package solutions

// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/

func MinimumLength(s string) int {
	if len(s) == 1 || len(s) == 0 {
		return len(s)
	}
	if s[0] != s[len(s)-1] {
		return len(s)
	}
	c := s[0]
	l := 0
	r := len(s) - 1
	for l <= r {
		if s[l] == c {
			l++
			continue
		}
		if s[r] == c {
			r--
			continue
		}
		break
	}
	return MinimumLength(string(s[l : r+1]))
}
