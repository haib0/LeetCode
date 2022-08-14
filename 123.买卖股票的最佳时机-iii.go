/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */
package leetcode

// @lc code=start
func maxProfit3(prices []int) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	// k: buy counts
	n, maxK := len(prices), 2
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, maxK+1)
		for k := 0; k <= maxK; k++ {
			dp[i][k] = make([]int, 2)
		}
	}

	for k := 0; k <= maxK; k++ {
		dp[0][k][1] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for k := 1; k <= maxK; k++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[n-1][maxK][0]
}

// @lc code=end
