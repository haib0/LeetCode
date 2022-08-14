/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

package leetcode

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	var ans int

	n := len(nums)
	var r int
	var backtracking func(start int)
	backtracking = func(start int) {
		if start == n {
			if r == target {
				ans++
			}
			return
		}

		r += nums[start]
		backtracking(start + 1)
		r -= nums[start]

		r -= nums[start]
		backtracking(start + 1)
		r += nums[start]
	}

	backtracking(0)
	return ans
}

// @lc code=end
