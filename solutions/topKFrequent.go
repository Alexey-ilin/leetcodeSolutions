package solutions

import "container/heap"

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0] > pq[j][0]
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([2]int))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func TopKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	pq := make(PriorityQueue, 0, len(nums))
	heap.Init(&pq)
	for _, n := range nums {
		count[n]++
	}
	for k, v := range count {
		heap.Push(&pq, [2]int{v, k})
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(&pq).([2]int)
		res[i] = item[1]
	}
	return res
}
