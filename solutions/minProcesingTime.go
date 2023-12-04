package solutions

import "sort"

func MinProcessingTime(processorTime []int, tasks []int) int {
	var res int
	sort.Sort(sort.IntSlice(processorTime))
	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))
	for i := 0; i < len(tasks); i += 4 {
		if processorTime[i/4]+tasks[i] > res {
			res = processorTime[i/4] + tasks[i]
		}
	}
	return res
}
