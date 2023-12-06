package solutions_test

import (
	"math"
	"testing"

	"github.com/Alexey-ilin/leetcodeSols/solutions"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

func TestLongestPalindromicSubstring(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{name: "Ordinary test 1", input: "babad", output: "aba"},
		{name: "Ordinary test 2", input: "cbbd", output: "bb"},
		{name: "All diff letters test", input: "abcdef", output: "a"},
		{name: "All identical letters test", input: "aaaaaaa", output: "aaaaaaa"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := solutions.LongestPalindromicSubstring(tc.input)
			if res != tc.output {
				t.Errorf("\n LongestPalindromicSubstring(%s) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestMinOpsToReduceToZero(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{name: "Ordinary test 1", input: 39, output: 3},
		{name: "Ordinary test 2", input: 54, output: 3},
		{name: "Big input test", input: 124532, output: 7},
		{name: "Huge input test", input: 3414531877481923, output: 17},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(t.Name(), func(t *testing.T) {
			t.Parallel()
			res := solutions.MinOpsToReduceToZero(tc.input)
			if res != tc.output {
				t.Errorf("\n MinOpsToReduceToZero(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestMaximalSquare(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  [][]byte
		output int
	}{
		{
			name:   "Ordinary test",
			input:  [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
			output: 4,
		},
		{
			name:   "Small test",
			input:  [][]byte{{'0', '1'}, {'1', '0'}},
			output: 1,
		},
		{
			name:   "Matrix with len of 1",
			input:  [][]byte{{'0'}},
			output: 0,
		},
		{
			name:   "All zeros",
			input:  [][]byte{{'0', '0', '0'}, {'0', '0', '0'}, {'0', '0', '0'}},
			output: 0,
		},
		{
			name:   "All ones",
			input:  [][]byte{{'1', '1', '1'}, {'1', '1', '1'}, {'1', '1', '1'}},
			output: 9,
		},
		{
			name:   "Large matrix: 200x200, all ones",
			input:  generateMatrix(200, 200, '1'),
			output: 40000,
		},
		{
			name:   "Large matrix: 300x300, alternating pattern",
			input:  generateAlternatingMatrix(300, 300),
			output: 1,
		},
		{
			name:   "Large matrix: 250x250, border pattern",
			input:  generateBorderMatrix(250, 250),
			output: 1,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.MaximalSquare(tc.input)
			if res != tc.output {
				t.Errorf("\n MaximalSquare(%s) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestMinFallingPathSum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{name: "matrix 3x3", input: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, output: 13},
		{name: "matrix 2x2", input: [][]int{{-19, 57}, {-40, -5}}, output: -59},
		{name: "matrix 2x1", input: [][]int{{-19, 57}}, output: -19},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.MinFallingPathSum(tc.input)
			if res != tc.output {
				t.Errorf("\n MinFallingPathSum(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestTriangle(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{name: "height 4 triangle", input: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, output: 11},
		{name: "height 1 triangle", input: [][]int{{-10}}, output: -10},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.MinimumTotalTriangle(tc.input)
			if res != tc.output {
				t.Errorf("\n Triangle(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{name: "3x3 grid", input: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, output: 2},
		{name: "2x2 grid", input: [][]int{{0, 1}, {0, 0}}, output: 1},
		{name: "2x2 grid (no paths)", input: [][]int{{0, 1}, {1, 0}}, output: 0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.UniquePathsWithObstacles(tc.input)
			if res != tc.output {
				t.Errorf("\n UniquePathsWithObstacles(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestMinumumPathSum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{name: "3x3 grid", input: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, output: 7},
		{name: "2x3 grid", input: [][]int{{1, 2, 3}, {4, 5, 6}}, output: 12},
		{name: "3x3 grid (0-s path)", input: [][]int{{0, 1, 2}, {0, 3, 4}, {0, 0, 0}}, output: 0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.MinPathSum(tc.input)
			if res != tc.output {
				t.Errorf("\n MinimumPathSum(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestMinCostClimbingStairs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{name: "3 step stairs", input: []int{10, 15, 20}, output: 15},
		{name: "10 step stairs", input: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, output: 6},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.MinCostClimbingStairs(tc.input)
			if res != tc.output {
				t.Errorf("\n MinCostClimbingStairs(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestClimbingStairs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{name: "2 step stairs", input: 2, output: 2},
		{name: "3 step stairs", input: 3, output: 3},
		{name: "45 step stairs", input: 45, output: 1836311903},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.ClimbStairs(tc.input)
			if res != tc.output {
				t.Errorf("\n ClimbStairs(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestMinProcessingTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		processorTime []int
		tasks         []int
		output        int
	}{
		{name: "test 1", processorTime: []int{8, 10}, tasks: []int{2, 2, 3, 1, 8, 7, 4, 5}, output: 16},
		{name: "test 2", processorTime: []int{10, 20}, tasks: []int{2, 3, 1, 2, 5, 8, 4, 3}, output: 23},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.MinProcessingTime(tc.processorTime, tc.tasks)
			if res != tc.output {
				t.Errorf("\n MinProcessingTime(%v, %v) gives: \n %v \n expected: \n %v", tc.processorTime, tc.tasks, res, tc.output)
			}
		})
	}
}

func TestNumOfGoodPairs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{name: "ordinary test", input: []int{1, 2, 3, 1, 1, 3}, output: 4},
		{name: "no pairs", input: []int{2, 3, 4, 1, 5}, output: 0},
		{name: "all same", input: []int{1, 1, 1, 1}, output: 6},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.NumIdenticalPairs(tc.input)
			if res != tc.output {
				t.Errorf("\n NumOfGoodPairs(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestDiffWaysToAddParans(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  string
		output []int
	}{
		{name: "ordinary test", input: "2-1-1", output: []int{0, 2}},
		{name: "ordinary test 2", input: "2*3-4*5", output: []int{-34, -14, -10, -10, 10}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.DiffWaysToCompute(tc.input)
			less := func(a, b int) bool { return a < b }
			equalIgnoreOrder := cmp.Diff(res, tc.output, cmpopts.SortSlices(less)) == ""
			if !equalIgnoreOrder {
				t.Errorf("\n DiffWaysToAddParans(%s) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestIncreasingTripletSubsequence(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  []int
		output bool
	}{
		{name: "all incresing", input: []int{1, 2, 3, 4, 5}, output: true},
		{name: "all decreasing", input: []int{5, 4, 3, 2, 1}, output: false},
		{name: "arbitrary", input: []int{2, 1, 5, 0, 4, 6}, output: true},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.IncreasingTriplet(tc.input)
			if res != tc.output {
				t.Errorf("\n IncreasingTripletSubsequence(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}

func TestSelfCrossing(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  []int
		output bool
	}{
		{name: "ordinary self crossing", input: []int{2, 1, 1, 2}, output: true},
		{name: "no crossing", input: []int{1, 2, 3, 4}, output: false},
		{name: "edge crossing", input: []int{1, 1, 1, 2, 1}, output: true},
		// {name: "end = start", input: []int{1, 1, 2, 1, 1}, output: true},
		// {name: "test 9", input: []int{3, 3, 4, 2, 2}, output: false},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := solutions.IsSelfCrossing(tc.input)
			if res != tc.output {
				t.Errorf("\n IsSelfCrossing(%v) gives: \n %v \n expected: \n %v", tc.input, res, tc.output)
			}
		})
	}
}
