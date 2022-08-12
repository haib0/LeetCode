/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */
package leetcode

// @lc code=start

// 数学法：逆向思维
func superEggDrop(k int, n int) int {
	// dp[k][m] = n: k 个鸡蛋, 扔 m 次,
	// 最多能确切测试一栋 n 层的楼
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for m := 1; m <= n; m++ {
		for i := 1; i <= k; i++ {
			// 扔一次鸡蛋: + 1
			// - 鸡蛋碎了: + dp[i-1][m-1]
			// - 鸡蛋没碎: + dp[i][m-1]
			dp[i][m] = 1 + dp[i-1][m-1] + dp[i][m-1]
		}
		// ans: 找到满足 dp[k][m] >= n 的最小 m
		if dp[k][m] >= n {
			return m
		}
	}

	return -1
}

// @lc code=end
