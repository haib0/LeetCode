/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */
package leetcode

import (
	"sort"
)

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	// 按每个区间的结束排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans := 0
	r := intervals[0][1]
	for _, v := range intervals[1:] {
		if v[0] < r {
			ans++
		} else {
			r = v[1]
		}
	}
	return ans
}

// @lc code=end
