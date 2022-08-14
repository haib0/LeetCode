/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */
package leetcode

// @lc code=start
func rob213(nums []int) int {
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
	dpWithout0 := make([]int, n)
	dpWithout0[0] = 0
	dpWithout0[1] = nums[1]
	for i := 2; i < n; i++ {
		if dp[i-1] < dp[i-2]+nums[i] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
		if dpWithout0[i-1] < dpWithout0[i-2]+nums[i] {
			dpWithout0[i] = dpWithout0[i-2] + nums[i]
		} else {
			dpWithout0[i] = dpWithout0[i-1]
		}
	}
	ans := dpWithout0[n-1]
	if ans < dp[n-2] {
		ans = dp[n-2]
	}
	return ans
}

// @lc code=end
