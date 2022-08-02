/*
 * @lc app=leetcode.cn id=1099 lang=golang
 *
 * [1099] 较小的三数之和
 */
package leetcode

import "sort"

// @lc code=start

//  给你⼀个整数数组 nums 和整数 k,返回最⼤和 sum,
//  满⾜ 存在 i < j 使得 nums[i] + nums[j] = sum 且 sum < k
//  如果没有满⾜此等式的 i, j 存在, 则返回 -1
func twoSumLessThanK(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := -1
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < k {
			if sum > ans {
				ans = sum
			}
			left++
		} else {
			right--
		}
	}
	return ans
}

// @lc code=end
