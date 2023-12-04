package solutions

var memoCl = map[int]int{
	0: 1,
	1: 1,
}

func ClimbStairs(n int) int {
	val, ok := memoCl[n]
	if ok {
		return val
	}
	f := ClimbStairs(n-1) + ClimbStairs(n-2)
	memoCl[n] = f
	return f
}
