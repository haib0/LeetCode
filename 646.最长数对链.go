/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */
package leetcode

import "sort"

// @lc code=start
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	var ans = 1
	n := len(pairs)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans
}

// @lc code=end
