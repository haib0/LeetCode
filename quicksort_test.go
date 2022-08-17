package leetcode_test

import (
	"testing"

	leetcode "github.com/haib0/leetcode/go"
)

func TestQuickSort(t *testing.T) {
	var (
		s1 = []int{1, 2, 3, 0, 2}
		s2 = []int{}
		s3 = []float64{1, -1}
	)

	leetcode.QuickSort(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	t.Log(s1)

	leetcode.QuickSort(s2, func(i, j int) bool {
		return s2[i] <= s2[j]
	})
	t.Log(s2)

	leetcode.QuickSort(s3, func(i, j int) bool {
		return s3[i] <= s3[j]
	})
	t.Log(s3)
}
