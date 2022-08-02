/*
 * @lc app=leetcode.cn id=1541 lang=golang
 *
 * [1541] 平衡括号字符串的最少插入次数
 */
package leetcode

// @lc code=start

func minInsertions(s string) int {
	ans := 0

	need := 0
	for _, v := range s {
		switch v {
		case '(':
			need += 2
			if need%2 != 0 {
				ans++
				need--
			}
		case ')':
			need--
			if need == -1 {
				need = 1
				ans++
			}
		default:
			return -1
		}
	}

	return need + ans
}

// @lc code=end
