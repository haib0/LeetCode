/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */
package leetcode

// @lc code=start
func canWinNim(n int) bool {
	return n%4 != 0
}

// @lc code=end
