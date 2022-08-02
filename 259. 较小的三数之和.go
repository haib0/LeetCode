/*
 * @lc app=leetcode.cn id=259 lang=golang
 *
 * [259] 较小的三数之和
 */
package leetcode

import "sort"

// @lc code=start

// 给定⼀个⻓度为 n 的整数数组和⼀个⽬标值 target，
// 寻找能够使条件 nums[i] + nums[j] + nums[k] < target 成⽴的
// 三元组 i, j, k 个数 (0 <= i < j < k < n)
func threeSumSmaller(nums []int, target int) int {
	ans := 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i, v := range nums {
		k := target - v
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] < k {
				// 注意！
				ans += right - left
				left++
			} else {
				right--
			}
		}
	}

	return ans
}

// @lc code=end
