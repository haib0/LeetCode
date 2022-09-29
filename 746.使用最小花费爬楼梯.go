/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */
package leetcode

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	var minInt = func(nums ...int) int {
		if len(nums) == 0 {
			panic(nil)
		}

		x := nums[0]
		for _, v := range nums[1:] {
			if v < x {
				x = v
			}
		}

		return x
	}

	var n = len(cost)
	if n == 0 {
		return 0
	}
	var dp = make([]int, n+1) // dp[i+1] is the min cost to i
	dp[1] = cost[0]
	for i := 2; i <= n; i++ {
		dp[i] = minInt(dp[i-2], dp[i-1]) + cost[i-1]
	}
	return minInt(dp[n], dp[n-1])
}

// @lc code=end
