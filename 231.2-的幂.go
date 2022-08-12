/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2 的幂
 */
package leetcode

// @lc code=start

// 位运算: n&(n-1)
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

// @lc code=end
