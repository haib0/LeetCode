/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */
package leetcode

// @lc code=start
func trailingZeroes(n int) int {
	ans := 0


	for n >= 5 {
		n = n / 5
		ans += n
	}

	return ans
}
// @lc code=end

