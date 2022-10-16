/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */
package leetcode

// @lc code=start
func integerBreak(n int) int {
	var maxInt = func(nums ...int) int {
		if len(nums) == 0 {
			panic(nil)
		}
		x := nums[0]
		for _, v := range nums[1:] {
			if v > x {
				x = v
			}
		}
		return x
	}

	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			dp[i] = maxInt(dp[i], dp[j]*(i-j), j*(i-j))
		}
	}
	return dp[n]
}

// @lc code=end
