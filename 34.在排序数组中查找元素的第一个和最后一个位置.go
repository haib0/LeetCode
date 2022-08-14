/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
package leetcode

// @lc code=start
func searchRange(nums []int, target int) []int {

	return []int{leftBound(nums, target), rightBound(nums, target)}
}

func leftBound(nums []int, target int) int {
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

func rightBound(nums []int, target int) int {
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

		left = mid + 1
	}

	right--
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

// @lc code=end
