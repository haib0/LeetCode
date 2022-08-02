/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */
package leetcode

import (
	"strconv"
)

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		n := len(stack)
		var x, y int
		if n > 1 {
			x, y = stack[n-2], stack[n-1]
		}

		switch token {
		case "+":
			stack[n-2] = x + y
		case "-":
			stack[n-2] = x - y
		case "*":
			stack[n-2] = x * y
		case "/":
			stack[n-2] = x / y
		default:
			i, err := strconv.Atoi(token)
			if err != nil {
				return -1
			}
			stack = append(stack, i)
			continue
		}

		stack = stack[:n-1]
	}

	return stack[0]
}

// @lc code=end
