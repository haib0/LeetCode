/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 */
package leetcode

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	var minInt = func(nums ...int) int {
		if len(nums) == 0 {
			panic(nil)
		}

		x := nums[0]
		for _, v := range nums[1:] {
			if v < x {
				x = v
			}
		}

		return x
	}

	const MAX = 100000
	n := len(matrix)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+2)
		dp[i][0] = MAX
		dp[i][n+1] = MAX
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = matrix[0][j-1]
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = matrix[i][j-1] + minInt(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])
		}
	}

	return minInt(dp[n-1]...)
}

// @lc code=end
