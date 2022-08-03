/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package leetcode

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = nums[1]
	if nums[0] > dp[1] {
		dp[1] = nums[0]
	}
	for i := 2; i < n; i++ {
		if dp[i-1] < dp[i-2]+nums[i] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}

// @lc code=end
