/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */
package leetcode

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	var ans int
	dp := 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp++
			ans += dp
		} else {
			dp = 0
		}
	}
	return ans
}

// @lc code=end
