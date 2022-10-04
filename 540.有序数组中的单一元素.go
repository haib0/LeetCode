/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 */
package leetcode

// @lc code=start
func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if mid%2 != 0 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

// @lc code=end
