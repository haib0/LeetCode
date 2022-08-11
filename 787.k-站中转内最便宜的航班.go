/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 */
package leetcode

// @lc code=start
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
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

	const INF = 10000*101 + 1

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+2)
		for j := 0; j < k+2; j++ {
			dp[i][j] = INF
		}
	}

	dp[src][0] = 0

	for ki := 1; ki < k+2; ki++ {
		for _, v := range flights {
			from, to, cost := v[0], v[1], v[2]
			dp[to][ki] = minInt(dp[to][ki], dp[from][ki-1]+cost)
		}
	}

	ans := minInt(dp[dst]...)
	if ans == INF {
		return -1
	}
	return ans
}

// @lc code=end
