/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */
package leetcode

import "math"

// @lc code=start
func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		s := i*i + j*j
		if s == c {
			return true
		}
		if s < c {
			i++
		} else {
			j--
		}
	}
	return false
}

// @lc code=end
