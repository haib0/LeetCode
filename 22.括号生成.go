/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

package leetcode

// @lc code=start
func generateParenthesis(n int) []string {
	var ans []string

	var backtrace func(left, right int, trace string)
	backtrace = func(left, right int, trace string) {
		if left < right || left > n || right > n {
			return
		}
		if left+right == 2*n {
			ans = append(ans, trace)
			return
		}

		backtrace(left+1, right, trace+"(")
		backtrace(left, right+1, trace+")")
	}

	backtrace(0, 0, "")
	return ans
}

// @lc code=end
