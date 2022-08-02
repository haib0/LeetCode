/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package leetcode

// @lc code=start
func isValid(s string) bool {
	var stack []rune

	pair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			if v == pair[stack[len(stack)-1]] {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

// @lc code=end
