/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */
package leetcode

import "sort"

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var ans = [][]int{[]int{}}
	var viewed = make([]bool, len(nums))
	var backtrack func(start int, trace []int)
	backtrack = func(start int, trace []int) {
		for i := start; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !viewed[i-1] {
				continue
			}
			ans = append(ans, append([]int{}, append(trace, nums[i])...))
			viewed[i] = true
			backtrack(i+1, append(trace, nums[i]))
			viewed[i] = false
		}
	}

	backtrack(0, []int{})
	return ans
}

// @lc code=end
