/*
 * @lc app=leetcode.cn id=921 lang=golang
 *
 * [921] 使括号有效的最少添加
 */
package leetcode

// @lc code=start
func minAddToMakeValid(s string) int {
	ans := 0

	need := 0
	for _, v := range s {
		switch v {
		case '(':
			need++
		case ')':
			need--
			if need == -1 {
				need = 0
				ans++
			}
		default:
			return -1
		}
	}

	return need + ans
}

// @lc code=end
