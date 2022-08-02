/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */
package leetcode

// @lc code=start
func canJump(nums []int) bool {
	var maxPos int
	n := len(nums)
	// note that exclude nums[n-1]
	for i := 0; i < n-1; i++ {
		if maxPos < i+nums[i] {
			maxPos = i + nums[i]
		}
		if i >= maxPos {
			return false
		}
	}

	return true
}

// @lc code=end
