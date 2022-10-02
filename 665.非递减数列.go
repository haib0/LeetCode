/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */
package leetcode

// @lc code=start
func checkPossibility(nums []int) bool {
	var count int
	for i := 1; i < len(nums) && count <= 1; i++ {
		if nums[i] >= nums[i-1] {
			continue
		}
		// need to change a number
		// nums[i]=nums[i-1] or nums[i-1]=nums[i] ?
		if i-2 >= 0 && nums[i] < nums[i-2] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
		count++

	}
	return count <= 1
}

// @lc code=end
