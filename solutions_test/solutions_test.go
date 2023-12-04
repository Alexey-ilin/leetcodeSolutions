package solutions_test

import (
	"math"
	"testing"

	"github.com/Alexey-ilin/leetcodeSols/solutions"
)

// unit tests for leetcode solutions

func TestPlayWithMaxGold(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		grid [][]int
		res  int
	}{
		{name: "3x3 test", grid: [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}, res: 24},
		{name: "3x5 test", grid: [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}, res: 28},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.GetMaximumGold(tc.grid)
			if res != tc.res {
				t.Errorf("\n GetMaximumGold(%v) gives: \n %v \n expected: \n %v", tc.grid, res, tc.res)
			}
		})
	}
}

func TestKorOfArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{name: "Test 1", nums: []int{7, 12, 9, 8, 9, 15}, k: 4, res: 9},
		{name: "Test 2", nums: []int{2, 12, 1, 11, 4, 5}, k: 6, res: 0},
		{name: "Test 0 len", nums: []int{}, res: 0},
		{name: "Test 4", nums: []int{10, 8, 5, 9, 11, 6, 8}, k: 1, res: 15},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := solutions.FindKOr(tc.nums, tc.k)
			if res != tc.res {
				t.Errorf("\n KorOfArray(%v, %v) gives: \n %v \n expected: \n %v", tc.nums, tc.k, res, tc.res)
			}
		})
	}
}

func TestMinLenAfterDelSimEnds(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    string
		res  int
	}{
		{name: "Test 1", s: "ca", res: 2},
		{name: "Test 2", s: "cabaabac", res: 0},
		{name: "Test 3", s: "aabccabba", res: 3},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := solutions.MinimumLength(tc.s)
			if res != tc.res {
				t.Errorf("\n MinLenAfterDelSimEnds(%s) gives: \n %v \n expected: \n %v", tc.s, res, tc.res)
			}
		})
	}
}

func TestMinSumMountTriplets(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		nums []int
		res  int
	}{
		{name: "Test 1", nums: []int{8, 6, 1, 5, 3}, res: 9},
		{name: "Test 2", nums: []int{5, 4, 8, 7, 10, 2}, res: 13},
		{name: "Test 3", nums: []int{6, 5, 4, 3, 4, 5}, res: -1},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := solutions.MinimumSumMountainTriplets(tc.nums)
			if res != tc.res {
				t.Errorf("\n MinSumMountTriplets(%v) gives: \n %v \n expected: \n %v", tc.nums, res, tc.res)
			}
		})
	}
}

func TestMinMovesToEqualArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		nums []int
		res  int
	}{
		{name: "Test 1", nums: []int{1, 2, 3}, res: 3},
		{name: "Test 2", nums: []int{1, 1, 1}, res: 0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := solutions.MinMovesToEqualArray(tc.nums)
			if res != tc.res {
				t.Errorf("\n MinMovesToEqualArray(%v) gives: \n %v \n expected: \n %v", tc.nums, res, tc.res)
			}
		})
	}
}

func TestRandPointInACircle(t *testing.T) {
	rad := 1
	x_center := 0
	y_center := 0
	t.Parallel()
	circle := solutions.NewCircle(float64(rad), float64(x_center), float64(y_center))
	test := struct {
		name string
		n    int
	}{name: "Trying to get n random points in a circle", n: 10}
	for i := 1; i <= test.n; i++ {
		i := i
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := circle.RandPoint()
			if (math.Pow(res[0], 2) + math.Pow(res[1], 2)) > float64(rad) {
				t.Errorf("\n Getting %v random point in a circle is wrong", i)
			}
		})
	}
}
func TestLastVisitedIntegers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  []string
		output []int
	}{
		{name: "Test 1", input: []string{"1", "2", "prev", "prev", "prev"}, output: []int{2, 1, -1}},
		{name: "Test 2", input: []string{"1", "prev", "2", "prev", "prev"}, output: []int{1, 2, 1}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := solutions.LastVisitedIntegers(tc.input)
			for i := range res {
				if res[i] != tc.output[i] {
					t.Errorf("\n LastVisitedIntegers(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
				}
			}
		})
	}
}

func TestCountSymmetricIntegers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		low  int
		high int
		res  int
	}{
		{name: "Test 1", low: 1, high: 100, res: 9},
		{name: "Test 2", low: 1200, high: 1230, res: 4},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := solutions.CountSymmetricIntegers(tc.low, tc.high)
			if res != tc.res {
				t.Errorf("\n CountSymmetricIntegers(%v, %v) gives: \n %v \n expected: \n %v", tc.low, tc.high, res, tc.res)
			}
		})
	}
}

func TestEqualSubstringWithinBudget(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		s       string
		t       string
		maxCost int
		res     int
	}{
		{name: "Ordinary test", s: "abcd", t: "bcdf", maxCost: 3, res: 3},
		{name: "No changes with suitable cost", s: "abcd", t: "cdef", maxCost: 3, res: 1},
		{name: "Cannot make any changes", s: "abcd", t: "acde", maxCost: 0, res: 1},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := solutions.EqualSubstringWithinBudget(tc.s, tc.t, tc.maxCost)
			if res != tc.res {
				t.Errorf("\n EqualSubstringWithinBudget(%s, %s) gives: \n %v \n expected: \n %v", tc.s, tc.t, res, tc.res)
			}
		})
	}
}
