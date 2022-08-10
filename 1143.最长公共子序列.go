/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */
package leetcode

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
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

	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = maxInt(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[m][n]
}

// @lc code=end
