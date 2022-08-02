/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
package leetcode

// @lc code=start
func maxSubArray(nums []int) int {
	max, dp := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if dp <= 0 {
			dp = nums[i]
		} else {
			dp += nums[i]
		}

		if dp > max {
			max = dp
		}
	}
	return max
}

// @lc code=end
