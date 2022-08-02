/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package leetcode

// @lc code=start
func moveZeroes(nums []int) {
	low, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[low] = nums[fast]
			low++
		}
		fast++
	}
	for low < len(nums) {
		nums[low] = 0
		low++
	}
}

// @lc code=end
