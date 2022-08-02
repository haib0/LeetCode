/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */
package leetcode

// @lc code=start
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	x_a := abs(ax2 - ax1)
	y_a := abs(ay2 - ay1)

	x_b := abs(bx2 - bx1)
	y_b := abs(by2 - by1)

	l_x := x_a + x_b
	l_y := y_a + y_b

	x_1 := abs(bx1 - ax1)
	x_2 := abs(bx2 - ax2)

	y_1 := abs(by1 - ay1)
	y_2 := abs(by2 - ay2)

	l_xx := x_1 + x_2
	l_yy := y_1 + y_2

	x := (l_x - l_xx) / 2
	y := (l_y - l_yy) / 2

	area := x * y
	if x <= 0 || y <= 0 {
		area = 0
	}

	return x_a*y_a + x_b*y_b - area
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// @lc code=end
