/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */
package leetcode

// @lc code=start
func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	intSize := 32 << (^uint(0) >> 63)
	MinInt := -1 << (intSize - 1)
	m1, m2, m3 := MinInt, MinInt, MinInt
	for _, n := range nums {
		if n > m1 {
			m3 = m2
			m2 = m1
			m1 = n
		} else if n < m1 && n > m2 {
			m3 = m2
			m2 = n
		} else if n < m2 && n > m3 {
			m3 = n
		}
	}
	if m2 == MinInt || m3 == MinInt {
		return m1
	}
	return m3
}

// @lc code=end
