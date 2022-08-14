/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */
package leetcode

// @lc code=start
func integerReplacement(n int) int {
	min := func(i1, i2 int) int {
		if i1 < i2 {
			return i1
		}
		return i2
	}

	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	} else {
		return 1 + min(integerReplacement(n+1), integerReplacement(n-1))
	}
}

// @lc code=end
