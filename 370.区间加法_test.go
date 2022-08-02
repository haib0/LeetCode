package leetcode

import (
	"testing"
)

func TestGetModifiedArray(t *testing.T) {
	length := 5
	updates := [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}
	want := []int{-2, 0, 3, 5, 3}

	r := getModifiedArray(length, updates)
	t.Logf("Want: %v\nGot: %v", want, r)
}
