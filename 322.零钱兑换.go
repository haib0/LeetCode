/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package leetcode

// @lc code=start
func coinChange(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dp = make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		// max count of coins eqs amount
		dp[i] = amount + 1
	}
	for am := 1; am <= amount; am++ {
		for _, c := range coins {
			if c <= am {
				dp[am] = min(dp[am], dp[am-c]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// @lc code=end
