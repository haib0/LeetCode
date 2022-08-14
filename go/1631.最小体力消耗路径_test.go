package leetcode

import (
	"testing"
)

func TestMinimumEffortPath(t *testing.T) {
	heights := [][]int{{1, 10, 6, 7, 9, 10, 4, 9}}
	x := minimumEffortPath(heights)
	t.Log(x)
}
