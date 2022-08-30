/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */
package leetcode

import "sort"

// @lc code=start
func permuteUnique(nums []int) [][]int {
	// sort first to avoid duplicate
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	var vis = make([]bool, n)
	var trace = []int{}
	var backtracking func(idx int)
	backtracking = func(idx int) {
		if idx == n {
			// need deep copy!
			ans = append(ans, append([]int{}, trace...))
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			// to avoid duplicate
			if i > 0 && nums[i] == nums[i-1] && !vis[i-1] {
				continue
			}
			trace = append(trace, nums[i])
			vis[i] = true
			backtracking(idx + 1)
			trace = trace[:len(trace)-1]
			vis[i] = false
		}
	}
	backtracking(0)
	return ans
}

// @lc code=end
