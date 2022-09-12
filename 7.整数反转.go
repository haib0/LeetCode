/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
package leetcode

import "math"

// @lc code=start
func reverse(x int) int {
	var ans int
	for x != 0 {
		ans = ans*10 + x%10
		if ans < math.MinInt32 || ans > math.MaxInt32 {
			return 0
		}
		x = x / 10
	}
	return ans
}

// @lc code=end
