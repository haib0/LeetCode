/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */
package leetcode
// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2

		if nums[mid] < target {
			left = mid + 1
			continue
		}

		if nums[mid] > target {
			right = mid
			continue
		}

		right = mid
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}
// @lc code=end

