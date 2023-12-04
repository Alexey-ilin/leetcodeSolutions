package solutions_test

import (
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
