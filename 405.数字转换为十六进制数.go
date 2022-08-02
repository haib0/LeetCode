/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] 数字转换为十六进制数
 */
package leetcode

import "fmt"

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = 4294967295 + num + 1
	}
	result := ""
	for num >= 16 {
		switch num % 16 {
		case 10:
			result = "a" + result
		case 11:
			result = "b" + result
		case 12:
			result = "c" + result
		case 13:
			result = "d" + result
		case 14:
			result = "e" + result
		case 15:
			result = "f" + result
		default:
			result = fmt.Sprint(num%16) + result
		}
		num /= 16
	}
	switch num {
	case 0:
		break
	case 10:
		result = "a" + result
	case 11:
		result = "b" + result
	case 12:
		result = "c" + result
	case 13:
		result = "d" + result
	case 14:
		result = "e" + result
	case 15:
		result = "f" + result
	default:
		result = fmt.Sprint(num) + result
	}

	return result
}

// @lc code=end
