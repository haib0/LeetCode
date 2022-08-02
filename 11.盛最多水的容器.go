/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package leetcode

// @lc code=start
func maxArea(height []int) int {
	ans := 0

	left, right := 0, len(height)-1

	for left < right {
		h := height[left]
		if height[left] > height[right] {
			h = height[right]
		}

		area := (right - left) * h
		if area > ans {
			ans = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}

// @lc code=end
