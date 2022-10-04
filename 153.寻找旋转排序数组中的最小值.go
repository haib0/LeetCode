/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
package leetcode

// @lc code=start
func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] < nums[hi] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return nums[lo]
}

// @lc code=end
