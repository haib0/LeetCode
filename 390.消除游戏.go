/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] 消除游戏
 */
package leetcode

// @lc code=start
func lastRemaining(n int) int {
	a1, step, k := 1, 1, 0
	for n > 1 {
		if n%2 == 0 && k%2 != 0 {
			// 偶数个,并且从尾部开始删除
		} else {
			a1 += step
		}
		step *= 2
		n /= 2
		k++
	}
	return a1
}

// @lc code=end
