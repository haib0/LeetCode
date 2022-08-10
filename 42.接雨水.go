/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package leetcode

// @lc code=start
func trap(height []int) int {
	var maxInt = func(nums ...int) int {
		if len(nums) == 0 {
			panic(nil)
		}

		x := nums[0]
		for _, v := range nums[1:] {
			if v > x {
				x = v
			}
		}

		return x
	}

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

// @lc code=end
