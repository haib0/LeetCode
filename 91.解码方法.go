/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
package leetcode

import "strconv"

// @lc code=start
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if s[i-1] == '0' {
			if s[i-2] != '1' && s[i-2] != '2' {
				return 0
			}
			dp[i-1] = 0
			dp[i] = dp[i-2]
			continue
		}
		dp[i] = dp[i-1]

		x, _ := strconv.Atoi(s[i-2 : i])
		if x >= 1 && x <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

// @lc code=end
