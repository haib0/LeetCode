package leetcode

import "testing"

func TestCountRangeSum(t *testing.T) {
	nums := []int{0}
	lower, upper := 0, 0

	c := countRangeSum(nums, lower, upper)

	if c != 1 {
		t.Error(c)
	}

	nums1 := []int{-2, 5, -1}
	lower1, upper1 := -2, 2

	c1 := countRangeSum(nums1, lower1, upper1)

	if c1 != 3 {
		t.Error(c1)
	}
}
