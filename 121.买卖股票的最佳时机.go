/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package leetcode

// @lc code=start
func maxProfit(prices []int) int {
	var ans int
	minb := prices[0]
	for _, v := range prices {
		if ans < v-minb {
			ans = v - minb
		}
		if v < minb {
			minb = v
		}
	}
	return ans
}

// @lc code=end
