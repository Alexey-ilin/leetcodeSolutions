package solutions

func EqualSubstringWithinBudget(s string, t string, maxCost int) int {
	var l int
	maxLen := 0
	curCost := 0
	for r := range s {
		curCost += Abs(int(rune(s[r]) - rune(t[r])))
		for l < len(s) && curCost > maxCost {
			curCost -= Abs(int(rune(s[l]) - rune(t[l])))
			l++
		}
		if d := r - l + 1; d > maxLen {
			maxLen = d
		}
	}
	return maxLen
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
