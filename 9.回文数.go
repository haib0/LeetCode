/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
package leetcode

// @lc code=start
func isPalindromeNum(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var rN int
	for x > rN {
		rN = rN*10 + x%10
		x /= 10
	}
	return x == rN || x == rN/10
}

// @lc code=end
