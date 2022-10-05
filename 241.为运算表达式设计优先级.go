/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */
package leetcode

import "strconv"

// @lc code=start
func diffWaysToCompute(expression string) []int {
	var ans []int
	for i, v := range expression {
		if v == '+' || v == '-' || v == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch v {
					case '+':
						ans = append(ans, l+r)
					case '-':
						ans = append(ans, l-r)
					case '*':
						ans = append(ans, l*r)
					}
				}
			}
		}
	}
	if len(ans) == 0 {
		x, _ := strconv.Atoi(expression)
		ans = append(ans, x)
	}
	return ans
}

// @lc code=end
