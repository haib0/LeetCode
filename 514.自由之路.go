/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 */
package leetcode

import "math"

// @lc code=start
func findRotateSteps(ring string, key string) int {
	min := func(s ...int) int {
		a := s[0]
		for _, v := range s[1:] {
			if v < a {
				a = v
			}
		}
		return a
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	n, m := len(ring), len(key)

	pos := [26][]int{}
	for i, v := range ring {
		pos[v-'a'] = append(pos[v-'a'], i)
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = min(p, n-p) + 1
	}

	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, p := range pos[key[i-1]-'a'] {
				dp[i][j] = min(
					dp[i][j],
					dp[i-1][p]+min(abs(p-j), n-abs(p-j))+1,
				)
			}
		}
	}

	return min(dp[m-1]...)
}

// @lc code=end
