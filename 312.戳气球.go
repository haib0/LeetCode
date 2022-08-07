/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */
package leetcode

// @lc code=start
func maxCoins(nums []int) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}

		return x
	}

	points := append(append([]int{1}, nums...), 1)

	n := len(points)
	// dp[i][j]: max points can get with all point in (i, j) is shooten.
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(
					dp[i][j],
					dp[i][k]+points[i]*points[k]*points[j]+dp[k][j],
				)
			}
		}
	}

	return dp[0][n-1]
}

// @lc code=end
