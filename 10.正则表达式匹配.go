/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */
package leetcode

// @lc code=start

func isMatch(s string, p string) bool {
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	m, n := len(s), len(p)
	var dp [][]bool
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]bool, n+1))
	}

	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if matches(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				if matches(i, j) {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			}

		}
	}

	return dp[m][n]
}

// @lc code=end
