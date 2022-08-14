/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */
package leetcode

// @lc code=start
func jump(nums []int) int {
	var step int
	var end, maxPos int
	n := len(nums)
	// note that exclude nums[n-1]
	for i := 0; i < n-1; i++ {
		if maxPos < i+nums[i] {
			maxPos = i + nums[i]
		}
		if i == end {
			step++
			end = maxPos
		}
	}
	return step
}

// @lc code=end
