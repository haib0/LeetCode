/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */
package leetcode

// @lc code=start
func sortColors(nums []int) {
	var swap = func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}
	var p0, p1, p2 = 0, 0, len(nums) - 1
	for p1 <= p2 {
		if nums[p1] == 0 {
			swap(p1, p0)
			p0++
			p1++
		} else if nums[p1] == 2 {
			swap(p1, p2)
			p2--
		} else {
			p1++
		}
	}
}

// @lc code=end
