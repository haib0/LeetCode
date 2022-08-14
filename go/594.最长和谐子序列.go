/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */
package leetcode

// @lc code=start
func findLHS(nums []int) int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	result := 0
	for _, n := range nums {
		s, ok := m[n+1]
		if ok {
			if s+m[n] > result {
				result = s + m[n]
			}
		}
	}
	return result
}

// @lc code=end
