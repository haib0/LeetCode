package leetcode

import (
	"testing"
)

func TestNumDistinctIslands(t *testing.T) {
	grid := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}

	ans := numDistinctIslands(grid)
	if ans != 3 {
		t.Errorf("Wrong, expect 3 but got %d", ans)
	}
}
