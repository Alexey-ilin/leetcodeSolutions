package solutions

import "strconv"

func LastVisitedIntegers(words []string) []int {
	res := []int{}
	consPrevs := 0
	numsSeen := []int{}
	for i := range words {
		if words[i] == "prev" {
			consPrevs++
			aIdx := len(numsSeen) - consPrevs
			toAppend := -1
			if aIdx >= 0 {
				toAppend = numsSeen[aIdx]
			}
			res = append(res, toAppend)
		} else {
			consPrevs = 0
			num, _ := strconv.Atoi(words[i])
			numsSeen = append(numsSeen, num)
		}
	}
	return res
}
