/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */
package leetcode

// @lc code=start

// 位运算: 异或
func missingNumber(nums []int) int {
	var ans int
	ans ^= len(nums)
	for i, v := range nums {
		ans ^= i ^ v
	}
	return ans
}

// @lc code=end
