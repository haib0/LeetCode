/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
package leetcode

import "sort"

// @lc code=start

func longestValidParentheses(s string) int {
	n := len(s)
	var stack []int
	dp := make([]int, n+1)
	for i, v := range s {
		switch v {
		case '(':
			stack = append(stack, i)
			dp[i+1] = 0

		case ')':
			nstack := len(stack)
			if nstack == 0 {
				dp[i+1] = 0
				continue
			}

			il := stack[nstack-1]
			stack = stack[:nstack-1]

			len := i - il + 1 + dp[il]
			dp[i+1] = len

		default:
			return -1
		}
	}

	sort.Slice(dp, func(i, j int) bool {
		return dp[i] > dp[j]
	})

	return dp[0]
}

// @lc code=end
