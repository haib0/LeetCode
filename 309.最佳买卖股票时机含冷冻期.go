/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */
package leetcode

// @lc code=start
func maxProfitf(prices []int) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[1][0] = max(0, prices[1]-prices[0])
	dp[1][1] = max(-prices[0], -prices[1])
	for i := 2; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-2][0]-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}

// @lc code=end
