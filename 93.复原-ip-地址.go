/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */
package leetcode

import (
	"strconv"
)

// @lc code=start
func restoreIpAddresses(s string) []string {
	// x is an ip piece that not start with 0.
	var isValid = func(x string) bool {
		i, err := strconv.Atoi(x)
		if err != nil {
			return false
		}
		if i < 0 || i > 255 {
			return false
		}
		return true
	}
	var ans []string
	var backtrack func(k, start int, curr string)
	backtrack = func(k, start int, curr string) {
		if k == 4 || start == len(s) {
			if k == 4 && start == len(s) {
				ans = append(ans, curr)
			}
			return
		}
		if s[start] == '0' {
			next := curr + "0"
			if k < 3 || start < len(s)-1 {
				next += "."
			}
			backtrack(k+1, start+1, next)
			return
		}
		for i := start; i < len(s) && i < start+3; i++ {
			x := s[start : i+1]
			if isValid(x) {
				next := curr + x
				if k < 3 || i < len(s)-1 {
					next += "."
				}
				backtrack(k+1, i+1, next)
			}
		}
	}

	backtrack(0, 0, "")
	return ans
}

// @lc code=end
