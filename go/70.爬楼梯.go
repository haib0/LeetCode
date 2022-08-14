/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
package leetcode

// @lc code=start
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	pp, p := 1, 2
	for i := 3; i <= n; i++ {
		r := pp + p
		pp = p
		p = r
	}
	return p
}

// @lc code=end
