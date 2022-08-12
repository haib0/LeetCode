/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */
package leetcode

// @lc code=start

// 位运算: 异或
func singleNumber(nums []int) int {
	var ans int
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

// @lc code=end
