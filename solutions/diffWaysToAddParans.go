package solutions

// https://leetcode.com/problems/different-ways-to-add-parentheses

import (
	"strconv"
)

func recr(expression string) []int {
	if n, err := strconv.Atoi(expression); err == nil {
		return []int{n}
	}
	var res = []int{}
	for i := range expression {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			curAns := 0
			left := expression[:i]
			operand := expression[i]
			right := expression[i+1:]
			leftRes := recr(left)
			rightRes := recr(right)
			for _, lr := range leftRes {
				for _, rr := range rightRes {
					if operand == '+' {
						curAns = lr + rr
					} else if operand == '-' {
						curAns = lr - rr
					} else {
						curAns = lr * rr
					}
					res = append(res, curAns)
				}
			}

		}
	}
	return res
}

func DiffWaysToCompute(expression string) []int {
	return recr(expression)
}
