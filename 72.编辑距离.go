/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
package leetcode

// @lc code=start
func minDistance(word1 string, word2 string) int {
	min3 := func(a, b, c int) int {
		if a < b && a < c {
			return a
		} else if b < c {
			return b
		} else {
			return c
		}
	}

	m, n := len(word1), len(word2)
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min3(
					dp[i-1][j]+1,
					dp[i][j-1]+1,
					dp[i-1][j-1]+1,
				)
			}
		}
	}

	return dp[m][n]
}

// @lc code=end
