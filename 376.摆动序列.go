/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */
package leetcode

// @lc code=start
func wiggleMaxLength(nums []int) int {
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		}
		if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up < down {
		return down
	}
	return up
}

// @lc code=end
