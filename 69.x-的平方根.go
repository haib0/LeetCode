/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
package leetcode

// @lc code=start
func mySqrt(x int) int {
	i, j := 1, x
	for i <= j {
		mid := (i + j) / 2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		} else if sqrt < mid {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return j
}

// @lc code=end
