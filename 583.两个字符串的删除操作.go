/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */
package leetcode

// @lc code=start
func minDistance583(word1 string, word2 string) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i, c1 := range word1 {
		for j, c2 := range word2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = 1 + min(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[m][n]
}

// @lc code=end
