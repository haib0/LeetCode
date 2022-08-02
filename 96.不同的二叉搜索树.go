/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

package leetcode

// @lc code=start

func numTrees(n int) int {
	var dp = make([]int, n+1)

	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for ii := 0; ii < i; ii++ {
			dp[i] += dp[ii] * dp[i-ii-1]
		}
	}

	return dp[n]
}

// @lc code=end
