package leetcode

import "testing"

// 给定⼀个⻓度为 n 的整数数组和⼀个⽬标值 target，
// 寻找能够使条件 nums[i] + nums[j] + nums[k] < target 成⽴的
// 三元组 i, j, k 个数 (0 <= i < j < k < n)

func TestThreeSumSmaller(t *testing.T) {
	nums := []int{-2, 0, 1, 3}
	target := 2
	want := 2

	ans := threeSumSmaller(nums, target)

	if ans != want {
		t.Errorf("Input nums=%v, k=%d. Want %d, but got %d", nums, target, want, ans)
	}
}
