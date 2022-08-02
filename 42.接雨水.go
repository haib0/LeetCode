/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package leetcode

// @lc code=start
func trap(height []int) int {
	l, r := 0, len(height)-1
	maxl, maxr := height[l], height[r]

	ans := 0
	for l < r {
		if maxl < maxr {
			ans += maxl - height[l]
			l++
			maxl = maxInt(maxl, height[l])
		} else {
			ans += maxr - height[r]
			r--
			maxr = maxInt(maxr, height[r])
		}
	}
	return ans
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}

	return x
}

// @lc code=end
