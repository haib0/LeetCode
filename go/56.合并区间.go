/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
package leetcode

import (
	"sort"
)

// @lc code=start
func merge56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	start, end := intervals[0][0], intervals[0][1]
	for _, v := range intervals {
		if v[0] <= end {
			if end < v[1] {
				end = v[1]
			}
		} else {
			ans = append(ans, []int{start, end})
			start = v[0]
			end = v[1]
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}

// @lc code=end
