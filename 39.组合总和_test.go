package leetcode

import "testing"

func TestCombinationSum(t *testing.T) {
	cand := []int{2, 3, 6, 7}
	target := 7
	ans := combinationSum(cand, target)
	t.Log(ans)
}
