package leetcode

import "testing"

//  给你⼀个整数数组 nums 和整数 k,返回最⼤和 sum,
//  满⾜ 存在 i < j 使得 nums[i] + nums[j] = sum 且 sum < k
//  如果没有满⾜此等式的 i, j 存在, 则返回 -1

func TestTwoSumLessThanK(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{34, 23, 1, 24, 75, 33, 54, 8}, 60, 58},
		{[]int{10, 20, 30}, 15, -1},
	}

	for _, tt := range tests {
		t.Run("Test", func(t *testing.T) {
			ans := twoSumLessThanK(tt.nums, tt.k)
			if ans != tt.want {
				t.Errorf("Input nums=%v, k=%d. Want %d, but got %d", tt.nums, tt.k, tt.want, ans)
			}
		})
	}
}
