/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
package leetcode

// @lc code=start
func searchInsert(nums []int, target int) int {
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
	// if left >= len(nums) || nums[left] != target {
	// 	return -1
	// }
	return left
}

// @lc code=end
