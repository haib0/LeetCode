/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */
package leetcode

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dh := 0
	for _, n := range nums {
		if n != nums[dh] {
			nums[dh+1] = n
			dh++
		}
	}
	return dh + 1
}

// @lc code=end
