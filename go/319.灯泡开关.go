/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 */
package leetcode

import "math"

// @lc code=start
func bulbSwitch(n int) int {
	/*
		第i个灯泡，其被切换的次数等于约数的个数
		e.g. i=16, 16 = 1*16 = 2*8=4*4, 约数有1, 2, 4, 8, 16, 奇数个,最后为开启状态
		故最后开启状态灯泡的个数等于n以内完全平方数的个数, 即sqrt(n)向下取整
	*/
	return int(math.Sqrt(float64(n)))
}

// @lc code=end
